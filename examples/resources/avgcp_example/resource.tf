terraform {
  required_providers {
    avgcp = {
      source  = "terraform.local/local/avgcp"
      version = "1.0.0"
    }
  }
}

resource "avgcp_example" "example" {
  bucket = "some-value"
  path = "some-value"
}

output "signed_url" {
  value = avgcp_example.example.signed_url
}