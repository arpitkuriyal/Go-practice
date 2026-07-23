# `http.NewServeMux` versus `http.DefaultServeMux`

`http.HandleFunc` registers on the global `http.DefaultServeMux`. Passing `nil` to `http.ListenAndServe` uses that global mux.

```go
http.HandleFunc("/", home)
http.ListenAndServe(":8080", nil) // uses DefaultServeMux
```

For a real application, create the mux explicitly:

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /", home)

server := &http.Server{Addr: ":8080", Handler: mux}
```

This avoids global route state, supports multiple independent servers, and lets each test construct a fresh router. The explicit form is the standard production default.
