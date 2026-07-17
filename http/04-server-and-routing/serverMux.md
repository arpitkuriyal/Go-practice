# Why Use ServeMux When `http.HandleFunc()` Exists?

## Short Answer

`http.HandleFunc()` already uses a router internally called:

```go id="t3f8jn"
http.DefaultServeMux
```

So when you write:

```go id="r6v2pk"
http.HandleFunc("/", home)
http.ListenAndServe(":8080", nil)
```

It means:

* Register route on default global router
* `nil` tells server to use `DefaultServeMux`

---

## Internal Equivalent

```go id="y7q3mw"
mux := http.DefaultServeMux
mux.HandleFunc("/", home)

http.ListenAndServe(":8080", mux)
```

---

# Why Use `http.NewServeMux()`?

## 1. Better Structure

Create your own router instance.

```go id="w1k9he"
mux := http.NewServeMux()
mux.HandleFunc("/", home)
mux.HandleFunc("/users", users)
```

Cleaner for medium/large apps.

---

## 2. Avoid Global State

`http.HandleFunc()` uses shared global router.

Custom mux keeps routes local to your app.

---

## 3. Easier Testing

You can test routes independently.

```go id="j8r0fs"
mux := http.NewServeMux()
```

---

## 4. Multiple Routers

Useful for different servers:

```go id="f9m2xa"
apiMux := http.NewServeMux()
adminMux := http.NewServeMux()
```

---

## 5. Production Style

```go id="c4t7po"
server := &http.Server{
	Addr:    ":8080",
	Handler: mux,
}
```

Preferred in real projects.

---

# When to Use What?

```text id="q5w8nd"
Small practice project -> http.HandleFunc()
Real backend project   -> http.NewServeMux()
```

---

# Interview Line

`http.HandleFunc()` uses the global `DefaultServeMux`. `http.NewServeMux()` creates an explicit router instance that is cleaner, safer, and better for scalable applications.
