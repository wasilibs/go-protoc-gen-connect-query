import "./textcoding.js";

import { protocGenConnectQuery } from "@connectrpc/protoc-gen-connect-query/dist/cjs/src/protoc-gen-connect-query-plugin.js";

import { runQuickJs  } from "./run-quickjs.js";

runQuickJs(protocGenConnectQuery);
