resource "aws_kinesis_stream" "source_stream" {
  name             = "benthos-kinesis-source"
  shard_count      = 1
  retention_period = 24

  stream_mode_details {
    stream_mode = "PROVISIONED"
  }
}

resource "aws_kinesis_stream" "input_stream" {
  name             = "benthos-kinesis-input"
  shard_count      = 1
  retention_period = 24

  stream_mode_details {
    stream_mode = "PROVISIONED"
  }
}

resource "aws_kinesis_stream" "destination_stream" {
  name             = "benthos-kinesis-output"
  shard_count      = 1
  retention_period = 24

  stream_mode_details {
    stream_mode = "PROVISIONED"
  }
}

resource "aws_ecs_cluster" "benthos_cluster" {
  name = "benthos-cluster"
}


resource "aws_ecs_task_definition" "benthos_processing_task" {
  family                   = "benthos-stream-processing"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = 256
  memory                   = 512
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  task_role_arn            = aws_iam_role.ecs_task_role.arn
  container_definitions = jsonencode([{
    name      = "benthos-container"
    image     = var.benthos_image
    essential = true
    portMappings = [{
      protocol      = "tcp"
      containerPort = var.benthos_container_port
      hostPort      = var.benthos_host_port
    }]
  }])
}

resource "aws_ecs_service" "benthos_service" {
  name                               = "benthos-service"
  cluster                            = aws_ecs_cluster.benthos_cluster.id
  task_definition                    = aws_ecs_task_definition.benthos_processing_task.arn
  desired_count                      = 2
  deployment_minimum_healthy_percent = 50
  deployment_maximum_percent         = 200
  launch_type                        = "FARGATE"
  scheduling_strategy                = "REPLICA"

  network_configuration {
    security_groups  = ["sg-e0755186"]
    subnets          = ["subnet-0c4cb057", "subnet-9fbd13f8"]
    assign_public_ip = true
  }


  lifecycle {
    ignore_changes = [task_definition, desired_count]
  }
}


