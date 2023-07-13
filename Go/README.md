# TIL: Go

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

Generates build for any operating system from any operating system. Currently the Go compiler is done in Go, so it is a self-compiling language.

## Package Imports

    import ( "fmt" )
    
Note that "fmt" refers to a native Go package, and it has most of the functionality needed for implementation, avoiding the need to install third-party packages. Most of the other packages for Go are available on Github itself, unlike Docker, for example, which has its own Hub.

## Exporting variables and functions

Go works with methods, variables and functions that can be exported or not, this means that every time I access an external package, I will only have access to the methods that are exported. But what defines if a method is exported or not? For this, the first letter used in the declaration of these methods must be capitalized, so to define something "private", just declare it starting with a lowercase letter.

## Error Handling
In Go, the error will always be something explicit when working with functions. We can have multiple values ​​in the return of a function, mostly one of these returns is an "error", which brings with it information about the errors that have occurred in the function in question.

## Good habits
- Whenever exporting methods, variables and/or functions, it is interesting to add a comment/documentation about the method to be exported, in order to provide documentation about the responsibilities of the exported target. You should always start the comment with the name of the method, that way it recognizes the documentation format.

## Index
- [Introduction](https://github.com/alonsofritz/TIL/tree/master/Go/hello-go)
- [Error Handling](https://github.com/alonsofritz/TIL/tree/master/Go/error-handling-go)

##
[![en](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/alonsofritz/TIL/tree/master/Go/README.md) [![pt-br](https://img.shields.io/badge/lang-pt--br-green.svg)](https://github.com/alonsofritz/TIL/tree/master/Go/README.pt-br.md)