# NOTES

## Reflection

Ref: https://blog.golang.org/laws-of-reflection

Ref: https://pkg.go.dev/reflect?utm_source=godoc

Reflection in computing is the ability of a program to examine its own structure, particularly through types; it's a form of metaprogramming. It's also a great source of confusion.

The typical use is to take a value with static type `interface{}` and extract its dynamic type information by calling
`TypeOf`, which returns a `Type`.

In short: it "allows" duck typing static typing. The quotes are not mistaken. This ought to be used very carefully and consciously.

## Channel iteration

A neat idiom to consume a dynamically inspected channel:

```go
val := reflectValueOf(x)
if val.Kind() == reflect.Ptr {
  val = val.Elem()
}

for v, ok := val.Recv(); ok; v, ok = val.Recv() {
  doSomething(v.Interface())
}
```


## Function returns

Ref: https://pkg.go.dev/reflect?utm_source=godoc#Value.Call

We can return an arbitrarily-lenghty output from an inspected function:

```go
fn := func() int int string {
  return 4, 9, "Dio"
}

val := reflect.ValueOf(fn)
result := val.Call(nil)

for _, res := range result {
  switch res.Kind() {
  case reflect.Int:
    fmt.Println(res.Int())
  case reflect.String:
    fmt.Println(res.String())
  }
}
```
