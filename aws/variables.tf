variable "project_name" {
  type        = string
  description = "The name of the project, used to name resources uniquely"
  default     = "zenlytic-code-interpreter"
}

variable "aws_region" {
  type = string
}

variable "github_organization" {
  type    = string
  default = "Zenlytic"
}

variable "github_repository" {
  type    = string
  default = "code-interpreter-infrastructure"
}

variable "github_branch" {
  type    = string
  default = "zenlytic"
}

variable "domain_name" {
  type        = string
  description = "The domain name where e2b will run"
}

variable "prefix" {
  type        = string
  description = "The prefix to use for all resources in this module"
  default     = "e2b"
}

variable "labels" {
  description = "The labels to attach to resources created by this module"
  type        = map(string)
  default = {
    "app"       = "e2b"
    "terraform" = "true"
  }
}

variable "terraform_state_bucket" {
  description = "The name of the bucket to store terraform state in"
  type        = string
  default     = "zenlytic-code-interpreter-terraform-state"
}

variable "terraform_state_dynamodb_table" {
  description = "The name of the dynamodb table to lock terraform state"
  type        = string
  default     = "zenlytic-code-interpreter-terraform-state-lock"
}

variable "template_bucket_region" {
  type        = string
  description = "The region of the FC template bucket"
}

variable "template_bucket_name" {
  type        = string
  description = "The name of the FC template bucket"
}
