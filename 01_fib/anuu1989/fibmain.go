package main 

import (
	"fmt"
	"io"
	"os"
)
var out io.Writer = os.Stdout
func fib(n int){

	var maxnumber int = n
	var previousNumber int = 0
	var nextNumber int = 1

	for i :=1 ; i <= maxnumber ; i++{
		fmt.Fprintln(out, previousNumber)
		sum := previousNumber + nextNumber
		previousNumber = nextNumber
		nextNumber = sum 
	}


}



func main(){
	fib(7)
}




