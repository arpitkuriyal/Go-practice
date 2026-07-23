# Requests and Responses: Headers, Parameters, Bodies, and Statuses

## revision

| Part | Use | Go access |
| --- | --- | --- |
| Path parameter | Identifies a resource: `/users/42` | `r.PathValue("id")` |
| Query parameter | Filters/options: `?page=2&limit=20` | `r.URL.Query().Get("page")` |
| Header | Metadata: auth, format, cache policy | `r.Header.Get("Authorization")` |
| Body | Payload for create/update operations | `r.Body` |
| Context | Cancellation/deadline for request work | `r.Context()` |

## Headers

Headers are case-insensitive metadata. Common examples:

| Header | Meaning |
| --- | --- |
| `Content-Type` | Format of this body: `application/json`. |
| `Accept` | Formats the client can receive. |
| `Authorization` | Credentials, commonly `Bearer <token>`. |
| `Location` | URL of a resource created by a `201` response. |
| `Cache-Control` | Caching policy. |
| `Set-Cookie` | Instructs the client to store a cookie. |

Set response headers before the response starts:

```go
w.Header().Set("Content-Type", "application/json; charset=utf-8")
w.Header().Set("Cache-Control", "no-store")
```

## Status codes worth knowing

| Code | Use it when |
| --- | --- |
| `200 OK` | Successful read or response with a body. |
| `201 Created` | A resource was created; include `Location` when possible. |
| `202 Accepted` | Work was accepted but completes asynchronously. |
| `204 No Content` | Success with no body. |
| `400 Bad Request` | Malformed JSON, invalid syntax, or malformed parameters. |
| `401 Unauthorized` | Authentication is missing or invalid. |
| `403 Forbidden` | Identity is known but lacks permission. |
| `404 Not Found` | Route or resource does not exist. |
| `405 Method Not Allowed` | Path exists but method is unsupported; set `Allow`. |
| `409 Conflict` | State conflicts, such as duplicate unique email. |
| `422 Unprocessable Content` | Syntax is valid but business validation fails. |
| `429 Too Many Requests` | Rate limit exceeded. |
| `500 Internal Server Error` | Unexpected server failure; do not expose internals. |

## Parameters: path versus query

```text
GET /users/42?include=posts
     └── resource ID  └── option for the representation
```

Path parameters identify the resource. Query parameters filter, paginate, sort, or choose a representation. Parse and validate query values rather than trusting them:

```go
limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
if err != nil || limit < 1 || limit > 100 {
	http.Error(w, "limit must be between 1 and 100", http.StatusBadRequest)
	return
}
```

## Bodies

The request body is a stream. Bound untrusted input and close only bodies you create as an HTTP client; Go's server manages `r.Body` for handlers.

```go
r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MiB limit
```

For JSON, decode once into a request type, validate it, and return a consistent JSON response. See the JSON lesson for the full pattern.

## Response order: a common trap

```go
w.Header().Set("Content-Type", "application/json") // first
w.WriteHeader(http.StatusCreated)                     // second
json.NewEncoder(w).Encode(value)                       // last
```

The first `Write` or `WriteHeader` commits status and headers. A second `WriteHeader` does not replace the first one.

## Interview answer

“Path parameters name a resource; query parameters modify the request. I validate both, set headers before writing a response, choose a status that describes the outcome, and pass `r.Context()` to downstream work.”
