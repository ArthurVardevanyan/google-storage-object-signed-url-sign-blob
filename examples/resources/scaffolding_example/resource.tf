terraform {
  required_providers {
    scaffolding = {
      source  = "terraform.local/local/scaffolding"
      version = "1.0.0"
    }
  }
}

resource "scaffolding_example" "example" {
  bucket = "some-value"
  path = "some-value"
}

output "signed_url" {
  value = scaffolding_example.example.signed_url
}