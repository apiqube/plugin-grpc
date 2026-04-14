package main

func info() PluginInfo {
	return PluginInfo{
		Name:        "grpc",
		Version:     "0.1.0",
		Description: "gRPC executor — unary, server/client/bidi streaming.",
		Protocols:   []string{"grpc", "grpcs"},
		Fields: map[string]FieldSpec{
			"call": {
				Type:        "string",
				Required:    true,
				Description: "Fully-qualified method: package.Service/Method",
			},
			"payload": {
				Type:        "any",
				Required:    false,
				Description: "Request message (unary or client stream first message).",
			},
			"metadata": {
				Type:        "map",
				Required:    false,
				Description: "gRPC metadata (request headers).",
			},
			"proto": {
				Type:        "string",
				Required:    false,
				Description: "Path to .proto file for schema; otherwise uses server reflection.",
			},
			"stream": {
				Type:        "string",
				Required:    false,
				Description: "Stream type: unary (default), server, client, bidi.",
			},
			"messages": {
				Type:        "array",
				Required:    false,
				Description: "For client streaming — sequence of messages to send.",
			},
			"exchange": {
				Type:        "array",
				Required:    false,
				Description: "For bidi streaming — sequence of {send: ..., receive: ...} steps.",
			},
		},
	}
}
