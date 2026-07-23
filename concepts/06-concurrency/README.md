# 06. Concurrency: Doing Work at the Same Time

A goroutine is a function running concurrently with other Go code.

```go
go fmt.Println("runs in another goroutine")
```

The main function can finish before that goroutine gets a chance to run. Use a `WaitGroup` when you need to wait for known work.

```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	fmt.Println("finished")
}()
wg.Wait()
```

## Channels: goroutines talking to each other

A channel lets one goroutine send a value and another receive it.

```go
ch := make(chan int)
go func() { ch <- 42 }()

value := <-ch
fmt.Println(value)
```

An unbuffered channel needs a sender and receiver to meet. A buffered channel can hold a limited number of values first:

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
```

## Closing and ranging

The sender that knows no more values will be sent closes the channel:

```go
close(ch)
```

`for value := range ch` stops only after the channel is closed. Receivers normally do not close channels, because a sender could panic by sending after close.

## Shared values need protection

This is unsafe when several goroutines run it:

```go
count++ // read, add, write: not one atomic step
```

Use a mutex for shared data:

```go
mu.Lock()
count++
mu.Unlock()
```

Use `WaitGroup` to wait; use `Mutex` to protect data. They solve different problems.

## Which tool should I choose?

| Need | Start with |
| --- | --- |
| Wait for known goroutines | `sync.WaitGroup` |
| Send jobs or results between goroutines | Channel |
| Protect one shared value or data structure | `sync.Mutex` |
| Stop work after timeout/cancellation | `context.Context` |

## Rules and traps

- Do not use `time.Sleep` to wait for work; it guesses instead of synchronizing.
- Call `wg.Add` before starting the goroutine.
- Go 1.22+ gives each `:=` loop iteration a new loop variable. Passing the value as a function argument is still very clear.
- Run `go test -race ./...` to find data races.

## Interview answer

“Goroutines run concurrent work. I use a `WaitGroup` to wait, channels to pass values or jobs, and a mutex for shared mutable state. I use the race detector to verify concurrent code.”

## Run

```bash
go run ./concepts/06-concurrency
```
