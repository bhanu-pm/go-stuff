package main

import "fmt"
import "unicode/utf8"


func main(){
	var variable_name int = 42
	fmt.Println(variable_name)
	
	var float_var float32 = 1.234588989
	fmt.Println(float_var)

	var float_var_precise float64 = 1.234588989
	fmt.Println(float_var_precise)

	var intnum int = 2
	var floatnum float32 = 123.33
	var result float32 = floatnum + float32(intnum)
	fmt.Println(result)

	var int1 int = 5
	int2 := 2
	var result_number_div = int1/int2
	var result_number_remainder = int1 % int2
	fmt.Println(result_number_div)
	fmt.Println(result_number_remainder)

	var name string = "Bam" + " " + "Bam"
	fmt.Println(name)

	fmt.Println(utf8.RuneCountInString(name))

	const myconst string = "Bhanu-pm"
	fmt.Println(myconst)
}