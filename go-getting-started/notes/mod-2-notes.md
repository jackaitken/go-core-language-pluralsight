# Module 2 Notes

In this module we're going to discuss where Go came from, where it excels, and where its weak points are. Then we're going to build a simple Hello, World application to get us started.

## What problems does Go solve?

First, let's discuss some of the problems that the developers were trying to solve when they wrote Go. If we look at some of the languages that Google was using before Go, we'll see some advantages and disadvantages:

1. C++
    C++, has very high performance and type safety, however it can have very slow compile times, and the syntax is pretty complex.
2. Java
    With Java, some of these problems were alleviated, as it has a faster compile time, and it has good type safety. However, as Java grew, the ecosystem became a bit complicated. There are so many different ways to solve the same problem, and more often than not, those solutions don't play well together, so in many cases you need to build custom tool chains for each Java application. This also means that a new developer on your project has to get up to speed on that specific tool chain.
3. Python
    Python is very easy to use, but being that is a dynamically typed language, it lacks type safety, so types can change over time if those types aren't used consistently. It's also pretty slow compared to Java and C++.

So none of these languages was a cure all for Google's problems. So let's look at why Go solved these problems:

1. **Fast compilation**: Go has a much faster compilation time, which allows developers to be quicker in development, so there isn't as much stop and start with changes to code.
2. **Fully compiled**: Go is also a fully compiled language, meaning that it can take full advantage of the resources available on the platform where the application is being executed.
3. **Strongly typed**: With strong typing, we're able to get the editor support for our types. Go however, will do its best to infer types while we're writing code.
4. **Go is concurrent by default**: With the three languages that we discussed above, there is a concept of threading, so we can have task running in parallel. But, with these languages, the tools to manage that concurrency are managed by third party libraries written after the language was created. When these languages were created multi-core processors weren't as popular.
5. **Garbage Collected**: Go is garbage collected, but there was some contention at the beginning of its development. Whenever a language does some sort of garbage collection, there is some type of pause. For languages that aren't very performant, the garbage collection event can be a bit of a problem. With Go, the advantages of garbage collection were used, but there has been a huge focus to ensure that the time necessary to perform this garbage collection is minimal.
6. **Simplicity as a core value**: This is not to say that Go is easy to write, but there has been a lot of effort to make Go simpler. For example, garbage collection done by the language makes things easier by allowing us to not worry about memory management. It uses strong typing, so that we can get immediate feedback from our editor, which simplifies the language greatly and simplifies the management of the language.

## What is Go good at?

The first two that we'll explore are *web services* and *web applications*. The former meaning that we're delivering some sort of data, and the latter handles delivering web pages to some web browser. Go is good at both of these because of its built in concurrency model. Go's standard library was built with network awareness. 

Go is also used frequently for *task automation*, which is typically left to bash scripts, or scripting languages. It turns out that Go's syntax is just about as light as many of those scripting languages. 

## Demo

We can go to (The Go Playground)[https://play.golang.org], to test out some simple applications. Here is what our simple Hello World application looks like:

```go
package main

// This is an import block. With import blocks we can imported packages
// without repeating the 'import' keyword.
import (
    "fmt"
)

/*
This is the main function, which, when part of the main package, indicates the entry point of an application.
*/

func main() {
    fmt.Println("Hello, World")
}
```

At the beginning of the program we have a package statement. Each Go program will start with a package statement. This helps to identify where it sits in our application. 

Then next we have an `import` block. Applications will almost always need some sort of additional package. In this example we've imported the `fmt` package from the standard library. Line 7 contains the `main` function. With every executable program we have in Go, there will be a `main` function that represents the start of our program.

Inside of our `main` function we use the `Println` method inside of the `fmt` package. We pass in a string to be printed out.

We can also see that single line comments are delinated by `//`. Multi-line comments, like JavaScript are delinated by using `/* */`.

Now interestingly, if we were to comment out the code inside of our `main` function, we would get a compile time error. This is because we would have imported the `fmt` package without using it. Every Go file that imports a package, must use that package. This is enforced to make our applications more maintainable. If a package is imported and not used it can cause confusion, so this is a compile time error.

Let's say that we uncomment out our code in the `main` function, but this time we remove the tabs so that our code is left justified without any whitespace. This will not create a compile time error, because whitespace is not enforced by Go. However it is strongly encouraged. 

There are however, some formatting rules that are strongly enforced by the compiler. For example, let's try running this:

```go
func main() 
{
    fmt.Println("Hello World")
}
// error: missing function body
// error: syntax error: unexpected semicolon or newline before {
```

Now the reason for the second error, is that Go will automatically insert semi colons. We don't have to worry about that as developers. But what happens when our format looks like it does above, since we don't have our opening curly brace on the same line as our function delcaration, the Go compiler does this:

```go
func main();
{
    fmt.Println("Hello World")
}
```

So here it sees that we have a function declaration without a body, our first error. And then we have a function body without any kind of declaration. Now, trying to format this, our formatter won't exactly know what to do, but it will guide us to the solution.




