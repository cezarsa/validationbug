package v1

import (
	"testing"

	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

func TestFunc(t *testing.T) {
	testEnv := &envtest.Environment{
		CRDDirectoryPaths: []string{"crd"},
	}
	_, err := testEnv.Start()
	defer testEnv.Stop()
	if err != nil {
		t.Fatal(err)
	}
}
