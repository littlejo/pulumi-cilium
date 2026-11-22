// Package examples deploys the example programs on the cluster of the current Kubernetes context.
//
// The cluster must not have a CNI installed. The tests use the provider binary built in ../bin and the
// SDKs built in ../sdk, so run `make provider` and the `make build_<language>` targets first.
package examples

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/assertpreview"
	"github.com/pulumi/providertest/pulumitest/opttest"
)

// defaultCiliumVersion is the version installed when CILIUM_VERSION is not set.
// renovate: datasource=github-releases depName=cilium/cilium
const defaultCiliumVersion = "1.20.5"

func TestMain(m *testing.M) {
	// The programs run in place, so they can reference the SDKs with relative paths, and the provider is
	// picked up from PATH: attaching it with opttest.LocalProviderPath would rewrite their Pulumi.yaml.
	// `pulumi install` would also try to download it, so the tests install the dependencies only.
	bin, err := filepath.Abs(filepath.Join("..", "bin"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if _, err := os.Stat(filepath.Join(bin, "pulumi-resource-cilium")); err != nil {
		fmt.Fprintln(os.Stderr, "provider binary not found, run `make provider` first:", err)
		os.Exit(1)
	}
	if err := os.Setenv("PATH", bin+string(os.PathListSeparator)+os.Getenv("PATH")); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func TestYaml(t *testing.T) {
	testExample(t, "yaml")
}

func TestNodejs(t *testing.T) {
	testExample(t, "nodejs", opttest.RequireYarnLinks(false))
}

func TestPython(t *testing.T) {
	testExample(t, "python")
}

func TestGo(t *testing.T) {
	testExample(t, "go")
}

// testExample installs Cilium with the example program in dir, checks that a second run has nothing to
// change, then destroys the stack.
func testExample(t *testing.T, dir string, opts ...opttest.Option) {
	version := os.Getenv("CILIUM_VERSION")
	if version == "" {
		version = defaultCiliumVersion
	}

	test := pulumitest.NewPulumiTest(t, dir, append([]opttest.Option{opttest.TestInPlace(), opttest.SkipInstall()}, opts...)...)
	install := exec.Command("pulumi", "install", "--no-plugins")
	install.Dir = test.WorkingDir()
	if out, err := install.CombinedOutput(); err != nil {
		t.Fatalf("pulumi install: %v\n%s", err, out)
	}
	test.SetConfig(t, "ciliumVersion", version)

	up := test.Up(t)
	if got := up.Outputs["version"].Value; got != version {
		t.Errorf("installed Cilium version is %v, want %s", got, version)
	}

	assertpreview.HasNoChanges(t, test.Preview(t))
}
