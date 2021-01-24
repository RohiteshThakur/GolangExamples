// update $GOPATH so external packages lands in this project.
// go get -u "github.com/fatih/structs"

package main

import (
    "fmt"

    structs "github.com/fatih/structs"
)

type A struct {
    Foo string
    Bar int
}

func main() {
    names := structs.Names(&A{})                   // if you don't know objects type just pass the address of object e.g. structs.Names(&A)
    fmt.Println(names) // ["Foo", "Bar"]
}