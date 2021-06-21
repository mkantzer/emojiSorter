variable "region" {
  type    = string
  default = "us-east-1"
}

variable "environment" {
  type = string
}

variable "image_name" {
  type        = string
  description = "Image name, not including the ECR prefix. Format: mkantzer/emojiSorter/emojiSorter"
}

variable "image_tag" {
  type        = string
  description = "Tag on docker image to run"
  validation {
    condition     = var.image_tag != "latest"
    error_message = "Don't you dare use latest."
  }
}

