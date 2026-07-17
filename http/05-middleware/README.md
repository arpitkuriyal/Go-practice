# Middleware in Go HTTP (Very Important)

## Definition

Middleware is a function that runs **before and/or after** your main handler.

Used for common features like:

* Logging
* Authentication
* Recovery
* CORS
* Rate limiting

---

# Mental Model

```text id="a8x2lm"
Request
  ↓
Middleware
  ↓
Handler
  ↓
Response
```

Multiple middleware can be chained.

```text id="6r0v7j"
Request
↓
Logging
↓
Auth
↓
Handler
↓
Response
```

---

# Basic Middleware Pattern

```go id="k2p9wb"
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// before handler
		fmt.Println("Before")

		next.ServeHTTP(w, r)

		// after handler
		fmt.Println("After")
	})
}
```

---

# 1. Logging Middleware

## Purpose

Track incoming requests.

Useful for:

* debugging
* monitoring
* analytics

---

## Example

```go id="s5n3hd"
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
```

---

## Output

```text id="p1n0rk"
GET /users
POST /login
```

---

# 2. Auth Middleware

## Purpose

Protect private routes.

Check:

* token
* session
* API key

---

## Example

```go id="e3r4xm"
func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Unauthorized", 401)
			return
		}

		next.ServeHTTP(w, r)
	})
}
```

---

## Flow

```text id="g5f8tm"
No token -> blocked
Valid token -> continue
```

---

# 3. Recovery Middleware

## Purpose

Prevent server crash from panic.

Without recovery:

```text id="yw8m3s"
panic in one request -> server may crash
```

---

## Example

```go id="n6u2pd"
func recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", 500)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
```

---

# 4. CORS Middleware

## Purpose

Allow frontend from another origin.

Example:

```text id="a9w6lt"
Frontend: localhost:3000
Backend : localhost:8080
```

Browser blocks unless CORS allowed.

---

## Example

```go id="t0m7eq"
func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		next.ServeHTTP(w, r)
	})
}
```

---

# 5. Rate Limiting

## Purpose

Limit too many requests.

Protect from:

* spam
* abuse
* brute force
* overload

---

## Example Logic

```text id="d8q2nv"
Only 100 requests/minute per IP
```

---

## Simple Example

```go id="j4y9po"
if requests > 100 {
	http.Error(w, "Too Many Requests", 429)
	return
}
```

---

# Chaining Middleware

```go id="c7r1mk"
handler := logging(auth(recovery(http.HandlerFunc(home))))
```

Flow:

```text id="r5x3uz"
logging -> auth -> recovery -> home
```

---

# Full Example

```go id="v1k4tb"
mux := http.NewServeMux()
mux.HandleFunc("/", home)

handler := logging(auth(mux))

http.ListenAndServe(":8080", handler)
```

---

# Why Middleware Important?

Because you avoid repeating code in every handler.

Bad:

```text id="m9s0fj"
auth check in every route
logging in every route
```

Good:

```text id="m1p6da"
One middleware applies everywhere
```

---

# Quick Memory Trick

```text id="w2h8rk"
Logging      -> track requests
Auth         -> protect routes
Recovery     -> catch panic
CORS         -> allow frontend origin
Rate Limit   -> block abuse
```

---

# Interview Lines

* Middleware wraps handlers to run common logic.
* It executes before/after main handler.
* Used for logging, auth, recovery, CORS, rate limiting.
* Middleware improves reusable and clean architecture.
