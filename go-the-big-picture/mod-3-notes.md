# Module 3

## Philosphy and Values

Looking at Go's philosophy and values, there are a few things to highlight:

**Simplicity**
The design and syntax of the language are all about simplicity. Simple != easy however, it means that the Go ecosystem should be easy to approach.

**It has been designed with network aware and concurrent applications in mind**
At the time of Go's development, the hot languages were Python, Java etc. which were developed before concurrent applications were being used. Go realized the importance of concurrency and the ability for computers with multiple cores, running multiple threads at once. Also, at some level, with every application that we design today, there is some aspect of using a network, and Go was also developed with that in mind.

**Out of the box experience**
One of the nice parts about Go is that it is in many ways a one-stop-shop to get up and running with Go.

**Go is cross-platform by nature**
We no longer have apps that are designed specifically for an OS. In order to reach the widest number of users. So it was developed in a world where cross platform applications are the norm.

**Backward compatability**
If a language is being actively developed and improved, it's important that it have backward compatability. Within limits, backward compatability is honored as the language grows and evolves.

## How does Go embrace simplicity

So many languages start with this as a goal, but over time the addition of new features, and the need to make things more powerful can drown that simplicity. Go has embraced simplicity as a core philosophy.

Here is an example of some (non-valid) Go code:

```go
i := 1
println(i++) //
println(++i) //
```

The first function call uses a postfix increment operator, whereas the second uses a prefix increment operator.

Now looking at both of these lines by themselves like this, we can probably imagine what they might do. But what if we're looking through a function that has 15-20 lines? When we glance past this we might not notice something like this that could lead to subtle bugs. 

These operators do significantly different things, and it can add issues to our code. 

So the problem here is that, the increment and decrement expressions in many languages are easily misinterpreted. **Go fixes this by changing increment and decrement operators from expressions to statements**. The key difference between an expression and a statement is:

**Statement**
A statement is evaluated entirely as one unit

**Expression**
An expression is a component of a statement

So then the valid Go representation of our code above is this:

```go
1 := 1
i++
println(i) // 2
i++
println(i) // 3
```

So we still initalize our variable, then separately, we increment our variable, then we print out the value of that variable.

Not only is this easier to read, but in fact we are not allowed to write Go as we tried above because `i` cannot be incremented as part of a larger statement. *If we're going to use an increment operation, then that **is** the statment*. 

So, this does of course involve more lines of code, but it comes at the advantage of knowing exactly what the intention of this code is.

We can see another example of Go's simplicity in how it handle looping constructs.

```go
// loop with incrementor
for i:= 0; i < 5; i++ ...
```

This is similar to a lot of the loops that we see in other languages

```go
// loop until condition
for i < 5 ...
```

This is a for loop as well, but all we have here is a condition, and we just test for that condition. Once that condition is met then we exit

```go
// infinite loop
for ...
```

This accomplishes the same thing as `while (true)` in other languages, expect we don't need anything more than just the `for` keyword.

```go
// loop over a collection
for user := range users ...
```

Here we just use a for loop again to loop over a collection, but we use a variable that is easy to read.

So this is all to say that **all loops in Go are for loops**.

## Network aware and Concurrent Apps

We mentioned that Go was developed with network awareness and concurrency in mind, so how did they do that?

Well, in the core and standard library are the net and *net/http packages*. From a concurrency standpoint, there are *goroutines*, which are lightweight threads which allow for a huge amount of concurrency. Then, in order to manage communication between those threads that we have running, there are *channels*.

The *net and net/http packages* allow us to create web servers using only the standard library.

*goroutines* allow us to start thousands of concurrent tasks with minimal resources. They are an abstraction over a processer thread. Go also has a scheduler that watches all of the goroutines and it will assign those goroutines to different processer threads. 

*channels* allows us to safely communicate between channels. The difficult part with concurrency isn't the concurrency part, it's the sharing of data. So with channels the sender knows that it sent the message successfully, and the receiver knows that it received it successfully. There is no worry about synchronizing, because the channel handles all of that. 

## Out of the box experience

Something that the Go developers focused on was the idea of getting up and running on Go as soon as possible. They focused on making features and other important parts of language accessible, simply through the standard library. 

**The Go standard library** is not as powerful as some other languages, but it is focused on providing the core and most powerful concepts to build applications. For example, in the standard library we have built in:

- string manipulation
- data compression (creating zip files)
- file manipulation
- API's for networking
- a completely developed testing suite
- much more...

The interesting bullet here is the testing one, which is certainly not a standard library feature of most languages. 

**The Go CLI** brings everything that you're going to need to create, build, and test Go applications. Included in that standard CLI is:

- project initialization
- building your application
- code generation (codegen)
- retrieve dependencies
- testing
- application profiling
- documentation
- report language bugs

## Cross platform 

If we were working with a language like C++, there is a lot of work that needs to happen in order to deploy that application to Windows vs. Mac vs. Linux. With Go, this cross platform nature is trivial to implement, because it's already baked into the language. 

Depending on what OS we're building on, we just change some environment variables and then we're all set. We change the `GOOS` and `GOARCH` variables.

## Backward compatability

This is a huge consideration for most organizations, and luckily, Go puts a lot of thought into this. Imagine that we're developing a new application, and a new version of the language comes out. What if there are new langauge features that the team wants to adopt? With some languages, if we upgraded to that new version, everything in our current application might break. 

Go is different however. In the Go docs, they specify that any Go code written with Go 1 let's say, will always compile within the lifetime of that specification. This is the Go Compatability Promise.

Now, there are some exceptions to this promise. *One being security*. If there is a security issue that is discovered in the code, this is an exception to the promise, and they reserve the right to break that feature to fix the security issue. 

Another being *unspecified behavior*. 

Another is *specification errors*. 

Finally, *bugs* are also an exception to this promise.

What is important is that the Go development team will not break backward compatability, just because there is a new design philosophy. For example, a new way to handle strings. The decisions that would break backward compatability in that way, are reserved for the next major version of a language.

## Looking at Go applications and Go's primary use cases

1. **web services**, and moving data back and forth from different servers
2. **web applications**, and trying to deliver html to the user
3. **devops**, and tools like Go and Docker
4. **GUI/Thick-client** applications
5. **machine learning**
6. **others**. Go is still a young language, so people are trying to find new use cases for it

The top three in this list however are where Go is really shining.




