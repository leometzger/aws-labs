data "aws_iam_policy_document" "assume_role" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}


data "archive_file" "lambda" {
  type        = "zip"
  source_file = "lambda/bin/kinesis_consumer"
  output_path = "kinesis-consumer.zip"
}

