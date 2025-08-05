package main 
import "fmt"


type intFloatStr interface{
	// combo of int float string addition or concatenation based on datatype
	int | float64 | string 
}
func main(){
	// fmt.Println("ag"+"let")
	addn(30, 9, "integer")
	addn(3.0, 0.9, "float")
	addn("3", "9", "string")
}

// C is the combo type 
func addn[C intFloatStr](var1, var2 C, var3 string){
	result := var1 + var2
	fmt.Printf("%v is the datatype for the following result: %v \n", var3, result)
}