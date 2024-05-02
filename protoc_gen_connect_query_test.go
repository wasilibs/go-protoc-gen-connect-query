package protoc_gen_connect_es

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestBuf(t *testing.T) {
	if err := os.RemoveAll(filepath.Join("out", "buf")); err != nil {
		t.Fatalf("failed to remove build directory: %v", err)
	}

	output := bytes.Buffer{}
	cmd := exec.Command("go", "run", "github.com/bufbuild/buf/cmd/buf@v1.30.0", "generate")
	cmd.Stderr = &output
	cmd.Stdout = &output
	cmd.Dir = "testdata"
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to run buf: %v\n%v", err, output.String())
	}

	for _, path := range []string{
		filepath.Join("out", "buf", "es", "helloworld-Greeter_connectquery.d.ts"),
		filepath.Join("out", "buf", "es", "helloworld-Greeter_connectquery.js"),
		filepath.Join("out", "buf", "ts", "helloworld-Greeter_connectquery.ts"),
	} {
		if _, err := os.Stat(path); err != nil {
			t.Errorf("failed to stat %v: %v", path, err)
		}
	}
}
