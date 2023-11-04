terraform {
  required_providers {
    avgcp = {
      source  = "terraform.local/local/avgcp"
      version = "1.0.0"
    }
  }
}

resource "avgcp_gcs_object_url_sign_blob" "example" {
  google_access_id = "tf-avgcp@homelab-X.iam.gserviceaccount.com"
  bucket           = "terraform_provider_avgcp"
  path             = "test"
}

output "signed_url" {
  value = avgcp_gcs_object_url_sign_blob.example.signed_url
}