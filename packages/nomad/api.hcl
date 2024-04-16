variable "gcp_zone" {
  type    = string
}

variable "image_name" {
  type    = string
  default = ""
}

variable "api_port_name" {
  type    = string
  default = ""
}

variable "api_port_number" {
  type    = number
  default = 0
}

variable "nomad_token" {
  type    = string
  default = ""
}

variable "nomad_address" {
  type    = string
  default = ""
}

variable "postgres_connection_string" {
  type    = string
  default = ""
}

variable "posthog_api_key" {
  type    = string
  default = ""
}

variable "environment" {
  type    = string
  default = ""
}

variable "analytics_collector_host" {
  type    = string
  default = ""
}

variable "analytics_collector_api_token" {
  type    = string
  default = ""
}

variable "otel_tracing_print" {
  type    = string
  default = ""
}

variable "loki_address" {
  type = string
  default = ""
}

variable "orchestrator_address" {
  type    = string
  default = ""
}

variable "template_manager_address" {
  type    = string
  default = ""
}

job "orchestration-api" {
  datacenters = [var.gcp_zone]

  priority = 90

  group "api-service" {
    network {
      port "api" {
        static = var.api_port_number
      }
    }

    service {
      name = "api"
      port = var.api_port_number

      check {
        type     = "http"
        name     = "health"
        path     = "/health"
        interval = "20s"
        timeout  = "5s"
        port     = var.api_port_number
      }
    }

    task "start" {
      driver = "docker"

      resources {
        memory     = 2048
        memory_max = 2048
        cpu        = 1024
      }

      env {
        ORCHESTRATOR_ADDRESS          = var.orchestrator_address
        TEMPLATE_MANAGER_ADDRESS      = var.template_manager_address
        NOMAD_ADDRESS                 = var.nomad_address
        NOMAD_TOKEN                   = var.nomad_token
        POSTGRES_CONNECTION_STRING    = var.postgres_connection_string
        ENVIRONMENT                   = var.environment
        POSTHOG_API_KEY               = var.posthog_api_key
        ANALYTICS_COLLECTOR_HOST      = var.analytics_collector_host
        ANALYTICS_COLLECTOR_API_TOKEN = var.analytics_collector_api_token
        LOKI_ADDRESS                  = var.loki_address
        OTEL_TRACING_PRINT            = var.otel_tracing_print
      }

      config {
        network_mode = "host"
        image        = var.image_name
        ports        = [var.api_port_name]
        args = [
          "--port", "${var.api_port_number}",
        ]
      }
    }
  }
}
