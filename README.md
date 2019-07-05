# try
A go package that offers a try/catch statement block.

[Documentation](https://godoc.org/github.com/thestrukture/try)


Here is a sample program making use of the try package. In this sample, variables in scope are accessed to demonstrate how this package interacts with your program.

```
package main

import (
	"github.com/thestrukture/try"
	"fmt"
	"errors"
)


func main(){

	count := 0

	try.Run(func(){
		count++
		count++

		// Will perform an error check, panic will
		// not be invoked if the interface passed is nil.
		try.Throw(errors.New("Crashed!"))

		fmt.Println(count);
	}).Catch(func(e interface{}){

		fmt.Println("Error : ", e)

	}).Finally(func(){
		fmt.Println("Done")
	})
}
```
