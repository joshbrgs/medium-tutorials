import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import {
    aws_events as events,
    aws_iam as iam,
    aws_s3 as s3,
    CfnOutput,
    Duration,
    Names,
    RemovalPolicy,
} from "aws-cdk-lib";
import * as dsf from "@cdklabs/aws-data-solutions-framework";

export class SparkApplicationStackFactory
    extends dsf.utils.ApplicationStackFactory {
    /**
     * Creates a stack for a Spark EMR application.
     */
    createStack(scope: Construct, stage: dsf.utils.CICDStage): cdk.Stack {
        return new ApplicationStack(scope, "EmrApplicationStack", stage);
    }
}

export class ApplicationStack extends cdk.Stack {
    constructor(
        scope: Construct,
        id: string,
        private stage?: dsf.utils.CICDStage,
        props?: cdk.StackProps,
    ) {
        super(scope, id, props);

        // Create Data Lake Storage
        const storage = new dsf.storage.DataLakeStorage(
            this,
            "DataLakeStorage",
            {
                removalPolicy: RemovalPolicy.DESTROY,
            },
        );

        // Create Data Lake Catalog
        const catalog = new dsf.governance.DataLakeCatalog(
            this,
            "DataLakeCatalog",
            {
                dataLakeStorage: storage,
                databaseName: "spark_data_lake",
                removalPolicy: RemovalPolicy.DESTROY,
            },
        );

        // Import Source Bucket
        const sourceBucket = s3.Bucket.fromBucketName(
            this,
            "SourceBucket",
            "redshift-demos",
        );

        // Copy Taxi Data
        this.copyTaxiData(storage, sourceBucket);

        // Create Execution Role
        const processingExecRole = this.createExecutionRole(storage, catalog);

        // Create Spark Runtime
        const sparkRuntime = this.createSparkRuntime();

        // Package Spark Application
        const sparkApp = this.packageSparkApplication(processingExecRole);

        // Schedule Spark Job
        const sparkJob = this.scheduleSparkJob(
            storage,
            catalog,
            sparkRuntime,
            sparkApp,
            processingExecRole,
        );

        // Athena Workgroup
        this.createAthenaWorkgroup();

        // Outputs
        new CfnOutput(this, "ProcessingStateMachineArn", {
            value: sparkJob.stateMachine!.stateMachineArn,
        });
    }

    private copyTaxiData(
        storage: dsf.storage.DataLakeStorage,
        sourceBucket: s3.IBucket,
    ): void {
        const dataSpecs = [
            {
                id: "YellowDataCopy",
                prefix: "data/NY-Pub/year=2016/month=1/type=yellow",
                targetPrefix: "yellow-trip-data/",
            },
            {
                id: "GreenDataCopy",
                prefix: "data/NY-Pub/year=2016/month=1/type=green",
                targetPrefix: "green-trip-data/",
            },
        ];

        dataSpecs.forEach((spec) => {
            new dsf.utils.S3DataCopy(this, spec.id, {
                sourceBucket,
                sourceBucketPrefix: spec.prefix,
                sourceBucketRegion: "us-east-1",
                targetBucket: storage.silverBucket,
                targetBucketPrefix: spec.targetPrefix,
                removalPolicy: RemovalPolicy.DESTROY,
            });
        });
    }

    private createExecutionRole(
        storage: dsf.storage.DataLakeStorage,
        catalog: dsf.governance.DataLakeCatalog,
    ): cdk.aws_iam.IRole {
        const role = dsf.processing.SparkEmrServerlessRuntime
            .createExecutionRole(this, "ProcessingExecRole");

        // Grant Permissions
        storage.goldBucket.grantReadWrite(role);
        storage.silverBucket.grantRead(role);

        const account = cdk.Stack.of(this).account;
        const region = cdk.Stack.of(this).region;
        const targetDb = catalog.goldCatalogDatabase.databaseName;
        const targetTable = "aggregated_trip_distance";

        role.addToPrincipalPolicy(
            new iam.PolicyStatement({
                effect: iam.Effect.ALLOW,
                actions: [
                    "glue:CreateTable",
                    "glue:GetTable",
                    "glue:GetTables",
                    "glue:BatchGetPartition",
                    "glue:GetDatabase",
                    "glue:GetDatabases",
                ],
                resources: [
                    `arn:aws:glue:${region}:${account}:catalog`,
                    `arn:aws:glue:${region}:${account}:database/default`,
                    `arn:aws:glue:${region}:${account}:database/${targetDb}`,
                    `arn:aws:glue:${region}:${account}:table/${targetDb}/${targetTable}`,
                ],
            }),
        );

        return role;
    }

