package tests

import (
	"encoding/json"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// AsMap pars block of JSON into a Go Map. Fails the test if the JSON is invalid.
func AsMap(t *testing.T, jsonString string) map[string]interface{} {
	var theMap map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &theMap); err != nil {
		t.Fatal(err)
	}
	return theMap
}

// TfOptions TF options that can be used by all tests
var TfOptions = &terraform.Options{
	TerraformDir: "../../",
	Upgrade:      true,
	Vars:         map[string]interface{}{},
}
