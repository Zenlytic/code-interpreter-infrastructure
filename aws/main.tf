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
