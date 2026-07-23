# 03. Requests and Responses: Reading and Writing Data

## A request has several parts

```text
GET /greet?name=Arpit HTTP/1.1
Authorization: Bearer token

(optional request body)
```

| Part | Example | Why it exists |
| --- | --- | --- |
| Method | `GET` | Says what action is wanted. |
| Path | `/greet` | Says which endpoint is wanted. |
| Query | `?name=Arpit` | Optional extra input. |
| Headers | `Authorization: ...` | Metadata such as credentials or body format. |
| Body | JSON data | Payload, commonly for `POST` and updates. |

## Your first request-data example

Run [`request-response-example.go`](request-response-example.go):

```bash
go run ./http/03-requests-and-responses/request-response-example.go
```

Then visit `http://localhost:8080/greet?name=Arpit`.

The handler reads a query value:

```go
name := r.URL.Query().Get("name")
```

If the query is missing, it uses a default. Always validate user input before using it.

## Writing the response

The response has three parts too:

```text
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8

Hello, Arpit!
```

| Part | Meaning |
| --- | --- |
| Status code | Result of the request, such as `200 OK`. |
| Headers | Metadata about the response, such as its format. |
| Body | Text, JSON, HTML, or other returned data. |

In Go, use this order:

```go
w.Header().Set("Content-Type", "text/plain; charset=utf-8")
w.WriteHeader(http.StatusOK)
fmt.Fprintln(w, "Hello")
```

Set headers first. Writing a body automatically sends `200 OK` if no status has been set yet.

## Status codes to learn first

| Code | Meaning | Common example |
| --- | --- | --- |
| `200 OK` | Request succeeded | Read a user. |
| `201 Created` | A new resource was created | Create a user. |
| `400 Bad Request` | Client sent malformed input | Broken JSON. |
| `401 Unauthorized` | Login/token is missing or invalid | No valid Bearer token. |
| `403 Forbidden` | Logged-in user lacks permission | Normal user tries admin action. |
| `404 Not Found` | No matching route or resource | Unknown user ID. |
| `500 Internal Server Error` | Unexpected server failure | Database failed unexpectedly. |

## Path versus query parameters

```text
/users/42?include=posts
 └── path value └── query value
```

The path normally identifies a resource. The query changes how to fetch it—filters, search terms, pagination, or optional details.

In modern `net/http`, a route such as `GET /users/{id}` reads its path value with `r.PathValue("id")`.

## Next level

For API input, JSON is the usual request body format. The JSON lesson adds decoding, validation, and safe responses.

## Interview answer

“A request contains a method, path, query, headers, and optional body. I read those from `*http.Request`; then I set response headers, write a suitable status code, and write the body through `ResponseWriter`.”
