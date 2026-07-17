# JSON APIs in Go

JSON APIs are the most common backend pattern.

Frontend / mobile apps send JSON to server, and server returns JSON.

```text id="u4m7pk"
Client ⇄ JSON ⇄ Go Backend
```

---

# Topics Covered

1. Encode JSON
2. Decode JSON
3. Validation
4. Error Handling

---

# 1. Encode JSON

## Definition

Convert Go struct/map into JSON response.

---

## Example Struct

```go id="h2f8qy"
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

---

## Send JSON Response

```go id="j6p1na"
func handler(w http.ResponseWriter, r *http.Request) {

	user := User{
		Name: "Arpit",
		Age:  21,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}
```

---

## Output

```json id="k3z9xr"
{
  "name": "Arpit",
  "age": 21
}
```

---

## Why Use Encoder?

* Writes directly to response
* Efficient
* Standard way in Go

---

# 2. Decode JSON

## Definition

Read incoming JSON request body into Go struct.

---

## Request JSON

```json id="f8d2ms"
{
  "name": "Arpit",
  "age": 21
}
```

---

## Decode Example

```go id="x9c4ql"
func createUser(w http.ResponseWriter, r *http.Request) {

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	fmt.Println(user.Name, user.Age)
}
```

---

## Why `&user`?

Need pointer so decoder fills struct fields.

---

# 3. Validation

## Definition

Check incoming data before using it.

Never trust client input.

---

## Example Rules

* Name required
* Age > 0
* Email valid
* Password length >= 8

---

## Manual Validation

```go id="s1v7de"
if user.Name == "" {
	http.Error(w, "Name required", 400)
	return
}

if user.Age <= 0 {
	http.Error(w, "Invalid age", 400)
	return
}
```

---

## Why Important?

Without validation:

```text id="q7m4ow"
empty name
negative age
bad email
broken data in DB
```

---

# 4. Error Handling

## Definition

Return clean errors with proper status codes.

---

## Common Status Codes

```text id="z5j8ri"
400 Bad Request
401 Unauthorized
404 Not Found
409 Conflict
422 Unprocessable Entity
500 Internal Server Error
```

---

## Example JSON Error Response

```go id="n4t1ux"
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(400)

json.NewEncoder(w).Encode(map[string]string{
	"error": "invalid input",
})
```

---

## Output

```json id="d3f9ab"
{
  "error": "invalid input"
}
```

---

# Full Example API

```go id="w8r2kc"
package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func createUser(w http.ResponseWriter, r *http.Request) {

	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	if user.Name == "" || user.Age <= 0 {
		http.Error(w, "Validation failed", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/users", createUser)
	http.ListenAndServe(":8080", nil)
}
```

---

# Best Practices

## Always Set Content-Type

```go id="g1y8fd"
w.Header().Set("Content-Type", "application/json")
```

---

## Close Unknown Fields (Advanced)

```go id="f2m6lo"
decoder := json.NewDecoder(r.Body)
decoder.DisallowUnknownFields()
```

Reject extra unexpected fields.

---

## Return Structured Errors

```json id="r9n2xa"
{
  "error": "email required"
}
```

Better than plain text.

---

# Quick Memory Trick

```text id="p6q1mz"
Encode   -> Go → JSON
Decode   -> JSON → Go
Validate -> check input
Errors   -> proper response
```

---

# Interview Lines

* Use `json.NewEncoder(w).Encode()` for responses.
* Use `json.NewDecoder(r.Body).Decode()` for requests.
* Validate all client input.
* Return proper status codes and JSON error messages.
