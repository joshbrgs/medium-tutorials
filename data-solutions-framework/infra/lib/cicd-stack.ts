import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import { RemovalPolicy } from "aws-cdk-lib";
import * as iam from "aws-cdk-lib/aws-iam";
import * as pipelines from "aws-cdk-lib/pipelines";
import * as dsf from "@cdklabs/aws-data-solutions-framework";
import { SparkApplicationStackFactory } from "./application-stack";

export class CICDPipelineStack extends cdk.Stack {
    constructor(scope: Construct, id: string, props?: cdk.StackProps) {
        super(scope, id, props);

        new dsf.processing.SparkEmrCICDPipeline(this, "SparkCICDPipeline", {
            sparkApplicationName: "SparkTest",
            // Pass the factory class to dynamically pass the Application Stack
            applicationStackFactory: new SparkApplicationStackFactory(),
            // Path of the CDK TypeScript application to be used by the CICD build and deploy phases
            cdkApplicationPath: "infra",
            // Path of the Spark application to be built and unit tested in the CICD
            sparkApplicationPath: "spark",
            // Path of the bash script responsible to run integration tests
            integTestScript: "./infra/resources/integ-test.sh",
            // Environment variables used by the integration test script
            integTestEnv: {
                STEP_FUNCTION_ARN: "ProcessingStateMachineArn",
            },
            // Additional permissions to give to the CICD to run the integration tests
            integTestPermissions: [
                new iam.PolicyStatement({
                    actions: [
                        "states:StartExecution",
                        "states:DescribeExecution",
                    ],
                    resources: ["*"],
                }),
            ],
            source: pipelines.CodePipelineSource.connection(
                "github.com/datasolutionsframework",
                "main",
                {
                    connectionArn:
                        "arn:aws:codeconnections:region:xxxxxxxxxxx:connection/xxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
                },
            ),
            removalPolicy: RemovalPolicy.DESTROY,
        });
    }
}
