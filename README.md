# Matrix Discovery

A little go web server to serve the URIs `/.well-known/matrix/server` and `/.well-known/matrix/client`.

## Configuration

This configuration can be set in the YAML files for Kubernetes or Swarm.

| Env var                            | Description                                                                      | Default             |
| ---------------------------------- | -------------------------------------------------------------------------------- | ------------------- |
| DISCOVERY_CORS_ALLOW_ORIGINS       |                                                                                  | `*`                 |
| DISCOVERY_CORS_ALLOW_METHODS       |                                                                                  | `GET`               |
| DISCOVERY_MATRIX_H_SERVER          | The server name to delegate server-server communciations to, with optional port. | ``                  |
| DISCOVERY_MATRIX_M_HOMESERVER      | The base URL for the homeserver for client-server connections.                   | ``                  |
| DISCOVERY_MATRIX_M_IDENTITY_SERVER | The base URL for the identity server for client-server connections.              | `https://vector.im` |
