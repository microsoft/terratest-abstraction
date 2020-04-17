package integrationtest

import (
	"context"
	"fmt"
	"os"
	"testing"

	"samples/azure/tests"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/network/mgmt/network"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"

	"github.com/microsoft/terratest-abstraction/integration"
	"github.com/stretchr/testify/require"
)

// ServicePrincipalAuthorizer - Configures an authorizer for the Azure SDK that can use the same
// environment variables as was used for the terraform deployment
func ServicePrincipalAuthorizer(t *testing.T) autorest.Authorizer {
	oauthConfig, err := adal.NewOAuthConfig(azure.PublicCloud.ActiveDirectoryEndpoint, os.Getenv("ARM_TENANT_ID"))
	require.Nil(t, err)
	token, err := adal.NewServicePrincipalToken(*oauthConfig, os.Getenv("ARM_CLIENT_ID"), os.Getenv("ARM_CLIENT_SECRET"), azure.PublicCloud.ResourceManagerEndpoint)
	require.Nil(t, err)
	return autorest.NewBearerAuthorizer(token)
}

func getVNETSubnets(t *testing.T, resourceGroupName, vnetName string) []string {
	subnetsClient := network.NewSubnetsClient(os.Getenv("ARM_SUBSCRIPTION_ID"))
	subnetsClient.Authorizer = ServicePrincipalAuthorizer(t)

	subnets, err := subnetsClient.List(context.Background(), resourceGroupName, vnetName)
	require.Nil(t, err)

	subnetNames := []string{}
	for subnets.NotDone() {
		subnetPageResults := subnets.Values()
		for _, result := range subnetPageResults {
			subnetNames = append(subnetNames, *result.Name)
		}
		require.Nil(t, subnets.Next())
	}

	return subnetNames
}

// verifies that the VNET is configured with the correct number of subnets
func testVNET(t *testing.T, output integration.TerraformOutput) {
	vnetName, ok := output["vnet_name"].(string)
	require.True(t, ok, "vnet_name was unexpectedly not a string value")

	resourceGroupName, ok := output["resource_group_name"].(string)
	require.True(t, ok, "resource_group_name was unexpectedly not a string value")

	// this operation will query Azure to identify all of the subnets within the named VNET
	subnetIDs := getVNETSubnets(t, resourceGroupName, vnetName)

	// assert the correct number of subnets have been provisioned
	require.Equal(t, 2, len(subnetIDs), fmt.Sprintf("Expected 2 subnets but found %v", len(subnetIDs)))
}

func TestTemplateIntegration(t *testing.T) {
	// This is the number of expected Terraform outputs that should be present after the `apply` command.
	expectedTerraformOutputCount := 2

	// This is a list of testing functions that should be invoked. Each of them will need to have the following interface:
	//	func(goTest *testing.T, output integration.TerraformOutput)
	//
	// It is expected that these functions make assertions about the data inside the Terraform output.
	//
	// Note: typical deployments will have many more tests that validate different components of the deployment. In this
	// example, we just have one test to run.
	outputValidations := []integration.TerraformOutputValidation{
		testVNET,
	}

	testFixture := integration.IntegrationTestFixture{
		GoTest:                t,
		TfOptions:             tests.TfOptions,
		ExpectedTfOutputCount: expectedTerraformOutputCount,
		TfOutputAssertions:    outputValidations,
	}
	integration.RunIntegrationTests(&testFixture)
}
