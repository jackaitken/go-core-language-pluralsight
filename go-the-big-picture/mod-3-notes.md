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
//loop over a collection
for user := range users ...
```

Here we just use a for loop again to loop over a collection, but we use a variable that is easy to read.

So this is all to say that **all loops in Go are for loops**.