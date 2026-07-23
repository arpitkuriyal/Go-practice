# Go `net/http`: Handlers, ServeMux, and Routing

## revision

```text
request → ServeMux matches method/path → handler → ResponseWriter
```

| Type | Job |
| --- | --- |
| `http.Handler` | Interface with `ServeHTTP(http.ResponseWriter, *http.Request)`. |
| `http.HandlerFunc` | Function adapter that implements `http.Handler`. |
| `http.ResponseWriter` | Writes response headers, status, and body. |
| `*http.Request` | Holds incoming request data and context. |
| `http.ServeMux` | Standard-library router that maps patterns to handlers. |

## Handlers

Any type implementing `ServeHTTP` is a handler. A function is commonly converted with `http.HandlerFunc`:

```go
func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

var handler http.Handler = http.HandlerFunc(healthz)
```

Use a struct handler when it needs dependencies such as a service, logger, or database:

```go
type userHandler struct{ service UserService }

func (h userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// h.service uses r.Context()
}
```

## Explicit routing with `ServeMux`

Create a mux rather than registering on global `http.DefaultServeMux`:

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /healthz", healthz)
mux.HandleFunc("GET /users/{id}", getUser)
mux.HandleFunc("POST /users", createUser)

server := &http.Server{Addr: ":8080", Handler: mux}
```

Go 1.22+ supports method-qualified patterns and path values. Read a value with `r.PathValue("id")`. `ServeMux` selects the most specific matching pattern and handles a known path with a wrong method as `405 Method Not Allowed`.

## ResponseWriter rules

```go
w.Header().Set("Content-Type", "application/json; charset=utf-8")
w.WriteHeader(http.StatusCreated)
_, err := w.Write(body)
```

- Set headers first.
- `WriteHeader` commits the status; if omitted, the first `Write` commits `200 OK`.
- Do not write a body for `204 No Content`.
- Return after handling an error to avoid writing a second response.

## Request essentials

```go
method := r.Method
path := r.URL.Path
query := r.URL.Query().Get("page")
token := r.Header.Get("Authorization")
ctx := r.Context()
```

Treat request input as untrusted. Validate path/query values, bound bodies, and propagate `ctx` to work that can block.

## Testing a route

```go
req := httptest.NewRequest(http.MethodGet, "/users/42", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)

if rec.Code != http.StatusOK { /* fail test */ }
```

An explicit mux makes this test independent of global state.

## Interview answer

“`http.Handler` is Go’s core server abstraction. `HandlerFunc` adapts ordinary functions, and `ServeMux` routes method/path patterns to handlers. I use an explicit mux, path values, request contexts, and `httptest` for isolated tests.”
