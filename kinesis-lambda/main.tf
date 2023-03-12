resource "aws_kinesis_stream" "kinesis_stream" {
  name             = "lambda-integration-kinesis-stream"
  shard_count      = 1
  retention_period = 24

  stream_mode_details {
    stream_mode = "PROVISIONED"
  }
}

resource "aws_lambda_function" "kinesis_stream_consumer" {
  function_name    = "kinesis-consumer"
  runtime          = "go1.x"
  handler          = "kinesis_consumer"
  memory_size      = 128
  role             = aws_iam_role.iam_for_kinesis_consumer.arn
  filename         = data.archive_file.lambda.output_path
  source_code_hash = data.archive_file.lambda.output_base64sha256
}


resource "aws_lambda_event_source_mapping" "incoming_events_to_ingest" {
  event_source_arn  = aws_kinesis_stream.kinesis_stream.arn
  function_name     = aws_lambda_function.kinesis_stream_consumer.arn
  starting_position = "LATEST"
}
