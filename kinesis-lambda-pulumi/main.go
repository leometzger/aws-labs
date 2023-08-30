package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/kinesis"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		function, err := createLambda(ctx)
		if err != nil {
			return err
		}

		stream, err := kinesis.NewStream(ctx, "lambda-integration-kinesis-stream", &kinesis.StreamArgs{
			Name:            pulumi.String("labs-kinesis-lambda-stream"),
			ShardCount:      pulumi.Int(1),
			RetentionPeriod: pulumi.Int(24),
			StreamModeDetails: kinesis.StreamStreamModeDetailsArgs{
				StreamMode: pulumi.String("PROVISIONED"),
			},
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
