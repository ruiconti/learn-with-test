# Notes

## Pointers

Ref: https://gobyexample.com/pointers

In Go, when you call a function or a method the **arguments are copied**.

```go
type Entity struct {
  value   int
}

func (e Entity) SetUp {
  e.value = 10
}

func (e Entity) RunIt() int {
  return e.value + 5
}

func test() {
  a := Entity{}

  a.SetUp()  // SetUp is executed on a copy of `a`
  a.RunIt()  // Outputs 5
}
```

To actually update the struct that called the method, we need that struct's reference.

```go
func (e *Entity) SetUp
func (e Entity) RunIt
```

We didn't need to dereference that pointer in a method, as we must in C.

```c
int *a;
printf("%p", a);   // 0x0
*a = 4;
printf("%d", *a);  // 4
```

In Go, struct pointers are [automatically dereferenced](automatically dereferenced).

## Structs

We can create new types from existing ones using

```go
type BiggerInt int

b := BiggerInt(23828838300002991021)
```

## Errors

Ref: https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully

In order to properly return errors, `import errors`. It is a non-special struct:

```go
func (a ABC) Quack(arg int) error {
  if arg > 10 {
    return error.New("Invalid integer")
  }
  return true 
}
```

So if a function might raise an error, you should *always* attempt to get it and check if it's
an error. Which is actually a little awkward.

Pointers can also be `nil`. That doesn't smell good as well. 
