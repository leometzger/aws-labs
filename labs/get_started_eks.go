package labs

import (
	"github.com/pulumi/pulumi-eks/sdk/go/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetStartedEKSOptions struct {
	ClusterName string
}

type GetStartedEKSOutput struct{}

func NewGetStartedEks(ctx *pulumi.Context, options *GetStartedEKSOptions) (*GetStartedEKSOutput, error) {
	cluster, err := eks.NewCluster(ctx, "eks-cluster", &eks.ClusterArgs{
		Tags: pulumi.StringMap{
			"Name": pulumi.String(options.ClusterName),
			"Type": pulumi.String("labs"),
		},
	})
	if err != nil {
		return nil, err
	}

	ctx.Export("kubeconfig", cluster.Kubeconfig)
	return nil, nil
}
