package main 

import "fmt"

func fib(n int){

	var maxnumber int = n
	var previousNumber int = 0
	var nextNumber int = 1

	for i :=1 ; i <= maxnumber ; i++{
		fmt.Println(previousNumber)
		sum := previousNumber + nextNumber
		previousNumber = nextNumber
		nextNumber = sum 

	}


}



func main(){
	fib(10)
}




