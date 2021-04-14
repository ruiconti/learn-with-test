# Notes

Very relevant reference: https://research.swtch.com/interfaces

## Methods

Ref: https://golang.org/ref/spec#Method_declarations

Method is a function with a receiver. The receiver is the instance on which this function is being called.

Syntax is:

```go
func (receiverName ReceiverType) MethodName(args) ReturnType {
  // body
}
```

It's interesting that Go doesn't "clutter" methods inside a single entity like Python:

```python
class Entity:
  a: int
  
  def calculate(tax):
    return a * 100 * tax
```

Whilst in Go:

```go
type Entity {
  A  int
}

func (e Entity) Calculate(tax int) int {
  return e.A * tax * 100
}
```

Which I must be honest that looks cleaner and more straightforward. Although less visually intuitive.

But it's up to programmer to properly organize and keep things tidy and modules cohesive.

#### Convention

The receiver is always referred to Struct's initial lowercase letter:

```go
func (e Entity) Calculate(a int[][])
func (r Rectangle) Area()
func (s Server) Serve(port int)
...
```

## Interfaces

This is a very powerful feature that requires minimum verbatim because it is something akin to a 
typed duck-typing:

```go
func (r Rectangle) Area() float64
func (c Circle) Area() float64

type Shape interface {
  Area() float64
}
```

In this examples, `Shape` interface will implicitly resolves both `Circle` and `Rectangle`.

Go does this by implementing [parametric polymorphism](https://en.wikipedia.org/wiki/Parametric_polymorphism).

## Side-notes

#### testing.Helper

[testing.Helper](https://golang.org/pkg/testing/#B.Helper) marks the calling function (which is calling Helper()) as a test helper function.

That way, when printing file and line information, that helper function will be skipped.

#### Table driven tests

Table based tests can be done using "anonymous struct":

```go
areaTests := []struct {
    shape Shape
    want  float64
}{
    {Rectangle{12, 6}, 72.0},
    {Circle{10}, 314.1592653589793},
}

for _, st := range areaTests {
    got := st.shape.Area()
    if got != st.want {
        t.Errorf("got %.2f want %.2f", got, st.want)
    }
}
```
