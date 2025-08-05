package main
import (
	"errors"
	"fmt"
)


func main(){
	var numerator int = 78
	var denominator int = 2
	quotient, remainder, err := intDivision(numerator, denominator)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("Q=%v, R=%v for numerator=%v and denominator=%v", quotient, remainder, numerator, denominator)
}

func intDivision(numerator, denominator int)(int, int, error){
	var err error
	if denominator == 0{
		err = errors.New("Denominator can't be zero for divisions bruh")
		return 0, 0, err
	}
	quotient := numerator/denominator
	remainder := numerator%denominator
	return quotient, remainder, err
}