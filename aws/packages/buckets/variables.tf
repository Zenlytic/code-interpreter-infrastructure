variable "project_name" {
  type        = string
  description = "The name of the project, used to name resources uniquely"
}

variable "service_role_name" {
  type = string
}

variable "labels" {
  description = "The labels to attach to resources created by this module"
  type        = map(string)
}

variable "fc_template_bucket_region" {
  type        = string
  description = "The region of the FC template bucket"
}

variable "fc_template_bucket_name" {
  type        = string
  description = "The name of the FC template bucket"
}
