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

resource "random_string" "s" {
  length = 16
}
