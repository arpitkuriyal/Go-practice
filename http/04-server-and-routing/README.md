# 04. Routing: Send Each URL to the Right Function

 : what is routing?

Routing answers one question:

> “This request came for this URL. Which Go function should run?”

```text
GET /       → home function
GET /about  → about function
```

In Go, `http.ServeMux` is the built-in router. “Mux” means multiplexer: it chooses one handler from many handlers.

## `http.NewServeMux()` from zero

```go
mux := http.NewServeMux()
```

This line creates an empty router. Next, add routes:

```go
mux.HandleFunc("/", home)
mux.HandleFunc("/about", about)
```

Finally, give the router to the server:

```go
http.ListenAndServe(":8080", mux)
```

Now the full flow is:

```text
browser requests /about
        ↓
mux finds the /about route
        ↓
about(w, r) runs
        ↓
about writes the response
```

Run [`http-core-example.go`](http-core-example.go), then open `/` and `/about`:

```bash
go run ./http/04-server-and-routing/http-core-example.go
```

## What is a handler?

A handler is anything that can serve an HTTP request. The core interface is:

```go
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}
```

Most beginners use a normal function:

```go
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}
```

`HandleFunc` automatically turns that function into a handler. Later, a struct can implement `ServeHTTP` when it needs dependencies such as a database or service.

## Why `NewServeMux` instead of `http.HandleFunc`?

This also works:

```go
http.HandleFunc("/", home)
http.ListenAndServe(":8080", nil)
```

It uses a global router called `http.DefaultServeMux`. That is fine for a tiny first program. In a larger application, create your own mux so routes are explicit, isolated, and easy to test:

```go
mux := http.NewServeMux()
mux.HandleFunc("/", home)
server := &http.Server{Addr: ":8080", Handler: mux}
```

Read [ServeMux explained](serverMux.md) for a step-by-step comparison.

## Next level: method and path routes

Go 1.22+ understands method/path patterns:

```go
mux.HandleFunc("GET /users/{id}", getUser)
mux.HandleFunc("POST /users", createUser)
```

For `GET /users/42`, read the changing part with:

```go
id := r.PathValue("id") // "42"
```

Start with fixed routes first. Add method patterns and path values when you understand methods and request data.

## Interview answer

“A router maps an incoming URL to a handler. In Go, `http.NewServeMux` creates an explicit router, `HandleFunc` connects paths to functions, and the mux is passed to the HTTP server.”
