package math

/**
*	Go works with methods, variables and functions that can be exported or not,
*	this means that every time I access an external package, I will only have access to the methods that are exported.
*	But what defines if a method is exported or not? For this, the first letter used in the declaration of these methods must be capitalized,
*	so to define something "private", just declare it starting with a lowercase letter.
**/

// SumX : Adds a variable with the value 10
func SumX(a int) int {
	return a + 10
}
