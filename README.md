# plugin-grpc

> gRPC executor plugin for [ApiQube](https://github.com/apiqube).

[![License](https://img.shields.io/badge/License-MIT-blue?style=flat-square)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Scaffold-yellow?style=flat-square)]()

Handles tests targeting `grpc://` or `grpcs://` URLs. Supports all four stream types:
unary, server-streaming, client-streaming, and bidirectional.

## Protocols

`grpc`, `grpcs`

## Manifest Fields

| Field      | Type   | Required | Description |
|------------|--------|----------|-------------|
| `call`     | string | yes      | `package.Service/Method` |
| `payload`  | any    | no       | Request message |
| `metadata` | map    | no       | gRPC metadata headers |
| `proto`    | string | no       | Path to `.proto` file (else reflection) |
| `stream`   | string | no       | `unary` / `server` / `client` / `bidi` |
| `messages` | array  | no       | For client streaming |
| `exchange` | array  | no       | For bidi streaming |

## Example

```yaml
target: grpc://localhost:9090

tests:
  - name: Get user
    call: users.UserService/GetUser
    payload: { id: 1 }
    expect:
      status: ok
```

## Build

```bash
tinygo build -o plugin-grpc.wasm -target=wasi ./
```

## License

[MIT](LICENSE)
