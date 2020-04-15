# Terratest Abstraction

This Go package offers abstractions over the popular [Terratest](https://github.com/gruntwork-io/terratest) library in order to abstract some common testing patterns that have been identified through multiple projects that deploying [Infrastructure as Code](https://docs.microsoft.com/en-us/azure/devops/learn/what-is-infrastructure-as-code) (*IAC*) using [Terraform](https://www.terraform.io/). The abstractions offered here can be used along side existing Terratest code and are quite easy to drop into existing projects. Feedback and OSS contributions are welcome!

## When to use?

Teams that rely on automated deployments of any kind demand robust automated validation of those environments in order to have confidence in changes. The results of these automated changes will allow software and operations engineers to sleep well at night, knowing that large classes of defects will be caught by repeatable and automated checks against their runtime systems. These concepts apply equally to software and infrastructure deployments.

`TerratestAbstraction` offers an intuitive interface to testing Terraform deployments that offers a more declarative approach to writing test cases. There is support for testing the [Terraform Plan](https://www.terraform.io/docs/commands/plan.html) as a *pre-deployment unit test* and [Terraform Output](https://www.terraform.io/docs/commands/output.html) as a *post-deployment integration test*.


## Getting started

**Importing the library**

The package can be imported (like any other go package) by adding an import statement to your test file:

```go
```

**Writing unit tests**

**Writing integration tests**

**Automating in CICD pipelines**


# Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
