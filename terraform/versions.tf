terraform {
  # further configuration in envs
  backend "s3" {}

  required_version = ">= 0.15.1"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 3.44.0, < 4.0.0"
    }
  }
}

provider "aws" {
  region  = var.region
  profile = var.environment
  default_tags {
    tags = {
      Owner       = "mikekantzer"
      Usage       = "proof-of-concept"
      Environment = var.environment
      Repo        = "emojiSorter"
    }
  }
  # assume_role {
  #   role_arn     = "arn:aws:iam::ACCOUNT_ID:role/ROLE_NAME"
  #   session_name = "SESSION_NAME"
  #   external_id  = "EXTERNAL_ID"
  # }
}
