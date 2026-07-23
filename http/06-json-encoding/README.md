# JSON APIs: Decode, Validate, Respond Safely

## revision

```text
limit body → decode one JSON value → reject unknown fields → validate → call service → write JSON
```

Use distinct request and response types when the API contract differs from the internal model. JSON tags define the wire names.

```go
type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
```

## Writing JSON

Set headers before writing status/body. A helper keeps all responses consistent:

```go
func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		// The response may already be committed: log internally.
	}
}
```

Use `json.NewEncoder(w)` rather than creating an intermediate string. Do not send a body with `204 No Content`.

## Reading JSON safely

```go
func decodeCreateUser(w http.ResponseWriter, r *http.Request) (createUserRequest, error) {
	var input createUserRequest
	decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20)) // 1 MiB
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&input); err != nil {
		return input, err
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return input, errors.New("body must contain exactly one JSON value")
	}
	return input, nil
}
```

This pattern limits a malicious body, rejects misspelled fields, and rejects `{} {}` or other trailing JSON. Decode into a pointer so the decoder can populate the value.

## Validation and error mapping

Decoding checks syntax; validation checks meaning. Keep them separate:

```go
if strings.TrimSpace(input.Name) == "" {
	writeJSON(w, http.StatusUnprocessableEntity, map[string]string{
		"error": "name is required",
	})
	return
}
```

| Problem | Typical response |
| --- | --- |
| Malformed JSON / bad query syntax | `400 Bad Request` |
| Valid JSON but invalid fields | `422 Unprocessable Content` |
| Duplicate or incompatible current state | `409 Conflict` |
| Missing/invalid credentials | `401 Unauthorized` |
| Unexpected internal failure | `500 Internal Server Error` |

Return a stable public error shape such as `{"error":"name is required"}`. Log internal errors with enough context, but do not expose implementation details.

## Presence versus zero values

For a create request, a zero value is often enough. For `PATCH`, distinguish omitted fields from explicit zero values with pointers:

```go
type updateUserRequest struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}
```

`nil` means absent; `Age != nil && *Age == 0` means the client explicitly supplied zero.

## Interview traps

- `json.Decoder.Decode` accepts one JSON value but does not alone reject trailing values.
- `DisallowUnknownFields` is useful for strict public APIs, but consider backward compatibility before enabling it.
- `omitempty` changes the response contract; do not use it only to hide a bug.
- Content type is metadata, not validation. Still decode and validate the actual body.

## One-line interview answer

“For a Go JSON endpoint, I limit the body, decode exactly one strict value, validate business rules, map failures to stable statuses and JSON errors, and set headers before writing the response.”
