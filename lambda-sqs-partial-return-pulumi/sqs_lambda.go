package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/lambda"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func NewSqsLambda(ctx *pulumi.Context) (*lambda.Function, *sqs.Queue, error) {
	queue, err := sqs.NewQueue(ctx, "lambda-sqs-partial-return", &sqs.QueueArgs{})
	if err != nil {
		return nil, nil, err
	}

	role, err := createBasicLambdaSqsRole(ctx)
	if err != nil {
		return nil, nil, err
	}

	function, err := createLambda(ctx, role)
	if err != nil {
		return nil, nil, err
	}

	_, err = lambda.NewEventSourceMapping(ctx, "sqs-lambda", &lambda.EventSourceMappingArgs{
		EventSourceArn: queue.Arn,
		FunctionName:   function.Name,
	})
	if err != nil {
		return nil, nil, err
	}

	return function, queue, nil
}

func createLambda(ctx *pulumi.Context, role *iam.Role) (*lambda.Function, error) {
	assetArchive := pulumi.NewAssetArchive(map[string]interface{}{
		".": pulumi.NewFileArchive("./lambda"),
	})
	function, err := lambda.NewFunction(ctx, "lambda-partial-return", &lambda.FunctionArgs{
		Name:    pulumi.String("partial-return-lambda"),
		Runtime: pulumi.String("go0.x"),
		Handler: pulumi.String(""),
		Code:    assetArchive,
		Role:    role.Arn,
	})

	if err != nil {
		return nil, err
	}

	return function, nil
}

func createBasicLambdaSqsRole(ctx *pulumi.Context) (*iam.Role, error) {
	assumeRolePolicy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
		Statements: []iam.GetPolicyDocumentStatement{
			{
				Actions: []string{"sts:AssumeRole"},
				Principals: []iam.GetPolicyDocumentStatementPrincipal{
					{
						Type:        "Service",
						Identifiers: []string{"lambda.amazonaws.com"},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return nil, err
	}

	funcRole, err := iam.NewRole(ctx, "role-kinesis-consumer", &iam.RoleArgs{
		Name:             pulumi.String("lambda-kinesis-processor-role"),
		AssumeRolePolicy: pulumi.String(assumeRolePolicy.Json),
	})
	if err != nil {
		return nil, err
	}

	_, err = iam.NewPolicyAttachment(ctx, "lambda-policy", &iam.PolicyAttachmentArgs{
		Name:      pulumi.String("rpa"),
		Roles:     pulumi.All(funcRole),
		PolicyArn: iam.ManagedPolicyAWSLambdaBasicExecutionRole,
	})
	if err != nil {
		return nil, err
	}

	_, err = iam.NewPolicyAttachment(ctx, "lambda-kinesis-attach", &iam.PolicyAttachmentArgs{
		Name:      pulumi.String("lambda-sqs-policy"),
		Roles:     pulumi.All(funcRole),
		PolicyArn: iam.ManagedPolicyAWSLambdaSQSQueueExecutionRole,
	})
	if err != nil {
		return nil, err
	}

	return funcRole, nil
}
