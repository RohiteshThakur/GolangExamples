package main

/*
	Note: In Order to use reflect package, the Methods of Struct MUST be exportable i.e. should start with capital letter (e.g. Addln).
		  else, program will run but return nothing because methods will deemed to be local and not visible externally.
*/

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age int
}

func (p Person) Addln (ln string) string {					/* function attached to an ob object -> Method */
	return (p.Name + " " + ln)
}
func (p Person) Person_age() int {
	return (p.Age)
}

func main (){
	p1 := Person{"Ro", 30}
	fmt.Println(p1.Addln("Thakur"))
	fmt.Println(p1.Person_age())

	perType := (reflect.TypeOf(Person{}))
	for i := 0; i < perType.NumMethod(); i++ {				/* Get number of methods in an object */
		method := perType.Method(i)							/* collect method name */
		fmt.Println(method.Name)							// Print it 
	}
}
