# Production HTTP Checklist

Use this as a last-minute checklist when discussing or building a Go HTTP service.

## Routing and handlers

- Prefer an explicit `http.NewServeMux()` over global `http.DefaultServeMux`.
- Modern `net/http` supports method and path patterns: `mux.HandleFunc("GET /users/{id}", handler)`.
- Read a path parameter with `r.PathValue("id")`.
- Use `405 Method Not Allowed` for a known path with an unsupported method.

## Request handling

- Read request cancellation and deadlines from `r.Context()`; pass that context to database and outbound calls.
- Bound request bodies using `http.MaxBytesReader`.
- Decode into a struct, use `Decoder.DisallowUnknownFields()`, and reject multiple JSON values.
- Validate decoded input before using it. Use `400 Bad Request` for malformed JSON and `422 Unprocessable Content` for syntactically valid but invalid input.

## Responses

- Set headers before calling `WriteHeader` or writing a body.
- A body write implicitly sends `200 OK` if no status has been written.
- Use a consistent JSON error shape such as `{"error":"message"}`.
- Use `201 Created` for a new resource; add a `Location` header when the resource URL is known.
- `401 Unauthorized` means authentication is missing or invalid; `403 Forbidden` means authentication succeeded but permission is denied.

## Middleware and security

- Keep middleware order intentional: recovery should normally wrap the rest; logging should observe the final response; authentication should short-circuit before protected handlers.
- Parse `Authorization: Bearer <token>` rather than accepting any non-empty header.
- Handle CORS preflight (`OPTIONS`) explicitly. Do not use `Access-Control-Allow-Origin: *` with cookies or other credentials.
- A recovery middleware cannot replace careful error handling, especially after a response has started.

## Server lifecycle and tests

- Configure `ReadHeaderTimeout`, plus sensible read, write, and idle timeouts.
- Handle `ListenAndServe` errors and treat `http.ErrServerClosed` as normal shutdown.
- On `SIGINT`/`SIGTERM`, call `Server.Shutdown` with a timeout to finish in-flight requests.
- Test handlers and middleware with `httptest`; use `go test -race ./...` when state is shared.

## Runnable example

`server.go` provides a small `POST /users` JSON API and `GET /healthz`. It demonstrates the points above using only the Go standard library.

```bash
go test ./http/07-production-checklist
go run ./http/07-production-checklist

curl -i -X POST http://localhost:8080/users \
  -H 'Content-Type: application/json' \
  -d '{"name":"Arpit","email":"arpit@example.com"}'
```
