package main

import (
	"aws-labs-pulumi/labs"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	kinesisLambda := labs.GetLambdaInfo(labs.KinesisConsumerLambda)
	sqsLambda := labs.GetLambdaInfo(labs.SqsPartialReturn)

	pulumi.Run(func(ctx *pulumi.Context) error {
		labs.NewKinesisLambda(ctx, &labs.KinesisLambdaOptions{
			StreamName:        "labs-kinesis-stream",
			ShardCount:        1,
			LambdaName:        kinesisLambda.Name,
			LambdaHandlerName: kinesisLambda.HandlerName,
			LambdaPath:        kinesisLambda.Path,
		})

		labs.NewSQSLambda(ctx, &labs.SQSLambdaOptions{
			QueueName:         "labs-sqs-queue",
			LambdaName:        sqsLambda.Name,
			LambdaHandlerName: sqsLambda.HandlerName,
			LambdaPath:        sqsLambda.Path,
		})

		return nil
	})
}
