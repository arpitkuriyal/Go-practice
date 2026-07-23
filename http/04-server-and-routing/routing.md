# Routing Quick Practice

## Fixed routes

```go
mux.HandleFunc("/", home)
mux.HandleFunc("/about", about)
```

Use fixed routes for pages and endpoints with no changing path part.

## Routes with a value

```go
mux.HandleFunc("GET /users/{id}", getUser)

func getUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintln(w, "user ID:", id)
}
```

`/users/42` makes `id` equal to `"42"`. It is still a string, so parse and validate it before using it as a numeric ID.

## Path or query?

```text
/users/42?include=posts
```

- `42` is a **path value**: it identifies the user.
- `include=posts` is a **query value**: it changes what details are returned.

## A useful API shape

```text
GET  /users          list users
POST /users          create a user
GET  /users/{id}     get one user
```

Keep routes small and let each handler do one job.
