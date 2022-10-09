# Module 6 Notes

In this module we're going to explore how to create functions and methods. Then we'll also explore interfaces to introduces polymorphism into our programs.

Let's take a closer look at our `main` function:

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, World!")
}
```

We see that we begin our function declaration with the `func` keyword. That's the function declaration for our main function. Then comes the name of the function. `main` in particular is a special name in Go, as it describes an entry point, but we can use any valid symbol as a function name in Go.

Then there are the `()` where we'll house parameters of course, and an opening and closing curly braces. In Go, curly braces are always required, even with single line functions.

Let's create our own function that simulates starting our web server.

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, World!")
    startWebServer()
}

func startWebServer() {
	fmt.Println("Starting server...")
	// do server things
	fmt.Println("Server started")
}
```

So this function that we've created is really a simulation of what our web server is meant to do. At this point our function is not being invoked, so to do that, we'll call the function with `startWebServer()`.

Now we can pass data into the function to actually use it.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	port := 3000
    retries := 2
	startWebServer(port, retries)
}

func startWebServer(port, numberOfRetries int) {
	fmt.Printf("Starting server on port %d\n", port)
	// do server things
	fmt.Println("Server started")
}
```

Here' we've passed the port number into our function. Inside of the function declaration we've defined a parameter called `port` and also specified the type of the parameter as type `int`.

We can also see that we've specifed a second parameter here called `numberOfRetries`. Because both of our parameters are of type `int`, we only need to declare that once.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	port := 3000
    retries := 2
	startWebServer(port, retries)
}

func startWebServer(port, numberOfRetries int) {
    return "Server succesfully started on port %d", port
}
```

To return data from our function we need to specify what type of data to return from our function, we need to specify the return type of our function.

```go
package main

import (
	"fmt"
)

func main() {
		fmt.Println("Hello, World!")
		port := 3000
		isStarted := startWebServer(port)
}

func startWebServer(port int) bool {
		fmt.Println("Starting server...")
		fmt.Println("Server started on port", port)
		return true
}
```

So here we're returning a boolean from the function `startWebServer`. But there is something to point out here. Returning a boolean from a function doesn't really do much. Returning true won't help us if something went wrong. Instead, what would be more valuable is to return an error from this function, then we could determine what went wrong. So here is a better idea:

```go
package main

import (
	"fmt"
)

func main() {
		fmt.Println("Hello, World!")
		port := 3000
		err := startWebServer(port)
		fmt.Println(err)
}

func startWebServer(port int) error {
		fmt.Println("Starting server...")
		fmt.Println("Server started on port", port)
		return nil
}

// nil
```

So we're going to see this paradigm a lot in Go, where some functions will return an error, then the calling function can decide what to do with that error. We don't throw a lot of exceptions in Go, we instead return error values, because those can be just as valid as a successful function.

Now above, we returned `nil`, but that also isn't too idomatic. So instead what we'll do is this:

```go
package main

import (
	"fmt"
	"errors"
)

func main() {
		fmt.Println("Hello, World!")
		port := 3000
		err := startWebServer(port)
		fmt.Println(err)
}

func startWebServer(port int) error {
		fmt.Println("Starting server...")
		fmt.Println("Server started on port", port)
		return errors.New("SOMETHING WENT WRONG")
}

// SOMETHING WENT WRONG
```

Now what if we want to install multiple pieces of information from a function?

```go
package main

import (
	"fmt"
	"errors"
)

func main() {
		fmt.Println("Hello, World!")
		port := 3000
		p, err := startWebServer(port)
}

func startWebServer(port int) (int, error) {
		fmt.Println("Starting server...")
		fmt.Println("Server started on port", port)
		return port, errors.New("SOMETHING WENT WRONG")
}
```

So here we have a variable `p`, but as we know, Go is going to throw an error if we have an unused variable. Luckily, Go allows us to use something called a write only variable. To use a write-only variable, we can use an `_`.

```go
package main

import (
	"fmt"
	"errors"
)

func main() {
		fmt.Println("Hello, World!")
		port := 3000
		_, err := startWebServer(port)
		fmt.Println(err)
}

func startWebServer(port int) (int, error) {
		fmt.Println("Starting server...")
		fmt.Println("Server started on port", port)
		return port, errors.New("SOMETHING WENT WRONG")
}
```
Now that we have this, let's start building out our models to give it some more behavior.

