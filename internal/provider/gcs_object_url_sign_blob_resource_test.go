// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGcsObjectUrlSignBlobResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccGcsObjectUrlSignBlobResourceConfig(os.Getenv("GCP_SERVICE_ACCOUNT"), "terraform_provider_avgcp", "test"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("avgcp_gcs_object_url_sign_blob.test", "google_access_id", os.Getenv("GCP_SERVICE_ACCOUNT")),
					resource.TestCheckResourceAttr("avgcp_gcs_object_url_sign_blob.test", "bucket", "terraform_provider_avgcp"),
					resource.TestCheckResourceAttr("avgcp_gcs_object_url_sign_blob.test", "path", "test"),
					//resource.TestMatchResourceAttr("avgcp_gcs_object_url_sign_blob.test", "signed_url", &regexp.Regexp{}),
				), // ImportState testing
			},
			{
				ResourceName:                         "avgcp_gcs_object_url_sign_blob.test",
				ImportState:                          true,
				ImportStateVerify:                    false, // true
				ImportStateVerifyIdentifierAttribute: "signed_url",
			},

			// Update and Read testing
			{
				Config: testAccGcsObjectUrlSignBlobResourceConfig(os.Getenv("GCP_SERVICE_ACCOUNT"), "terraform_provider_avgcp", "test"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("avgcp_gcs_object_url_sign_blob.test", "google_access_id", os.Getenv("GCP_SERVICE_ACCOUNT")),
					resource.TestCheckResourceAttr("avgcp_gcs_object_url_sign_blob.test", "bucket", "terraform_provider_avgcp"),
					resource.TestCheckResourceAttr("avgcp_gcs_object_url_sign_blob.test", "path", "test"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccGcsObjectUrlSignBlobResourceConfig(GoogleAccessID string, Bucket string, Path string) string {
	return fmt.Sprintf(`
resource "avgcp_gcs_object_url_sign_blob" "test" {
  google_access_id = %[1]q
  bucket = %[2]q
  path = %[3]q
}
`, GoogleAccessID, Bucket, Path)
}
