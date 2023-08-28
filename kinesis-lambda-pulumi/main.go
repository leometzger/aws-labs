package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/kinesis"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stream, err := kinesis.NewStream(ctx, "lambda-integration-kinesis-stream", &kinesis.StreamArgs{
			ShardCount:      pulumi.Int(1),
			RetentionPeriod: pulumi.Int(24),
			StreamModeDetails: kinesis.StreamStreamModeDetailsArgs{
				StreamMode: pulumi.String("PROVISIONED"),
			},
		})
		if err != nil {
			return err
		}

		funcRole, err := iam.NewRole(ctx, "role-kinesis-consumer", &iam.RoleArgs{})
		if err != nil {
			return err
		}
		iam.NewPolicyAttachment(ctx, "lambda-policy", &iam.PolicyAttachmentArgs{
			Name:      pulumi.String("rpa"),
			Roles:     pulumi.All(funcRole),
			PolicyArn: iam.ManagedPolicyAWSLambdaBasicExecutionRole,
		})
		iam.NewPolicyAttachment(ctx, "lambda-kinesis-attach", &iam.PolicyAttachmentArgs{
			Name:      pulumi.String("lambda-kinesis-policy"),
			Roles:     pulumi.All(funcRole),
			PolicyArn: iam.ManagedPolicyAWSLambdaKinesisExecutionRole,
		})

		assets := make(map[string]interface{})
		assets["."] = pulumi.NewFileArchive("./lambda")
		function, err := lambda.NewFunction(ctx, "lambda-bla", &lambda.FunctionArgs{
			Code:    pulumi.NewAssetArchive(assets),
			Runtime: pulumi.String("go1.x"),
			Role:    funcRole.Arn,
		})
		if err != nil {
			return err
		}

		_, err = lambda.NewEventSourceMapping(ctx, "kinesis-lambda", &lambda.EventSourceMappingArgs{
			EventSourceArn:   stream.Arn,
			FunctionName:     function.Name,
			StartingPosition: pulumi.String("LATEST"),
		})
		if err != nil {
			return err
		}

		return nil
	})
}
