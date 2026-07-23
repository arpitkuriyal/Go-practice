# Middleware: Cross-Cutting HTTP Behavior

## Mental model

Middleware takes a handler and returns a handler:

```go
type Middleware func(http.Handler) http.Handler
```

It can act before a handler, after it returns, or stop the request completely.

```text
request → recovery → logging → auth → application handler → response
```

## Basic pattern

```go
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s finished in %s", r.Method, r.URL.Path, time.Since(start))
	})
}
```

`next.ServeHTTP` is the point where the request continues. Omitting it intentionally short-circuits the request.

## Common middleware

| Middleware | Responsibility | Important detail |
| --- | --- | --- |
| Recovery | Contain unexpected panics | Cannot undo a response already written. |
| Logging/metrics | Record method, path, status, latency | Do not log tokens or credentials. |
| Authentication | Establish identity | Return `401` for missing/invalid credentials. |
| Authorization | Check permission | Return `403` after identity is known. |
| CORS | Control browser cross-origin access | Handle `OPTIONS` preflight; do not use `*` with credentials. |
| Rate limit | Protect capacity/abuse | Return `429` and consider `Retry-After`. |
| Request IDs | Correlate logs and traces | Store with a typed context key. |

## Ordering matters

With nested calls, the outer middleware runs first on the way in and last on the way out:

```go
handler := recovery(logging(auth(mux)))
```

- `recovery` should normally be outermost so it covers the other middleware.
- `logging` should wrap protected routes if you want to record rejected requests too.
- `auth` must return immediately after writing an error; otherwise the protected handler runs.

## Context values

Middleware often adds request-scoped information:

```go
type contextKey string
const requestIDKey contextKey = "request_id"

ctx := context.WithValue(r.Context(), requestIDKey, requestID)
next.ServeHTTP(w, r.WithContext(ctx))
```

Use private typed keys and only cross-cutting request values. Do not put required business arguments or database handles in a context.

## Response-status logging trap

`ResponseWriter` does not expose the final status after the handler writes it. A production logger usually wraps it to record the first status passed to `WriteHeader` and treats a first `Write` as `200`. Be careful to preserve optional interfaces when writing a general-purpose wrapper.

## Recovery caveat

Recovery is a safety net, not normal error handling. If a downstream handler has started the response, sending a clean `500` may be impossible. Log the panic internally; do not return panic details to the client.

## Interview answer

“Middleware decorates handlers with shared behavior. It is composable, can short-circuit requests, and its order is part of correctness—especially recovery, logging, authentication, and CORS.”
