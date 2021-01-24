package main

import (
	"fmt"
	"encoding/json"
	//"reflect"
)


// JSON.Unmarshal(): Takes JSON data and returns Go Data Structures.

// From JSON to Dictionary. Reads 
func FromJSON_To_Map() map[string]interface{} {				// Parses JSON data in returns dictionary

	//byt := []byte(`{"num":6.13,"strs":["a","b"]}`)			// Note: This is a Byte array. So we can pass this directly to json.Unmarshal()
	str := `{"num":6.13,"strs":["a","b"]}`
	var dat map[string]interface{}							// map whose keys are string and values are interface.

	/*
	err := json.Unmarshal(byt, &dat)
	if (err != nil){
		panic(err)
	}
	*/
	srr := json.Unmarshal([]byte(str), &dat)
	if (srr != nil){
		panic(srr)
	}
	return dat
}


// From JSON to Struct. This methods helps to tag and pick values from JSON, we are interested in. 
// Note: JSON data is in dict format and of type string and not of type byte Array.
type data struct {
	Page	int			`json:"num"`
	Fruits	[]string	`json:"fruits"`		// These tags MUST match with incoming JSON fields or else, returned data will be nil. (no errors thrown)
}	

func FromJSON_To_Struct() data {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`		// Note: this is "NOT" a Byte array. So we will have to cast it before passing.
	dat := data{}
	err := json.Unmarshal([]byte(str), &dat)
	if (err != nil){
		panic(err)
	}
	return dat
}


// Parsing JSON data which is an array of dictionaries.

type array_data struct {
	Page	int			`json:"page"`
	Fruits	[]string	`json:"fruits"`		// These tags MUST match with incoming JSON fields or else, returned data will be nil. (no errors thrown)
}	

func FromJSONArray_To_Struct() []array_data {				// We can certainly pass an Array of struct as return variable.
	str := `[{"page": 1, "fruits": ["apple", "peach"]},		
			 {"page": 2, "fruits": ["mango", "banana"]},
			 {"page": 6, "fruits": ["melon", "grapes"]}]`	// No "," in the last dict element.

	dat := []array_data{}						// Instantiate and initialize a variable of type []array_data.
	err := json.Unmarshal([]byte(str), &dat)	// JSON passed a Byte array. So we will have to cast "str" before passing.
	if (err != nil){
		panic(err)
	}
	return dat									// Note this is an Array of Struct: array_data.
}



func main(){
	return_data := FromJSON_To_Map()
	fmt.Println(return_data)
	fmt.Println(return_data["num"])
	fmt.Println(return_data["strs"])
	//fmt.Println(return_data["strs"].([]string)[0]) -> runtime error (interface conversion: interface {} is []interface {}, not []string)
	fmt.Println(return_data["strs"].([]interface{})[0])		// Accessing dict values which are array.
	fmt.Println()

	response := FromJSON_To_Struct()
	fmt.Println(response)
	fmt.Println(response.Page)
	fmt.Println(response.Fruits)
	fmt.Println(response.Fruits[0])

	ret := FromJSONArray_To_Struct()
	fmt.Println(ret)				// This is an slice of Struct array_data
	fmt.Println(ret[0])				// Because its a slice, we can index it.
	fmt.Println(ret[0].Page)		// To display particular value, select a particular struct in array and access it element using tags
	fmt.Println(ret[0].Fruits)

	}
	