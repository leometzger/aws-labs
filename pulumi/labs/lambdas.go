package labs

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
		kinesisLambdaPath, err := filepath.Abs("bin/kinesis.zip")
		checkError(err)

		return &LambdaInfo{
			Name:        "aws-labs-kinesis-consumer",
			Path:        kinesisLambdaPath,
			HandlerName: "bootstrap",
		}
	}

	if implementation == SqsPartialReturn {
		sqsPartialReturnPath, err := filepath.Abs("bin/sqs.zip")
		checkError(err)

		return &LambdaInfo{
			Name:        "aws-labs-sqs-partial-return",
			Path:        sqsPartialReturnPath,
			HandlerName: "bootstrap",
		}
	}

	return nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
