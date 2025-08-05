package main
import "fmt"


type Number interface{
	int | float64
}

func main(){
	addition(13, 26)
	addition(1.3, 2.6)
}

func addition[intFloat Number](num1 intFloat, num2 intFloat){
	fmt.Println(num1 + num2)
}