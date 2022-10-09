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
## Maps

Maps allow us to store keys with value types. With them, since we're working with two data types, we need to be a bit more verbose in our initialization of maps. Here is the short declaration syntax:

```go
m := map[string]int{"foo": 42}
fmt.Println(m)
fmt.Println(m["foo"])

m["foo"] = 41
fmt.Println(m["foo"])

delete(m, "foo")
fmt.Println(m)

// map[foo:42]
// 42
// 41
// map[]
```

With maps, we declare them using the `map` keyword. Then next to the `map` keyword we use square brackets to define our data type for our keys. Then next to that we set the data type of our values. So in this case above we're mapping strings to integers, strings being our keys, and integers being our values. Next to that we can also initialize our map with some information.

Below that we printed out our map, and then we retrieve a specific value by passing in our key.

After that we reassociated the key `foo` to another number, and printed that out as well.

Then finally we use the `delete` function to delete a key, value pair from our map. `delete` takes two arguments, the map, and then the key that we want to delete.

## Structs

Struct's are the last collection type that we'll look at, and they're unique in that we're able to associate disparate data types together. With arrays and slices, our values had to be of a certain type. With maps, our keys had to be of one type, and our values had to be of one type. With a struct we can associate any type of data together, but the fields that we set in our struct are fixed at compile time and cannot be changed later in our program.

With arrays and slices, we were able to define the types and initialize the values at the same time. With structs we don't have that same kind of flexibility. We need to define the struct in one step and then initialize it in a second step.

```go
package main

import (
    "fmt"
)

func main() {
    type user struct {
        ID int
        FirstName string
        LastName string
    }

    var u user
    fmt.Println(u)

    u.ID = 1
    u.FirstName = "Jack"
    u.LastName = "Aitken"
    fmt.Println(u)

    fmt.Println(u.FirstName)

    u2 := user{ ID: 1, 
        FirstName: "Jack", 
        LastName: "Aitken",
    }
    fmt.Println(u2)

    // {0  }
    // {1 Jack Aitken}
    // Jack
    // {1 Jack Aitken}
}

```

Here we've defined a type called `user` and the type of `user` is a struct. In our first print, we get a pretty boring result. This output indicates that our struct does exist. We have a `0`, and two blank spaces.

What happens here is that when we initialize the variable `u`, each field is initialized to what's called its **zero value**. The zero value for an integer is a 0, the zero value for a string is the empty string. So our first print here, prints the struct, with the zero values for our fields.

After that we use the dot syntax to set the values in the fields that we've defined, so that after, our fields have been initialized to some value. We can see that in our second print.

We can also use the dot to retrieve specific values from our struct, like maps.

Now there is a shorter syntax that we can use to initialize our structs once they've been defined. We can see an example of this above. We can see that we have a trailing comma after our `LastName` field. This needs to be there because of Go' automatic colon insertion. With this multiline initialization syntax, if we don't have the trailing comma, Go will assume our last field (`LastName`), is the end of the statement, and it will add a semi-colon, so we'll get a compile time error:

```go
u2 := user{ ID: 1,
    FirstName: "Jack",
    LastName: "Aitken"
    }

fmt.Println(u2)

// syntax error: unexpected newline, expecting comma or }
```

*A note on scoping*: In the previous example, we declared our struct in the `main` function. That means that we can only create objects of that type in that function. If we wanted to declare something outside of the main function, then we could also move it up to the package level.

## Building out a web service

Okay now we're going to turn our attention to a web service that we're building in our `main.go` file.

One important piece of language: Everything within a directory that holds our `go.mod` file is called a **module**, and within a module we have **packages**, which are components that help to make up some concept within our module.

Inside of our module, we created a package called `models`. Again, a package is just a directory. Inside of this directory, we created a file called `user.go` and specified the package name as `models`. It's important that the package name is the same as our directory name. 

```go
// models/user.go

package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users []*User
    nextID = 1
)
```

So first we named our package `models`, and then we defined a struct called `User` at the package level. After that we declared some variables in a block.

The first variable that we created is called `users`. It is a slice that contains pointers to `User` objects. The reason to use pointers, is that we're going to be able to manipulate these user objects throughout our program without having to copy that data in multiple places. 

Then we declare a variable called `nextID`. This will serve the purpose of something like a primary key since we're not going to be using a database.

Back in our `main.go` file we're going to use our `User` struct to create an object.

So now here is what our `main.go` file looks like:

```go
package main

import (
	"fmt"

	"github.com/jackaitken/go-core-language-pluralsight/go-getting-started/models/go-getting-started/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Arlo",
		LastName:  "Aitken",
	}

	fmt.Println(u)
}
```

We import our models package, and then we're able to use it in our `main` package.

