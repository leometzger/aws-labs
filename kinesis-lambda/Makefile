.PHONY: build 
build:
	cd lambda; GOOS=linux GOARCH=amd64 go build -o "bin/kinesis_consumer" main.go;


.PHONY: init
init:
	terraform init; terraform fmt; terraform validate


.PHONY: validate
validate:
	terraform validate


.PHONY: deploy
deploy:
	make build
	terraform apply


.PHONY: destroy
destroy:
	terraform destroy
