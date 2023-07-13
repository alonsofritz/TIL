package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://googlecom.br")

	/** Note that we have a request that returns a 'res' type and another 'err' type referring to the possible error in the request.
	*	To catch the error, we'll leave the Try/Catch curse behind, and in Go the following block becomes familiar,
	*	since it's the "default flow" used to handle errors within Go.
	**/
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res.Header)

	resTest, errTest := sum(7, 2)

	if errTest != nil {
		log.Fatal(errTest.Error())
	}

	fmt.Println(resTest)

	/**
	*	BLANK IDENTIFIER:
	*	If we want to ignore the error return, that is, not treat it, we can ignore it in the process of returning the function,
	*	using what we call Blank Identifier, which basically "throws away" the returned variable. The Black Identifier is represented by " _ "
	**/
	resSum, _ := sum(7, 2)
	fmt.Println(resSum)

}

func sum(x int, y int) (int, error) {
	res := x + y
	if res > 10 {
		return 0, errors.New("Total greater than 10")
	}

	return res, nil
}
