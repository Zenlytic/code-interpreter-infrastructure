data "aws_iam_policy_document" "ec2_assume_role_policy_doc" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "infra_instances_service_role" {
  name               = "${var.prefix}-infra-instances-service-role"
  assume_role_policy = data.aws_iam_policy_document.ec2_assume_role_policy_doc.json
}

resource "aws_iam_instance_profile" "infra_instances_service_profile" {
  name = "${var.prefix}-infra-instances-service-profile"
  role = aws_iam_role.infra_instances_service_role.name
}

resource "aws_ssm_parameter" "cloudflare_api_token" {
  name  = "/${var.prefix}/cloudflare-api-token"
  type  = "SecureString"
  value = " " # To be set manually during setup

  lifecycle {
    ignore_changes = [value]
  }
}

resource "random_uuid" "consul_acl_token" {}

resource "aws_ssm_parameter" "consul_acl_token" {
  name  = "/${var.prefix}/consul-secret-id"
  type  = "SecureString"
  value = random_uuid.consul_acl_token.result
}

resource "random_uuid" "nomad_acl_token" {}

resource "aws_ssm_parameter" "nomad_acl_token" {
  name  = "/${var.prefix}/nomad-secret-id"
  type  = "SecureString"
  value = random_uuid.nomad_acl_token.result
}

resource "aws_ssm_parameter" "grafana_api_key" {
  name  = "/${var.prefix}/grafana-api-key"
  type  = "SecureString"
  value = " " # To be set manually during setup

  lifecycle {
    ignore_changes = [value]
  }
}

resource "aws_ssm_parameter" "analytics_collector_host" {
  name  = "/${var.prefix}/analytics-collector-host"
  type  = "SecureString"
  value = " " # To be set manually during setup

  lifecycle {
    ignore_changes = [value]
  }
}

resource "aws_ssm_parameter" "analytics_collector_api_token" {
  name  = "/${var.prefix}/analytics-collector-api-token"
  type  = "SecureString"
  value = " " # To be set manually during setup

  lifecycle {
    ignore_changes = [value]
  }
}

data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "ssm_read_policy_doc" {
  statement {
    effect = "Allow"
    actions = [
      "ssm:GetParameter",
      "ssm:GetParameters"
    ]
    resources = [
      "arn:aws:ssm:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:parameter/${var.prefix}/*"
    ]
  }
}

resource "aws_iam_policy" "ssm_read_policy" {
  name        = "${var.prefix}-ssm-read-policy"
  description = "Allow reading from SSM Parameter Store"

  policy = data.aws_iam_policy_document.ssm_read_policy_doc.json
}

resource "aws_iam_role_policy_attachment" "ssm_read_attachment" {
  role       = aws_iam_role.infra_instances_service_role.name
  policy_arn = aws_iam_policy.ssm_read_policy.arn
}

resource "aws_ecr_repository" "orchestration_repository" {
  name = "e2b-orchestration"

  tags = var.labels
}

data "aws_iam_policy_document" "ecr_pull_policy_doc" {
  statement {
    effect = "Allow"
    actions = [
      "ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage",
      "ecr:BatchCheckLayerAvailability"
    ]
    resources = [
      aws_ecr_repository.orchestration_repository.arn
    ]
  }

  statement {
    effect = "Allow"
    actions = [
      "ecr:GetAuthorizationToken"
    ]
    # A wildcard is unavoidable, as the GetAuthorizationToken API doesn't support
    # resource-level permissions
    resources = ["*"]
  }
}

resource "aws_iam_policy" "ecr_pull_policy" {
  name        = "${var.prefix}-ecr-pull-policy"
  description = "Allow pulling from ECR repositories"

  policy = data.aws_iam_policy_document.ecr_pull_policy_doc.json
}

resource "aws_iam_role_policy_attachment" "ecr_pull_attachment" {
  role       = aws_iam_role.infra_instances_service_role.name
  policy_arn = aws_iam_policy.ecr_pull_policy.arn
}
