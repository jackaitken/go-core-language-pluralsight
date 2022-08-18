# Module 4

Let's explore a little bit more about what it is like to use the Go language. Instead of talking about theory and going through slides, we're going to look at a couple of demos. First, we're going to look at a command line application. This will be located in the `go-cli` directory.

Then we're going to look at what it's like to create web services with Go. 

## CLI applicaton

The goal with this app is to create something that will take our logs and filter out all of the log messages except the ones that say `ERROR`.

Here is the first stage of our app:

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("myapp.log")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	r := bufio.NewReader(f)

	for {
		s, err := r.ReadString(('\n'))
		if err != nil {
			break
		}
		if strings.Contains(s, "ERROR") {
			fmt.Println(s)
		}
	}
}
```

First we declare our package name, which is `main`. Then we import the packages that we need. 

Then we enter the `main` function. First we use the `os` package and the `Open` method available on the package to open our logs. `Open` returns two values, the file (`f`), and an error handler (`err`). 

`err` is a pointer value, so we use `nil` to check if this variable is non-null. If it is, then we'll log out the error. 

`defer f.Close()` can be used right after our file is opened to defer the closing of the file once we have exited our `main` function. 

Then we use the `bufio` package. `NewReader` returns a new reader object. Here we store it in the variable `r`.

Then we enter an infinite loop and use the `ReadString` method on on reader. We pass the newline in as delimiter, which tells ReadString to read until that delimter, then store the string in the variable `s`. Again, if we find an error, then we'll break out of the loop, but if not we'll continue.

We then use the `strings` package and it's `Contains` method. `Contains` will return `true` if in the string `s`, there exists a substring `'ERROR'`. If we do encounter that substring, then we print out the entire line.

But what if we wanted to allow the user to pass some values into the function so that they could filter on a specific log file, and the keyword that they're looking for?

Here is our updated CLI application, which allows command line arguments:

```go
func main() {
	path := flag.String("path", "myapp.log", "The path to the log file being analyzed")
	level := flag.String("level", "ERROR", "Log level to search for. Log level to search for. Options are DEBUG, INFO, ERROR, and CRITICAL")

	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	r := bufio.NewReader(f)

	for {
		s, err := r.ReadString(('\n'))
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}
```

We see that a couple of things have changed. We're now using a Go package called `flag`. These are command line flags. With the flag, we need to tell it what type of data that we're expecting with the flag. Here we define that the value will be a string. This is helpful for us to strictly enforce what kind of data can be passed for our command line arguments. This means that we can pass in our arguments like this:

```bash
go run . -level ERROR
```

There are three arguments that we've also passed. 
1. the name of the flag ('path', 'level')
2. the default value ('myapp.log', 'ERROR')
3. some information about the flag

To get some documentation for our flags that we've defined we can use the `-help` command:

```bash
go run . -help
```

Running that we get this helpful output:

```bash
  -level string
        Log level to search for. Log level to search for. Options are DEBUG, INFO, ERROR, and CRITICAL (default "ERROR")
  -path string
        The path to the log file being analyzed (default "myapp.log")
```

After our flags have been defined, we need to call `flag.Parse()` in order to parse the arguments. Go will populate the `path` and `level` variables with the command line arugments, or the default values. This is all an effort to make Go as efficient as possible.

Then, we can also see that we've changed the argument that we pass to `os.Open`. We now set the argument to `*path`. Since the value that is returned by `flag.String` is a pointer to our actual string value. We've done the same thing with our argument to `r.ReadString`, this now takes the pointer to our level string.

With these updates, we're able to make our program more flexible, so that a user can actually search for the log level that they want.

## Creating a web service

Here is the simple web service that we wrote:

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	healthHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "applicaton/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "OK"}`))
	}

	pingHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "applicaton/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "status:OK", "data": {"message": "pong"} }`))
	}

	greetUserHandler := func(w http.ResponseWriter, req *http.Request) {
		names := req.URL.Query()["name"]
		var name string
		if len(names) == 1 {
			name = names[0]
		}

		m := map[string]string{"name": name}
		enc := json.NewEncoder(w)
		enc.Encode(m)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}

	http.HandleFunc("/", greetUserHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/ping", pingHandler)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
```

Let's first break it into its component parts.

We have defined three routes and three handlers for those routes. The routes and their handlers are:
1. "/", greetUserHandler
2. "/health", healthHandler
3. "/ping", pingHandler

We'll first describe the routes that we've defined and then we can discuss the handlers that are actually called when those routes are hit.

We are using the `net/http` package, which is available in the standard library. In that package we use a method called `HandleFunc`, which takes as argument, the path, and the handler. I've decided to move the handler logic into their own functions. Let's discuss the handlers now

The first two handlers that we have defined `healthHandler` and `pingHandler` are quite similar so we can describe both of them below.

When the `"/health"` route, or the `"/ping"` route is hit, it will call its respective handler. The handler is a function, which takes two arguments, a response writer object and a request object. We've named these parameters `w`, and `req` respectively. 

Inside of our function we use a few methods available on the response writer object to set the content type and the status code. Then finally we use this: 

```go
w.Write([]byte(`{ "status:OK", "data": {"message": "pong"} }`))
```

to actually write our response. The `[]byte()` part is important. The reason for this is that we're writing our response to a writer interface, when we use `ResponseWriter`, which always works with collections of bytes, because computers work with collections of bytes.

If we wrote a simple program to print out the 'Hello World' in byte form:

```go
package main

import "fmt"

func main() {
	st := []byte("Hello, World")
	fmt.Println(st)
}
```

We would get this output:

```bash
[72 101 108 108 111 44 32 87 111 114 108 100]
```

These are the ASCII encodings for 'Hello World'.

Now let's talking about the `greetUserHandler`:

```go
greetUserHandler := func(w http.ResponseWriter, req *http.Request) {
    names := req.URL.Query()["name"]
    var name string
    if len(names) == 1 {
        name = names[0]
    }

    m := map[string]string{"name": name}
    enc := json.NewEncoder(w)
    enc.Encode(m)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
}
```

When we pass in a query parameter to our URL, we'll be able to access those parameters on the `http.Request` object. We get the query parameters by using `req.URL.Query()`. 

This function will return a map of data that contains all of the query parameters that we passed in. We use `["name"]` to pull out the query parameters called `name`. Now, importantly, because this is a query parameter, there could be multiple parameters called `name`, which is why the name of variable is `names`. This could be a collection of more than 1 name.

Then we declare a variable called `name` and then assign it to the first name that's passed in as a parameter.

The three lines that follow that `if` statement handle the JSON encoding for our response. 

First what we do is create a map, which is a collection of key, value pairs. Maps in Go can be a wide variety of data types, but for us, we're going to create a map of strings to strings:

```go
m := map[string]string
```

This says create a map where the keys are strings and the values are strings.

Then the next part sets an initial key, value pair.

```go
m := map[string]string{"name": name}
```

Then we need to actually encode this map to a JSON string. We do this be creating a JSON encoder from the `json` package, and we'll use the `NewEncoder` method on that package. The `NewEncoder` method requires as its first argument some kind of writer. In this case we'll pass our `http.ResponseWriter` object. This is because the encoded JSON will be passed directly to the writer that we pass in. This will return for us an encoder that can write to the writer that we passed in.

```go
enc := json.NewEncoder(w)
```

Then finally we encode the map. So we do this:

```go
enc.Encode(m)
```

So here is all of that together again:

```go
m := map[string]string{"name": name}
enc := json.NewEncoder(w)
enc.Encode(m)
```

Now when we hit the `/` route, a user can pass in a name like this `/name=Jack` and we'll get a response like this:

```json
{"name": "Jack"}
```





