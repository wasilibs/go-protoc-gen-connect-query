package main

import (
	"github.com/goyek/x/boot"
	"github.com/wasilibs/tools/tasks"
)

func main() {
	tasks.Define(tasks.Params{
		LibraryName: "protoc-gen-connect-query",
		LibraryRepo: "connectrpc/connect-query",
		GoReleaser:  true,
	})
	boot.Main()
}
