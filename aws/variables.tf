variable "project_name" {
  type        = string
  description = "The name of the project, used to name resources uniquely"
  default     = "zenlytic-code-interpreter"
}

variable "aws_region" {
  type = string
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

variable "template_bucket_region" {
  type        = string
  description = "The region of the FC template bucket"
}

variable "template_bucket_name" {
  type        = string
  description = "The name of the FC template bucket"
}
