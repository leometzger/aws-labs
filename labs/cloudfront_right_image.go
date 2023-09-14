package labs

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudfront"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudfrontRightImageOptions struct{}

type CloudfrontRightImageOutput struct{}

func NewCloudfrontRightImage(ctx *pulumi.Context, options *CloudfrontRightImageOptions) (*CloudfrontRightImageOutput, error) {
	bucketV2, err := s3.NewBucketV2(ctx, "bucketV2", &s3.BucketV2Args{
		Tags: pulumi.StringMap{
			"Name": pulumi.String("My bucket"),
		},
	})
	if err != nil {
		return nil, err
	}
	_, err = s3.NewBucketAclV2(ctx, "bAcl", &s3.BucketAclV2Args{
		Bucket: bucketV2.ID(),
		Acl:    pulumi.String("private"),
	})
	if err != nil {
		return nil, err
	}
	s3OriginId := "dsad"

	_, err = cloudfront.NewDistribution(ctx, "s3Distribution", &cloudfront.DistributionArgs{
		Origins: cloudfront.DistributionOriginArray{
			&cloudfront.DistributionOriginArgs{
				DomainName: bucketV2.BucketRegionalDomainName,
				OriginId:   pulumi.String(s3OriginId),
			},
		},
		Enabled:           pulumi.Bool(true),
		IsIpv6Enabled:     pulumi.Bool(true),
		Comment:           pulumi.String("Some comment"),
		DefaultRootObject: pulumi.String("index.html"),
		Tags: pulumi.StringMap{
			"Environment": pulumi.String("production"),
		},
		ViewerCertificate: &cloudfront.DistributionViewerCertificateArgs{
			CloudfrontDefaultCertificate: pulumi.Bool(true),
		},
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
