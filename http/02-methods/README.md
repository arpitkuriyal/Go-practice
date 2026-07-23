# HTTP Methods: Intent, Safety, and Idempotency

## revision

| Method | Normal intent | Safe* | Idempotent** | Typical success |
| --- | --- | --- | --- | --- |
| `GET` | Read a representation | Yes | Yes | `200 OK` |
| `HEAD` | Read headers only | Yes | Yes | `200 OK` |
| `POST` | Create or trigger an action | No | Usually no | `201 Created` / `202 Accepted` |
| `PUT` | Replace a known resource | No | Yes | `200 OK` / `204 No Content` |
| `PATCH` | Partially update a resource | No | Depends on operation | `200 OK` / `204 No Content` |
| `DELETE` | Remove a resource | No | Usually yes | `204 No Content` |

*Safe means the request should not change server state.
**Idempotent means repeating the same request has the same intended final server state; its response can still differ.

## Resource-oriented routes

```text
GET    /users          list users
POST   /users          create a user
GET    /users/{id}     get one user
PUT    /users/{id}     replace one user
PATCH  /users/{id}     change selected fields
DELETE /users/{id}     remove one user
```

This is a useful convention, not a law. A command that does not fit a resource can be explicit: `POST /reports/{id}:publish`.

## `PUT` versus `PATCH`

- `PUT` sends the full desired representation. Omitted fields commonly mean replacement/default values.
- `PATCH` sends only changes. Define its semantics carefully; a patch such as “increment balance” is not idempotent.

## Why idempotency matters

A timeout does not prove that the server did not process a request. Clients, proxies, and load balancers can retry. Retrying the same `POST /orders` might create duplicate orders, while repeating a `PUT /users/7` with the same representation should leave the final state unchanged.

For a retryable create operation, consider an idempotency key stored with the result.

## Method handling in Go

Modern Go `ServeMux` can match method and path together:

```go
mux.HandleFunc("GET /users/{id}", getUser)
mux.HandleFunc("POST /users", createUser)
```

For a shared endpoint, switch explicitly and return `405` with an `Allow` header:

```go
switch r.Method {
case http.MethodGet:
	// list
case http.MethodPost:
	// create
default:
	w.Header().Set("Allow", "GET, POST")
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
```

## Interview traps

- `401 Unauthorized` means authentication is missing or invalid; `403 Forbidden` means authentication succeeded but permission is denied.
- `DELETE` can return `404` on a repeated call and still be idempotent: the target remains absent.
- A `GET` with side effects is unsafe and may be cached, prefetched, or retried unexpectedly.

## One-line interview answer

“Methods express intent. I make reads safe, use idempotent operations when retries are expected, distinguish replacement (`PUT`) from partial update (`PATCH`), and return `405` for unsupported methods.”
