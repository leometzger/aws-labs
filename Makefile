LAMBDA_GO_ARCH ?= amd64

build-sqs:
	cd aws/lambdas/sqs-partial-return && CGO_ENABLED=0 GOOS=linux GOARCH=$(LAMBDA_GO_ARCH) go build -tags lambda.norpc -o bin/bootstrap main.go;
	zip -j bin/sqs.zip aws/lambdas/sqs-partial-return/bin/bootstrap;


build-kinesis:
	cd aws/lambdas/kinesis-consumer && CGO_ENABLED=0 GOOS=linux GOARCH=$(LAMBDA_GO_ARCH) go build -tags lambda.norpc -o bin/bootstrap main.go;
	zip -j bin/kinesis.zip aws/lambdas/kinesis-consumer/bin/bootstrap;


build:
	mkdir -p bin;
	make build-sqs && make build-kinesis; 

deploy:
	cd pulumi; pulumi up;


destroy:
	cd pulumi && pulumi destroy --remove;
