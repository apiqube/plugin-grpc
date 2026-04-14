package main

// info returns the plugin metadata: name, version, manifest fields,
// and event schemas. This is everything the engine needs to know about
// this plugin without loading or running it.
//
// The engine calls plugin_info() once at load time; the result drives:
//   - manifest field validation (users can only write fields we declare)
//   - frontend introspection (what events can this plugin emit?)
//   - typed event dispatch via Dispatcher.SubscribePluginTyped
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

		// Events this plugin can emit at runtime.
		// Frontends subscribe via Dispatcher using the fully-qualified name
		// "grpc.<Kind>" — for example "grpc.StreamMessageReceived".
		Events: map[string]EventSpec{
			"StreamMessageReceived": {
				Description: "Emitted each time a server-streaming or bidi call receives a message.",
				Fields: map[string]FieldSpec{
					"stream_id": {
						Type:        "string",
						Required:    true,
						Description: "Unique identifier of the stream this message belongs to.",
					},
					"sequence": {
						Type:        "number",
						Required:    true,
						Description: "1-based index of this message within the stream.",
					},
					"bytes": {
						Type:        "number",
						Required:    true,
						Description: "Wire size of the message payload in bytes.",
					},
				},
			},

			"StreamOpened": {
				Description: "Emitted when a streaming call is successfully established.",
				Fields: map[string]FieldSpec{
					"stream_id": {
						Type:        "string",
						Required:    true,
						Description: "Unique identifier assigned to the new stream.",
					},
					"type": {
						Type:        "string",
						Required:    true,
						Description: "Stream type: server, client, or bidi.",
					},
					"method": {
						Type:        "string",
						Required:    true,
						Description: "Fully-qualified gRPC method being called.",
					},
				},
			},

			"StreamClosed": {
				Description: "Emitted when a streaming call completes or is cancelled.",
				Fields: map[string]FieldSpec{
					"stream_id": {
						Type:        "string",
						Required:    true,
						Description: "Unique identifier of the stream that closed.",
					},
					"status": {
						Type:        "string",
						Required:    true,
						Description: "gRPC status code at close: ok, cancelled, unavailable, etc.",
					},
					"message_count": {
						Type:        "number",
						Required:    true,
						Description: "Total messages exchanged before close.",
					},
					"duration_ms": {
						Type:        "number",
						Required:    true,
						Description: "Total stream lifetime in milliseconds.",
					},
				},
			},
		},
	}
}
