variable "project_name" {
  type        = string
  description = "The name of the project, used to name resources uniquely"
}

variable "github_organization" {
  description = "The name of the github organization"
  type        = string
}

variable "github_repository" {
  description = "The name of the repository"
  type        = string
}

variable "github_branch" {
  description = "The name of the branch to deploy"
  type        = string
}

variable "prefix" {
  description = "The prefix to use for all resources in this module"
  type        = string
}

variable "terraform_state_bucket" {
  description = "The name of the bucket to store terraform state in"
  type        = string
}

variable "terraform_state_dynamodb_table" {
  description = "The name of the dynamodb table to lock terraform state"
  type        = string
}

variable "fc_versions_bucket" {
  description = "The name of the bucket to store build fc versions"
  type        = string
}

variable "kernel_bucket" {
  description = "The name of the bucket to store built kernels"
  type        = string
}

variable "public_builds_bucket" {
  description = "The name of the bucket to store public builds"
  type        = string
}

variable "domain_name" {
  type = string
}
