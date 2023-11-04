// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
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
				Config: testAccGcsObjectUrlSignBlobResourceConfig("one", "one"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("avgcp_gcs_object_url_sign_blob.test", "bucket", "one"),
					resource.TestCheckResourceAttr("avgcp_gcs_object_url_sign_blob.test", "path", "one"),
					resource.TestCheckResourceAttr("avgcp_gcs_object_url_sign_blob.test", "signed_url", "gcs_object_url_sign_blob-signed_url"),
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
				Config: testAccGcsObjectUrlSignBlobResourceConfig("two", "two"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("avgcp_gcs_object_url_sign_blob.test", "bucket", "two"),
					resource.TestCheckResourceAttr("avgcp_gcs_object_url_sign_blob.test", "path", "two"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccGcsObjectUrlSignBlobResourceConfig(Bucket string, Path string) string {
	return fmt.Sprintf(`
resource "avgcp_gcs_object_url_sign_blob" "test" {
  bucket = %[1]q
  path = %[2]q
}
`, Bucket, Path)
}
