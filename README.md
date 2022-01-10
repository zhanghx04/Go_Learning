# Go_Learning

## basic
file extension : .go

to run:
    go run file.go

## how to decalre
### int or char
```go
var x int = 5
// or
x := 5

//Note: if use :=, then Go can automatically figure out the type
```

### array - 0 based
```go
var a [5]int
the initialized value in a is [0, 0, 0, 0, 0]
// or
a := [5]int {1, 2, 3, 4, 5}

// change value with index
a[2] = 8

// append
a = append(a, 15)
```

### map
```go
    // map[key_type]value_type
    vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3

    // to delete
    delete(vertices, "triangle")

	fmt.Println(vertices)
```

### for loop
```go
    for i := 0; i < 5; i++ {
		fmt.Println(i)
	}


    arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}


    m := make(map[string]string)
	m["a"] = "apple"
	m["b"] = "banana"

	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}
```

### create another function
```go
package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result := sum(4, 6)
	fmt.Println(result)

	res, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func sum(x int, y int) int {
	return x + y
}

// error here means sqrt function could print out error
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined for negative numbers")
	}
	// nil here is like null but not like null, zero values
	return math.Sqrt(x), nil
}

```

### struct
```go
package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Jake", age: 23}
	fmt.Println(p)
	fmt.Println(p.name)
}
```

### pointer
When pass the pointer to a function, we need to send the address to the function. Also, in the function, we need to add a * in front of the pointer variable
```go
package main

import (
	"fmt"
)

func main() {
	i := 7
	fmt.Println(i)
	fmt.Println(&i)

	inc(&i)
	fmt.Println(i)
}

func inc(x *int) {
	*x++
}
```