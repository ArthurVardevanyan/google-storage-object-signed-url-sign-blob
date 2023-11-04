// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccExampleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccExampleResourceConfig("one", "one"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("avgcp_example.test", "bucket", "one"),
					resource.TestCheckResourceAttr("avgcp_example.test", "path", "one"),
					resource.TestCheckResourceAttr("avgcp_example.test", "signed_url", "example-signed_url"),
				), // ImportState testing
			},
			{
				ResourceName:                         "avgcp_example.test",
				ImportState:                          true,
				ImportStateVerify:                    false, // true
				ImportStateVerifyIdentifierAttribute: "signed_url",
			},

			// Update and Read testing
			{
				Config: testAccExampleResourceConfig("two", "two"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("avgcp_example.test", "bucket", "two"),
					resource.TestCheckResourceAttr("avgcp_example.test", "path", "two"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccExampleResourceConfig(Bucket string, Path string) string {
	return fmt.Sprintf(`
resource "avgcp_example" "test" {
  bucket = %[1]q
  path = %[2]q
}
`, Bucket, Path)
}
