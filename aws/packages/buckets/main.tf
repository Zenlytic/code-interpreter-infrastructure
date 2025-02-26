resource "aws_s3_bucket" "loki_storage_bucket" {
  bucket = "${var.project_name}-loki-storage"

  tags = var.labels
}

resource "aws_s3_bucket_lifecycle_configuration" "loki_storage_lifecycle" {
  bucket = aws_s3_bucket.loki_storage_bucket.id

  rule {
    id     = "delete-old-logs"
    status = "Enabled"

    expiration {
      days = 8
    }
  }
}

resource "aws_s3_bucket_public_access_block" "loki_storage_block" {
  bucket = aws_s3_bucket.loki_storage_bucket.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket" "setup_bucket" {
  bucket = "${var.project_name}-instance-setup"

  tags = var.labels
}

resource "aws_s3_bucket_public_access_block" "setup_bucket_block" {
  bucket = aws_s3_bucket.setup_bucket.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket" "fc_kernels_bucket" {
  bucket = "${var.project_name}-fc-kernels"

  tags = var.labels
}

resource "aws_s3_bucket_public_access_block" "fc_kernels_bucket_block" {
  bucket = aws_s3_bucket.fc_kernels_bucket.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket" "fc_versions_bucket" {
  bucket = "${var.project_name}-fc-versions"

  tags = var.labels
}

resource "aws_s3_bucket_public_access_block" "fc_versions_bucket_block" {
  bucket = aws_s3_bucket.fc_versions_bucket.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket" "fc_env_pipeline_bucket" {
  bucket = "${var.project_name}-fc-env-pipeline"

  # Use a provider configured with the template bucket region
  provider = aws.fc_template_bucket_region

  tags = var.labels
}

resource "aws_s3_bucket_public_access_block" "fc_env_pipeline_bucket_block" {
  bucket = aws_s3_bucket.fc_env_pipeline_bucket.id

  provider = aws.fc_template_bucket_region

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket" "fc_template_bucket" {
  bucket = var.fc_template_bucket_name

  provider = aws.fc_template_bucket_region

  tags = var.labels
}

resource "aws_s3_bucket_intelligent_tiering_configuration" "fc_template_bucket_tiering" {
  bucket = aws_s3_bucket.fc_template_bucket.id
  name   = "EntireTemplateBucket"

  provider = aws.fc_template_bucket_region

  status = "Enabled"

  tiering {
    access_tier = "ARCHIVE_ACCESS"
    days        = 90
  }

  tiering {
    access_tier = "DEEP_ARCHIVE_ACCESS"
    days        = 180
  }
}

resource "aws_s3_bucket_lifecycle_configuration" "fc_template_bucket_lifecycle" {
  bucket = aws_s3_bucket.fc_template_bucket.id

  provider = aws.fc_template_bucket_region

  rule {
    id     = "soft-delete"
    status = "Enabled"

    expiration {
      days = 7
    }
  }
}

