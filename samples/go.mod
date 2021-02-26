module samples

go 1.12

replace github.com/microsoft/terratest-abstraction => ../

require (
	github.com/Azure/azure-sdk-for-go v46.0.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.5
	github.com/Azure/go-autorest/autorest/adal v0.9.2
	github.com/gruntwork-io/terratest v0.32.8
	github.com/hashicorp/terraform-json v0.8.0
	github.com/microsoft/terratest-abstraction v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.5.1
)
