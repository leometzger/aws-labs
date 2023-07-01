resource "aws_iam_role" "iam_for_kinesis_consumer" {
  name               = "iam_for_kinesis_consumer"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

# Attach Basic Lambda
resource "aws_iam_role_policy_attachment" "attach_logs_policy" {
  role       = aws_iam_role.iam_for_kinesis_consumer.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

# Attach Use Kinesis
resource "aws_iam_role_policy_attachment" "attach_kinesis_consumer_policy" {
  role       = aws_iam_role.iam_for_kinesis_consumer.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaKinesisExecutionRole"
}

resource "aws_iam_role" "ecs_task_role" {
  name               = "benthos-container-role"
  assume_role_policy = data.aws_iam_policy_document.ecs_task_role.json
}

resource "aws_iam_policy" "kinesis_consumer" {
  name   = "kinesis_consumer"
  policy = data.aws_iam_policy_document.kinesis_consumer.json
}

resource "aws_iam_policy" "dynamodb_access" {
  name        = "benthos-task-policy-dynamodb"
  description = "Policy that allows access to DynamoDB"
  policy      = data.aws_iam_policy_document.dynamodb.json
}

resource "aws_iam_role" "ecs_task_execution_role" {
  name               = "benthos-ecs-task-execution-role"
  assume_role_policy = data.aws_iam_policy_document.ecs_task_role.json
}

# ECS permissions to the execution task
resource "aws_iam_role_policy_attachment" "ecs_task_execution_role_policy_attachment" {
  role       = aws_iam_role.ecs_task_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

# DynamoDB access to benthos
resource "aws_iam_role_policy_attachment" "ecs_task_role_policy_attachment" {
  role       = aws_iam_role.ecs_task_role.name
  policy_arn = aws_iam_policy.dynamodb_access.arn
}

# Kinesis stream access to benthos
resource "aws_iam_role_policy_attachment" "attach_consume_kinesis" {
  role       = aws_iam_role.ecs_task_role.name
  policy_arn = aws_iam_policy.kinesis_consumer.arn
}
