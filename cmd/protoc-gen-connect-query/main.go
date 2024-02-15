package main

import (
	"os"

	"github.com/wasilibs/go-protoc-gen-connect-query/internal/runner"
	"github.com/wasilibs/go-protoc-gen-connect-query/internal/wasm"
)

func main() {
	os.Exit(runner.Run("protoc-gen-connect-query", os.Args[1:], wasm.ProtocGenConnectQuery, os.Stdin, os.Stdout, os.Stderr, "."))
}
