package main
import "fmt"


func main() {
	var name string = "Bhanu"
	printMe(name)
	add(5, 2)
}

func printMe(val string) {
	fmt.Println("Hello!")
	fmt.Println(val)
}

func add(num1 int, num2 int){
	fmt.Println(num1 + num2)
}