resource "scaffolding_example" "example" {
  configurable_attribute = "some-value"
}
terraform {
  required_providers {
    scaffolding = {
      source  = "terraform.local/local/scaffolding"
      version = "1.0.0"
    }
  }
}


