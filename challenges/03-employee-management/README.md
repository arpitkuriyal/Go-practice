# Employee Management

## Challenge

Build an in-memory employee manager that adds, removes, finds employees by ID, and calculates an average salary.

## Concepts practised

| Concept | Where it appears |
| --- | --- |
| Structs | `Employee` models related employee fields. |
| Methods | `Manager` owns operations on its employee slice. |
| Pointer receiver | Methods update the original manager. |
| Slices | Add and remove employees. |
| Linear search | Find an employee by ID. |

## Core patterns

Append adds an employee:

```go
m.Employees = append(m.Employees, employee)
```

Remove the matching element and return immediately:

```go
m.Employees = append(m.Employees[:i], m.Employees[i+1:]...)
```

Finding returns `nil` when an ID is absent, which makes absence explicit to callers.

## Complexity

| Operation | Time |
| --- | --- |
| Add | Amortized `O(1)` |
| Find by ID | `O(n)` |
| Remove by ID | `O(n)` |
| Average salary | `O(n)` |

For many ID lookups, consider `map[int]Employee`; then decide how to preserve ordering and enforce consistency.

## Edge cases to discuss

- Duplicate IDs: should `AddEmployee` reject or replace them?
- Missing ID: removal should be a no-op or return an error.
- Empty manager: average salary needs a defined result. Avoid division by zero (`NaN`) by returning an error, a boolean, or a documented zero value.
- `FindEmployeeByID` returns a pointer to the range copy, not the element inside the slice. It is fine for reading; do not use it to mutate the stored employee.

## Interview answer

“The manager owns a slice of employee values. I use pointer receivers for mutating methods, linear search for a small in-memory collection, and define explicit behavior for duplicate IDs and an empty average.”

## Run

```bash
go run ./challenges/03-employee-management/data-management.go
```
