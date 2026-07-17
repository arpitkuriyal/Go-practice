# Short Notes — HTTP Request / Response Lifecycle

## Core Flow

Client sends Request → Server receives → Route to handler → Process logic → Send Response

---

## In Go

```go
func handler(w http.ResponseWriter, r *http.Request)
```

* `w` = `ResponseWriter` → used to send response back
* `r` = `*Request` → contains incoming request data

---

## Why `*` in `r` and not in `w`

### `r *http.Request`

* `Request` is a large struct with many fields (URL, Header, Body, Method, Context, etc.)
* Passing pointer avoids copying whole struct.
* Handler can access same request object efficiently.
* Standard library uses pointer because request data is complex.

### `w http.ResponseWriter`

* `ResponseWriter` is an **interface**, not a struct.
* Interfaces already reference underlying concrete value internally.
* No need to use `*`.
* We call methods like:

```go
w.Write()
w.Header()
```

### Easy Memory Trick

* Large struct → usually pointer (`*Request`)
* Interface → no pointer needed (`ResponseWriter`)

---

## Important Points

Server starts with:

```go
http.ListenAndServe(":8080", nil)
```

* Server keeps listening continuously.
* Handler runs only when request arrives.
* Each request gets its own `handler(w, r)` call.

---

## Request Contains

* Method (`GET`, `POST`)
* URL
* Headers
* Body
* Query params
* Path
* Cookies

---

## Response Contains

* Status code
* Headers
* Body

---

## Interview Line

HTTP works in request-response model: client sends request, server processes it, returns response.
