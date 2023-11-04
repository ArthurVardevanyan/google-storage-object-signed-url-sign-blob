// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGcsObjectUrlSignBlobDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccGcsObjectUrlSignBlobDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.avgcp_gcs_object_url_sign_blob.test", "signed_url", "gcs_object_url_sign_blob-signed_url"),
				),
			},
		},
	})
}

const testAccGcsObjectUrlSignBlobDataSourceConfig = `
data "avgcp_gcs_object_url_sign_blob" "test" {
  google_access_id = "some"
  bucket = "example"
  path = "some-value"
}
`
