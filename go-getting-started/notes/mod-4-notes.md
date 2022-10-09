# Module 4 Notes

We're going to get started working with primitive data types. We're going to learn:

1. How to declare our variables, and learn what primitives are available in Go.
    - We'll see that with declarations, we have a few different ways to declare a variable in Go.
2. We'll learn about how Go works with pointers.
    - Pointers are definitely used in Go, but a lot of the power and potential hazard with pointers, has been stripped away in Go. This gives us the primary advantages of pointers, without the pitfalls associated.
3. Then we'll finish by talking about constants.
    - With most languages, constants are just things that we can set and forget about, but with Go, we'll see that there is a lot of subtle power that we have with constants.

## Variable Declaration and Primitive Data Types

The first method of variable declaration that we'll discuss looks something like this:

```go
var i int
i = 42
fmt.Println(i)
```

So we use the `var` keyword, then we specify the name of the variable `i`, then we give it a type `int`.

Then we can initialize that variable to a value `42`.

Then we can print that value out.

This is pretty verbose, but there are times that we need to take care of variable declaration and variable initialization in different places. This might be the case with something like a for loop.

However, we may not want to declare variables with this much specificity all of the time, so we can do this instead:

```go
var f float32 = 3.14
fmt.Println(f)
```

This allows us to set the type, and initialize it on the same line.

If we want to declare a floating point value in Go we actually have to specify the size of our float. So we have `float32`'s and `float64`'s. These represent 32 and 64 bit floating point integers.

But this is still kind of verbose. Sometimes we do need to do this, but most times we just want to declare a variable and initialize it. We can use this:

```go
firstName := "Jack"
fmt.Println(firstName)
```

Using the `:` allows Go to use the implicit initialization syntax. This means that Go is going to decide what the data type is based on the value that is assigned to the variable.

So we have three options essentially:
1. declare the variable and its type on one line, then initialize it on another
2. delcare the variable and its type *and* initialize the variable on the same line, or
3. we can use the implicity initialization syntax and let Go determine the type of our variable.

Most of the time, we're going to be using the 3rd option. Now if we were to declare one of these variables, but not actually use it, then we'll get a compile time error:

```go
firstName := "Jack"
// fmt.Println(firstName)
// error: firstName declared and not used
```

So if we have a local variable, we have to use it in some meaningful way.

Some other data types that we have available to us are:

- Booleans
```go
b := true
```

- Complex data types
```go
c := complex(3, 4)
fmt.Println(c)
// (3 + 4i)

r, i := real(c), imag(c)
fmt.Println(r, i)
// 3 4
```

This allows us to use complex mathematics to make use of real and imaginary numbers. Once we use this data type, we can also use the `real` function and the `imag` function to pull out the real and imaginary numbers. This also demonstrates Go's ability to handle multiple assignments.

## Pointer data type

The variables that we've talked about so far are called **value types**. So when we did this:

```go
i := 42
```

The variable `i` points to the number `42` in memory.

However, in Go we also have the **pointer data type** in Go. So instead of allowing the variable to reference something directly, we're going to initialize a variable to reference an address where a value is stored.

```go
var firstName *string
fmt.Println(firstName)
// <nil>
```

`nil` is Go's terminology, and it means that we have a pointer that doesn't reference anything right now. So what if we wanted to assign a string to the variable `firstName`?

```go
var firstName *string
firstName = "Jack"
// error: cannot use "Jack" (type string) as type *string in assignment
```

This won't work because we're trying to assign a string data type to pointer type. Well what if we try to *dereference* the pointer? 

```go
var firstName *string
*firstName = "Jack"
// panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x1004a4020]
```

We're derefencing the pointer by reaching through the pointer and getting the memory address stored there. We can see that attempting this gives us an error. **The reason for this error is that we're trying to assign a string to an uninitialized pointer**. Since the pointer is uninitialized, Go will not let us assign anything to that pointer. Go needs to ensure that there is a specific memory address that has been set aside to store this string, but as of yet, there is no memory address, it's currently just `nil`.

So we need to initialize the pointer by using the built in `new` function.

```go
var firstName *string = new(string)
*firstName = "Jack"
fmt.Println(firstName)
// 0x14000104210
```

So now, we see that we have an address in memory. This is memory address where the string "Jack" is being held. To get the actual value at this memory address, we need to dereference the pointer.

```go
var firstName *string = new(string)
*firstName = "Jack"
fmt.Println(*firstName)
// Jack
```

So all that we need to know when working with pointers in Go, is that to create a pointer, we precede the data type with an `*`, and then we dereference them by preceding the variable name with a `*`. The first use is a *pointer operator*, the second use of `*` is the *dereferencing operator*.

