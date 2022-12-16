# Interfaces

## Check that type must implement interface

To do that add global variable like this

var _ json.Marshaler = (*RawMessage)(nil)

## Links

- https://go.dev/doc/effective_go#interfaces
- https://go.dev/doc/effective_go#embedding