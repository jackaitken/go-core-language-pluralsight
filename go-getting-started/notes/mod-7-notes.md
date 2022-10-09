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
    printLn(i)
    i++
} 

// 0
// 1
// 2
// 3
// 4
```

A note to make here is that we used the built in `printLn` function. This is not part of the `fmt` package, and is less sophisticated, and normally only used for debugging.

To break out of a loop early we can use the `break` statement.

```go
var i int
for i < 5 {
    printLn(i)
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
    printLn(i)
    i++
    if i == 3 {
        continue
    }
    printLn("continuing...")
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
    printLn(i)
}
```

So here, we've created a post clause that will handle the incrementation for us. Something big to note here is that when we use a post clause like this, we have to have 3 terms in the for loop. This indicates to the compiler that we're using a loop till condition with a post clause. So this code above will not compile. Instead we can use that third term to initialize our variable:

```go
for i := 0; i < 5; i++ {
    printLn(i)
}
```
Like in JavaScript, the variable `i` here is scoped to that block, so it can't be used elsewhere