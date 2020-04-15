# hello-world

The code in here is used to show off some simple use cases of the test harness.

> This assumes that your environment is configured with the values found in [`.envrc.template`](./.envrc.template)

## Run the unit tests

```bash
go test -v $(go list ./... | grep unit)
```

## Run the deployment

```bash
terraform apply
```

## Run the integration tests

```bash
go test -v $(go list ./... | grep integration)
```

## Cleanup

```bash
terraform destroy
```
