package unit

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// the workspace cannot be deleted in the case that the "before" workspace is the same
// as the one used for the test. If this test passes, then the behavior is correct
func TestSupportForUnitTestWithNoWorkspaceChange(t *testing.T) {
	tfOptions := &terraform.Options{
		TerraformDir: "testing-tf/",
		Upgrade:      true,
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

func TestTerraformCommandStdout(t *testing.T) {
	tfOptions := &terraform.Options{
		TerraformDir: "testing-tf/",
		Upgrade:      true,
		Vars: map[string]interface{}{
			"length": 15,
		},
	}

	testFixture := UnitTestFixture{
		GoTest:    t,
		TfOptions: tfOptions,
		CommandStdoutAssertions: []TerraformCommandStdoutValidation{
			func(t *testing.T, output string, err error) {
				if err == nil {
					t.Fatal("should err")
				}
				if !strings.Contains(output, "Invalid value for variable") {
					t.Fatal("should contain 'Invalid value for variable'")
				}
			},
		},
	}

	RunUnitTests(&testFixture)
}
