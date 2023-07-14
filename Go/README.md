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

## Functions
When declaring functions, the reserved word *func* is used, as we can see in the example:

    func  main() {
	    //code here
    }
    
Note that the previous function does not need and therefore has no return type. For a function to have a return, it is necessary to declare it previously in the scope of the function, as in the following case:

    func  sum(a int, b int) int {
	    return a + b
    }

Otherwise, even if there is a *return* when compiling, we will have an error, as in the example:

    func  sumError(a int, b int) {
	    return a + b
    }
    
Note that the difference between the functions happens right after the declaration of the function's parameters and before opening the curly braces, this is where we declare the function's return type.

Now that we know how to declare a function with a return, we can understand that in *Go* our functions can have multiple returns, as follows:

    func  sumAndReturnString(a int, b int) (int, string) {
	    // Retorna respectivamente um int e uma string
	    return a + b, "Added!"
    }

In Go we also have what we call *Named Return*, which is nothing more than previously defining a *name* for your return. This can be done by defining a name to the left of the return type declaration, in this way, having a *Named Return* it is not necessary to explain what is returned in the function, only the reserved word *return* is enough, as Go itself understands what is expected from the return of that function, as in the example:

    func  sumNamedReturn(a int, b int) (result int) {
	    result = a + b
	    return
    }

Another variation that we have are the *Variatic Functions* where we can pass an indefinite amount of parameters to the function, but pay attention that **all must be of the type defined in the function declaration,** see below:

      // FUNCTION
        func  sumAll(x ...int) int {
    	    result :=  0
    	    
    	    for _, v :=  range x {
    		    result += v
    	    }
    	    
    	    return result
        }

The one responsible for allowing the passing of **n** parameters to a function is the section `(x ...int)`, where **x** works as an array of values of the defined type, in this case *int*. The rest of the function uses a loop to walk through the values of **x** and thus add all the values passed to the function, it is noted that we have the use of a **blank identifier** o ** _** which in this case is used to discard the key/index returned by the **x** loop. To call such a function, do the following:

    resultAll :=  sumAll(10, 20, 3, 5, 6, 4)

We can also have a variable that is implicitly assigned with the type of a function, which returns the type `func() int` which in turn will return an `int` type as shown below:

    resultFunc :=  func(x ...int) func() int {
	    resultTemp :=  0
	    for _, v :=  range x {
		    resultTemp += v
	    }
	    return  func() int {
		    return resultTemp * resultTemp
	    }
    }

To better understand what this function returns we can do:

    fmt.Println(resultFunc(54, 54, 54, 54))
   and
   
    fmt.Println(resultFunc(54, 54, 54, 54)())

In the first case we have a different return, that is, instead of obtaining the final value that should be returned by the internal function within the function, we actually obtain the memory position of this return. In the second case, as we are calling the function, passing the parameters and finally executing the function, it is understood that we are executing the function inside the function, and therefore the return we get in this way is the expected one, that is, a return of type *int* with the operations of values passed by parameter carried out.
Undoubtedly, something very important in Go is the "transformation" of a function into a method, it should be noted that Go has an object orientation itself, we have what we call GoWay, that is, the Go way of building things, that way, to create **methods** we need two things, firstly a Struct and secondly a function linked to that Struct, but how? See below:

    type  Car  struct {
	    Name string
    }

We have a *Car* struct with a String type Name property. Now to create the method, we need to declare a function practically as we normally do, with the only difference that this time we will do a *Binding* type to bind the function to the previous Struct, and thus become a Struct method, look:

    func (c Car) drive() {
	    fmt.Println(c.Name, "Drive")
    }

Notice that the structure for declaring the function changes slightly, so that we have the *func* keyword and then the binding itself, when we do `(c Car)` before naming the function.

## Good habits
- Whenever exporting methods, variables and/or functions, it is interesting to add a comment/documentation about the method to be exported, in order to provide documentation about the responsibilities of the exported target. You should always start the comment with the name of the method, that way it recognizes the documentation format.

## Index
- [Introduction](https://github.com/alonsofritz/TIL/tree/master/Go/hello-go)
- [Error Handling](https://github.com/alonsofritz/TIL/tree/master/Go/error-handling-go)
- [Functions](https://github.com/alonsofritz/TIL/tree/master/Go/functions-go)

##
[![en](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/alonsofritz/TIL/tree/master/Go/README.md) [![pt-br](https://img.shields.io/badge/lang-pt--br-green.svg)](https://github.com/alonsofritz/TIL/tree/master/Go/README.pt-br.md)