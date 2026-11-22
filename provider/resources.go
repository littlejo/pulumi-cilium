// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cilium maps the Cilium Terraform provider to a Pulumi provider.
package cilium

import (
	_ "embed"
	"path"

	"github.com/littlejo/pulumi-cilium/provider/pkg/version"
	cilium "github.com/littlejo/terraform-provider-cilium/cilium"
	pf "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
)

//go:embed cmd/pulumi-resource-cilium/bridge-metadata.json
var bridgeMetadata []byte

const (
	pkgName = "cilium"
	mainMod = "index"
)

// Provider returns the mapping between the Terraform provider and the Pulumi package.
func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:                 pf.ShimProvider(cilium.New(version.Version)()),
		Name:              pkgName,
		DisplayName:       "Cilium",
		Publisher:         "littlejo",
		LogoURL:           "https://raw.githubusercontent.com/littlejo/pulumi-cilium/main/docs/cilium.png",
		PluginDownloadURL: "github://api.github.com/littlejo/pulumi-cilium",
		Description:       "A Pulumi package for creating and managing Cilium resources",
		Keywords:          []string{"pulumi", "cilium", "category/network"},
		License:           "Apache-2.0",
		Homepage:          "https://github.com/littlejo/pulumi-cilium",
		Repository:        "https://github.com/littlejo/pulumi-cilium",
		Version:           version.Version,
		GitHubOrg:         "littlejo",
		MetadataInfo:      tfbridge.NewProviderMetadata(bridgeMetadata),
		Resources: map[string]*tfbridge.ResourceInfo{
			// The install resource is named "cilium" upstream, without the prefix the other tokens are derived from.
			"cilium": {Tok: tfbridge.MakeResource(pkgName, mainMod, "Install")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName:          "@littlejo/cilium",
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			PackageName:          "littlejo_cilium",
			PythonRequires:       ">=3.10",
			RespectSchemaVersion: true,
			PyProject:            struct{ Enabled bool }{true},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				"github.com/littlejo/pulumi-cilium/sdk",
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				pkgName,
			),
			GenerateResourceContainerTypes: true,
		},
	}

	// Map every other resource and data source from its Terraform name: "cilium_kubeproxy_free" becomes
	// "cilium:index/kubeproxyFree:KubeproxyFree" and "cilium_helm_values" becomes "cilium:index/getHelmValues:getHelmValues".
	prov.MustComputeTokens(tokens.SingleModule(pkgName+"_", mainMod, tokens.MakeStandard(pkgName)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
