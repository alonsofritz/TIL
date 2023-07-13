// Set package name
package main

/**
*	Package and lib imports:
*	Note that "fmt" refers to a native Go package, and it has most of the functionality needed for implementation,
*	avoiding the need to install third-party packages. Most of the other packages for Go are available on Github itself,
*	unlike Docker, for example, which has its own Hub.
**/
import (
	"fmt"

	"github.com/alonsofritz/TIL/tree/master/Go/hello-go/math"
	"github.com/google/uuid"
)

/**
*	To declare functions, use the word "func"
**/
func main() {
	fmt.Println("Hello Go" + uuid.NewString())

	// Variable declaration
	var a string
	a = "Test 00"

	// Variable declaration by inference
	b := "Test 01"
	b = "Test 02"

	c := 3.14
	d := false
	e := `Funny
	Text
	with
	ln
	`

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// or

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)

	/**
	*	Note that inside the "math" folder we have two files, respectively, "operations.go" and "operationsX.go",
	*	both in their declaration were defined as belonging to the "math" package (the folder name is not the package name ,
	*	it's just good practice). Note that even calling the same package, I can access the methods of the files even if they are separated,
	*	because what makes up a package is its declaration and not a file itself.
	 */
	result := math.Sum(1, 1)
	fmt.Printf("%T \n", result)
	fmt.Printf("%v \n", result)

	resultX := math.SumX(1)
	fmt.Printf("%T \n", resultX)
	fmt.Printf("%v \n", resultX)

	resultY := math.A
	fmt.Printf("%v \n", resultY)
}
