package unittest

import (
	"samples/azure/tests"
	"testing"

	"github.com/hashicorp/terraform-json"
	"github.com/microsoft/terratest-abstraction/unit"
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
			"resource_group_name": "MyTestResourceGroup",
			"name":                "virtualNetwork1",
			"tags": {
				"environment": "production"
			},
			"address_space": [
				"10.0.0.0/16"
			],
			"dns_servers": [
				"10.0.0.4",
				"10.0.0.5"
			],
			"subnet": [
				{
					"address_prefix": "10.0.1.0/24",
					"name":           "MyTestSubnet1"
				},{
					"address_prefix": "10.0.3.0/24",
					"name":           "MyTestSubnet2"
				}
			]
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

func TestPlanOutputs(t *testing.T) {

	testFixture := unit.UnitTestFixture{
		GoTest:                t,
		TfOptions:             tests.TfOptions,
		ExpectedResourceCount: 3,
		PlanAssertions: []unit.TerraformPlanValidation{
			func(t *testing.T, plan tfjson.Plan) {
				change := plan.OutputChanges["resource_group_name"]
				if change.After != "MyTestResourceGroup" {
					t.Fatal("output validation failed")
				}
			},
		},
	}

	unit.RunUnitTests(&testFixture)
}
