# NOTES

## Golang

We can use `...` as a way to pass zero or more arguments to a function. A function that
accepts this sort of parameter is called *variadic* function.

It is somewhat analogous to Python's `*` expression.

```go
func PrintVariadic(letters ...byte) {
  fmt.Println(letters)
}

func main() {
  anArr := []byte{'X', 'S', 'V', 'I'}
  fmt.Println(anArr)       // [88 83 86 73]
  PrintVariadic(anArr...)  // [88 83 86 73]
```

## Property based testing

Ref: https://increment.com/testing/in-praise-of-property-based-testing/

Property based tests are based on domain axioms. Turns out that deriving axioms on 
computer programs turns out to be rather difficult. It requires a deeper understanding of
the problem-space and its relationships.

But doing so makes us more confident that the behaviour is properly implemented.
They're often about input-output space boundaries.

