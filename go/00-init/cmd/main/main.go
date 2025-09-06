package main //special package, tell compiler main() is entry-pt

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

var log = fmt.Println

func basic() {
	fmt.Println("Hello")

	// var num int16 = 32767+1 --compile time overflow error

	var num int16 = 32767
	num += 1 //overflow that may happen at runtimes not detected
	fmt.Println(num)

	var num1 float32 = 12345678.9 //approx to 12345679.00, use 64 but for more precision
	fmt.Println(num1)

	var res float32 = float32(num) + num1 //have to type conv, since num1 is f32
	fmt.Println(res)

	fmt.Println(`hello
	there`)

	fmt.Println(len("Γ"))                    //2 bytes
	fmt.Println(utf8.RuneCountInString("Γ")) //get actual number of letters

	var myrune rune = 'a' //sngle quotes
	fmt.Println(myrune)   // 97 - ascii val of 'a'

	// var var1, var2 int16 = 2, 4; //works too

	const newConst string = "dfs"
	// const must be declared & *initialized*
	fmt.Println(newConst)

}

func div(numerator int, denom int) (int, int, error) {
	// can return multiple vals

	var err error //default nil
	if denom == 0 {
		err = errors.New("cant div by zero")
		return 0, 0, err
	}

	res := numerator / denom
	fmt.Printf("Result %v", res)
	remain := numerator % denom
	return res, remain, err
}

func arrays() {
	// only declare
	// var arrae []int32

	// declare n init , ... is for infer size
	var arr = [...]int32{1, 2, 3}

	fmt.Println(arr[:])  //gives all elem, arr[inclusive:exclusive]
	fmt.Println(&arr[1]) // gives address

	var x []string = []string{"sfs"} // slice
	fmt.Println(x)
	slice := append(x, "xys")
	log(slice)

	//append using spread operator
	newSliceAgain := append(x, slice...) //copy using spread operator
	log(newSliceAgain)

	var newSlice []string = make([]string, 4, 10) //len, capacity
	//an empty slice with capacity 10
	fmt.Println(newSlice)
	log(len(newSlice)) //set len of 4 and capacity 10 (hence wont realloc unless reaches end of space)

}

func slices() {
	mySlice := []string{"First", "Second", "Third"}
	newSlice := make([]string, 2) //less space , hence will copy only 2 elements, when u append to slice, itll hv to find a new contig memory
	copy(newSlice, mySlice)
	fmt.Println(newSlice)
}

func strings() {

}

func main() {
	// basic()

	/* fn and error handling
	{
		var err error
		res, remainder, err := div(5, 0)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("div: %v remainder: %v\n", res, remainder)
	}
	*/

	// arrays()

	// slices()

	// for loops-----------
	i := 0
	for {
		if i > 10 {
			break
		}
		if i == 4 {
			log("Nice")
		}
		i++
	}

	for i := 0; i < 5; i++ {
		// do somthing
	}

}
