LAMBDA_GO_ARCH ?= amd64

.PHONY: build-sqs-lambda
build-sqs-lambda:
	cd lambdas/sqs-partial-return; 
	CGO_ENABLED=0 GOOS=linux GOARCH=$(LAMBDA_GO_ARCH) go build -tags lambda.norpc -o bin/sqs/bootstrap main.go;
	zip -j bin/sqs.zip bin/sqs/bootstrap;


.PHONY: build-kinesis-consumer-lambda
build-kinesis-consumer-lambda:
	cd lambdas/kinesis-consumer; 
	CGO_ENABLED=0 GOOS=linux GOARCH=$(LAMBDA_GO_ARCH) go build -tags lambda.norpc -o bin/kinesis/bootstrap main.go; 
	zip -j bin/kinesis.zip ./bin/kinesis/bootstrap;


.PHONY: build
build:
	make build-sqs-lambda && make build-kinesis-consumer-lambda;


.PHONY: deploy
deploy:
	make build; pulumi up;


.PHONY: destroy
destroy:
	pulumi destroy --remove;
