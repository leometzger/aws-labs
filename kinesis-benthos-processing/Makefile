.PHONY: build 
build:
	echo "build";


.PHONY: run-local 
run-local:
	cd lambda; docker run --network host --rm -v $(pwd)/config.yaml:/benthos.yaml jeffail/benthos;


.PHONY: init
init:
	terraform init; terraform fmt; terraform validate


.PHONY: validate
validate:
	terraform validate


.PHONY: deploy
deploy:
	make build;
	terraform apply


.PHONY: destroy
destroy:
	terraform destroy
