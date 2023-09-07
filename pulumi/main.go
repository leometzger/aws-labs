package main

import (
	"aws-labs-pulumi/labs"
	"aws-labs-pulumi/lambdas"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	kinesisLambda := lambdas.GetLambdaInfo(lambdas.KinesisConsumerLambda)
	sqsLambda := lambdas.GetLambdaInfo(lambdas.SqsPartialReturn)

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
