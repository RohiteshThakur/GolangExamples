package main

import (
	"fmt"
	"encoding/json"
	//"reflect"
)

type Book struct {
	Title	string
	Author	string
	Pages	int
}

// JSON.Marshal(): Takes Struct and returns a byte array. 
func (b Book) ToJSON() []byte {									// No need to pass object as method is attached to Book struct anyway.						
	byte_array, err := json.Marshal(b)
	if (err != nil){
		panic(err)
	}
	return (byte_array)
}



func main() {
	b := []Book{
		Book{Title: "Python", Author: "Ro", Pages:500},
		Book{Title: "Golang", Author: "RoT", Pages: 400},			// Notice the "," in the end. This is not JSON but a struct.
	}
	//b := Book{Title: "Python", Author: "Ro", Pages:500}
	// result := reflect.TypeOf(b)    -> []uint8  (Because its a Byte array of integers)
	fmt.Println(string(b[1].ToJSON()))
	
}
