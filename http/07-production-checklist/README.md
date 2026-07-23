# Production HTTP: Final Revision Checklist

Use this as the last pass before an interview or before exposing a Go HTTP service.

## Request lifecycle

```text
route → middleware → bounded/validated input → service with r.Context()
      → consistent response → logging/metrics → graceful shutdown
```

## Routing and handlers

- Prefer `http.NewServeMux()` over `http.DefaultServeMux`.
- Use modern patterns: `mux.HandleFunc("GET /users/{id}", handler)` and read `r.PathValue("id")`.
- Let resource paths identify nouns; use query parameters for filters, pagination, and representation options.
- Return `405 Method Not Allowed` (with `Allow`) when a path exists but the method is unsupported.
- Handlers run concurrently. Protect any shared mutable state and run `go test -race`.

## Input and context

- Use `r.Context()` for database calls, outbound HTTP calls, and any work that should stop on disconnect/shutdown.
- Bound request bodies with `http.MaxBytesReader`.
- Decode into a request struct, call `DisallowUnknownFields` when strictness fits the API, and reject trailing JSON values.
- Separate parsing (`400`) from semantic validation (`422`) and state conflicts (`409`).
- Never trust identifiers, headers, query parameters, or JSON fields without validation.

## Responses

- Set headers before `WriteHeader` or `Write`; a body write otherwise commits `200 OK`.
- Return a consistent JSON error contract, for example `{"error":"message"}`.
- Use `201 Created` for new resources and set `Location` when the resource URL is known.
- Use no body for `204 No Content`.
- `401` means authentication failed; `403` means the authenticated principal lacks permission.

## Middleware and security

- Make ordering intentional: recovery typically outermost; logging should observe the result; auth must short-circuit protected routes.
- Parse `Authorization: Bearer <token>` rather than accepting a non-empty header.
- Handle CORS `OPTIONS` preflight. Never combine `Access-Control-Allow-Origin: *` with credentials.
- Add rate limits, request IDs, and structured logs as appropriate for public services.
- Recovery handles unexpected panics but cannot repair a response already started.

## Server lifecycle

- Configure `ReadHeaderTimeout`, and sensible read, write, and idle timeouts.
- Treat `http.ErrServerClosed` as expected during shutdown.
- On `SIGINT`/`SIGTERM`, call `Server.Shutdown` with a bounded context to drain in-flight requests.
- Give outbound HTTP clients a timeout too; `http.DefaultClient` has no deadline.

## Testing

- Test handlers and middleware with `httptest.NewRequest` and `httptest.NewRecorder`.
- Test failure paths: invalid JSON, missing auth, validation, body too large, method mismatch, cancellation.
- Run `go test ./...` and `go test -race ./...` for shared state.

## Runnable example

[`server.go`](server.go) provides `POST /users` and `GET /healthz` using only the standard library. It demonstrates explicit routing, bounded strict JSON decoding, validation, response helpers, timeouts, and graceful shutdown.

```bash
go test ./http/07-production-checklist
go run ./http/07-production-checklist

curl -i -X POST http://localhost:8080/users \
  -H 'Content-Type: application/json' \
  -d '{"name":"Arpit","email":"arpit@example.com"}'
```
