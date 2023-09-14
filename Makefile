LAMBDA_GO_ARCH ?= amd64

.PHONY: build-sqs-lambda
build-sqs:
	cd aws/lambdas/sqs-partial-return && CGO_ENABLED=0 GOOS=linux GOARCH=$(LAMBDA_GO_ARCH) go build -tags lambda.norpc -o bin/bootstrap main.go;
	zip -j pulumi/bin/sqs.zip aws/lambdas/sqs-partial-return/bin/bootstrap;


.PHONY: build-kinesis-consumer-lambda
build-kinesis:
	cd aws/lambdas/kinesis-consumer && CGO_ENABLED=0 GOOS=linux GOARCH=$(LAMBDA_GO_ARCH) go build -tags lambda.norpc -o bin/bootstrap main.go;
	zip -j pulumi/bin/kinesis.zip aws/lambdas/kinesis-consumer/bin/bootstrap;

.PHONY: build
build:
	mkdir -p pulumi/bin;
	make build-sqs && make build-kinesis; 

.PHONY: deploy
deploy:
	make build; cd pulumi; pulumi up;


.PHONY: destroy
destroy:
	pulumi destroy --remove;