    private createSparkRuntime(): dsf.processing.SparkEmrServerlessRuntime {
        return new dsf.processing.SparkEmrServerlessRuntime(
            this,
            "SparkProcessingRuntime",
            {
                name: "TaxiAggregation",
                runtimeConfiguration: [
                    {
                        classification: "spark-defaults",
                        properties: {
                            "spark.hadoop.hive.metastore.client.factory.class":
                                "com.amazonaws.glue.catalog.metastore.AWSGlueDataCatalogHiveClientFactory",
                            "spark.sql.catalogImplementation": "hive",
                        },
                    },
                ],
                removalPolicy: RemovalPolicy.DESTROY,
            },
        );
    }

    private packageSparkApplication(
        executionRole: cdk.aws_iam.IRole,
    ): dsf.processing.PySparkApplicationPackage {
        const sparkApp = new dsf.processing.PySparkApplicationPackage(
            this,
            "SparkApp",
            {
                entrypointPath: "./../spark/src/agg_trip_distance.py",
                applicationName: "TaxiAggregation",
                dependenciesFolder: "./../spark",
                venvArchivePath: "/venv-package/pyspark-env.tar.gz",
                removalPolicy: RemovalPolicy.DESTROY,
            },
        );

        sparkApp.artifactsBucket.grantReadWrite(executionRole);
        return sparkApp;
    }

    private scheduleSparkJob(
        storage: dsf.storage.DataLakeStorage,
        catalog: dsf.governance.DataLakeCatalog,
        sparkRuntime: dsf.processing.SparkEmrServerlessRuntime,
        sparkApp: dsf.processing.PySparkApplicationPackage,
        executionRole: iam.IRole,
    ): dsf.processing.SparkEmrServerlessJob {
        const sparkParams = [
            ` --conf spark.emr-serverless.driverEnv.YELLOW_SOURCE=s3://${storage.silverBucket.bucketName}/yellow-trip-data/`,
            ` --conf spark.emr-serverless.driverEnv.GREEN_SOURCE=s3://${storage.silverBucket.bucketName}/green-trip-data/`,
            ` --conf spark.emr-serverless.driverEnv.TARGET_DB=${catalog.goldCatalogDatabase.databaseName}`,
            ` --conf spark.emr-serverless.driverEnv.TARGET_TABLE=aggregated_trip_distance`,
        ].join("");

        const schedule = this.stage !== dsf.utils.CICDStage.PROD
            ? undefined
            : events.Schedule.rate(Duration.days(1));

        return new dsf.processing.SparkEmrServerlessJob(
            this,
            "SparkProcessingJob",
            {
                name: `taxi-agg-job-${Names.uniqueResourceName(this, {})}`,
                applicationId: sparkRuntime.application.attrApplicationId,
                executionRole,
                sparkSubmitEntryPoint: sparkApp.entrypointUri,
                sparkSubmitParameters: sparkApp.sparkVenvConf + sparkParams,
                removalPolicy: RemovalPolicy.DESTROY,
                schedule,
            },
        );
    }

    private createAthenaWorkgroup(): void {
        new dsf.consumption.AthenaWorkGroup(this, "InteractiveQuery", {
            name: "discovery",
            resultLocationPrefix: "athena-results",
            publishCloudWatchMetricsEnabled: false,
            engineVersion: dsf.consumption.EngineVersion.ATHENA_V3,
            recursiveDeleteOption: true,
            removalPolicy: RemovalPolicy.DESTROY,
            resultsRetentionPeriod: Duration.days(1),
        });
    }
}
