package main 
import "fmt"


func main(){
	var intlist [3]int32
	var intlist2 = [...]int32{0, 9}
	var intlist9 [9]int32 = [9]int32{0, 1, 2, 3, 4, 5, 6, 7, 8}
	intlist10 := [...]int32{0, 2, 4, 6, 8}

	fmt.Println(intlist)
	fmt.Println(intlist2)
	fmt.Println(intlist9)
	fmt.Println(intlist10)

	fmt.Println("\nTesting indexes")

	fmt.Printf("item  at index %v i.e at %v th position is %v", 4, 5, intlist9[4])
	fmt.Printf("\nAccess smaller parts of arrays (slice) only gets the headers, not a new copy")
	// slice
	smallerList := intlist9[1:4]
	fmt.Printf("\nslice from idx 1 to 4 in intlist9 is %v", smallerList)

	//change slice changes original list as it has just the pointers of the original list, not a new copy
	fmt.Println("changing an ele in slice to 99 changes the original array")
	smallerList[2] = 99
	fmt.Printf("slice changed to %v and the original array is changed to %v", smallerList, intlist9)

	// addresses (pointers) in arrays
	variablee := 39
	fmt.Printf("\nVariable is %v and its address(pointer) is %v", variablee, &variablee)
	fmt.Printf("\narray is %v and its address (of first element) is %v and address of idx 2 is %v", intlist9, &intlist9[0], &intlist9[2])
}