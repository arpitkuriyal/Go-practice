# Routing in Go HTTP

Routing means:

> Matching incoming URL path to the correct handler.

When request comes:

```text id="v8r3n1"
/users
/about
/products/10
```

Router decides which function should run.

---

# 1. Static Routes

## Definition

Fixed URL paths.

Examples:

```text id="a5n2qk"
/
///about
/contact
/login
```

These paths do not change.

---

## Go Example

```go id="m2c7p9"
mux := http.NewServeMux()

mux.HandleFunc("/", home)
mux.HandleFunc("/about", about)
mux.HandleFunc("/contact", contact)
```

---

## Meaning

```text id="x4f1wr"
/about request -> about handler
/contact request -> contact handler
```

---

## Use Cases

* Homepage
* About page
* Login page
* Health check

---

# 2. Dynamic Routes

## Definition

Routes containing variable values.

Examples:

```text id="q8j4sy"
/users/10
/users/25
/products/99
/posts/500
```

Here IDs change dynamically.

---

## Why Needed?

Instead of creating:

```text id="pw6m0t"
/users/1
/users/2
/users/3
```

Use one route pattern:

```text id="g7w1dn"
/users/{id}
```

---

## Standard net/http (manual parse)

```go id="k1m8te"
func user(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	fmt.Fprint(w, "User ID:", id)
}
```

---

## Router Libraries

Using Gin or Chi:

```go id="h2n4vf"
r.GET("/users/:id", handler)
```

---

# 3. Multiple Handlers

## Definition

Different routes use different handlers.

```go id="t9y1op"
mux.HandleFunc("/", home)
mux.HandleFunc("/users", users)
mux.HandleFunc("/orders", orders)
mux.HandleFunc("/login", login)
```

---

## Flow

```text id="e5d7wa"
/users  -> users handler
/login  -> login handler
/orders -> orders handler
```

---

## Why Important?

Keeps code modular.

Bad:

```go id="g0q9kt"
one giant handler for everything
```

Good:

```go id="p7r2lm"
separate handlers per feature
```

---

# 4. Subrouters Patterns

## Definition

Group related routes under common prefix.

Examples:

```text id="z4f6mj"
/api/users
/api/orders
/api/products

/admin/users
/admin/settings
```

---

## Why Useful?

Organized route structure.

```text id="x6v3pq"
/api/*    -> public API
/admin/*  -> admin panel
/v1/*     -> version 1 API
```

---

## Example with ServeMux

```go id="u3k8ra"
apiMux := http.NewServeMux()
apiMux.HandleFunc("/users", users)
apiMux.HandleFunc("/orders", orders)

mainMux := http.NewServeMux()
mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))
```

---

## With Chi / Gin Easier

```go id="m8s4wb"
r.Route("/api", func(r chi.Router) {
	r.Get("/users", users)
	r.Get("/orders", orders)
})
```

---

# Full Example

```go id="r1v7np"
mux := http.NewServeMux()

mux.HandleFunc("/", home)          // static
mux.HandleFunc("/about", about)   // static
mux.HandleFunc("/users/", user)   // dynamic manual
mux.HandleFunc("/login", login)   // multiple handler

http.ListenAndServe(":8080", mux)
```

---

# Quick Memory Trick

```text id="k2d9go"
Static route   -> fixed path
Dynamic route  -> variable path
Multiple handlers -> many routes many funcs
Subrouter      -> grouped routes by prefix
```

---

# Interview Lines

* Routing maps URL paths to handlers.
* Static routes are fixed paths.
* Dynamic routes contain variables like IDs.
* Multiple handlers keep code modular.
* Subrouters group related endpoints under prefixes.
