terraform {
  required_version = ">= 1.5.0, < 1.6.0"
  backend "s3" {
    bucket         = "zenlytic-code-interpreter-terraform-state"
    key            = "terraform/orchestration/state"
    dynamodb_table = "zenlytic-code-interpreter-terraform-state-lock"
    region         = "us-east-1"
    encrypt        = true
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.88.0"
    }
    docker = {
      source  = "kreuzwerker/docker"
      version = "3.0.2"
    }
    nomad = {
      source  = "hashicorp/nomad"
      version = "2.1.0"
    }
    github = {
      source  = "integrations/github"
      version = "5.42.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.5.1"
    }
  }
}

data "aws_caller_identity" "current" {}

data "aws_ecr_authorization_token" "token" {
  registry_id = data.aws_caller_identity.current.account_id
}

provider "aws" {
  region = var.aws_region
}

provider "docker" {
  registry_auth {
    address  = data.aws_ecr_authorization_token.token.proxy_endpoint
    username = data.aws_ecr_authorization_token.token.user_name
    password = data.aws_ecr_authorization_token.token.password
  }
}

module "init" {
  source = "./packages/init"

  labels = var.labels
  prefix = var.prefix
}

module "buckets" {
  source = "./packages/buckets"

  project_name      = var.project_name
  service_role_name = module.init.service_role_name

  fc_template_bucket_name   = length(var.template_bucket_name) > 0 ? var.template_bucket_name : "${var.project_name}-fc-templates"
  fc_template_bucket_region = var.template_bucket_region

  labels = var.labels
}

module "github_tf" {
  source = "./github-tf"

  project_name = var.project_name

  github_organization = var.github_organization
  github_repository   = var.github_repository
  github_branch       = var.github_branch

  domain_name                    = var.domain_name
  terraform_state_bucket         = var.terraform_state_bucket
  terraform_state_dynamodb_table = var.terraform_state_dynamodb_table
  kernel_bucket                  = module.buckets.fc_kernels_bucket_name
  fc_versions_bucket             = module.buckets.fc_versions_bucket_name
  public_builds_bucket           = module.buckets.public_builds_bucket_name

  prefix = var.prefix
}
