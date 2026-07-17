# HTTP Methods — GET, POST, PUT, DELETE

## Purpose

HTTP Methods tell the server what action the client wants to perform.

---

## Common Methods

### GET

Used to fetch/read data.

```text
GET /users
GET /products
```

* Should not modify server data
* Can be cached
* Used for viewing/searching

---

### POST

Used to create new data.

```text
POST /users
POST /login
```

* Usually sends data in request body
* Used for signup, login, create resource

---

### PUT

Used to update/replace existing data.

```text
PUT /users/5
PUT /profile
```

* Usually sends updated data in body
* Replaces existing resource

---

### DELETE

Used to remove data.

```text 
DELETE /users/5
DELETE /posts/10
```

* Deletes existing resource

---

## REST API Pattern

```text 
GET     /users      -> list users
GET     /users/1    -> get one user
POST    /users      -> create user
PUT     /users/1    -> update user
DELETE  /users/1    -> delete user
```

---

## Go Example

```go
func users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "Get users")
	case "POST":
		fmt.Fprint(w, "Create user")
	case "PUT":
		fmt.Fprint(w, "Update user")
	case "DELETE":
		fmt.Fprint(w, "Delete user")
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}
```

---

## Important Concepts

### Idempotent

**Idempotent** means:

> Sending the same request multiple times produces the same final server state as sending it once.

It does **not** always mean same response body.
It means the end result on server remains unchanged after first successful request.

---

## Why It Matters

* Safe retries when network fails
* Load balancers may retry requests
* Clients can resend requests without duplicating data
* Important in distributed systems

---

## Methods Explanation

### GET

Used to read data.

```text
GET /users/1
GET /products
```

Calling 1 time or 10 times does not change database state.

Same final result: data is only fetched.

---

### PUT

Used to replace/update resource.

```text
PUT /users/1
{
  "name": "Arpit"
}
```

Send once:

```text
User name = Arpit
```

Send again with same body:

```text
User name still = Arpit
```

No extra duplicate resource created.

---

### DELETE (Usually)

Used to remove resource.

```text
DELETE /users/1
```

First call:

```text
User deleted
```

Second call:

```text
User already missing
```

Even if second response may be `404 Not Found`, final state is same:

```text
User does not exist
```

So usually considered idempotent.

---

### POST

Used to create new resource.

```text
POST /orders
```

First request:

```text
Order #101 created
```

Second same request:

```text 
Order #102 created
```

Now two resources exist.

So repeated request changes state again.

Not idempotent.

---

## Quick Memory Trick

```text 
GET     -> Read again = same
PUT     -> Set again = same
DELETE  -> Remove again = same end state
POST    -> Create again = duplicate/new data
```

---

## One Liner

Idempotent means repeating the same request keeps the final server state unchanged. GET, PUT, DELETE are idempotent; POST usually is not.


HTTP methods define action type: GET reads, POST creates, PUT updates, DELETE removes.
