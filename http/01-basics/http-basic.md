# 01. HTTP Basics: Your First Go Server

Start here if HTTP is new.

## What is HTTP?

HTTP is how a client and a server talk over the web.

```text
Browser / app                    Go server
     |                                |
     |  request: "please give /"      |
     |------------------------------->|
     |                                | runs Go code
     |  response: "hello world"       |
     |<-------------------------------|
```

For example, when you open `https://example.com`, your browser sends an HTTP request. The server sends an HTTP response containing a status, headers, and usually a body.

## The first handler

Open [`http-example.go`](http-example.go). Its important function is:

```go
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}
```

| Part | Simple meaning |
| --- | --- |
| Handler | A Go function that runs when a request arrives. |
| `r` | The request sent by the client. Read its method, URL, headers, and body. |
| `w` | The response going back to the client. Write text, headers, and a status through it. |

`r` is a pointer because `http.Request` is a large struct. `w` is an interface, so it is passed directly.

## How the example works

```go
http.HandleFunc("/", handler)     // connect URL / to handler
http.ListenAndServe(":8080", nil) // start a server on port 8080
```

1. `HandleFunc` says: “when somebody requests `/`, call `handler`.”
2. `ListenAndServe` starts the server.
3. Visit `http://localhost:8080` in the browser.
4. Go calls `handler`, which writes `hello world` back.

`localhost` means “this computer.” `:8080` means port 8080. A port lets multiple network programs run on one computer.

## Run it

```bash
go run ./http/01-basics/http-example.go
```

Then open `http://localhost:8080` or run:

```bash
curl http://localhost:8080/
```

Stop the server with `Ctrl+C`.

## What comes next?

- The request has a **method** such as `GET` or `POST` → [02 Methods](../02-methods/README.md)
- It also has headers, a path, query values, and sometimes a body → [03 Requests and Responses](../03-requests-and-responses/README.md)

## First interview answer

“HTTP uses a request-response model. In Go, a handler receives an `http.ResponseWriter` to send the response and an `*http.Request` to read what the client sent.”
