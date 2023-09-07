package labs

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GoLambdaOptions struct {
	Name        string
	HandlerName string
	Archive     pulumi.Archive
	Role        *iam.Role
}

func NewGoLambda(ctx *pulumi.Context, options *GoLambdaOptions) (*lambda.Function, error) {
	function, err := lambda.NewFunction(ctx, options.Name, &lambda.FunctionArgs{
		Name:          pulumi.String(options.Name),
		Handler:       pulumi.String(options.HandlerName),
		Runtime:       pulumi.String("provided.al2"),
		Code:          options.Archive,
		Role:          options.Role.Arn,
		Architectures: pulumi.StringArray{pulumi.String("x86_64")},
	})
	if err != nil {
		return nil, err
	}

	return function, nil
}

// Creates a New role with basic permissions to execute a lambda
// function
func NewLambdaRole(ctx *pulumi.Context, resourceName string, roleName string) (*iam.Role, error) {
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

	role, err := iam.NewRole(ctx, roleName, &iam.RoleArgs{
		Name:             pulumi.String(roleName),
		AssumeRolePolicy: pulumi.String(assumeRolePolicy.Json),
	})
	if err != nil {
		return nil, err
	}

	_, err = iam.NewPolicyAttachment(ctx, roleName+"lambda-basic-execution-policy", &iam.PolicyAttachmentArgs{
		Name:      pulumi.String("lambda-basic-execution-policy"),
		Roles:     pulumi.All(role),
		PolicyArn: iam.ManagedPolicyAWSLambdaBasicExecutionRole,
	})
	if err != nil {
		return nil, err
	}

	return role, nil
}
