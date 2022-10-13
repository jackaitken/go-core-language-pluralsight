# Controlling Program Flow

We're going to talk now about controlling our program with loops and conditionals.

In Go we only have for loops, there are no while loops, but we can using for loops to accomplish the same thing that while loops do in other languages.

Then we'll talk about branching including:
- `if` statements
- panics, and
- switch statements

## Types of Loops

Every loop is a for loop, but in order to change the behavior, we can use 1 of 4 different types of for loops.

1. Loop till condition
2. Loop till condition with post clause
3. Infinite loops
4. Looping over collections

## Looping till condition

Here is what this looks like:

```go
var i int
for i < 5 {
    println(i)
    i++
} 

// 0
// 1
// 2
// 3
// 4
```

A note to make here is that we used the built in `println` function. This is not part of the `fmt` package, and is less sophisticated, and normally only used for debugging.

To break out of a loop early we can use the `break` statement.

```go
var i int
for i < 5 {
    println(i)
    i++
    if i == 3 {
        break
    }
} 

// 0
// 1
// 2
```

Related to the `break` statement is the `continue` statement:

```go
var i int 
for i < 5 {
    println(i)
    i++
    if i == 3 {
        continue
    }
    println("continuing...")
}
// 0
// continuing...
// 1 
// continuing...
// 2
// 3
// continuing...
// 4
// continuing...
```

A continue is similar to the `break` except it only breaks one iteration of the loop. After it hits a `continue`, it will leave the process and start over.

## Loop till condition with a post clause

With this variation of a for loop we can add a statement into our for loop clause that will be executed everytime after our loop completes an iteration.

```go
var i int
for i < 5; i++ {
    println(i)
}
```

So here, we've created a post clause that will handle the incrementation for us. Something big to note here is that when we use a post clause like this, we have to have 3 terms in the for loop. This indicates to the compiler that we're using a loop till condition with a post clause. So this code above will not compile. Instead we can use that third term to initialize our variable:

```go
for i := 0; i < 5; i++ {
    println(i)
}
```
Like in JavaScript, the variable `i` here is scoped to that block, so it can't be used elsewhere

## Infinite loops

There are a couple of ways to create an infinite loops in Go. There's a pretty way, and then there's an ugly way. Let's look at the ugly way:

```go
var i int 
for ; ; {
    if i == 5 {
        break
    }
    println(i)
    i++
}

// 0
// 1
// 2
// 3
// 4
```

So now technically this is an infinite loop. We pulled out the test term, the initializing term, and the incrementing term, but technically, this is an infinite loop.

Let's look at another approach, which is a bit prettier:

```go
var i int
for {
    if i == 5 {
        break
    }
    println(i)
    i++
}
```

This is allowed in go, where we simply omit the semi-colons for a much prettier looking infinite loop.

## Looping over a collection

We of course could use a normal for loop iterate for the length of the array: 

```go
slice := []int{1, 2, 3}
for i := 0; i < len(slice); i++ {
    println(slice[i])
}

// 0
// 1
// 2
```

But we could also use a different syntax, which is a bit more expressive:

```go
slice := []int{1, 2, 3}
for i, v := range slice {
    println(i, v)
}

// 0 1
// 1 2
// 2 3
```

So we've added the `range` keyword. What this tells the compiler, is that we're going to be passing in a collection type, and it will return out to us two variables:
1. the index, and
2. the value at that index

We can also use `range` with maps:

```go
wellKnownPorts := map[string]int{"http": 80, "https": 443}
for k, v := range wellKnownPorts {
    println(k, v)
}
// http 80
// https 443
```

So now what if we don't need one of these values? If we only need to loop over the keys, they we can eliminate the `v`:

```go
for k := range(wellKnownPorts)
```

Normally ignoring a second return value is something that we're not allowed to do, but this is a special case with the compiler because it knows about range, and allows us to to ignore it.

If we only need the values, then we do this:

```go
for _, v := range wellKnownPorts
```

