# HTTP Authentication Middleware

## Challenge

Write `net/http` middleware that validates a Bearer token, returns consistent JSON errors, and makes the authenticated user available to the next handler.

## Request flow

```text
request → read Authorization header → validate Bearer token
        → reject with 401 OR add user to context → next handler
```

The middleware has the standard decorator shape:

```go
func Middleware(tokens map[string]string) func(http.Handler) http.Handler
```

## Token parsing

The expected header is:

```text
Authorization: Bearer <token>
```

`strings.CutPrefix` ensures the scheme is present instead of treating any non-empty `Authorization` header as valid. Missing and unknown tokens both return `401 Unauthorized`.

## Context propagation

After successful authentication, create a derived request and pass it onward:

```go
ctx := context.WithValue(r.Context(), userKey, user)
next.ServeHTTP(w, r.WithContext(ctx))
```

Use a private typed context key to avoid collisions. Context values are appropriate for request-scoped identity; they are not a replacement for ordinary required function parameters.

## Response rules

Errors use a stable JSON shape:

```json
{"error":"invalid bearer token"}
```

Set `Content-Type` before `WriteHeader`, then return immediately. Continuing after a rejection could invoke the protected handler.

## Security notes

- This token map is a learning example, not a production credential store.
- Never log raw Bearer tokens.
- Validate token signature, expiry, audience, and issuer when using signed tokens.
- Authentication establishes identity (`401` on failure); authorization checks permission (`403` on failure).
- Protect shared token state if it can change concurrently.

## Interview answer

“Middleware wraps the next handler. It validates the Bearer scheme and token, short-circuits failures with a JSON `401`, and attaches the authenticated identity to a derived request context for downstream handlers.”

## Test

```bash
go test -race ./challenges/09-http-auth-middleware
```
