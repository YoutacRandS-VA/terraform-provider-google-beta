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

package beyondcorp

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceBeyondcorpAppConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceBeyondcorpAppConnectorCreate,
		Read:   resourceBeyondcorpAppConnectorRead,
		Update: resourceBeyondcorpAppConnectorUpdate,
		Delete: resourceBeyondcorpAppConnectorDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBeyondcorpAppConnectorImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `ID of the AppConnector.`,
			},
			"principal_info": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Principal information about the Identity of the AppConnector.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_account": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `ServiceAccount represents a GCP service account.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"email": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Email address of the service account.`,
									},
								},
							},
						},
					},
				},
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An arbitrary user-provided name for the AppConnector.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Resource labels to represent user provided metadata.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The region of the AppConnector.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Represents the different states of a AppConnector.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceBeyondcorpAppConnectorCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandBeyondcorpAppConnectorDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	principalInfoProp, err := expandBeyondcorpAppConnectorPrincipalInfo(d.Get("principal_info"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("principal_info"); !tpgresource.IsEmptyValue(reflect.ValueOf(principalInfoProp)) && (ok || !reflect.DeepEqual(v, principalInfoProp)) {
		obj["principalInfo"] = principalInfoProp
	}
	labelsProp, err := expandBeyondcorpAppConnectorEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{region}}/appConnectors?app_connector_id={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AppConnector: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AppConnector: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating AppConnector: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/appConnectors/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = BeyondcorpOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating AppConnector", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create AppConnector: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/appConnectors/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating AppConnector %q: %#v", d.Id(), res)

	return resourceBeyondcorpAppConnectorRead(d, meta)
}

func resourceBeyondcorpAppConnectorRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{region}}/appConnectors/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AppConnector: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("BeyondcorpAppConnector %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading AppConnector: %s", err)
	}

	if err := d.Set("display_name", flattenBeyondcorpAppConnectorDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppConnector: %s", err)
	}
	if err := d.Set("labels", flattenBeyondcorpAppConnectorLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppConnector: %s", err)
	}
	if err := d.Set("principal_info", flattenBeyondcorpAppConnectorPrincipalInfo(res["principalInfo"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppConnector: %s", err)
	}
	if err := d.Set("state", flattenBeyondcorpAppConnectorState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppConnector: %s", err)
	}
	if err := d.Set("terraform_labels", flattenBeyondcorpAppConnectorTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppConnector: %s", err)
	}
	if err := d.Set("effective_labels", flattenBeyondcorpAppConnectorEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppConnector: %s", err)
	}

	return nil
}

func resourceBeyondcorpAppConnectorUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AppConnector: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandBeyondcorpAppConnectorDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	principalInfoProp, err := expandBeyondcorpAppConnectorPrincipalInfo(d.Get("principal_info"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("principal_info"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, principalInfoProp)) {
		obj["principalInfo"] = principalInfoProp
	}
	labelsProp, err := expandBeyondcorpAppConnectorEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{region}}/appConnectors/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AppConnector %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("principal_info") {
		updateMask = append(updateMask, "principalInfo")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
		})

		if err != nil {
			return fmt.Errorf("Error updating AppConnector %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating AppConnector %q: %#v", d.Id(), res)
		}

		err = BeyondcorpOperationWaitTime(
			config, res, project, "Updating AppConnector", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceBeyondcorpAppConnectorRead(d, meta)
}

func resourceBeyondcorpAppConnectorDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AppConnector: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{region}}/appConnectors/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	log.Printf("[DEBUG] Deleting AppConnector %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "AppConnector")
	}

	err = BeyondcorpOperationWaitTime(
		config, res, project, "Deleting AppConnector", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting AppConnector %q: %#v", d.Id(), res)
	return nil
}

func resourceBeyondcorpAppConnectorImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/appConnectors/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/appConnectors/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBeyondcorpAppConnectorDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppConnectorLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenBeyondcorpAppConnectorPrincipalInfo(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["service_account"] =
		flattenBeyondcorpAppConnectorPrincipalInfoServiceAccount(original["serviceAccount"], d, config)
	return []interface{}{transformed}
}
func flattenBeyondcorpAppConnectorPrincipalInfoServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["email"] =
		flattenBeyondcorpAppConnectorPrincipalInfoServiceAccountEmail(original["email"], d, config)
	return []interface{}{transformed}
}
func flattenBeyondcorpAppConnectorPrincipalInfoServiceAccountEmail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppConnectorState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppConnectorTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenBeyondcorpAppConnectorEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandBeyondcorpAppConnectorDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectorPrincipalInfo(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedServiceAccount, err := expandBeyondcorpAppConnectorPrincipalInfoServiceAccount(original["service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAccount"] = transformedServiceAccount
	}

	return transformed, nil
}

func expandBeyondcorpAppConnectorPrincipalInfoServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEmail, err := expandBeyondcorpAppConnectorPrincipalInfoServiceAccountEmail(original["email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["email"] = transformedEmail
	}

	return transformed, nil
}

func expandBeyondcorpAppConnectorPrincipalInfoServiceAccountEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectorEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