So here we use a write only variable: `_`. This means that we can't actually read from it, so the value is essentially ignored. We need to use this of course because if we had a second variable that we didn't use we would get a compile time error.

## Branching constructs

## panics

First we're going to talk about `panic`'s. Panic's are similar to exceptions in other languages. A panic will happen when our program has no idea how to proceed, in this case it will create a panic. Normally we can return errors from functions so that the calling function knows how to handle it, and can do something with that error.

Eventually, though our application may hit a state where it cannot proceed at all. If for example we're relying on a database connection for an application, and the connection cannot be established, then a panic would be raised.


```go
func main() {
    println("starting web server")

    panic("could not start server")

    println("server started")
}

// panic: could not start server
```

`panic` is a built-in function, which we can pass a string to, which when triggered, will exit the program entirely. We'll also get some further information that will show us where the panic occured. Before the panic occurs, we'll also get an output from Go vet that will tell us that since we have a hardcoded panic, the line `println("server started")` is unreachable.

Now it is possible to recover from panics, but the scope of that is beyond this course.

## if statement

```go
type User struct {
    ID          int
    FirstName   string
    LastName    string
}

func main() {
    u1 := User{
        ID:         1,
        FirstName:  "Arthur",
        LastName:   "Dent",
    }

    u2 := User{
        ID:         2, 
        FirstName:  "Ford",
        LastName:   "Prefect",
    }

    if u1 == u2 {
        println("Object is the same")
    } else if u1.FirstName == u2.FirstName {
        println("Same first names")
    } else {
        println("Object is not the same")
    }
}
```

Here we check that user 1 is equal to user 2. We can also check for inequality using `!=`. Since structs are value types in go we can check for direct equality.

Now that we understand the basic concepts of branching, we can add those to our web server program.
## Back to our web service

Inside of our models package, we have our `users.go` file. We added some functions, so here's how it looks now with some notes to describe what's happening.

```go
package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

/*
AddUser checks first that the user passed in is set to it's zero value
if not then return an error. We also return an empty User struct as we must
have a second return value here.
*/
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new user must not include id")
	}
 	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

/*
GetUserById takes an id as argument, iterates over the users slice
and returns the user if we find it and nil for the error. 
If it doesn't then it returns a formatted error string
*/
func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	
	return User{}, fmt.Errorf("User ID '%v' not found", id)
} 

/*
UpdateUser takes as argument a user, then it will iterate 
through our list of users. At each user,
it will check if a user in the list has the same id as the user passed
in. If so, then it will reassign the user at that index to the address of the new user
struct passed in. Then returns the new user. If we don't find a user, then
it will return an error.
*/
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
} 

/*
RemoveUserById takes an id as argument and then will iterate through the 
users list. It will check the id's and if we get a match, it will remove that
user from the list. It does this by reassigning the users list to a new list 
which is a slice from everything up to i, and everything starting after i. 
It combines these two and reassigns users.
*/
func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' not found", id)
}
```

## Switch statements

Let's build a scenario in which we could use a switch statement:

```go
package main

type HTTPRequest struct {
	Method string
}

func main()  {
    r := HTTPRequest{ Method: "GET" }
	
	switch r.Method {
	case "GET": 
		    println("Get request")
	case "DELETE":
		    println("Delete request")
	case "POST":
		    println("Post request")
	case "PUT":
		    println("Put request")
	default:
		    println("Unhandled method")
    }
}
```

Here is an example of a switch statement. Now this isn't valid Go 
code, but it's just there to show what we can do with switch 
statements. Something to mention is that unlike JavaScript, Go 
does not have implicit fallthroughs, which means that we don't 
need to provide break statements. Instead, Go has implicit breaks. 

We can also use the `default` keyword to define a statement that 
will run if no condition is met.

Now let's add some switch statements into our web service.

```go
package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller"))
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

```

Now here is what our code looked like previously. Whenever we used 
our web service, it would just hand back a static response because 
we had no way to direct traffic. 

Here is how we're going to adjust this code:

