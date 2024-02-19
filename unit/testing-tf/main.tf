// terraform file used for unit tests
terraform {
  required_providers {
    random = {
      source  = "hashicorp/random"
      version = "3.1.0"
    }
  }
}

provider "random" {
}

variable "length" {
  type    = number
  default = 16
  validation {
    condition     = var.length == 16
    error_message = "Random string length must be 16."
  }
}

resource "random_string" "s" {
  length = var.length
}

resource "local_file" "json" {
  filename = "${path.module}/values.json"
  content = jsonencode({
    length = var.length
    nested = {
      value = var.length
    }
  })
}

resource "local_file" "yaml" {
  filename = "${path.module}/values.yaml"
  content = yamlencode({
    length = var.length
    nested = {
      value = var.length
    }
  })
}

locals {
  helm_values = <<VALUES
securityContext:
  fsGroup: 1001
installCRDs: true
VALUES
}

resource "helm_release" "cert_manager" {
  chart = "cert-manager"
  repository = "https://charts.jetstack.io"
  name = "cert-manager"
  values = [
    local.helm_values
  ]
}

output "random_string_result" {
  value = random_string.s.result
}