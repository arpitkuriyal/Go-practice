# HTTP Authentication Middleware

Bearer-token middleware that rejects missing or invalid tokens with JSON errors. A valid request receives its authenticated user through `context.Context`.

```bash
go test -race ./challenges/09-http-auth-middleware
```
