variable "aws_aws_region" {
  type        = string
  default     = "us-east-1"
  description = "AWS Region to apply the infrastructure"
}

variable "aws_account_id" {
  type        = string
  description = "AWS AccountID to be apply changes"
}

variable "aws_s3_sample_bucket" {
  type        = string
  description = "Bucket name"
}


locals {
  s3_origin_id = "MySampleOrigin"
}
