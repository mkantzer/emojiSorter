data "aws_caller_identity" "current" {}

locals {
  full_image_name = "${data.aws_caller_identity.current.account_id}.dkr.ecr.${var.region}.amazonaws.com/${var.image_name}"
  cleaner_name    = replace(trimprefix(var.image_name, "drizlyinc/"), "/", "-")
}

resource "aws_apprunner_service" "this" {
  service_name = "emojiSorter"

  source_configuration {
    authentication_configuration {
      access_role_arn = aws_iam_role.apprunner_access_role.arn
    }
    image_repository {
      image_configuration {
        port = 8080
        runtime_environment_variables = {
          ENV : var.environment
          ECHOSTRING : "BANANA"
        }
      }
      image_identifier      = "${local.full_image_name}:${var.image_tag}"
      image_repository_type = "ECR"
    }
  }


  # auto_scaling_configuration_arn = aws_apprunner_auto_scaling_configuration_version.this.arn
  health_check_configuration {
    healthy_threshold = 2
    interval          = 5
    protocol          = "HTTP"
    path              = "/healthz"
  }

  instance_configuration {
    instance_role_arn = aws_iam_role.apprunner_app_role.arn
  }
}


data "aws_iam_policy_document" "appruner_access_role" {
  statement {
    effect = "Allow"
    resources = [
      "arn:aws:ecr:${var.region}:${data.aws_caller_identity.current.account_id}:${var.image_name}",
      "*"
    ]
    actions = [
      "ecr:GetAuthorizationToken",
      "ecr:BatchCheckLayerAvailability",
      "ecr:GetDownloadUrlForLayer",
      "ecr:DescribeRepositories",
      "ecr:ListImages",
      "ecr:DescribeImages",
      "ecr:BatchGetImage",
    ]
  }
  statement {
    effect = "Allow"
    resources = [
      "*",
    ]
    actions = [
      "ecr:GetAuthorizationToken",
    ]
  }
}

data "aws_iam_policy_document" "apprunner_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["build.apprunner.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "apprunner_access_role" {
  name               = "${local.cleaner_name}-${var.region}-access-role"
  assume_role_policy = data.aws_iam_policy_document.apprunner_assume_role.json
  inline_policy {
    name   = "${local.cleaner_name}-${var.region}-access-role"
    policy = data.aws_iam_policy_document.appruner_access_role.json
  }
}

data "aws_iam_policy_document" "apprunner_app_role" {
  statement {
    effect = "Allow"
    resources = [
      "arn:aws:s3:::test"
    ]
    actions = [
      "s3:ListBucket",
    ]
  }
}

data "aws_iam_policy_document" "apprunner_app_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["tasks.apprunner.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "apprunner_app_role" {
  name               = "${local.cleaner_name}-${var.region}-app-role"
  assume_role_policy = data.aws_iam_policy_document.apprunner_app_assume_role.json
  inline_policy {
    name   = "${local.cleaner_name}-${var.region}-app-role"
    policy = data.aws_iam_policy_document.apprunner_app_role.json
  }
}
