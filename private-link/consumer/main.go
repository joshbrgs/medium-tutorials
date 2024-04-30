package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create VPC
		vpc, err := ec2.NewVpc(ctx, "consumer-vpc", &ec2.VpcArgs{
			CidrBlock:          pulumi.String("10.0.0.0/16"),
			EnableDnsHostnames: pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("consumer-vpc"),
			},
		})
		if err != nil {
			return err
		}

		// Create private subnet
		subnet, err := ec2.NewSubnet(ctx, "consumer-private-subnet", &ec2.SubnetArgs{
			VpcId:               vpc.ID(),
			CidrBlock:           pulumi.String("10.0.64.0/24"),
			MapPublicIpOnLaunch: pulumi.Bool(false),
			AvailabilityZone:    pulumi.String("us-east-1b"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("consumer-private-subnet"),
			},
		})
		if err != nil {
			return err
		}

		// Create a NAT Gateway for the private subnets.
		eip, err := ec2.NewEip(ctx, "myEip", &ec2.EipArgs{
			Vpc: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		ig, err := ec2.NewInternetGateway(ctx, "myInternetGateway", &ec2.InternetGatewayArgs{
			VpcId: vpc.ID(),
		})
		if err != nil {
			return err
		}

		nat, err := ec2.NewNatGateway(ctx, "myNatGateway", &ec2.NatGatewayArgs{
			SubnetId:     subnet.ID(),
			AllocationId: eip.ID(),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("consumer-ngw"),
			},
		})
		if err != nil {
			return err
		}

		// Create a route table for the private subnets with a default route through the NAT Gateway.
		privateRouteTable, err := ec2.NewRouteTable(ctx, "myPrivateRouteTable", &ec2.RouteTableArgs{
			VpcId: vpc.ID(),
			Routes: ec2.RouteTableRouteArray{
				&ec2.RouteTableRouteArgs{
					CidrBlock:    pulumi.String("0.0.0.0/0"),
					NatGatewayId: nat.ID(),
				},
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("consumer-rt"),
			},
		})
		if err != nil {
			return err
		}

		// Associate the route table with the private subnets.
		_, err = ec2.NewRouteTableAssociation(ctx, "myPrivateRouteTableAssociation1", &ec2.RouteTableAssociationArgs{
			RouteTableId: privateRouteTable.ID(),
			SubnetId:     subnet.ID(),
		})
		if err != nil {
			return err
		}

		// Create a new security group that permits outbound HTTPS traffic on port 443 and no inbound traffic
		endpointsg, err := ec2.NewSecurityGroup(ctx, "secgrp", &ec2.SecurityGroupArgs{
			Description: pulumi.String("Allow traffic from the ec2 instance to the endpoint cidr"),
			VpcId:       vpc.ID(),
			Egress: ec2.SecurityGroupEgressArray{
				&ec2.SecurityGroupEgressArgs{
					ToPort:     pulumi.Int(3306),
					FromPort:   pulumi.Int(3306),
					Protocol:   pulumi.String("tcp"),
					CidrBlocks: pulumi.StringArray{pulumi.String("10.0.0.0/16")},
				},
			},
			Ingress: ec2.SecurityGroupIngressArray{}, // No inbound rules

			Tags: pulumi.StringMap{
				"Name": pulumi.String("consumer-sg"),
			},
		})
		if err != nil {
			return err
		}

		// Create a new security group that permits outbound HTTPS traffic on port 443 and no inbound traffic
		sg1, err := ec2.NewSecurityGroup(ctx, "endpoint-sg", &ec2.SecurityGroupArgs{
			Description: pulumi.String("Allow inbound traffic to endpoints from VPC"),
			VpcId:       vpc.ID(),
			Ingress: ec2.SecurityGroupIngressArray{
				&ec2.SecurityGroupIngressArgs{
					ToPort:     pulumi.Int(443),
					FromPort:   pulumi.Int(443),
					Protocol:   pulumi.String("tcp"),
					CidrBlocks: pulumi.StringArray{pulumi.String("10.0.0.0/16")},
				},
			},
			Egress: ec2.SecurityGroupEgressArray{}, // No outbound rules

			Tags: pulumi.StringMap{
				"Name": pulumi.String("ssm-endpoint-sg"),
			},
		})
		if err != nil {
			return err
		}

		// Security Group for the application server
		ec2sg, err := ec2.NewSecurityGroup(ctx, "ec2SecurityGroup", &ec2.SecurityGroupArgs{
			Description: pulumi.String("Allow MySQL traffic to endpoint"),
			VpcId:       vpc.ID(),
		})
		if err != nil {
			return err
		}

		// Ingress rule for the RDS security group allowing MySQL traffic from the application server security group
		_, err = ec2.NewSecurityGroupRule(ctx, "ec2SecurityGroupEgress", &ec2.SecurityGroupRuleArgs{
			Type:                  pulumi.String("egress"),
			FromPort:              pulumi.Int(3306),
			ToPort:                pulumi.Int(3306),
			Protocol:              pulumi.String("tcp"),
			SecurityGroupId:       ec2sg.ID(),
			SourceSecurityGroupId: endpointsg.ID(),
			Description:           pulumi.String("Allow ec2 instance traffic to endpoint sg"),
		})
		if err != nil {
			return err
		}

		_, err = ec2.NewSecurityGroupRule(ctx, "ec2SecurityGroupEgressSSM", &ec2.SecurityGroupRuleArgs{
			Type:            pulumi.String("egress"),
			FromPort:        pulumi.Int(443),
			ToPort:          pulumi.Int(443),
			Protocol:        pulumi.String("tcp"),
			SecurityGroupId: ec2sg.ID(),
			CidrBlocks: pulumi.StringArray{
				pulumi.String("0.0.0.0/0"),
			},
			Description: pulumi.String("Allow ec2 instance traffic to ssm"),
		})
		if err != nil {
			return err
		}

		// Create Endpoint for Session Manager
		_, err = ec2.NewVpcEndpoint(ctx, "com.amazonaws.us-east-1.ssm", &ec2.VpcEndpointArgs{
			VpcId:             vpc.ID(),
			ServiceName:       pulumi.String("com.amazonaws.us-east-1.ssm"),
			VpcEndpointType:   pulumi.String("Interface"),
			PrivateDnsEnabled: pulumi.Bool(true),
			SecurityGroupIds:  pulumi.StringArray{sg1.ID()},
			SubnetIds:         pulumi.StringArray{subnet.ID()},
			// SecurityGroupIds: insert your security group IDs here if needed.
			Tags: pulumi.StringMap{
				"Name": pulumi.String("com.amazonaws.us-east-1.ssm"),
			},
		})
		if err != nil {
			return err
		}

		// Create Endpoint for Session Manager
		_, err = ec2.NewVpcEndpoint(ctx, "com.amazonaws.us-east-1.ec2messages", &ec2.VpcEndpointArgs{
			VpcId:             vpc.ID(),
			ServiceName:       pulumi.String("com.amazonaws.us-east-1.ec2messages"),
			VpcEndpointType:   pulumi.String("Interface"),
			PrivateDnsEnabled: pulumi.Bool(true),
			SubnetIds:         pulumi.StringArray{subnet.ID()},
			// SecurityGroupIds: insert your security group IDs here if needed.
			SecurityGroupIds: pulumi.StringArray{sg1.ID()},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("com.amazonaws.us-east-1.ec2messages"),
			},
		})
		if err != nil {
			return err
		}

		// Create Endpoint for Session Manager
		_, err = ec2.NewVpcEndpoint(ctx, "com.amazonaws.us-east-1.ssmmessages", &ec2.VpcEndpointArgs{
			VpcId:             vpc.ID(),
			ServiceName:       pulumi.String("com.amazonaws.us-east-1.ssmmessages"),
			VpcEndpointType:   pulumi.String("Interface"),
			PrivateDnsEnabled: pulumi.Bool(true),
			SubnetIds:         pulumi.StringArray{subnet.ID()},
			// SecurityGroupIds: insert your security group IDs here if needed.
			SecurityGroupIds: pulumi.StringArray{sg1.ID()},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("com.amazonaws.us-east-1.ssmmessages"),
			},
		})
		if err != nil {
			return err
		}

		// Create Endpoint for Session Manager
		_, err = ec2.NewVpcEndpoint(ctx, "com.amazonaws.us-east-1.ec2", &ec2.VpcEndpointArgs{
			VpcId:             vpc.ID(),
			ServiceName:       pulumi.String("com.amazonaws.us-east-1.ec2"),
			VpcEndpointType:   pulumi.String("Interface"),
			PrivateDnsEnabled: pulumi.Bool(true),
			SubnetIds:         pulumi.StringArray{subnet.ID()},
			// SecurityGroupIds: insert your security group IDs here if needed.
			SecurityGroupIds: pulumi.StringArray{sg1.ID()},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("com.amazonaws.us-east-1.ec2"),
			},
		})
		if err != nil {
			return err
		}

		// Create an IAM role for the Session Manager
		role, err := iam.NewRole(ctx, "sessionManagerRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(`{
			"Version": "2012-10-17",
			"Statement": [
				{
					"Action": "sts:AssumeRole",
					"Principal": {
						"Service": "ec2.amazonaws.com"
					},
					"Effect": "Allow",
					"Sid": ""
				}
			]
			}`),
		})
		if err != nil {
			return err
		}

		// Attach the AWS managed policy for Session Manager to the IAM role
		_, err = iam.NewRolePolicyAttachment(ctx, "sessionManagerPolicy", &iam.RolePolicyAttachmentArgs{
			Role:      role.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"),
		})
		if err != nil {
			return err
		}

		// Create an Instance Profile to attach the role to the EC2 instance
		instanceProfile, err := iam.NewInstanceProfile(ctx, "instanceProfile", &iam.InstanceProfileArgs{
			Role: role.Name,
		})
		if err != nil {
			return err
		}

		ami, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
			Filters: []ec2.GetAmiFilter{
				{
					Name:   "name",
					Values: []string{"amzn2-ami-hvm-*-x86_64-gp2"},
				},
			},
			Owners:     []string{"amazon"},
			MostRecent: pulumi.BoolRef(true),
		})
		if err != nil {
			return err
		}

		// Create EC2 instance
		instance, err := ec2.NewInstance(ctx, "ghost", &ec2.InstanceArgs{
			InstanceType:       pulumi.String("t2.micro"),
			SecurityGroups:     pulumi.StringArray{ec2sg.ID()},
			SubnetId:           subnet.ID(),
			Ami:                pulumi.StringPtr(ami.Id),
			IamInstanceProfile: instanceProfile.Name,
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Consumer-instance"),
			},
		})
		if err != nil {
			return err
		}

		// Output the instance public IP
		ctx.Export("instancePublicIp", instance.PublicIp)
		ctx.Export("ig ID", ig.ID())
		// Output the ID of the security group and the AMI
		ctx.Export("securityGroupId", ec2sg.ID())
		ctx.Export("amiId", pulumi.String(ami.Id))
		return nil
	})
}
