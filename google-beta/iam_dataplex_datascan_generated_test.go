// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccDataplexDatascanIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  acctest.GetTestProjectFromEnv(),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexDatascanIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataplex_datascan_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/dataScans/%s roles/viewer", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-datascan%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccDataplexDatascanIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_dataplex_datascan_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/dataScans/%s roles/viewer", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-datascan%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataplexDatascanIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  acctest.GetTestProjectFromEnv(),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccDataplexDatascanIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataplex_datascan_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/dataScans/%s roles/viewer user:admin@hashicorptest.com", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-datascan%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataplexDatascanIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  acctest.GetTestProjectFromEnv(),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexDatascanIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_dataplex_datascan_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_dataplex_datascan_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/dataScans/%s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-datascan%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataplexDatascanIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_dataplex_datascan_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/dataScans/%s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-datascan%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataplexDatascanIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataplex_datascan" "basic_profile" {
  location     = "us-central1"
  data_scan_id = "tf-test-datascan%{random_suffix}"

  data {
	  resource = "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare"
  }

  execution_spec {
    trigger {
      on_demand {}
    }
  }

  data_profile_spec {  
  }

  project = "%{project_name}"
}

resource "google_dataplex_datascan_iam_member" "foo" {
  project = google_dataplex_datascan.basic_profile.project
  location = google_dataplex_datascan.basic_profile.location
  data_scan_id = google_dataplex_datascan.basic_profile.data_scan_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccDataplexDatascanIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataplex_datascan" "basic_profile" {
  location     = "us-central1"
  data_scan_id = "tf-test-datascan%{random_suffix}"

  data {
	  resource = "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare"
  }

  execution_spec {
    trigger {
      on_demand {}
    }
  }

  data_profile_spec {  
  }

  project = "%{project_name}"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_dataplex_datascan_iam_policy" "foo" {
  project = google_dataplex_datascan.basic_profile.project
  location = google_dataplex_datascan.basic_profile.location
  data_scan_id = google_dataplex_datascan.basic_profile.data_scan_id
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_dataplex_datascan_iam_policy" "foo" {
  project = google_dataplex_datascan.basic_profile.project
  location = google_dataplex_datascan.basic_profile.location
  data_scan_id = google_dataplex_datascan.basic_profile.data_scan_id
  depends_on = [
    google_dataplex_datascan_iam_policy.foo
  ]
}
`, context)
}

func testAccDataplexDatascanIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataplex_datascan" "basic_profile" {
  location     = "us-central1"
  data_scan_id = "tf-test-datascan%{random_suffix}"

  data {
	  resource = "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare"
  }

  execution_spec {
    trigger {
      on_demand {}
    }
  }

  data_profile_spec {  
  }

  project = "%{project_name}"
}

data "google_iam_policy" "foo" {
}

resource "google_dataplex_datascan_iam_policy" "foo" {
  project = google_dataplex_datascan.basic_profile.project
  location = google_dataplex_datascan.basic_profile.location
  data_scan_id = google_dataplex_datascan.basic_profile.data_scan_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataplexDatascanIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataplex_datascan" "basic_profile" {
  location     = "us-central1"
  data_scan_id = "tf-test-datascan%{random_suffix}"

  data {
	  resource = "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare"
  }

  execution_spec {
    trigger {
      on_demand {}
    }
  }

  data_profile_spec {  
  }

  project = "%{project_name}"
}

resource "google_dataplex_datascan_iam_binding" "foo" {
  project = google_dataplex_datascan.basic_profile.project
  location = google_dataplex_datascan.basic_profile.location
  data_scan_id = google_dataplex_datascan.basic_profile.data_scan_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccDataplexDatascanIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataplex_datascan" "basic_profile" {
  location     = "us-central1"
  data_scan_id = "tf-test-datascan%{random_suffix}"

  data {
	  resource = "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare"
  }

  execution_spec {
    trigger {
      on_demand {}
    }
  }

  data_profile_spec {  
  }

  project = "%{project_name}"
}

resource "google_dataplex_datascan_iam_binding" "foo" {
  project = google_dataplex_datascan.basic_profile.project
  location = google_dataplex_datascan.basic_profile.location
  data_scan_id = google_dataplex_datascan.basic_profile.data_scan_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
