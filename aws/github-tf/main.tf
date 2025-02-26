terraform {
  required_providers {
    github = {
      source  = "integrations/github"
      version = "5.42.0"
    }
  }
}

resource "aws_ssm_parameter" "github_token" {
  name  = "/${var.prefix}/github-repo-token"
  type  = "SecureString"
  value = " " # To be set manually during setup

  lifecycle {
    ignore_changes = [value]
  }
}

data "aws_ssm_parameter" "github_token" {
  name = aws_ssm_parameter.github_token.name
}

provider "github" {
  owner = var.github_organization
  token = data.aws_ssm_parameter.github_token.value
}

data "aws_iam_policy_document" "github_actions_assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Federated"
      identifiers = [aws_iam_openid_connect_provider.github_actions.arn]
    }

    actions = ["sts:AssumeRoleWithWebIdentity"]

    condition {
      test     = "StringEquals"
      variable = "token.actions.githubusercontent.com:aud"
      values   = ["sts.amazonaws.com"]
    }

    condition {
      test     = "StringLike"
      variable = "token.actions.githubusercontent.com:sub"
      values   = ["repo:${var.github_organization}/${var.github_repository}:ref:refs/heads/${var.github_branch}"]
    }
  }
}

resource "aws_iam_role" "github_actions_service_role" {
  name               = "${var.prefix}-github-actions-role"
  assume_role_policy = data.aws_iam_policy_document.github_actions_assume_role.json
}

# Create OIDC provider for GitHub Actions
resource "aws_iam_openid_connect_provider" "github_actions" {
  url             = "https://token.actions.githubusercontent.com"
  client_id_list  = ["sts.amazonaws.com"]
  thumbprint_list = ["6938fd4d98bab03faadb97b34396831e3780aea1"]
}

# Attach identity-based policies to the GitHub Actions role
resource "aws_iam_role_policy_attachment" "github_actions_ecr_admin" {
  role       = aws_iam_role.github_actions_service_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryFullAccess"
}

resource "aws_iam_role_policy_attachment" "github_actions_ec2_admin" {
  role       = aws_iam_role.github_actions_service_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2FullAccess"
}

resource "aws_iam_role_policy_attachment" "github_actions_iam_admin" {
  role       = aws_iam_role.github_actions_service_role.name
  policy_arn = "arn:aws:iam::aws:policy/IAMFullAccess"
}

resource "aws_iam_role_policy_attachment" "github_actions_ssm" {
  role       = aws_iam_role.github_actions_service_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonSSMReadOnlyAccess"
}

resource "aws_iam_role_policy_attachment" "github_actions_s3_admin" {
  role       = aws_iam_role.github_actions_service_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonS3FullAccess"
}

data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create GitHub Actions secrets
resource "github_actions_secret" "service_role_arn" {
  repository      = var.github_repository
  secret_name     = "E2B_SERVICE_ROLE_ARN"
  plaintext_value = aws_iam_role.github_actions_service_role.arn
}

resource "github_actions_secret" "project_name" {
  repository      = var.github_repository
  secret_name     = "E2B_PROJECT_NAME"
  plaintext_value = var.project_name
}

resource "github_actions_secret" "aws_region" {
  repository      = var.github_repository
  secret_name     = "E2B_AWS_REGION"
  plaintext_value = data.aws_region.current.name
}

resource "github_actions_secret" "terraform_prefix" {
  repository      = var.github_repository
  secret_name     = "E2B_TERRAFORM_PREFIX"
  plaintext_value = "${var.prefix}-"
}

resource "github_actions_secret" "terraform_state_bucket" {
  repository      = var.github_repository
  secret_name     = "E2B_TERRAFORM_STATE_BUCKET"
  plaintext_value = var.terraform_state_bucket
}

resource "github_actions_secret" "terraform_state_dynamodb_table" {
  repository      = var.github_repository
  secret_name     = "E2B_TERRAFORM_STATE_DYNAMODB_TABLE"
  plaintext_value = var.terraform_state_dynamodb_table
}

resource "github_actions_secret" "domain_name" {
  repository      = var.github_repository
  secret_name     = "E2B_DOMAIN_NAME"
  plaintext_value = var.domain_name
}

# Grant GitHub Actions role access to S3 buckets via resource-based policies
data "aws_iam_policy_document" "public_builds_bucket_policy" {
  count = var.public_builds_bucket != null ? 1 : 0

  statement {
    effect = "Allow"

    principals {
      type        = "AWS"
      identifiers = [aws_iam_role.github_actions_service_role.arn]
    }

    actions = ["s3:*"]

    resources = [
      "arn:aws:s3:::${var.public_builds_bucket}",
      "arn:aws:s3:::${var.public_builds_bucket}/*"
    ]
  }
}

resource "aws_s3_bucket_policy" "public_builds_bucket_policy" {
  count = var.public_builds_bucket != null ? 1 : 0

  bucket = var.public_builds_bucket
  policy = data.aws_iam_policy_document.public_builds_bucket_policy[0].json
}

data "aws_iam_policy_document" "fc_kernels_bucket_policy" {
  statement {
    effect = "Allow"

    principals {
      type        = "AWS"
      identifiers = [aws_iam_role.github_actions_service_role.arn]
    }

    actions = ["s3:*"]

    resources = [
      "arn:aws:s3:::${var.kernel_bucket}",
      "arn:aws:s3:::${var.kernel_bucket}/*"
    ]
  }
}

resource "aws_s3_bucket_policy" "fc_kernels_bucket_policy" {
  bucket = var.kernel_bucket
  policy = data.aws_iam_policy_document.fc_kernels_bucket_policy.json
}

data "aws_iam_policy_document" "fc_versions_bucket_policy" {
  statement {
    effect = "Allow"

    principals {
      type        = "AWS"
      identifiers = [aws_iam_role.github_actions_service_role.arn]
    }

    actions = ["s3:*"]

    resources = [
      "arn:aws:s3:::${var.fc_versions_bucket}",
      "arn:aws:s3:::${var.fc_versions_bucket}/*"
    ]
  }
}

resource "aws_s3_bucket_policy" "fc_versions_bucket_policy" {
  bucket = var.fc_versions_bucket
  policy = data.aws_iam_policy_document.fc_versions_bucket_policy.json
}
