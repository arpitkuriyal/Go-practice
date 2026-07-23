# 05. Middleware: Code That Runs Around a Handler

## Start simple

Middleware is a wrapper around a handler. It can do something before the handler, after the handler, or stop the request before it reaches the handler.

```text
request → middleware → your handler → response
```

For example, logging middleware can print every incoming request without repeating logging code in every handler.

## The basic shape

```go
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // continue to the real handler
		fmt.Println("after")
	})
}
```

| Part | Meaning |
| --- | --- |
| `next` | The handler that should run after this middleware. |
| `next.ServeHTTP(w, r)` | Continue the request. |
| `return` before `next` | Stop the request, often after an error response. |

## Run the example

[`middleware-example.go`](middleware-example.go) contains logging, auth, recovery, and CORS examples:

```bash
go run ./http/05-middleware/middleware-example.go
```

Try a request without a token, then with one:

```bash
curl -i http://localhost:8080/
curl -i -H 'Authorization: Bearer demo-token' http://localhost:8080/
```

The example only checks whether a header exists. Real authentication must validate the token.

## Common middleware, one at a time

| Middleware | Beginner explanation |
| --- | --- |
| Logging | Record the request method, URL, and time taken. |
| Authentication | Check who the caller is. Stop with `401` if not authenticated. |
| Recovery | Stop one unexpected panic from crashing the request flow. |
| CORS | Tell browsers which frontend origins may call the server. |
| Rate limiting | Reject too many requests with `429`. |

## Combining middleware

```go
handler := logging(auth(mux))
http.ListenAndServe(":8080", handler)
```

The request enters `logging` first, then `auth`, then `mux`. On the way back, it returns in reverse order.

```text
in:  logging → auth → handler
out: handler → auth → logging
```

Order matters. Put recovery around the handlers it should protect. Put authentication before a protected handler. Always return after writing an auth error.

## Next level: context

After authentication, middleware can add the authenticated user to a derived request context and pass it to the next handler. Use a private typed key and only store request-scoped values such as identity or a request ID.

## Interview answer

“Middleware decorates a handler with shared behavior. It receives the next handler, runs before and/or after it, and can short-circuit a request. The order of middleware changes behavior, so I make it intentional.”
