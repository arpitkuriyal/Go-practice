# 06. Errors: Tell the Caller What Went Wrong

Go functions often return two values: the result and an error.

```go
user, err := findUser(id)
if err != nil {
	return err
}
```

`nil` means no error. A non-nil error explains why the operation failed.

## A known error: sentinel errors

Use `errors.New` for an expected condition callers may need to handle:

```go
var ErrInvalidEmail = errors.New("invalid email")
```

The caller can make a decision:

```go
if errors.Is(err, ErrInvalidEmail) {
	// ask the user to correct the email
}
```

## Add useful context with `%w`

```go
return fmt.Errorf("register user: %w", ErrInvalidEmail)
```

The message now says which operation failed, but `%w` keeps the original error available to `errors.Is`.

```go
errors.Is(err, ErrInvalidEmail) // true
```

Do not compare a wrapped error with `==`; that only compares the outer error.

## Errors with details

Sometimes a caller needs more than a yes/no category. This lesson’s `ValidationError` has a field and reason:

```go
var validationErr *ValidationError
if errors.As(err, &validationErr) {
	fmt.Println(validationErr.Field)
}
```

Use `errors.Is` to ask “is this kind of error?” Use `errors.As` to ask “does this error contain this type with extra details?”

## Good error habits

- Return errors to the layer that can handle them.
- Add context such as `"load user: %w"`, not vague text such as `"failed"`.
- Usually log an error once at the application boundary; logging it at every layer creates duplicates.
- Do not use `panic` for normal bad input, missing records, or network failures.
- Do not put secrets or internal infrastructure details in errors sent to users.

## Interview answer

“Errors are normal return values in Go. I use sentinel errors for known conditions, wrap errors with `%w` to add operation context, check categories with `errors.Is`, and use `errors.As` when callers need structured details.”

## Test

```bash
go test ./concepts/06-errors
```
