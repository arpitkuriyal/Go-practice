# 06. JSON APIs: Send and Receive Go Data

## Start simple

JSON is a common text format for API data.

```json
{
  "name": "Arpit",
  "age": 21
}
```

In Go, a struct describes this data:

```go
type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

The text inside the backticks is a JSON tag. It says the JSON field should be named `name`, not `Name`.

## First runnable JSON API

Run [`json-example.go`](json-example.go):

```bash
go run ./http/06-json-encoding/json-example.go
```

In a second terminal, send JSON with `curl`:

```bash
curl -i -X POST http://localhost:8080/users \
  -H 'Content-Type: application/json' \
  -d '{"name":"Arpit","age":21}'
```

Try invalid input too:

```bash
curl -i -X POST http://localhost:8080/users \
  -H 'Content-Type: application/json' \
  -d '{"name":"","age":0}'
```

## Decode: JSON to Go

```go
var input user
err := json.NewDecoder(r.Body).Decode(&input)
```

`r.Body` is the data sent by the client. `&input` is a pointer, allowing the decoder to fill in the struct fields. If decoding fails, return `400 Bad Request`.

## Validate after decoding

Valid JSON is not automatically valid application data. This JSON is correctly formatted but should be rejected:

```json
{"name":"","age":0}
```

```go
if input.Name == "" || input.Age <= 0 {
	http.Error(w, "name and positive age are required", http.StatusBadRequest)
	return
}
```

## Encode: Go to JSON

```go
w.Header().Set("Content-Type", "application/json; charset=utf-8")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(input)
```

Set the content type before writing the response. `201 Created` tells the client that a new resource was created.

## Next level: safer public APIs

Once the basic flow is clear, add these protections:

- Limit body size with `http.MaxBytesReader`.
- Use `decoder.DisallowUnknownFields()` to reject misspelled fields when a strict contract is appropriate.
- Reject trailing JSON values, such as `{...} {...}`.
- Return a consistent JSON error shape such as `{"error":"invalid input"}`.
- Use `422 Unprocessable Content` when JSON syntax is valid but business validation fails.

The production lesson contains the full safe version.

## Interview answer

“I decode the request body into a struct, validate the fields, set the JSON content type, write the correct status, and encode a response. For public APIs I also bound the body and reject unknown or trailing JSON.”
