module samples

go 1.12

replace github.com/microsoft/terratest-abstraction => ../

require (
	github.com/Azure/azure-sdk-for-go v38.1.0+incompatible
	github.com/Azure/go-autorest/autorest v0.9.3
	github.com/Azure/go-autorest/autorest/adal v0.8.1

	github.com/gruntwork-io/terratest v0.26.5
	github.com/microsoft/terratest-abstraction v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.5.1
)
