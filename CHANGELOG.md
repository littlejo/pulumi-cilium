# Changelog

## Unreleased

### Breaking changes

- The provider is built from version 0.4.0 of the `cilium` Terraform provider, which supports Cilium 1.17 and newer
  and installs Cilium 1.20 by default. See its
  [changelog](https://github.com/littlejo/terraform-provider-cilium/blob/main/CHANGELOG.md).
- `getHelmValues` no longer returns an `id`: the Terraform data source has no such attribute.
- The Python SDK requires Python 3.10 or newer.
- The C# and Java SDKs, which were never published, are no longer generated.

### Features

- `Install` has a new `resetThenReuse` property; `reusethenreuse` is deprecated.
- The Node.js SDK is published on npm as `@littlejo/cilium`.
- The Go SDK is tagged `sdk/vX.Y.Z`, so `go get github.com/littlejo/pulumi-cilium/sdk@vX.Y.Z` works.
- The README published on PyPI and npm describes this provider.
- The release checksums are signed with cosign, and the archives have build provenance attestations.

### Internal

- Pulumi Terraform bridge v3.143.0: the deprecated `pf` module is replaced by `pulumi-terraform-bridge/v3/pkg/pf`.
- Go 1.27, as required by the Terraform provider.
- The tokens of new Terraform resources are computed automatically.
- The Python SDK uses `pyproject.toml` instead of `setup.py`.
- CI lints the provider, checks the generated code is committed and deploys the examples of each language on kind,
  with several Cilium versions.
- PyPI and npm publication use trusted publishing.
