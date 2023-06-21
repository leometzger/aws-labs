resource "aws_iam_role" "iam_for_choose_image_lambda" {
  name               = "iam_for_choose_image_lambda"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

# Attach Basic Lambda
resource "aws_iam_role_policy_attachment" "attach_logs_policy" {
  role       = aws_iam_role.iam_for_choose_image_lambda.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
