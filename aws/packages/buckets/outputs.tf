output "loki_bucket_name" {
  value = aws_s3_bucket.loki_storage_bucket.id
}

output "cluster_setup_bucket_name" {
  value = aws_s3_bucket.setup_bucket.id
}

output "fc_env_pipeline_bucket_name" {
  description = "The name of the bucket to store the files for firecracker environment pipeline"
  value       = aws_s3_bucket.fc_env_pipeline_bucket.id
}

output "fc_kernels_bucket_name" {
  value = aws_s3_bucket.fc_kernels_bucket.id
}

output "fc_versions_bucket_name" {
  value = aws_s3_bucket.fc_versions_bucket.id
}

output "fc_template_bucket_name" {
  value = aws_s3_bucket.fc_template_bucket.id
}
