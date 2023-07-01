variable "aws_region" {
  type        = string
  default     = "us-west-1"
  description = "AWS Region to apply the infrastructure"
}


variable "benthos_image" {
  type        = string
  default     = "971064939130.dkr.ecr.us-west-1.amazonaws.com/benthos"
  description = "Benthos image"
}

variable "benthos_host_port" {
  type        = number
  default     = 4195
  description = "Benthos image"
}


variable "benthos_container_port" {
  type        = number
  default     = 4195
  description = "Benthos image"
}
