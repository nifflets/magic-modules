package resourcemanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccDataSourceGoogleFolderOrganizationPolicy_basic(t *testing.T) {
	folder := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))
	org := envvar.GetTestOrgFromEnv(t)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleFolderOrganizationPolicy_basic(org, folder),
				Check: acctest.CheckDataSourceStateMatchesResourceState(
					"data.google_folder_organization_policy.data",
					"google_folder_organization_policy.resource",
				),
			},
		},
	})
}

func testAccDataSourceGoogleFolderOrganizationPolicy_basic(org, folder string) string {
	return fmt.Sprintf(`
resource "google_folder" "orgpolicy" {
  display_name = "%s"
  parent       = "%s"
  deletion_protection = false
}

resource "google_folder_organization_policy" "resource" {
  folder     = google_folder.orgpolicy.name
  constraint = "serviceuser.services"

  restore_policy {
    default = true
  }
}

data "google_folder_organization_policy" "data" {
  folder     = google_folder_organization_policy.resource.folder
  constraint = "serviceuser.services"
}
`, folder, "organizations/"+org)
}