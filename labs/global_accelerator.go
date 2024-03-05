package labs

import (
	"log/slog"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/elb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GlobalAcceleratorLabOptions struct{}

type GlobalAcceleratorLabOutput struct{}

func NewGlobalAcceleratorLab(ctx *pulumi.Context, options *GlobalAcceleratorLabOptions) (*GlobalAcceleratorLabOutput, error) {
	slog.Info("Creating Global Accelerator Lab")

	vpc, err := ec2.LookupVpc(ctx, &ec2.LookupVpcArgs{
		Default: pulumi.BoolRef(true),
	})
	if err != nil {
		return nil, err
	}

	loadBalancer, err := elb.NewLoadBalancer(ctx, "ga-alb", &elb.LoadBalancerArgs{
		Name:     pulumi.String("GlobalAcceleratorLB"),
		Internal: pulumi.BoolPtr(false),
		Listeners: elb.LoadBalancerListenerArray{
			elb.LoadBalancerListenerArgs{},
		},
		HealthCheck: elb.LoadBalancerHealthCheckArgs{
			Target:   pulumi.String("/health"),
			Interval: pulumi.Int(5),
		},
	})
	if err != nil {
		return nil, err
	}

	return &GlobalAcceleratorLabOutput{}, err
}
