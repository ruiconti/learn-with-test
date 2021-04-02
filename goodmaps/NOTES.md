# Notes

## Maps

Ref: https://blog.golang.org/maps

### Declaration

```go
dict = map[string]int;
```

Key types [must be of a comparable type](https://golang.org/ref/spec#Comparison_operators).

**Gotchas!**

Even though we might be tempted to initialize an empty map, it is not a very good idea. Mainly because maps can 
hold `nil` values. And we might find ourselves trying to use a map that hasn't been properly initialized.

So, what's the Gopher way to initialize a data structure? 

```go
var dict = make(map[string]string)
var dict = map[string]string{}
```

Yes, by its de-facto generic data type factory or with an empty braces declaration.


### Behavior

#### Check if item exists

In order to check if a given key exists on a given map, none but the Gopher way: Wait for it in a tuple result :-)

```go
result, ok := dict["chave"]
if !ok {
  return "", errors.New("word not in map!")
}
```

The second value returns a boolean indicating if querying was successful.

### Extended structs of maps

Consider a type that extends a dictionary:

```go
type NewDict map[string]string
```

If we run

```go
d := NewDict{"a": "b"}

func (n NewDict) Update(k, v string) {
  n[k] = v
}

d.Update("c", "d")
d  // outputs {"a": "b", "c": "d"}
```

Note that no pointer was passed to `Update` method. However, `d` reference was updated as if a pointer to d was passed.

This behavior is thoroughly explained [here](https://dave.cheney.net/2021/01/05/a-few-bytes-here-a-few-there-pretty-soon-youre-talking-real-memory). Excelent resource.


## Errors

Ref: https://dave.cheney.net/2016/04/07/constant-errors

It really makes sense to write up our own errors. Why? To make 'em more reusable and immutable. 

There is a powerful way to extend `error` interfaces without, in fact, extending `error` type. We just need to adhere to its polymorphism.

```go
type SpecificErr string

func (e SpecificErr) Error() {
  return string(e)
}
```

"Automagically" makes an interface compliant to `error.Error()`. Beautiful.
