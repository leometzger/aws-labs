.PHONY: destroy<p align="center">
<img style="width: 360px" src="https://user-images.githubusercontent.com/15220162/224574915-6d0a36f6-debe-45a0-bb19-5baf60f1b97c.svg" alt="Labs illustration" />

</p>

# My personal AWS Labs

This repository is dedicated to learning and proof of concept purposes. You can run if you want it.

## Getting Started

To run these labs, follow the steps:

### Step 1: Install Pulumi

Pulumi is the primary tool used for developing and deploying these labs. If you haven't already installed Pulumi, you can do so by running the following command:

```sh
curl -sSL https://get.pulumi.com | sh
```

### Step 2: Install Go dependencies

```sh
go get
```

### Step 3: Configure the Lab

Within the main.go file and pulumi.yaml, make any necessary configurations for the lab you've chosen.
This may include specifying AWS regions, adjusting settings, or providing input parameters as required by the
lab instructions.

### Step 4: Deploy the stack on AWS

To deploy it, run the following command from your terminal:

```sh
make build && make deploy
```

This command will build the necessary resources and deploy the lab on AWS.
Make sure to follow any prompts or instructions that appear during the deployment process.

## Sugestions

You are welcome to provide feedback or suggestions. I will be happy to try to implement.
