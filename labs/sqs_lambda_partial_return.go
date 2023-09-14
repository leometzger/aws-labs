package labs

import (
	"aws-labs-pulumi/labs/labsx"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/lambda"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SQSLambdaOptions struct {
	QueueName         string
	LambdaName        string
	LambdaHandlerName string
	LambdaPath        string
}

type SQSLambdaOutput struct {
	Function *lambda.Function
	Queue    *sqs.Queue
}

func NewSQSLambda(ctx *pulumi.Context, options *SQSLambdaOptions) (*SQSLambdaOutput, error) {
	role, err := labsx.NewLambdaRole(ctx, "sqs-lambda-role", "labs-sqs-lambda-role")
	if err != nil {
		return nil, err
	}

	_, err = iam.NewPolicyAttachment(ctx, "lambda-kinesis-attach", &iam.PolicyAttachmentArgs{
		Name:      pulumi.String("lambda-sqs-policy"),
		Roles:     pulumi.All(role),
		PolicyArn: iam.ManagedPolicyAWSLambdaSQSQueueExecutionRole,
	})
	if err != nil {
		return nil, err
	}

	queue, err := sqs.NewQueue(ctx, options.QueueName, &sqs.QueueArgs{})
	if err != nil {
		return nil, err
	}

	function, err := labsx.NewGoLambda(ctx, &labsx.GoLambdaOptions{
		Role:        role,
		Name:        options.LambdaName,
		HandlerName: options.LambdaHandlerName,
		Archive:     pulumi.NewFileArchive(options.LambdaPath),
	})
	if err != nil {
		return nil, err
	}

	_, err = lambda.NewEventSourceMapping(ctx, "sqs-lambda-event-source-mapping", &lambda.EventSourceMappingArgs{
		EventSourceArn: queue.Arn,
		FunctionName:   function.Name,
		FunctionResponseTypes: pulumi.StringArray{
			pulumi.String("ReportBatchItemFailures"),
		},
	})
	if err != nil {
		return nil, err
	}

	return &SQSLambdaOutput{Function: function, Queue: queue}, nil
}
