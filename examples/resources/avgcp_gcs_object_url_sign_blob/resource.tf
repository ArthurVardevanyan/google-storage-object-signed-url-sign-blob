terraform {
  required_providers {
    avgcp = {
      source  = "terraform.local/local/avgcp"
      version = "1.0.0"
    }
  }
}

resource "avgcp_gcs_object_url_sign_blob" "example" {
  bucket = "some-value"
  path   = "some-value"
}

output "signed_url" {
  value = avgcp_gcs_object_url_sign_blob.example.signed_url
}