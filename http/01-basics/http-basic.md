# HTTP Basics: Request to Response

## revision

HTTP is a stateless request-response protocol. A client sends a request; the server selects a handler, performs work, and returns a response.

```text
Client → HTTP request → router → handler → application / database
Client ← HTTP response ← handler ← result / error
```

In Go, a handler has this shape:

```go
func handler(w http.ResponseWriter, r *http.Request)
```

| Value | Purpose |
| --- | --- |
| `r *http.Request` | The incoming method, URL, headers, body, cookies, and request context. |
| `w http.ResponseWriter` | The response status, headers, and body sent back to the client. |

`*http.Request` is a pointer because it is a substantial struct and should not be copied. `http.ResponseWriter` is already an interface, so handlers receive the interface value directly.

## Anatomy of a request

```http
GET /products?limit=10 HTTP/1.1
Host: api.example.com
Accept: application/json
Authorization: Bearer <token>
```

- **Method**: intended action, such as `GET` or `POST`.
- **Path**: resource location, such as `/products`.
- **Query**: optional modifiers, such as `limit=10`.
- **Headers**: metadata, such as authentication and content format.
- **Body**: payload, usually used for create or update requests.

## Anatomy of a response

```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8

{"items":[]}
```

A response consists of a status code, headers, and an optional body. In Go, set headers before calling `WriteHeader` or writing the body:

```go
w.Header().Set("Content-Type", "application/json; charset=utf-8")
w.WriteHeader(http.StatusOK)
_, _ = w.Write([]byte(`{"status":"ok"}`))
```

Writing a body first implicitly sends `200 OK`; changing headers or status after that is too late.

## Minimal server

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
})

server := &http.Server{Addr: ":8080", Handler: mux}
log.Fatal(server.ListenAndServe())
```

Use an explicit `ServeMux` instead of the global `http.DefaultServeMux` in application code. It keeps routes isolated and easy to test.

## Rules to remember

- A handler may run concurrently for many requests. Do not use unsynchronized shared mutable state.
- `r.Context()` is cancelled when the client disconnects or the server shuts down; pass it to database and outbound HTTP calls.
- Read a request body only as much as needed and set a size limit for untrusted input.
- Never call `log.Fatal` inside a handler: it terminates the process.

## Interview answer

“`net/http` dispatches each request to an `http.Handler`. The handler reads from `*http.Request` and writes status, headers, then a body through `http.ResponseWriter`. Request contexts carry cancellation through downstream work.”
