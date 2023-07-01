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

