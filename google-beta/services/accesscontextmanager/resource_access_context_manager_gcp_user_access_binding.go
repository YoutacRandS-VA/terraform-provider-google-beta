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

package accesscontextmanager

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceAccessContextManagerGcpUserAccessBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessContextManagerGcpUserAccessBindingCreate,
		Read:   resourceAccessContextManagerGcpUserAccessBindingRead,
		Update: resourceAccessContextManagerGcpUserAccessBindingUpdate,
		Delete: resourceAccessContextManagerGcpUserAccessBindingDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAccessContextManagerGcpUserAccessBindingImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"access_levels": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Required. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"`,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"group_key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Required. Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the G Suite Directory API's Groups resource. If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"`,
			},
			"organization_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Required. ID of the parent organization.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by RFC 3986 Section 2.3). Should not be specified by the client during creation. Example: "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceAccessContextManagerGcpUserAccessBindingCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	groupKeyProp, err := expandAccessContextManagerGcpUserAccessBindingGroupKey(d.Get("group_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("group_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(groupKeyProp)) && (ok || !reflect.DeepEqual(v, groupKeyProp)) {
		obj["groupKey"] = groupKeyProp
	}
	accessLevelsProp, err := expandAccessContextManagerGcpUserAccessBindingAccessLevels(d.Get("access_levels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("access_levels"); !tpgresource.IsEmptyValue(reflect.ValueOf(accessLevelsProp)) && (ok || !reflect.DeepEqual(v, accessLevelsProp)) {
		obj["accessLevels"] = accessLevelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AccessContextManagerBasePath}}organizations/{{organization_id}}/gcpUserAccessBindings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GcpUserAccessBinding: %#v", obj)
	billingProject := ""

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
		return fmt.Errorf("Error creating GcpUserAccessBinding: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = AccessContextManagerOperationWaitTimeWithResponse(
		config, res, &opRes, "Creating GcpUserAccessBinding", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create GcpUserAccessBinding: %s", err)
	}

	if err := d.Set("name", flattenAccessContextManagerGcpUserAccessBindingName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating GcpUserAccessBinding %q: %#v", d.Id(), res)

	return resourceAccessContextManagerGcpUserAccessBindingRead(d, meta)
}

func resourceAccessContextManagerGcpUserAccessBindingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AccessContextManagerBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("AccessContextManagerGcpUserAccessBinding %q", d.Id()))
	}

	if err := d.Set("name", flattenAccessContextManagerGcpUserAccessBindingName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading GcpUserAccessBinding: %s", err)
	}
	if err := d.Set("group_key", flattenAccessContextManagerGcpUserAccessBindingGroupKey(res["groupKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading GcpUserAccessBinding: %s", err)
	}
	if err := d.Set("access_levels", flattenAccessContextManagerGcpUserAccessBindingAccessLevels(res["accessLevels"], d, config)); err != nil {
		return fmt.Errorf("Error reading GcpUserAccessBinding: %s", err)
	}

	return nil
}

func resourceAccessContextManagerGcpUserAccessBindingUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	accessLevelsProp, err := expandAccessContextManagerGcpUserAccessBindingAccessLevels(d.Get("access_levels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("access_levels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, accessLevelsProp)) {
		obj["accessLevels"] = accessLevelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AccessContextManagerBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating GcpUserAccessBinding %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("access_levels") {
		updateMask = append(updateMask, "accessLevels")
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
			return fmt.Errorf("Error updating GcpUserAccessBinding %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating GcpUserAccessBinding %q: %#v", d.Id(), res)
		}

		err = AccessContextManagerOperationWaitTime(
			config, res, "Updating GcpUserAccessBinding", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceAccessContextManagerGcpUserAccessBindingRead(d, meta)
}

func resourceAccessContextManagerGcpUserAccessBindingDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{AccessContextManagerBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	log.Printf("[DEBUG] Deleting GcpUserAccessBinding %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "GcpUserAccessBinding")
	}

	err = AccessContextManagerOperationWaitTime(
		config, res, "Deleting GcpUserAccessBinding", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting GcpUserAccessBinding %q: %#v", d.Id(), res)
	return nil
}

func resourceAccessContextManagerGcpUserAccessBindingImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	name := d.Get("name").(string)

	if err := d.Set("name", name); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name)
	return []*schema.ResourceData{d}, nil
}

func flattenAccessContextManagerGcpUserAccessBindingName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessContextManagerGcpUserAccessBindingGroupKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessContextManagerGcpUserAccessBindingAccessLevels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandAccessContextManagerGcpUserAccessBindingGroupKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerGcpUserAccessBindingAccessLevels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
