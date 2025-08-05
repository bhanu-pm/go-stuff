//control stuctures go
package main
import (
	"fmt"
	"errors"
)


func main(){
	result := 13 * 3
	var err_val bool = false
	iferr := ifElseifStatements(result, err_val)
	normalSwitchStatement(result, iferr)
	conditionalSwitchStatement(result)
}

func ifElseifStatements(result int, err_val bool) (error){
	var err error
	if err_val{
		err = errors.New("There is an error created on purpose by the programmer.")
		return err
	}else{
		err = nil
	}
	if result == 0{
		fmt.Println("Evaluated with if-elseif function. Very small result.")
	}else if result == 1 || result == 2{
		fmt.Println("Evaluated with if-elseif function. 1 or 2.")
	}else if result == 39{
		fmt.Println("Evaluated with if-elseif function. Lucky number!")
	}else{
		fmt.Println("Evaluated with if-elseif function. Some not small number")
	}
	return err
}

func normalSwitchStatement(result int, err error){
	//Normal switch statement
	switch{
	case err != nil:
		fmt.Println("Eval with normal switch fn. There is some error in if Elseif fn")
	case result==0:
		fmt.Println("Eval with normal switch fn. Very small result.")
	case result==1 || result == 2:
		fmt.Println("Eval with normal switch fn. 1 or 2")
	case result==39:
		fmt.Println("Eval with normal switch fn. Lucky number")
	default:
		fmt.Println("Eval with normal switch fn. Something else")
	}
}

func condSwitchStatement(result int, err error){
	//Conditional swtich statement
	if err != nil{
		fmt.Println("Eval with conditional switch fn. There is some error in if Elseif fn")
	}else{
		switch result{
		case 0:
			fmt.Println("Eval with conditional switch fn. Very small result.")
		case 1:
			fmt.Println("Eval with conditional switch fn. 1")
		case 2:
			fmt.Println("Eval with conditional switch fn. 2")
		case 39:
			fmt.Println("Eval with conditional switch fn. Lucky Number!")
		}
	}
}