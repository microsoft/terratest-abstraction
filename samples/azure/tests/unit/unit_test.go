package unittest

import (
	"samples/azure/tests"
	"testing"

	"github.com/microsoft/terratestabstraction/unit"
)

func TestTemplateUnit(t *testing.T) {

	// This is a JSON representation of what the Terraform plan should produce if run against a "fresh" environment.
	// Some things to keep in mind:
	//	(1) The key will be the path to the resource. This will match the path found from `terraform state list`. For
	//		this template, it is simply `{RESOURCE_TYPE}.{RESOURCE_NAME}` but this will be more complex if you are using
	//		modules in your deployment!
	//	(2) The test will only verify that the data is specified here exists in the plan. If you omit fields from the
	//		description, they will be ignored.
	resourceDescription := unit.ResourceDescription{
		"azurerm_resource_group.rg": tests.AsMap(t, `{
			"location": "centralus",
			"name":     "MyTestResourceGroup"
		}`),
		"azurerm_network_security_group.nsg": tests.AsMap(t, `{
			"location":            "centralus",
			"name":                "MyTestResourceNSG",
			"resource_group_name": "MyTestResourceGroup"
		}`),
		"azurerm_virtual_network.vnet": tests.AsMap(t, `{
		}`),
	}

	// This is the number of expected Terraform resources being provisioned.
	//
	// Note: There may be more Terraform resources provisioned than Azure resources provisioned!
	expectedTerraformResourceCount := 3

	testFixture := unit.UnitTestFixture{
		GoTest:                          t,
		TfOptions:                       tests.TfOptions,
		ExpectedResourceCount:           expectedTerraformResourceCount,
		ExpectedResourceAttributeValues: resourceDescription,
	}

	unit.RunUnitTests(&testFixture)
}
