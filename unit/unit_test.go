package unit

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// the workspace cannot be deleted in the case that the "before" workspace is the same
// as the one used for the test
func TestSupportForUnitTestWitNoWorkspaceChange(t *testing.T) {
	tfOptions := &terraform.Options{
		TerraformDir: "testing-tf/",
		Upgrade:      true,
		Vars:         map[string]interface{}{},
	}

	testFixture := UnitTestFixture{
		GoTest:                          t,
		TfOptions:                       tfOptions,
		ExpectedResourceCount:           1,
		ExpectedResourceAttributeValues: nil,
		Workspace:                       "default",
	}

	RunUnitTests(&testFixture)
}