```go
// models/user.go

package models

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

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
```

Here is what our `models` package looks like now. We have a struct called `User`, which defines a user, then we have a variable called `users`, which is a slice that contains pointers to our User objects. Using a pointer here will allow us to actually manipulate the User objects in memory. We also have a `nextID` variable that we'll increment.

Then we have a function `GetUsers`, which will return our users slice.

Then we have a function `AddUser`, which takes as argument a User object, and then returns a new user. It also handles incrementing the nextID variable.

## Object oriented Go

Let's try and add some behaviors to our custom data types. Let's add a package called `controllers`. We need some way for users of our webservice to actually work with the users that we've created.

Here is what our controller package looks like:

```go
package controllers

import "net/http"

type userController struct {}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller"))
}
```

So we have a struct called `userController`, which does not have any fields in it, then we also have a method below. This is how we create methods in Go. Before the name of the function, we wrote `uc userController`. What this does is attach the function to the `userController` struct. 

The function is named `ServeHTTP` and it takes as argument an http response write and an http request. For now we're just going to write to the writer, a string coerced to bytes.

For right now our `userController` struct is empty. This can be helpful to do because we can use this empty struct to associate related behaviors by tying them to this struct, even if we're not storing data inside that object.

We are actually are going to add some fields inside of that struct however.

This `userController` will eventually have some routing responsibilities. It will handle two resource requests:
1. Get all of the users that we have inside of the `users` collection, and
2. Manipulate individual users.

Now in order to actually figure out which request it's working with we're going to use a regular expression. In order to do this, we're going to add a field called `userIDPattern`

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
```

Now as it stands, someone else could pass a regular expression in and be able to control how we use this controller, which isn't good. We want to strictly define how a user can use this controller.

Now, in order for someone to actually use this `userController` object, we need to provide them with a contstructor function.

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

}
```

So here is our constructor function. It is a Go convention to use `new` before the name of the function to signal to other users, that this is a constructor function.

Now we return a pointer here, because this contstructor function is just setting up the object for someone else to use, so we want it to reference the object that has been created to avoid copy operations.

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

So here is the implementation of our `newUserController` constructor function. We use a regular expression to locate a path that is something like `/users/:id`. It's going to ensure that the path has `/users` followed by some number. 

Something that's interesting, is that we're not returning the actual `userController`, we're returning the address of that userController object.

So at this point, the userController is struct, and we have a constructor function that actually allows us to create a new user controller.

## Interfaces

So now we want to take our http requests and route it to the userController that we've created. Let's talk about the `Handler` type. For right now we can ignore the regular expression, because at this stage it won't be doing anything. We don't have any logic at the current time that can handle a number a path like `/users/12`.

The `Handler` type in the `net/http` package defines an interface that takes as argument: `http.ResponseWriter`, and `*http.Request`. A handler responds to an HTTP request. So we created a method called `ServeHTTP`, which takes those two arguments, and satifies the `Handler` type.

Now that we have this handler configured, let's create another file that will handle the actual routing. This will be done in `front.go`.

```go
package controllers

import "net/http"

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
```

So this is in the controllers package as well, so it has access to the `newUserController` constructor function. Then it defines two routes:
1. `/users`, and
2. `/users/`

We also pass in as the second argument, the pointer to the controller that we set up as the handler. Since our handler implements the Handler type, it will be able to successfully take these http requests and respond to them.

Now finally, in our `main.go` file, where the program will kick off, we have this:

```go
package main

import (
	"net/http"

	"github.com/jackaitken/go-core-language-pluralsight/go-getting-started/models/go-getting-started/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
```

So we import the controllers package, and then call register controllers. This creates our handlers, and then once that's done we can start listening for requests.

In `ListenAndServe`, we just provide the port, and then the second argument is for the ServeMultiplexer, this the object that handles all of the requests coming in and it handles the high level routing. We're going to just use the default Serve Mux.

With all of this completed, we can start our server and see `Hello from the User Controller`.

