resource "aws_kinesis_stream" "source_stream" {
  name             = "benthos-kinesis-source"
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

resource "aws_lambda_function" "kinesis_stream_processing" {
  function_name    = "kinesis-processing"
  runtime          = "go1.x"
  handler          = "kinesis_consumer"
  memory_size      = 512
  role             = aws_iam_role.iam_for_kinesis_consumer.arn
  filename         = data.archive_file.lambda.output_path
  source_code_hash = data.archive_file.lambda.output_base64sha256
}

