package labs

import (
	"aws-labs-pulumi/labs/labsx"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/kinesis"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KinesisLambdaOptions struct {
	StreamName        string
	ShardCount        int
	LambdaName        string
	LambdaHandlerName string
	LambdaPath        string
}

type KinesisLambdaOutput struct {
	Stream   *kinesis.Stream
	Function *lambda.Function
}

func NewKinesisLambda(ctx *pulumi.Context, options *KinesisLambdaOptions) (*KinesisLambdaOutput, error) {
	role, err := labsx.NewLambdaRole(ctx, "kinesis-lambda-role", "kinesis-lambda-role")
	if err != nil {
		return nil, err
	}

	_, err = iam.NewPolicyAttachment(ctx, options.LambdaName+"-lambda-kinesis-execution", &iam.PolicyAttachmentArgs{
		Name:      pulumi.String("lambda-kinesis-policy"),
		Roles:     pulumi.All(role),
		PolicyArn: iam.ManagedPolicyAWSLambdaKinesisExecutionRole,
	})
	if err != nil {
		return nil, err
	}

	stream, err := kinesis.NewStream(ctx, "aws-labs-kinesis-stream", &kinesis.StreamArgs{
		Name:            pulumi.String(options.StreamName),
		ShardCount:      pulumi.Int(options.ShardCount),
		RetentionPeriod: pulumi.Int(24),
		StreamModeDetails: kinesis.StreamStreamModeDetailsArgs{
			StreamMode: pulumi.String("PROVISIONED"),
		},
	})
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

	_, err = lambda.NewEventSourceMapping(ctx, options.LambdaName+"-eventsource-kinesis", &lambda.EventSourceMappingArgs{
		EventSourceArn:   stream.Arn,
		FunctionName:     function.Name,
		StartingPosition: pulumi.String("LATEST"),
	})
	if err != nil {
		return nil, err
	}

	return &KinesisLambdaOutput{
		Stream:   stream,
		Function: function,
	}, nil
}
