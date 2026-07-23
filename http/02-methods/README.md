# 02. HTTP Methods: What Does the Client Want to Do?

## Start simple

The HTTP method tells the server what kind of action the client wants.

```text
GET  /users  → "give me users"
POST /users  → "create a user"
```

The most useful methods to learn first are:

| Method | Plain-English meaning | Typical use |
| --- | --- | --- |
| `GET` | Read something | List users or get one user. |
| `POST` | Create something | Create a user or submit a form. |
| `PUT` | Replace something | Replace all stored details for one user. |
| `PATCH` | Change part of something | Update only a user's name. |
| `DELETE` | Remove something | Delete one user. |

## First runnable example

Run [`methods-example.go`](methods-example.go):

```bash
go run ./http/02-methods/methods-example.go
```

In another terminal, try:

```bash
curl http://localhost:8080/users
curl -X POST http://localhost:8080/users
curl -X DELETE http://localhost:8080/users
```

The example uses `r.Method` to decide what to do:

```go
switch r.Method {
case http.MethodGet:
	// list users
case http.MethodPost:
	// create a user
default:
	// tell the client this URL does not support that method
}
```

Use constants such as `http.MethodGet` rather than spelling `"GET"` yourself.

## `405 Method Not Allowed`

`/users` can be a valid URL while `DELETE /users` is not a valid action in this API. Return `405` and tell the client which methods work:

```go
w.Header().Set("Allow", "GET, POST")
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
```

## Next level: routes and methods together

Go 1.22+ can match both a method and path:

```go
mux.HandleFunc("GET /users", listUsers)
mux.HandleFunc("POST /users", createUser)
mux.HandleFunc("GET /users/{id}", getUser)
```

This is a cleaner choice when different handlers own different operations. You will use it in the routing lesson.

## Important interview idea: idempotency

An operation is **idempotent** when repeating the same request has the same final result as doing it once.

```text
GET /users/1          read again: no data changed
PUT /users/1 {name}   set the same name again: final state is the same
DELETE /users/1       delete again: user is still absent
POST /users           create again: usually creates another user
```

So `GET`, `PUT`, and usually `DELETE` are idempotent; `POST` usually is not. This matters because a client may retry after a network timeout.

## Interview answer

“Methods express intent: `GET` reads, `POST` creates, `PUT` replaces, `PATCH` partially updates, and `DELETE` removes. I return `405` for unsupported methods and think about idempotency when retries are possible.”
