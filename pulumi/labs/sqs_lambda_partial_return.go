package labs

import (
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
	role, err := NewLambdaRole(ctx, "sqs-lambda-role", "labs-sqs-lambda-role")
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

	function, err := NewGoLambda(ctx, &GoLambdaOptions{
		Name:        options.LambdaName,
		Role:        role,
		HandlerName: "",
		Archive:     pulumi.NewFileArchive(""),
	})
	if err != nil {
		return nil, err
	}

	_, err = lambda.NewEventSourceMapping(ctx, "sqs-lambda-event-source-mapping", &lambda.EventSourceMappingArgs{
		EventSourceArn: queue.Arn,
		FunctionName:   function.Name,
	})
	if err != nil {
		return nil, err
	}

	return &SQSLambdaOutput{Function: function, Queue: queue}, nil
}
