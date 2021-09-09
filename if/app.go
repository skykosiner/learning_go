package main

import "fmt"

type Person struct {
    name string;
    age int;
    food int;
}

func main() {
    fmt.Println(Person.name)
}
