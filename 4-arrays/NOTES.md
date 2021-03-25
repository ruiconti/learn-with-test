# Notes

## Arrays 

- Have fixed size and must be passed within variable declaration.
- That's why they're not commonly used.
- However it's a useful and needed primitive that enables slices implementation.

```go
var a [4]int  // define an array that can hold 4 integers 
a[0] = 10
fmt.Printf("%d\n", a[1])  // 0
```

- Arrays are initialized with "type's zero" value. In an integer, zeroed.

## Slices

Slices can be arbitrarily long.

Declaration syntax is 

```go
a := []string{"this", "is", "a", "slice"}
b := [][]int{{0, 1}, {1, 0}}
```

### Slicing

Can also be formed by "slicing" an existing slice or array. It is done by 
specifying an half-open range with two indices.

Half-open range can be visualized as follows, considering an x-axis:

---|-----[o]======[ ]-----> x

Where [o] means an inclusive range edge, and [ ] means a non-inclusive range edge.

The start and end range are optional and default to both 0 and the length of the slice. 

**Slicing does not copy the slice's data. It creates a new slice that points to the original array.**

Therefore, it makes slice operations as efficient as manipulating array indices.

### Capacity

...

## Go syntax

### Variadic functions

Functions that can be called with an arbitrary number of trailing arguments.

Here it's also illustrated how to iterate over slices with `range`:

```go
func sum(nums ...int) {
  total := 0

  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
```

It's analogous to python's * function argument:

```python
def sum(*nums):
  total = 0
  for num in nums:
    total += num
  return total
```
