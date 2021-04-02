# Notes

## Dependency Injection

Ref: https://blog.drewolson.org/dependency-injection-in-go
Ref:  https://elliotchance.medium.com/a-new-simpler-way-to-do-dependency-injection-in-go-9e191bef50d5
Ref: Go Programming Language book

There is a neat way to test side-effects behavior, such as the output of a `fmt.Printf`.

It is by using dependency injection and referring to dependency virtual memory address.

```go
func PrintSomething(buff *bytes.Buffer) {
  fmt.Fprintf(buff, "WRITTEN")
}

buffer := bytes.Buffer{}
PrintSomething(&buffer)
buffer.String()  // WRITTEN
```

This is a convenient way to manage and have control of rather dangerous side-effects.


### Efficient io interfaces

It is stated that Go's interfaces are a rather powerful abstraction. Having that said
Pike and Thompson clearly tried to replicate file descriptor's awesome design. In a sense that
by using `io.Reader` and `io.Writer`, we hide implementation details of about **exactly what** we
are actually reading and writing to.

It is an approach that leads to more reusable and extensible code. It also defines a broad and easy to reason
interface between different modules and libraries.


