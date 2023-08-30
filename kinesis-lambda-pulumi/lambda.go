package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createLambda(ctx *pulumi.Context) (*lambda.Function, error) {
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
		Name:      pulumi.String("lambda-kinesis-policy"),
		Roles:     pulumi.All(funcRole),
		PolicyArn: iam.ManagedPolicyAWSLambdaKinesisExecutionRole,
	})
	if err != nil {
		return nil, err
	}

	assetArchive := pulumi.NewAssetArchive(map[string]interface{}{
		".": pulumi.NewFileArchive("./lambda"),
	})
	function, err := lambda.NewFunction(ctx, "lambda-kinesis-processor", &lambda.FunctionArgs{
		Name:    pulumi.String("lambda-kinesis-processor"),
		Runtime: pulumi.String("go1.x"),
		Handler: pulumi.String("kinesis_consumer"),
		Code:    assetArchive,
		Role:    funcRole.Arn,
	})
	if err != nil {
		return nil, err
	}

	return function, nil
}
