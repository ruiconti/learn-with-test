# Notes

## Goroutines 

Ref: https://gopherlabs.kubedaily.com/Concurrency/goroutines-explained.html
Ref: https://medium.com/rungo/anatomy-of-goroutines-in-go-concurrency-in-go-a4cb9272ff88
Ref: https://appliedgo.net/concurrencyslower/


1. Go implements concurrency execution with Goroutines. This is somewhat *conceptually* to a thread execution, but **way way way** cheaper. Goroutines are [really lightweight](https://medium.com/the-polyglot-programmer/what-are-goroutines-and-how-do-they-actually-work-f2a734f6f991).
2. Usually is done using as an argument anonymous functions.

### Behavior

Goroutines have a "similar" approach to the way forks works (Thompsonish?):

```go
for _, url := range(urls) {
  go func() {
    m[url] = parse(url)
  }
}
```

Every goroutine will write to the same `m[url]`. Hence, after running this, our map `m` will hold only the last value of `urls` slice.

## Channels

Ref: https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb
Ref: https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html
Ref: https://inconshreveable.com/07-08-2014/principles-of-designing-go-apis-with-channels/

Channels are a communication facility. Again, another analogy to UNIX. They're pretty similar to pipes. No, really, they really are, they use the same underlying principle: [CSP](https://swtch.com/~rsc/thread/).

Channels also resembles how Python's coroutines work in a way they send and receive values in an asynchronous fashion.

```go
ch := make(chan string)

ch <- "BUFFERING..."
read := <- ch
fmt.Println(read)   // BUFFERING...
another_read <- ch  // Blocks until something is written on ch
```

- They always have a type to signal which interface is being buffered (and ofc its size): `c chan int`
- It is used `c <- 4` to **send** data
- And it is used `b := <- c` to read data

## Anonymous Functions

Familiar to Python's lambdas? That's that.

```go
b := 5
func() {
  a := 3
  fmt.Printf("%d %d", a, b)
}()  // 3 5
```

- Executed as they're declared.
- They're compliant to closures i.e. can access the lexical scope in which they're declared.


## Goish

#### Benchmarking

Remember, we can benchmark our code :-) 

```bash
go test -bench=.
```

#### Race condition detector

Ref: https://blog.golang.org/race-detector

Wow. This is impressive. There's a native tool to detect race conditions at virtual addresses level.

This is awesome:

```bash
go test -race
```

