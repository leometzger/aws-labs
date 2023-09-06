package pulumi

import (
	"aws-labs-pulumi/labs"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		labs.NewKinesisLambda(ctx, &labs.KinesisLambdaOptions{
			StreamName:        "labs-kinesis-stream",
			ShardCount:        1,
			LambdaName:        "labs-kinesis-consumer",
			LambdaHandlerName: "kinesis_consumer",
			LambdaPath:        "./lambdas/kinesis-consumer",
		})

		labs.NewSQSLambda(ctx, &labs.SQSLambdaOptions{
			QueueName:         "labs-sqs-queue",
			LambdaName:        "labs-sqs-partial-return",
			LambdaHandlerName: "",
			LambdaPath:        "./lambdas/sqs-partial-return",
		})

		return nil
	})
}
