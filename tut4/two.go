package main
import "fmt"


func main(){
	// Slices stuff
	intSlice := []int32{4, 5, 6}
	fmt.Println("Slice capacity is the length of the array that is created in the background to create this slice")
	fmt.Printf("SLice is %v\nSlice length is %v\nSlice capacity is %v\n", intSlice, len(intSlice), cap(intSlice))

	// appending values to the slice
	fmt.Println("Appending values to the slice")
	intSlice = append(intSlice, 39)
	intSlice = append(intSlice, 99)
	fmt.Printf("New Slice is %v\nNew Slice length is %v\nNew slice capacity is %v", intSlice, len(intSlice), cap(intSlice))

	// appending multiple values to a slice using spread operator '...'
	slice2 := []int32{42, 39}
	intSlice = append(intSlice, slice2...)
	fmt.Printf("\nNew slice after appending another slice which is %v to the original slice using spread operator '...'", slice2)
	fmt.Printf("\nNew slice %v\nNew slice length %v\nNew slice capacity %v", intSlice, len(intSlice), cap(intSlice))
}