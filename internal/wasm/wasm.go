package wasm

import _ "embed"

//go:embed protoc-gen-connect-query.wasm
var ProtocGenConnectQuery []byte
