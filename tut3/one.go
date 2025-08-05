package main
import (
	"fmt"
	"errors"
)


func main(){
	var1 := 39
	var2 := 0
	var quotient, remainder, err = intDivision(var1, var2)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("quotient: %v, remainder: %v, for variables %v and %v", quotient, remainder, var1, var2)
}

func intDivision(var1, var2 int)(int, int, error){
	var err error
	if var2 == 0{
		err = errors.New("Denominator should not be zero!")
		return 0, 0, err
	}

	var quotient = var1 / var2
	var remainder = var1 % var2
	return quotient, remainder, err
}