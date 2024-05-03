package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
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
			CidrBlock:           pulumi.String("10.0.0.0/24"),
			MapPublicIpOnLaunch: pulumi.Bool(false),
			AvailabilityZoneId:  pulumi.String("use1-az2"),
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

		_, err = ec2.NewInternetGateway(ctx, "myInternetGateway", &ec2.InternetGatewayArgs{
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
		endpointsg, err := ec2.NewSecurityGroup(ctx, "endpoint-security-group", &ec2.SecurityGroupArgs{
			Description: pulumi.String("Allow traffic from Glue to the endpoint service"),
			VpcId:       vpc.ID(),
			Egress: ec2.SecurityGroupEgressArray{
				&ec2.SecurityGroupEgressArgs{
					ToPort:     pulumi.Int(5432),
					FromPort:   pulumi.Int(5432),
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

		// Security Group for the application server
		gluesg, err := ec2.NewSecurityGroup(ctx, "glue-security-group", &ec2.SecurityGroupArgs{
			Description: pulumi.String("Allow Glue traffic to endpoint"),
			VpcId:       vpc.ID(),
		})
		if err != nil {
			return err
		}

		_, err = ec2.NewSecurityGroupRule(ctx, "EndpointSecurityGroupIngress", &ec2.SecurityGroupRuleArgs{
			Type:                  pulumi.String("ingress"),
			FromPort:              pulumi.Int(5432),
			ToPort:                pulumi.Int(5432),
			Protocol:              pulumi.String("tcp"),
			SecurityGroupId:       endpointsg.ID(),
			SourceSecurityGroupId: gluesg.ID(),
			Description:           pulumi.String("Allow ec2 instance traffic to ssm"),
		})
		if err != nil {
			return err
		}
		// Ingress rule for the RDS security group allowing postgresql traffic from the application server security group
		_, err = ec2.NewSecurityGroupRule(ctx, "GlueSecurityGroupEgress", &ec2.SecurityGroupRuleArgs{
			Type:                  pulumi.String("egress"),
			FromPort:              pulumi.Int(5432),
			ToPort:                pulumi.Int(5432),
			Protocol:              pulumi.String("tcp"),
			SecurityGroupId:       gluesg.ID(),
			SourceSecurityGroupId: endpointsg.ID(),
			Description:           pulumi.String("Allow ec2 instance traffic to endpoint sg"),
		})
		if err != nil {
			return err
		}

		_, err = ec2.NewSecurityGroupRule(ctx, "GlueSecurityGroupEgress2", &ec2.SecurityGroupRuleArgs{
			Type:                  pulumi.String("egress"),
			FromPort:              pulumi.Int(0),
			ToPort:                pulumi.Int(65535),
			Protocol:              pulumi.String("tcp"),
			SecurityGroupId:       gluesg.ID(),
			SourceSecurityGroupId: gluesg.ID(),
			Description:           pulumi.String("Allow ec2 instance traffic to ssm"),
		})
		if err != nil {
			return err
		}

		_, err = ec2.NewSecurityGroupRule(ctx, "GlueSecurityGroupIngress", &ec2.SecurityGroupRuleArgs{
			Type:                  pulumi.String("ingress"),
			FromPort:              pulumi.Int(0),
			ToPort:                pulumi.Int(65535),
			Protocol:              pulumi.String("tcp"),
			SecurityGroupId:       gluesg.ID(),
			SourceSecurityGroupId: gluesg.ID(),
			Description:           pulumi.String("Allow ec2 instance traffic to ssm"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewSecurityGroupRule(ctx, "GlueSecurityGroupNatEgress", &ec2.SecurityGroupRuleArgs{
			Type:            pulumi.String("egress"),
			FromPort:        pulumi.Int(443),
			ToPort:          pulumi.Int(443),
			Protocol:        pulumi.String("tcp"),
			SecurityGroupId: gluesg.ID(),
			CidrBlocks:      pulumi.StringArray{pulumi.String("0.0.0.0/0")},
			Description:     pulumi.String("Allow ec2 instance traffic to ssm"),
		})
		if err != nil {
			return err
		}

		return nil
	})
}