resource "aws_s3_bucket_public_access_block" "fc_template_bucket_block" {
  bucket = aws_s3_bucket.fc_template_bucket.id

  provider = aws.fc_template_bucket_region

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

data "aws_iam_policy_document" "loki_storage_policy" {
  statement {
    actions = [
      "s3:PutObject",
      "s3:GetObject",
      "s3:DeleteObject",
      "s3:ListBucket"
    ]
    resources = [
      aws_s3_bucket.loki_storage_bucket.arn,
      "${aws_s3_bucket.loki_storage_bucket.arn}/*"
    ]
  }
}

resource "aws_iam_policy" "loki_storage_policy" {
  name   = "${var.project_name}-loki-storage-policy"
  policy = data.aws_iam_policy_document.loki_storage_policy.json
}

resource "aws_iam_role_policy_attachment" "loki_storage_policy_attachment" {
  role       = var.service_role_name
  policy_arn = aws_iam_policy.loki_storage_policy.arn
}

data "aws_iam_policy_document" "setup_bucket_policy" {
  statement {
    actions = [
      "s3:GetObject",
      "s3:ListBucket"
    ]
    resources = [
      aws_s3_bucket.setup_bucket.arn,
      "${aws_s3_bucket.setup_bucket.arn}/*"
    ]
  }
}

resource "aws_iam_policy" "setup_bucket_policy" {
  name   = "${var.project_name}-setup-bucket-policy"
  policy = data.aws_iam_policy_document.setup_bucket_policy.json
}

resource "aws_iam_role_policy_attachment" "setup_bucket_policy_attachment" {
  role       = var.service_role_name
  policy_arn = aws_iam_policy.setup_bucket_policy.arn
}

data "aws_iam_policy_document" "fc_kernels_bucket_policy" {
  statement {
    actions = [
      "s3:GetObject",
      "s3:ListBucket"
    ]
    resources = [
      aws_s3_bucket.fc_kernels_bucket.arn,
      "${aws_s3_bucket.fc_kernels_bucket.arn}/*"
    ]
  }
}

resource "aws_iam_policy" "fc_kernels_bucket_policy" {
  name   = "${var.project_name}-fc-kernels-bucket-policy"
  policy = data.aws_iam_policy_document.fc_kernels_bucket_policy.json
}

resource "aws_iam_role_policy_attachment" "fc_kernels_bucket_policy_attachment" {
  role       = var.service_role_name
  policy_arn = aws_iam_policy.fc_kernels_bucket_policy.arn
}

data "aws_iam_policy_document" "fc_versions_bucket_policy" {
  statement {
    actions = [
      "s3:GetObject",
      "s3:ListBucket"
    ]
    resources = [
      aws_s3_bucket.fc_versions_bucket.arn,
      "${aws_s3_bucket.fc_versions_bucket.arn}/*"
    ]
  }
}

resource "aws_iam_policy" "fc_versions_bucket_policy" {
  name   = "${var.project_name}-fc-versions-bucket-policy"
  policy = data.aws_iam_policy_document.fc_versions_bucket_policy.json
}

resource "aws_iam_role_policy_attachment" "fc_versions_bucket_policy_attachment" {
  role       = var.service_role_name
  policy_arn = aws_iam_policy.fc_versions_bucket_policy.arn
}

data "aws_iam_policy_document" "fc_env_pipeline_bucket_policy" {
  statement {
    actions = [
      "s3:GetObject",
      "s3:ListBucket"
    ]
    resources = [
      aws_s3_bucket.fc_env_pipeline_bucket.arn,
      "${aws_s3_bucket.fc_env_pipeline_bucket.arn}/*"
    ]
  }
}

resource "aws_iam_policy" "fc_env_pipeline_bucket_policy" {
  name     = "${var.project_name}-fc-env-pipeline-bucket-policy"
  policy   = data.aws_iam_policy_document.fc_env_pipeline_bucket_policy.json
  provider = aws.fc_template_bucket_region
}

resource "aws_iam_role_policy_attachment" "fc_env_pipeline_bucket_policy_attachment" {
  role       = var.service_role_name
  policy_arn = aws_iam_policy.fc_env_pipeline_bucket_policy.arn
  provider   = aws.fc_template_bucket_region
}

data "aws_iam_policy_document" "fc_template_bucket_policy" {
  statement {
    actions = [
      "s3:PutObject",
      "s3:GetObject",
      "s3:DeleteObject",
      "s3:ListBucket"
    ]
    resources = [
      aws_s3_bucket.fc_template_bucket.arn,
      "${aws_s3_bucket.fc_template_bucket.arn}/*"
    ]
  }
}

resource "aws_iam_policy" "fc_template_bucket_policy" {
  name     = "${var.project_name}-fc-template-bucket-policy"
  policy   = data.aws_iam_policy_document.fc_template_bucket_policy.json
  provider = aws.fc_template_bucket_region
}

resource "aws_iam_role_policy_attachment" "fc_template_bucket_policy_attachment" {
  role       = var.service_role_name
  policy_arn = aws_iam_policy.fc_template_bucket_policy.arn
  provider   = aws.fc_template_bucket_region
}

resource "aws_s3_bucket" "public_builds_storage_bucket" {
  # Only create the bucket if the project name contains "prod"
  count  = can(regex("prod", var.project_name)) ? 1 : 0
  bucket = "${var.project_name}-public-builds"

  tags = var.labels
}

resource "aws_s3_bucket_lifecycle_configuration" "public_builds_lifecycle" {
  count  = can(regex("prod", var.project_name)) ? 1 : 0
  bucket = aws_s3_bucket.public_builds_storage_bucket[0].id

  rule {
    id     = "delete-old-builds"
    status = "Enabled"

    expiration {
      days = 30
    }
  }
}

data "aws_iam_policy_document" "public_builds_bucket_policy" {
  count = can(regex("prod", var.project_name)) ? 1 : 0

  statement {
    actions   = ["s3:GetObject"]
    resources = ["${aws_s3_bucket.public_builds_storage_bucket[0].arn}/*"]
    principals {
      type        = "*"
      identifiers = ["*"]
    }
  }
}

resource "aws_s3_bucket_policy" "public_builds_bucket_policy" {
  count  = can(regex("prod", var.project_name)) ? 1 : 0
  bucket = aws_s3_bucket.public_builds_storage_bucket[0].id
  policy = data.aws_iam_policy_document.public_builds_bucket_policy[0].json
}
