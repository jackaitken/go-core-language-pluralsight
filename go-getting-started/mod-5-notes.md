# Module 5 Notes

## Working with Collections

We're going to be talking about 4 different collection types that are built into Go. There are highger level constructs that are available in third party libraries, like linked lists etc. but for now, we're going to just look at collections that are part of the standard library.

1. Arrays
    - We'll look at how we can use arrays to create a fixed size collection of similar object types. Likely though in day to day programming with Go, we wouldn't deal with arrays that often.
2. Slices
    - These are an evolution past arrays in Go. Arrays in Go are fixed size entities, whereas slices are dynamically sized. In most cases we'll be using slices for operations that in other languages we might use an array for.
3. Maps
    - Maps have a key, value relationship
4. Structs
    - In Go, we don't have the concept of classes, like we do in other languages, where we can encapsulate the behavior and actions of something. Structs represent that data side of a class. It's a fixed number of fields that have certain data types that describe one concept. So we might have a struct that defines the data associated with a user, like their name, email, phone number etc.

### Arrays

An array in Go, is a fixed size collection of similar data types. So we can have an array of integers, booleans, strings etc. All of the elements in the array must be of the same type.

```go
package main

import (
    "fmt"
)

func main() {
    var arr [3]int
    arr[0] = 1
    arr[1] = 2
    arr[2] = 3
    fmt.Println(arr)
    fmt.Println(arr[1])
    // [1 2 3]
    // 2
}
```

So here in our `main` function we declare a variable `arr` and then we explicitly set the size of the array and the data type. So we have an array of size 3, which will contain integers.

Arrays in Go are bounded so if we tried to reach for an index outside of the range that we've set, then we'll get a compile time error.

As we might expect, we can also simply create an array with the implicit initialization format. We can also initialize the elements of the array using curly braces.

```go
package main

import (
    "fmt"
)

func main() {
    arr := [3]int{1, 2, 3}
    fmt.Println(arr)
    fmt.Println(arr[1])
    // [1 2 3]
    // 2
}
```

### Slices

The previous array that we looked at can only ever be 3 elements. Its size cannot be increased or decreased. **Slices are actually built on top of arrays**. We can actually use an array to build a slice.

```go
package main

import (
    "fmt"
)

func main() {
    arr := 3[int]{1, 2, 3}

    slice := arr[:]

    fmt.Println(arr, slice)
    /// [1 2 3] [1 2 3] 
}
```

So here, we initialize an array of size 3, then below that, we're telling Go to create a slice using the `arr` array that contains all of the elements from that array. 

Now what if we were to change some of the data inside the slice and the array after they're initialized?

```go
package main

import (
    "fmt"
)

func main() {
    arr := 3[int]{1, 2, 3}

    slice := arr[:]

    arr[1] = 42
    slice[2] = 16 

    fmt.Println(arr, slice)
    /// [1 42 16] [1 42 16]
}
```

What we'll notice is that the data is changed in both `arr` and `slice`. **The reason being that the slice references the data that `arr` is referencing**. Because they're both referencing the same array in memory, and changes in one will be reflected in the other.

This can sometimes be useful, but more often than not it's going to be a bit cumbersome to create an array each time that we want to use a slice. This also limits the advantages of a slice, since we're bound to a fixed size array.

```go
    slice := []int{1, 2, 3}
    fmt.Println(slice)
    // [1 2 3]
```

So here, the only difference is that we have not specifed the size of this slice. The Go compiler will still create an array, but we've effectively told Go that we want *it* to handle the underlying array for us and worry about allocating the memory.

**So the point of all this is that a slice is not a fixed point collection**. We can show this by using the `append` function to add another element to our array.

```go
    slice := [int]{1, 2, 3}
    slice = append(slice, 4)

    fmt.Println(slice)
    // [1 2 3 4]
```

`append` takes as its first argument the slice, and then it takes one or more elements that we want to append to our slice. Here we're reassigning the `slice` variable to a slice containing the elements that we've added.

Under the hood, Go is managing the underlying array. When we first declare our slice, Go allocates memory for the size of our initialized slice, but as we start to add elements, the array will run out of room, so it will create a new underlying array, and copy the data from the previous array into the new underlying array. The slice now points to that new underlying array of data.

Let's also talk a bit more about the colon operator that we used earlier to create a slice of the array. We can also create slices of slices.

```go
    slice := []int{1, 2, 3}
    slice2 := slice[1:] // slice of `slice`, starting at index 1
    slice3 := slice[:2] // slice of `slice`, up to but not including index 2
    slice4 := slice [1:2] // slice of `slice`, from index 1 up to, but not including index 2

    fmt.Println(slice, slice2, slice3, slice4)
    // [1 2 3] [2 3] [1 2] [2]
```


