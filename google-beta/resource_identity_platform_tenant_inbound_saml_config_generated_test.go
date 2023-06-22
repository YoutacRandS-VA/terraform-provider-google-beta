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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccIdentityPlatformTenantInboundSamlConfig_identityPlatformTenantInboundSamlConfigBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"name":          "saml.tf-config-" + RandString(t, 10),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIdentityPlatformTenantInboundSamlConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIdentityPlatformTenantInboundSamlConfig_identityPlatformTenantInboundSamlConfigBasicExample(context),
			},
			{
				ResourceName:            "google_identity_platform_tenant_inbound_saml_config.tenant_saml_config",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"tenant"},
			},
		},
	})
}

func testAccIdentityPlatformTenantInboundSamlConfig_identityPlatformTenantInboundSamlConfigBasicExample(context map[string]interface{}) string {
	return tpgresource.Nprintf(`
resource "google_identity_platform_tenant" "tenant" {
  display_name  = "tenant"
}

resource "google_identity_platform_tenant_inbound_saml_config" "tenant_saml_config" {
  name         = "%{name}"
  display_name = "Display Name"
  tenant       = google_identity_platform_tenant.tenant.name
  idp_config {
    idp_entity_id = "tf-test-tf-idp%{random_suffix}"
    sign_request  = true
    sso_url       = "https://example.com"
    idp_certificates {
      x509_certificate = file("test-fixtures/rsa_cert.pem")
    }
  }

  sp_config {
    sp_entity_id = "tf-test-tf-sp%{random_suffix}"
    callback_uri = "https://example.com"
  }
}
`, context)
}

func testAccCheckIdentityPlatformTenantInboundSamlConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_identity_platform_tenant_inbound_saml_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{tenant}}/inboundSamlConfigs/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("IdentityPlatformTenantInboundSamlConfig still exists at %s", url)
			}
		}

		return nil
	}
}
