resource "aws_s3_bucket" "sample_bucket" {
  bucket = var.aws_s3_sample_bucket

  tags = {
    Name = "My bucket"
  }
}

resource "aws_s3_bucket_acl" "sample_bucket_acl" {
  bucket = aws_s3_bucket.sample_bucket.id
  acl    = "private"
}


resource "aws_cloudfront_distribution" "s3_distribution" {
  enabled             = true
  default_root_object = "index.html"

  origin {
    domain_name              = aws_s3_bucket.sample_bucket.bucket_regional_domain_name
    origin_access_control_id = aws_cloudfront_origin_access_control.default.id
    origin_id                = local.s3_origin_id
  }
}
