# Contributing

Please read and observe the [Code of Conduct](CODE-OF-CONDUCT.md).

This provider is generated from the [`cilium` Terraform provider](https://github.com/littlejo/terraform-provider-cilium)
with the [Pulumi Terraform bridge](https://github.com/pulumi/pulumi-terraform-bridge). Bugs and features of the
resources themselves belong to the Terraform provider; this repository maps it to Pulumi and publishes the SDKs.

## Development environment

The [dev container](.devcontainer/devcontainer.json) has everything installed. Otherwise you need:

- Go, at the version of [`provider/go.mod`](provider/go.mod)
- [Pulumi](https://www.pulumi.com/docs/install/)
- Node.js and Python 3 with the [`build`](https://pypi.org/project/build/) package, for the SDKs
- [golangci-lint](https://golangci-lint.run/) v2
- [kind](https://kind.sigs.k8s.io/) and Docker, for the examples

Linking the provider needs about 8 GB of memory.

## Building

`make help` lists the targets. The main ones:

| Target              | Does                                                                  |
|---------------------|-----------------------------------------------------------------------|
| `make schema`       | Regenerates `provider/cmd/pulumi-resource-cilium/schema.json`          |
| `make provider`     | Builds `bin/pulumi-resource-cilium`                                    |
| `make build_sdks`   | Regenerates and builds the Node.js, Python and Go SDKs in `sdk/`       |
| `make lint`         | Lints the provider code                                                |
| `make test_provider`| Runs the unit tests of the provider                                    |
| `make test`         | Deploys the examples on the current Kubernetes context                 |

The schema, the bridge metadata and the SDKs are committed. CI fails when they are not up to date, so after a
change of the provider or of its dependencies, run:

```bash
make schema build_sdks
```

and commit the result. Local builds use the version `0.0.0-alpha.0+dev`, which keeps the generated code identical
from one build to the next; the release workflow builds with the version of the tag.

## Updating the Terraform provider

Renovate opens a pull request when a new version of the Terraform provider is released. Check out its branch, run
`make schema build_sdks` and push the regenerated code.

A new resource or data source of the Terraform provider is mapped automatically from its name (see
[`provider/resources.go`](provider/resources.go)). `make test_provider` then fails until it is added to the tokens
pinned in [`provider/resources_test.go`](provider/resources_test.go): renaming a token breaks the programs that use
it.

## Testing the examples

The tests in [`examples`](examples) install Cilium with each example program on the cluster of the current
Kubernetes context, which must not have a CNI:

```bash
kind create cluster --config .github/config/acceptance_tests_kind_config.yaml
make provider build_sdks
cd examples && go test -v -run TestPython .
```

`CILIUM_VERSION` sets the version of Cilium to install.

## Releasing

Push a tag `vX.Y.Z` on `main`. The release workflow:

1. publishes the provider binaries in a GitHub release, with a signed checksum file and build provenance attestations;
2. tags the Go SDK `sdk/vX.Y.Z`;
3. publishes the Python SDK on PyPI and the Node.js SDK on npm, with trusted publishing.

Tags with a pre-release suffix, like `v1.2.0-rc1`, create a GitHub pre-release and are published on npm with the
`next` tag.
