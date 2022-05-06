package unit

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

// the workspace cannot be deleted in the case that the "before" workspace is the same
// as the one used for the test. If this test passes, then the behavior is correct
func TestSupportForUnitTestWithNoWorkspaceChange(t *testing.T) {
	tfOptions := &terraform.Options{
		TerraformDir: test_structure.CopyTerraformFolderToTemp(t, "testing-tf/", "."),
		Upgrade:      true,
	}

	testFixture := UnitTestFixture{
		GoTest:                          t,
		TfOptions:                       tfOptions,
		ExpectedResourceCount:           3,
		ExpectedResourceAttributeValues: nil,
		Workspace:                       "default",
	}

	RunUnitTests(&testFixture)
}

func TestSupportForJSONRawResourceAttribute(t *testing.T) {
	tfOptions := &terraform.Options{
		TerraformDir: test_structure.CopyTerraformFolderToTemp(t, "testing-tf/", "."),
	}

	testFixture := UnitTestFixture{
		GoTest:                t,
		TfOptions:             tfOptions,
		ExpectedResourceCount: 3,
		ExpectedRawResourceAttributes: []RawResourceAttribute{
			{
				ResourceAddress:   "local_file.json",
				ResourceAttribute: "content",
				Value: jsonToMap(t, `{
					"length": 16,
					"nested": {
						"value": 16
					}
				}`),
			},
		},
	}

	RunUnitTests(&testFixture)
}

func TestSupportForYAMLRawResourceAttribute(t *testing.T) {
	tfOptions := &terraform.Options{
		TerraformDir: test_structure.CopyTerraformFolderToTemp(t, "testing-tf/", "."),
	}

	testFixture := UnitTestFixture{
		GoTest:                t,
		TfOptions:             tfOptions,
		ExpectedResourceCount: 3,
		ExpectedRawResourceAttributes: []RawResourceAttribute{
			{
				ResourceAddress:       "local_file.yaml",
				ResourceAttribute:     "content",
				ResourceAttributeType: YAML,
				Value: jsonToMap(t, `{
					"length": 16,
					"nested": {
						"value": 16
					}
				}`),
			},
		},
	}

	RunUnitTests(&testFixture)
}

func TestTerraformCommandStdout(t *testing.T) {
	tfOptions := &terraform.Options{
		TerraformDir: test_structure.CopyTerraformFolderToTemp(t, "testing-tf/", "."),
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
