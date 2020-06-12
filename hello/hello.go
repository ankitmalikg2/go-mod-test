package hello

import (
	"fmt"

	"rsc.io/quote"
)

func Hello() string {
	fmt.Println("ankit")
	return quote.Hello()
}
