package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// kinesisLambda := labs.GetLambdaInfo(labs.KinesisConsumerLambda)
		// labs.NewKinesisLambda(ctx, &labs.KinesisLambdaOptions{
		// 	StreamName:        "labs-kinesis-stream",
		// 	ShardCount:        1,
		// 	LambdaName:        kinesisLambda.Name,
		// 	LambdaHandlerName: kinesisLambda.HandlerName,
		// 	LambdaPath:        kinesisLambda.Path,
		// })

		// sqsLambda := labs.GetLambdaInfo(labs.SqsPartialReturn)
		// labs.NewSQSLambda(ctx, &labs.SQSLambdaOptions{
		// 	QueueName:         "labs-sqs-queue",
		// 	LambdaName:        sqsLambda.Name,
		// 	LambdaHandlerName: sqsLambda.HandlerName,
		// 	LambdaPath:        sqsLambda.Path,
		// })

		return nil
	})
}
