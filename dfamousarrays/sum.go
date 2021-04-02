package arrays

/*
	Ref: https://blog.golang.org/go-slices-usage-and-internals
	Go's arrays are values. An array variable denotes the entire array; it is not a pointer to the first array element (as would be the case in C). This means that when you assign or pass around an array value you will make a copy of its contents. (To avoid the copy you could pass a pointer to the array, but then that's a pointer to an array, not an array.) One way to think about arrays is as a sort of struct but with indexed rather than named fields: a fixed-size composite value
*/

/*
	range iterates over elements in a data structure
	it provides an index and a value for each entry
*/

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, slice := range numbersToSum {
		sums = append(sums, Sum(slice))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, slice := range numbersToSum {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(slice[1:]))
		}
	}
	return
}

/*
	Slices do have a capacity i.e trying to access an index
	greater than a slice capacity will "throw" a runtime error
	It's important to note that these "unsafe" accesses do compile
	but are caught only on runtime
	A slice is a descriptor of an array segment.
	It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).

	Slices behaves "like" python lists in a sense that they can be
	appended

	make([]T, N) creates a slice of type T of size N
*/
