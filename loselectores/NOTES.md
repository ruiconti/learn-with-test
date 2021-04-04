# NOTES

## Basic of HTTP standard library

Ref http https://golang.org/pkg/net/http

Ref httptest: https://golang.org/pkg/net/http/httptest/#NewServer

## Gho-ul

### Defer

By using `defer`, we are simply making it explicit that what is after `defer` is the last thing
the current scope (or frame) we're in does. Hence, deferring to be done later :-)

```go
// Instead of
func ServeWWW() {
  server := http.NewServer(...)

  // 50 LOC later...

  server.Close()
}

// Instancing and cleaning are close together
func ServeWWW() {
  server := http.NewServer(...)
  defer server.Close()

  // Those same 50 LOC
}
```

Props for expressability and readability.

### Time + Channels

There are some cool utilities under `time` module. One of them uses a channel to do something after some time span:

```go
for {
  // some doing
  select {
  case <-time.After(30 * time.Second):
    fmt.Println("Timeout! Leaving.")
    break
  }
}
```

### Channels for fun

We can use channels for signaling when something happens. To do that, we must use it with selectors.

```go
func WhoHitMeFirst() string {
  select {
  case <-Hit("ALANA"):
    return "ALANA hit first"
  case <-Hit("THIAGO"):
    return "THIAGO hit first"
  }
}

func Hit(id string) {
  ch := make(chan struct{})
  go {
    time.Sleep(rand.Intn(100) * time.Millisecond)
    close(ch)
  }
  return ch
```

It is important to be aware of types' memory footprints. Since we're not really passing *content* within that channel, only a ping, a positive signal, 
it doesn't matter which type it is. Only that it is the smallest possible, which (surprisingly enough) is `struct{}`.

### Selecting

Ref, specification: https://golang.org/ref/spec#Select_statements

Ref, implementation: https://golang.org/src/runtime/select.go

Since communication on Go is done through channels, what is the proper interface to handle multiple channels?

Yes, I've heard you say: Selectors. It works as analogous to a `switch` statement. But in this case, instead of a static comparison, select will listen to channels
and selects based on whether a signal has arrived there.

Selecting is very used within `for` loops in a way to continually deal with incoming messages.

```go
in  := make(chan int)
out := make(chan int)

// writing end
for {
  go {
    time.Sleep(rand.Intn(100) * time.Millisecond)
    rand.Int() -> in
  }
}


// receiving end
for {
  select {
  case integerIn := <- in:
    fmt.Println("in:", integerIn)
    integerIn + rand.Int() -> out
  case integerOut := <- out:
    fmt.Println("out:", integerOut)
  default:
    fmt.Println("IDLE")
  }
}
```
