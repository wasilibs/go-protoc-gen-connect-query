# go-protoc-gen-connect-query

go-protoc-gen-connect-query is a distribution of [protoc-gen-connect-query][1] that can be built with Go. 
It does not actually reimplement any functionality of connect-query in Go, instead packaging it with the
lightweight JS runtime [QuickJS][6], and executing with the pure Go Wasm runtime [wazero][2].
This means that `go install` or `go run` can be used to execute it, with no need to rely on external
package managers such as NPM, on any platform that Go supports.

## Installation

Precompiled binaries are available in the [releases](https://github.com/wasilibs/go-protoc-gen-connect-query/releases).
Alternatively, install the plugin using `go install`.

```bash
$ go install github.com/wasilibs/go-protoc-gen-connect-query/cmd/protoc-gen-connect-query@latest
```

As long as `$GOPATH/bin`, e.g. `~/go/bin` is on the `PATH`, you can use it with protoc as normal.

```bash
$ protoc --es_out=out/es -Iprotos protos/helloworld.proto
```

For [buf][3] users, to avoid installation entirely, it can be convenient to use `go run` in `buf.gen.yaml`.

```yaml
version: v1
plugins:
  - plugin: connect-query
    out: out/es
    path: ["go", "run", "github.com/wasilibs/go-protoc-gen-connect-query/cmd/protoc-gen-connect-query@latest"]
```

This makes it possible to have full protobuf/gRPC generation with no installation of tools,
besides Go itself, on any platform that Go supports. The above examples use `@latest`, but it is
recommended to specify a version, in which case all of the developers on your codebase will use the
same version of the tool with no special steps.

See a full [example][5] in `go-protoc-gen-builtins`. To generate protos, enter the directory and run
`go run github.com/bufbuild/buf/cmd/buf@v1.30.0 generate`. As long as your machine has Go installed,
you will be able to generate protos. The first time using `go run` for a command, Go automatically builds
it making it slower, but subsequent invocations should be quite fast.

[1]: https://github.com/connectrpc/connect-query
[2]: https://wazero.io/
[3]: https://buf.build/
[4]: https://github.com/wasilibs/go-protoc-gen-builtins
[5]: https://github.com/wasilibs/go-protoc-gen-builtins/tree/main/example
[6]: https://bellard.org/quickjs/
