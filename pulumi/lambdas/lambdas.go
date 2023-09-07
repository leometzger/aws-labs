package lambdas

import (
	"path/filepath"
)

type LambdaImplementation int

const (
	KinesisConsumerLambda LambdaImplementation = iota
	SqsPartialReturn
)

type LambdaInfo struct {
	Name        string
	HandlerName string
	Path        string
}

func GetLambdaInfo(implementation LambdaImplementation) *LambdaInfo {
	if implementation == KinesisConsumerLambda {
		kinesisLambdaPath, err := filepath.Abs("./lambdas/kinesis-consumer")
		checkError(err)

		return &LambdaInfo{
			Name:        "aws-labs-kinesis-consumer",
			Path:        kinesisLambdaPath,
			HandlerName: "kinesis_consumer",
		}
	}

	if implementation == SqsPartialReturn {
		sqsPartialReturnPath, err := filepath.Abs("./lambdas/sqs-partial-return")
		checkError(err)

		return &LambdaInfo{
			Name:        "aws-labs-sqs-partial-return",
			Path:        sqsPartialReturnPath,
			HandlerName: "lambda_sqs_partial_return",
		}
	}

	return nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
