package main

import "fmt"

type Car struct {
	Name string
}

/**
*	By placing (c Car) before naming the function, what we did was a Binding between Function and Struct,
*	thus creating a method.
**/
func (c Car) drive() {
	fmt.Println(c.Name, "Drive")
}

func main() {

	myCar := Car{
		Name: "BMW",
	}
	myCar.drive()

	result, str := sumAndReturnString(10, 20)
	fmt.Println(result, str)

	/**
	*	When making a call to a function that returns (x ...int), we can pass any number of values ​​of the type defined by the parameter.
	**/
	resultAll := sumAll(10, 20, 3, 5, 6, 4)
	fmt.Println(resultAll)

	/**
	*	Function that returns a function
	**/
	resultFunc := func(x ...int) func() int {
		resultTemp := 0

		for _, v := range x {
			resultTemp += v
		}

		return func() int {
			return resultTemp * resultTemp
		}
	}

	/**
	*	Note that this way it does not return the result value, but the function itself.
	**/
	fmt.Println(resultFunc(54, 54, 54, 54))

	/**
	* 	So that we can return the value of the function, which in this case is of type int,
	*	we need to add () after passing parameters, in this way we are saying that we are executing the returned function,
	*	in this way it considers the value returned by it.
	**/
	fmt.Println(resultFunc(54, 54, 54, 54)())
}

/**
*	The following function does not have a return type definition,
*	so it is a function that does not return and cannot return anything,
*	notice that when defining a "return" for the function we get an error.
**/
func sumError(a int, b int) {
	return a + b
}

/**
*	Correct Way
**/
func sum(a int, b int) int {
	return a + b
}

/**
*	We can have returns with multiple values
**/
func sumAndReturnString(a int, b int) (int, string) {
	return a + b, "Added!"
}

/**
*	Named return: As we declare a name for the return type,
*	it is no longer necessary to specify what should be returned by the function, since it is implicitly defined.
**/
func sumNamedReturn(a int, b int) (result int) {
	result = a + b
	return
}

/**
*	Variatic Functions
*	By declaring the parameters received as (x ...int), we are saying that x will work as an array of integers,
*	so when calling the function in question we can pass an undetermined amount of values ​​through parameters.
**/
func sumAll(x ...int) int {
	result := 0

	/**
	*	The rest of the function uses a loop to walk through the values of "x"
	*	and thus add all the values passed to the function, it is noted that we have the use of a "blank identifier" the "_"
	*	which in this case is used to discard the key/index returned by the "x" loop.
	**/
	for _, v := range x {
		result += v
	}

	return result
}
