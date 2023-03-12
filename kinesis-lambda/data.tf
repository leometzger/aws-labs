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


data "aws_iam_policy_document" "allow_save_logs" {
  statement {
    effect    = "Allow"
    actions   = ["logs:CreateLogGroup"]
    resources = ["arn:aws:logs:${var.region}:${var.account_id}:*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = [
      "arn:aws:logs:${var.region}:${var.account_id}:log-group:/aws/lambda/${aws_lambda_function.kinesis_stream_consumer.function_name}:*"
    ]
  }
}


## Permissions to lambda retrieve the kinesis image
data "aws_iam_policy_document" "allow_retrieve_kinesis_images" {
  statement {
    effect    = "Allow"
    resources = ["arn:aws:kinesis:*:${var.account_id}:${aws_kinesis_stream.kinesis_stream.name}"]
    actions = [
      "ecr:BatchGetImage",
      "ecr:GetDownloadUrlForLayer"
    ]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}
