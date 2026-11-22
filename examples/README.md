# Examples

Each directory installs Cilium with one Pulumi language. The programs use the provider and the SDKs built from this
repository: the Node.js, Python and Go programs reference `sdk/` with relative paths. In your own programs, install
the published packages instead (see the [README](../README.md#installing)).

| Directory          | Language   |
|--------------------|------------|
| [`yaml`](yaml)     | YAML       |
| [`nodejs`](nodejs) | TypeScript |
| [`python`](python) | Python     |
| [`go`](go)         | Go         |

`examples_test.go` deploys them on the cluster of the current Kubernetes context: see
[CONTRIBUTING.md](../CONTRIBUTING.md#testing-the-examples).
