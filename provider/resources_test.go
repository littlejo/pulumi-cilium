package cilium

import (
	"testing"

	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
)

// Renaming a token breaks the programs and the stacks that use it, so the tokens are pinned here.
var expectedResources = map[string]string{
	"cilium":                        "cilium:index/install:Install",
	"cilium_clustermesh":            "cilium:index/clustermesh:Clustermesh",
	"cilium_clustermesh_connection": "cilium:index/clustermeshConnection:ClustermeshConnection",
	"cilium_config":                 "cilium:index/config:Config",
	"cilium_hubble":                 "cilium:index/hubble:Hubble",
	"cilium_kubeproxy_free":         "cilium:index/kubeproxyFree:KubeproxyFree",
}

var expectedDataSources = map[string]string{
	"cilium_helm_values": "cilium:index/getHelmValues:getHelmValues",
}

func TestResourceTokens(t *testing.T) {
	info := Provider()

	info.P.ResourcesMap().Range(func(name string, _ shim.Resource) bool {
		if _, ok := expectedResources[name]; !ok {
			t.Errorf("new Terraform resource %q: add its token to expectedResources", name)
		}
		return true
	})
	for name, tok := range expectedResources {
		res, ok := info.Resources[name]
		if !ok {
			t.Errorf("Terraform resource %q is not mapped", name)
			continue
		}
		if string(res.Tok) != tok {
			t.Errorf("Terraform resource %q is mapped to %q, want %q", name, res.Tok, tok)
		}
	}
}

func TestDataSourceTokens(t *testing.T) {
	info := Provider()

	info.P.DataSourcesMap().Range(func(name string, _ shim.Resource) bool {
		if _, ok := expectedDataSources[name]; !ok {
			t.Errorf("new Terraform data source %q: add its token to expectedDataSources", name)
		}
		return true
	})
	for name, tok := range expectedDataSources {
		ds, ok := info.DataSources[name]
		if !ok {
			t.Errorf("Terraform data source %q is not mapped", name)
			continue
		}
		if string(ds.Tok) != tok {
			t.Errorf("Terraform data source %q is mapped to %q, want %q", name, ds.Tok, tok)
		}
	}
}
