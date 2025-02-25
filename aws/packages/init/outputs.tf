output "service_profile_name" {
  value = aws_iam_instance_profile.infra_instances_service_profile.name
}

output "service_role_name" {
  value = aws_iam_role.infra_instances_service_role.name
}

output "consul_acl_token_secret" {
  value = aws_ssm_parameter.consul_acl_token.value
}

output "nomad_acl_token_secret" {
  value = aws_ssm_parameter.nomad_acl_token.value
}

output "grafana_api_key_secret_name" {
  value = aws_ssm_parameter.grafana_api_key.name
}

output "analytics_collector_host_secret_name" {
  value = aws_ssm_parameter.analytics_collector_host.name
}

output "analytics_collector_api_token_secret_name" {
  value = aws_ssm_parameter.analytics_collector_api_token.name
}

output "orchestration_repository_name" {
  value = aws_ecr_repository.orchestration_repository.name
}

output "cloudflare_api_token_secret_name" {
  value = aws_ssm_parameter.cloudflare_api_token.name
}
