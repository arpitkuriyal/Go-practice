# net/http Core Types

These are the most important building blocks of Go HTTP server.

```text id="q6iwod"
http.Handler
http.HandlerFunc
ResponseWriter
*Request
ServeMux
```

They work together when request comes to your server.

---

# Full Flow

```text id="9kpl5z"
Client Request
   ↓
ServeMux matches route
   ↓
Handler / HandlerFunc runs
   ↓
Uses Request data
   ↓
Writes response using ResponseWriter
```

---

# 1. http.Handler

## Definition

`http.Handler` is an interface.

```go id="pw5x8q"
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}
```

Any type with `ServeHTTP()` method becomes a handler.

---

## Example

```go id="jgk7ek"
type HomeHandler struct{}

func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
```

Now `HomeHandler` can serve requests.

---

## Use Case

* Custom handlers
* Middleware
* Routers
* Struct-based dependencies

---

# 2. http.HandlerFunc

## Definition

Adapter that converts normal function into Handler.

```go id="7c8f0v"
type HandlerFunc func(http.ResponseWriter, *http.Request)
```

This is why normal functions work in Go HTTP.

---

## Example

```go id="6j3n9x"
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}
```

Used as:

```go id="i9e8o1"
http.HandleFunc("/", home)
```

Internally Go converts function into `HandlerFunc`.

---

## Why Useful?

You don't need struct every time.

Use simple function for small routes.

---

# 3. ResponseWriter

## Definition

Used to send response back to client.

```go id="c9z2n1"
w http.ResponseWriter
```

It is an interface.

---

## Common Uses

### Write body

```go id="m6f1p7"
fmt.Fprint(w, "Hello")
```

### Set headers

```go id="x7h3l9"
w.Header().Set("Content-Type", "application/json")
```

### Set status code

```go id="r2j4q6"
w.WriteHeader(201)
```

---

## Why Interface?

Different internal implementations can still behave same way.

---

# 4. *Request

## Definition

Incoming HTTP request object.

```go id="j0v2m4"
r *http.Request
```

Contains all client request data.

---

## Why Pointer `*`?

* Large struct
* Avoid copying
* Efficient
* Shared request object

---

## Important Fields

### Method

```go id="h8s1q5"
r.Method
```

### URL Path

```go id="k7d2n3"
r.URL.Path
```

### Query Params

```go id="u3f8p1"
r.URL.Query().Get("page")
```

### Headers

```go id="w5m7r2"
r.Header.Get("Authorization")
```

### Body

```go id="v1n6c4"
r.Body
```

---

# 5. ServeMux

## Definition

Default Go router / multiplexer.

It matches URL path to handlers.

```go id="b9p2l7"
mux := http.NewServeMux()
```

---

## Example

```go id="e6r1m8"
mux.HandleFunc("/", home)
mux.HandleFunc("/about", about)

http.ListenAndServe(":8080", mux)
```

---

## Flow

```text id="q1m3s8"
/about request
↓
ServeMux checks routes
↓
Runs about handler
```

---

## Why Called Multiplexer?

Because it routes many paths to many handlers.

---


# Quick Memory Trick

```text id="u0s4p9"
Handler      -> object that handles request
HandlerFunc  -> function as handler
ResponseWriter -> sends response
Request      -> incoming request data
ServeMux     -> router
```

---

# Interview Lines

* `http.Handler` is core interface with `ServeHTTP()`.
* `http.HandlerFunc` converts functions into handlers.
* `ResponseWriter` writes response body, headers, status.
* `*Request` contains request info.
* `ServeMux` routes paths to handlers.
