LAMBDA_GO_ARCH ?= amd64

.PHONY: build-lambda-sqs
build-sqs:
	cd aws/lambdas/sqs-partial-return && CGO_ENABLED=0 GOOS=linux GOARCH=$(LAMBDA_GO_ARCH) go build -tags lambda.norpc -o bin/bootstrap main.go;
	zip -j bin/sqs.zip aws/lambdas/sqs-partial-return/bin/bootstrap;


.PHONY: build-lambda-kinesis
build-kinesis:
	cd aws/lambdas/kinesis-consumer && CGO_ENABLED=0 GOOS=linux GOARCH=$(LAMBDA_GO_ARCH) go build -tags lambda.norpc -o bin/bootstrap main.go;
	zip -j bin/kinesis.zip aws/lambdas/kinesis-consumer/bin/bootstrap;


.PHONY: build-lambdas
build:
	mkdir -p pulumi/bin;
	make build-sqs && make build-kinesis; 

.PHONY: deploy
deploy:
	cd pulumi; pulumi up;


.PHONY: destroy
destroy:
	cd pulumi && pulumi destroy --remove;
