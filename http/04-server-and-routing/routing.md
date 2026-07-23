# Routing Patterns: From Static Paths to APIs

## Static and dynamic routes

```go
mux.HandleFunc("GET /about", about)          // static
mux.HandleFunc("GET /users/{id}", getUser)   // dynamic
```

Static routes are fixed. Dynamic route segments capture a value:

```go
func getUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.NotFound(w, r)
		return
	}
}
```

Use path values to identify a resource and query values for filtering/pagination:

```text
GET /users/42?include=posts
     └─ resource  └─ representation option
```

## Method-aware patterns

Modern `net/http` can keep method selection in the router:

```go
mux.HandleFunc("GET /users", listUsers)
mux.HandleFunc("POST /users", createUser)
mux.HandleFunc("GET /users/{id}", getUser)
```

This avoids a large method switch and gives correct method handling. If one handler intentionally owns several methods, switch on `r.Method` and set `Allow` for `405` responses.

## Grouping a prefix

Use a child mux and `http.StripPrefix` for a small standard-library subrouter:

```go
api := http.NewServeMux()
api.HandleFunc("GET /users", listUsers)

root := http.NewServeMux()
root.Handle("/api/", http.StripPrefix("/api", api))
```

Keep API versions explicit when their contracts differ: `/api/v1/...`.

## Routing pitfalls

- Do not manually parse paths with `strings.TrimPrefix` when `PathValue` can express the route.
- Validate every path value before using it in a query or lookup.
- Avoid an all-purpose `"/"` handler that implements the whole API with nested `if` statements.
- Prefer resource nouns (`/users`) over verbs (`/getUsers`) for conventional CRUD endpoints.