In other languages, we can do things like pointer arithmetic, but in Go, we do not have the capability to do that.

Let's look at another operator that is used alongside pointers. That is the **address of** operator.

```go
firstName := "Jack"

ptr := &firstName
fmt.Println(ptr, *ptr)
// 0x14000090210 Jack
```

So the `&` operator, is the **address of** operator. It allows us to set a pointer that references a variable that we've already created. So we initialize a variable `firstName`, then we create a pointer `ptr`, and using the address of operator we're essentially saying, create a pointer that references this value. 

Then we can see when we print out the value of `ptr`, we get the memory address where `firstName` is stored, and then we can dereference `ptr` to get the value that is stored in that address.

```go
firstName := "Jack"

ptr := &firstName
fmt.Println(ptr, *ptr)

firstName = "Daphne"
fmt.Println(ptr, *ptr)
// 0x14000090210 Jack
// 0x14000090210 Daphne
```

So even if we change the value of `firstName`, we can see that the value changes, but the memory address stays the same.

## Constants

Like other languages, the main difference between constants and normal variables, is that constants are immutable.

```go
const pi = 3.1415
pi = 1.2
// error: cannot assign to pi
```

When we declare constants, we do have to give them an initial value, much like JavaScript. The other thing to be aware of is that **the value of a constant must be able to be determined at compile time**. So attempting to assign a constant to the value of a function return will not work since functions are not evaluated until the runtime in Go.

Let's say that we do something like this:

```go
const c = 3
fmt.Println(c + 3)

fmt.Println(c + 1.2)
// 6
// 4.2
```

Now this works fine, but what is important to recognize is the way that Go is interpreting the type of `c`. When the first print statement is run, Go is interpreting `c` as an integer as we're adding an integer to the value of `c`. However, when the second print statement is running, Go will interpret `c` to be of type float, since we're adding 1.2, a floating point integer, to the value of `c`. This is what's called an **implicitly typed constant**. The compiler is going to appropriately interpret the type each time it comes across it.

However if we constrain `c` a bit more by setting its type upon initialization, we'll see an error:

```go
const c int = 3
fmt.Println(c + 3)

fmt.Println(c + 1.2)

// error: constant 1.2 truncated to integer
```

This is because we can't add a floating point integer to an integer in Go. There need to be some type of explicit type coercion in order to make that work:

```go
const c int = 3
fmt.Println(c + 3)

fmt.Println(float32(c) + 1.2)
// 6
// 4.2
```

So this is all to say that if we just need to declare a constant and don't care as much what the type is, then we can just use the implicit type syntax. If the type of the constant does matter then we can set that explicity upon initialization.

So normally, this would be the long and short of constants, but in Go, there is more than just that. Let's look at **iota's** and **constant expressions**. 

In order to work with both of these, we need to pull our constants out of the `main` function we need to put these into the package level.

```go
package main

import (
    "fmt"
)

const (
    first = 1
    second = "second"
)

func main() {
    fmt.Println(first, second)
}

// 1 second
```

What we've done here is created a **constant block**. This scopes our constants to the entire package and can be used in a similar way that we might use the `import` block.

Now what is interesting, is that we can change assignments of these constants to the keyword `iota`.

```go
package main

import (
    "fmt"
)

const (
    first = iota
    second = iota
)

func main() {
    fmt.Println(first, second)
}

// 0 1
```

So we can see that our output is `0` and `1`. Each time it is reused, `iota` will increment by 1. So this means that we can use `iota`s so that we can have constants whose values change. As it stands right now, what we have above is a *constant expression*. Using `iota` by itself is an example of a constant expression. So we could do something like this:

```go
package main

import (
    "fmt"
)

const (
    first = iota + 6
    second = iota + 2
)

func main() {
    fmt.Println(first, second)
}

// 6 3
```

So we can see here that the constant `first` was used once so its initial value: 0, was incremented by 6. The same thing happened with the constant `second`.

We can also omit the `iota` that we declared with the constant `second`, since we've already declared a constant expression, and it will use the value of `first`.

```go
package main

import (
    "fmt"
)

const (
    first = iota + 6
    second
)

func main() {
    fmt.Println(first, second)
}

// 6 7
```

The same rules that we discussed with constant also apply to constant blocks. We can't set the value of a constant to the return value of a function.

```go
package main

import (
    "fmt"
)

const (
    first = iota
    second
)

const (
    third = iota
    fourth
)

func main() {
    fmt.Println(first, second, third, fourth)
}

// 0 1 0 1
```

Here is another interesting example of `iota`s. We can declare another constant block and assign it to `iota`, and *it will reset the value of `iota`*.

