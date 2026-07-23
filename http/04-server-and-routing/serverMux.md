# `http.NewServeMux()`

## The problem it solves

One server usually has more than one page or API endpoint:

```text
/          home page
/about     about page
/users     users API
```

The server needs something that reads the request path and picks the correct function. `ServeMux` does that job.

## Step 1: create the router

```go
mux := http.NewServeMux()
```

Think of `mux` as an empty map from URL patterns to handlers.

## Step 2: register routes

```go
mux.HandleFunc("/", home)
mux.HandleFunc("/about", about)
```

The first argument is the URL pattern. The second argument is the function to run.

## Step 3: start the server with the router

```go
http.ListenAndServe(":8080", mux)
```

The server receives requests; `mux` decides which handler receives each one.

## Why not pass `nil`?

When you pass `nil`, Go silently uses the global `http.DefaultServeMux`:

```go
http.HandleFunc("/", home)
http.ListenAndServe(":8080", nil)
```

This is okay for a short first example. `NewServeMux` is better for projects because:

- routes are visible in one place;
- tests can build a fresh router; and
- two servers can have different routes without sharing global state.

## Modern route patterns

After you understand basic paths, use method and path patterns:

```go
mux.HandleFunc("GET /users/{id}", getUser)
```

This only matches a `GET` request such as `/users/42`. Inside the handler:

```go
id := r.PathValue("id")
```

Use this instead of manually cutting strings from `r.URL.Path`.
