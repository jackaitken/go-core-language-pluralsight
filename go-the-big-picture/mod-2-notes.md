# Module 2

The goal of this course is not to dive very deeply into Golang. Instead, we're going to get an overview of what Golang is, what it is used for, and why someone might use it. 

In the next module, we're going to use Go to actually create some applications.

## What is Go?

First lets talk about who developed Go and the problems that they were trying to solve when they developed it. 

Go was developed by 3 engineers: Ken Thompson. Ken helped to implement the B and C languages, the original versions of Unix, and he's one of the creators of UTF-8. Another engineer, was Rob Pike, who also worked on Unix and UTF-8, and finally there is Robert Griesemer who created the Hotspot JVM.

Now these 3 engineers were working for Google at the time that they developed the language, and some things that we see that are implemented in Go were definitely targeted at solving some of the problems at Google, but Go is definitely not a Google project.

So what was the problem they were trying to solve by creating Go?

When they were first discussing the development of Go, they categorized existing languages into 3 broad categories:

- Efficient Compilation
- Efficient Execution
- Ease of Programming

C++ is a language that might fall into the realm of efficient execution, and Java might fall in between efficient compilation and efficient execution, but it's pretty hard to write. JavaScript has a very fast compilation time, and it's pretty easy to get started, but the execution efficiency can suffer. Python is another example of this.

So if none of these languages fit into these 3 categories neatly, they would need a new language.

## Language Characteristics

Let's first look at the syntax of Golang. 

**It has a strong statically typed system**
This means that it has well defined types at compile time, which allows for more tooling support from our IDE as we write our programs, and it allows our compiler to be more efficient when creating output.

**Go's syntax is inspired by C**
Upon first glance, we can see that Go has a lot of similarity to C, but the syntax for Go wasn't just blindly adopted from C. Instead, the developers tried to focus on how we write programs now, and make sure that the syntax flowed properly. So when we're looking at Go code, we'll see that there are some things that are in C, which are missing from Go, and some things will automatically be entered by the compiler. All of this is to help increase the usability of Go

**It is a multi-paradigm language**
Meaning that we don't have to just use OOP, or procedural programming, and almost all Go programs are a mix of procedural and OOP code. It all depends on the specific solution that the code is trying to solve. So it's encouraged to examine our use case for our code and determine whether we really need it to be object oriented, or if we can simply have procedural code.

**Go is a garbage collected language**
This means that we don't need to manage memory on our own when working with Go. This was part of a decision to make Go as robust as possible when creating applications. This helps Go programs scale more easily and reduce the mental overload that can occur.

**It is a fully compiled language**
Go applications are fully compiled down to an executable binary, which gives us a good opportunity to achieve pretty great performance.

**It focuses on rapid compilation**
This helps us stay in a nice flow while developing. If we're using TDD, then we need to have fast compilation so that we stay productive while writing code.

**We get a single binary output**
All of our Go code will be compressed into a single binary and compressed for us to deploy. Now importantly, this single binary compilation, is the default option, and there are other options available to us now. 

## How long has Go been around and where is it going in the future?

Here's a timeline of Go so far:

*2007*
The beginning of Go's design is begun
|
|
*2009*
The language was publicly announced
|
|
*2012*
This was the version 1.0 release of the language
|
|
*???*
Currently there is a lot of talk about a Go version 2 release. Currently it's in a discussion phase. 



