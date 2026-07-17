# HTTP Basics — Headers, Status Codes, Body, Query Params, Path Params

---

# 1. Headers

## Definition

Headers are **metadata** sent with request or response.

They describe extra information about the message.

```text id="zk9h7m"
Key: Value
```

---

## Request Header Example

```http id="y6n1l4"
GET /users HTTP/1.1
Authorization: Bearer token123
Content-Type: application/json
User-Agent: Chrome
```

---

## Response Header Example

```http id="4n57wy"
HTTP/1.1 200 OK
Content-Type: application/json
Set-Cookie: session=abc
```

---

## Common Headers

* `Authorization` → login/auth token
* `Content-Type` → body format
* `Accept` → expected response type
* `Set-Cookie` → save cookie
* `Cache-Control` → cache rules

---

## Go Example

```go id="ij3n7j"
token := r.Header.Get("Authorization")

w.Header().Set("Content-Type", "application/json")
```

---

# 2. Status Codes

## Definition

Status codes tell client what happened after request.

---

## Common Status Codes

### Success (2xx)

```text id="1m5nki"
200 OK         -> success
201 Created    -> new resource created
204 No Content -> success, no body
```

### Client Error (4xx)

```text id="s9j7m1"
400 Bad Request   -> invalid input
401 Unauthorized -> login required
403 Forbidden    -> no permission
404 Not Found    -> route/resource missing
405 Method Not Allowed
```

### Server Error (5xx)

```text id="4g7kg6"
500 Internal Server Error
502 Bad Gateway
503 Service Unavailable
```

---

## Go Example

```go id="jv4n0p"
w.WriteHeader(http.StatusCreated)
fmt.Fprint(w, "User created")
```

---

# 3. Body

## Definition

Body contains actual data sent in request or response.

---

## Request Body Example

```json id="x9x7hf"
{
  "name": "Arpit",
  "email": "arpit@mail.com"
}
```

Used in:

* POST
* PUT
* PATCH

---

## Response Body Example

```json id="tdn9s2"
{
  "message": "success"
}
```

---

## Go Read Body

```go id="k2v0wa"
data, _ := io.ReadAll(r.Body)
```

---

## Go Write Body

```go id="efp3f8"
fmt.Fprint(w, "Hello")
```

---

# 4. Query Params

## Definition

Extra values in URL after `?`

Used for:

* filtering
* search
* pagination
* sorting

---

## Example

```text id="lf3b2k"
GET /products?category=phone&page=2
```

Here:

* `category = phone`
* `page = 2`

---

## Go Example

```go id="99h4qf"
page := r.URL.Query().Get("page")
category := r.URL.Query().Get("category")
```

---

## Multiple Params

```text id="qq4g5p"
?search=go&sort=asc&limit=10
```

Use `&` to separate params.

---

# 5. Path Params

## Definition

Dynamic values inside URL path.

Used to identify specific resource.

---

## Example

```text id="2g2g35"
GET /users/10
GET /posts/55
GET /orders/999
```

Here:

* `10` = user id
* `55` = post id

---

## Meaning

```text id="t2j3li"
GET /users       -> all users
GET /users/10    -> user with id 10
```

---

## In Standard net/http

Need to parse manually:

```go id="2w4qk2"
id := strings.TrimPrefix(r.URL.Path, "/users/")
```

---

## In Routers Like Gin / Chi

```go id="9w7mye"
id := c.Param("id")
```

---

# Quick Difference

```text id="d6vt2g"
Path Param  = identifies resource
Query Param = modifies request
```

Example:

```text id="k2mylo"
/users/10?active=true
```

* `10` = path param
* `active=true` = query param

---

# Interview Lines

* Headers carry metadata like auth and content type.
* Status codes show result of request.
* Body contains actual payload.
* Query params are optional filters/options.
* Path params identify specific resource.
