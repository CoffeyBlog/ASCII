// Storing a string as a slice of bytes
// And the ASCII coding scheme

package main

import "fmt"

func main() {

	str := "Hello"  // declares hello to str
	fmt.Println(str) // prints hello / str
	fmt.Printf("%T\n", str) // shows type

	byteslice := []byte(str) // puts str inside a slice of bytes
	fmt.Println(byteslice) // prints the bytes
	fmt.Printf("%T\n", byteslice) // shows the bytes in their ASCII coding scheme

}





