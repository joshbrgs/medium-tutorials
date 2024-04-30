package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create VPC
		vpc, err := ec2.NewVpc(ctx, "producer-vpc", &ec2.VpcArgs{
			CidrBlock:          pulumi.String("10.0.0.0/16"),
			EnableDnsHostnames: pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("producer-vpc"),
			},
		})
		if err != nil {
			return err
		}

		// Create private subnet
		subnet, err := ec2.NewSubnet(ctx, "producer-private-subnet", &ec2.SubnetArgs{
			VpcId:               vpc.ID(),
			CidrBlock:           pulumi.String("10.0.128.0/24"),
			MapPublicIpOnLaunch: pulumi.Bool(false),
			AvailabilityZone:    pulumi.String("us-east-1b"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("producer-private-subnet"),
			},
		})
		if err != nil {
			return err
		}

		// Create private subnet
		subnet1, err := ec2.NewSubnet(ctx, "producer-private-subnet1", &ec2.SubnetArgs{
			VpcId:               vpc.ID(),
			CidrBlock:           pulumi.String("10.0.0.0/24"),
			MapPublicIpOnLaunch: pulumi.Bool(false),
			AvailabilityZone:    pulumi.String("us-east-1a"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("producer-private-subnet-2"),
			},
		})
		if err != nil {
			return err
		}

		// Create a subnet group for the RDS instance
		subnetGroup, err := rds.NewSubnetGroup(ctx, "db-sub-grp", &rds.SubnetGroupArgs{
			SubnetIds: pulumi.StringArray{
				subnet.ID(),
				subnet1.ID(),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("my-db-subnet-group"),
			},
		})
		if err != nil {
			return err
		}

		// Security Group for the RDS instance
		rdsSecurityGroup, err := ec2.NewSecurityGroup(ctx, "rdsSecurityGroup", &ec2.SecurityGroupArgs{
			Description: pulumi.String("Allow MySQL inbound traffic"),
			VpcId:       vpc.ID(),
		})
		if err != nil {
			return err
		}

		// Security Group for the application server
		proxySecurityGroup, err := ec2.NewSecurityGroup(ctx, "proxySecurityGroup", &ec2.SecurityGroupArgs{
			Description: pulumi.String("Allow MySQL traffic to RDS"),
			VpcId:       vpc.ID(),
		})
		if err != nil {
			return err
		}

		// Ingress rule for the RDS security group allowing MySQL traffic from the application server security group
		_, err = ec2.NewSecurityGroupRule(ctx, "rdsSecurityGroupIngress", &ec2.SecurityGroupRuleArgs{
			Type:                  pulumi.String("ingress"),
			FromPort:              pulumi.Int(3306),
			ToPort:                pulumi.Int(3306),
			Protocol:              pulumi.String("tcp"),
			SecurityGroupId:       rdsSecurityGroup.ID(),
			SourceSecurityGroupId: proxySecurityGroup.ID(),
			Description:           pulumi.String("MySQL access from proxy security group"),
		})
		if err != nil {
			return err
		}

		// Egress rule for the application server security group allowing MySQL traffic to the RDS security group
		_, err = ec2.NewSecurityGroupRule(ctx, "proxySecurityGroupEgress", &ec2.SecurityGroupRuleArgs{
			Type:                  pulumi.String("egress"),
			FromPort:              pulumi.Int(3306),
			ToPort:                pulumi.Int(3306),
			Protocol:              pulumi.String("tcp"),
			SecurityGroupId:       proxySecurityGroup.ID(),
			SourceSecurityGroupId: rdsSecurityGroup.ID(),
			Description:           pulumi.String("MySQL access to RDS security group"),
		})
		if err != nil {
			return err
		}

		// Ingress rule for the application server security group allowing MySQL traffic from the VPC CIDR
		_, err = ec2.NewSecurityGroupRule(ctx, "proxySecurityGroupIngress", &ec2.SecurityGroupRuleArgs{
			Type:            pulumi.String("ingress"),
			FromPort:        pulumi.Int(3306),
			ToPort:          pulumi.Int(3306),
			Protocol:        pulumi.String("tcp"),
			SecurityGroupId: proxySecurityGroup.ID(),
			CidrBlocks:      pulumi.StringArray{pulumi.String("10.0.0.0/16")},
			Description:     pulumi.String("MySQL access from within the VPC"),
		})
		if err != nil {
			return err
		}

		// Create a new AWS Secrets Manager secret to store the database credentials
		secret, err := secretsmanager.NewSecret(ctx, "secret", &secretsmanager.SecretArgs{
			Name: pulumi.String("dbPassword"),
		})
		if err != nil {
			return err
		}

		// Generate a random password and assign it to the secret
		secretVersion, err := secretsmanager.NewSecretVersion(ctx, "secret-version", &secretsmanager.SecretVersionArgs{
			SecretId:     secret.ID(),
			SecretString: pulumi.String("mySuperSecretPassword123!"),
		})
		if err != nil {
			return err
		}

		// Create an IAM role for the RDS proxy
		proxyRole, err := iam.NewRole(ctx, "proxy-role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(`{
				"Version": "2012-10-17",
				"Statement": [{
					"Effect": "Allow",
					"Principal": {
						"Service": "rds.amazonaws.com"
					},
					"Action": "sts:AssumeRole"
				}]
			}`),
		})
		if err != nil {
			return err
		}

		policyProxy, err := iam.NewPolicy(ctx, "proxy-secret-policy", &iam.PolicyArgs{
			Policy: pulumi.String(`{
                "Version": "2012-10-17",
                "Statement": [
                    {
                        "Effect": "Allow",
                        "Action": [
                            "rds:*"
                        ],
                        "Resource": "*"
                    },
                    {
                        "Effect": "Allow",
                        "Action": [
                            "secretsmanager:*"
                        ],
                        "Resource": "arn:aws:secretsmanager:us-east-1:*:secret:dbPassword*"
                    } 
                ]
            }`),
		})
		if err != nil {
			return err
		}

		// Attach the required policy to the role
		_, err = iam.NewRolePolicyAttachment(ctx, "proxy-role-attachment", &iam.RolePolicyAttachmentArgs{
			Role:      proxyRole.ID(),
			PolicyArn: policyProxy.Arn,
		})
		if err != nil {
			return err
		}

		// Create an RDS instance using the password from Secrets Manager
		rdsI, err := rds.NewInstance(ctx, "db", &rds.InstanceArgs{
			InstanceClass:     pulumi.String(rds.InstanceType_T3_Micro), // Smallest instance size (as of AWS' current offering)
			AllocatedStorage:  pulumi.Int(20),                           // Minimum allocated storage for an RDS instance in GB
			Engine:            pulumi.String("mysql"),
			EngineVersion:     pulumi.String("5.7"),
			DbSubnetGroupName: subnetGroup.Name,
			VpcSecurityGroupIds: pulumi.StringArray{
				rdsSecurityGroup.ID(), // Associate the security group with the RDS instance
			},
			Username: pulumi.String("admin"),
			Password: secretVersion.SecretString,
			// ... other configuration ...
		})
		if err != nil {
			return err
		}

		// Create an RDS proxy using the credentials from Secrets Manager and the IAM role
		rproxy, err := rds.NewProxy(ctx, "db-proxy", &rds.ProxyArgs{
			Name:              pulumi.String("db-proxy"),
			DebugLogging:      pulumi.Bool(false),
			EngineFamily:      pulumi.String("MYSQL"),
			IdleClientTimeout: pulumi.Int(1800),
			RequireTls:        pulumi.Bool(true),
			VpcSecurityGroupIds: pulumi.StringArray{
				proxySecurityGroup.ID(),
			},
			VpcSubnetIds: subnetGroup.SubnetIds,
			RoleArn:      proxyRole.Arn,
			Auths: rds.ProxyAuthArray{
				&rds.ProxyAuthArgs{
					AuthScheme: pulumi.String("SECRETS"),
					SecretArn:  secret.ID(),
				},
			},
			// ... other configuration ...

		})
		if err != nil {
			return err
		}

		// Create a custom target group for the RDS proxy
		tg, err := rds.NewProxyDefaultTargetGroup(ctx, "my-db-proxy-default-target-group", &rds.ProxyDefaultTargetGroupArgs{
			DbProxyName: rproxy.Name,
			ConnectionPoolConfig: &rds.ProxyDefaultTargetGroupConnectionPoolConfigArgs{
				// Set the connection pool configuration settings
				ConnectionBorrowTimeout:   pulumi.Int(120),
				MaxConnectionsPercent:     pulumi.Int(90),
				MaxIdleConnectionsPercent: pulumi.Int(10),
				// Additional settings can be added
			},
		})
		if err != nil {
			return err
		}

		// // Attach the RDS Proxy to the Instance
		_, err = rds.NewProxyTarget(ctx, "proxy-target", &rds.ProxyTargetArgs{
			DbProxyName:          rproxy.Name,
			DbInstanceIdentifier: rdsI.Identifier,
			TargetGroupName:      tg.Name,
		})
		if err != nil {
			return err
		}

		// Output the IDs of the security groups
		ctx.Export("rdsSecurityGroupId", rdsSecurityGroup.ID())
		ctx.Export("proxySecurityGroupId", proxySecurityGroup.ID())

		return nil
	})
}
