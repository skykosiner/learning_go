package main;

import (
    "fmt"
    hello "example.com/hello"
)

func main() {
    hallo := hello.Name("Yoni");
    fmt.Println(hallo)

    var n [10]int
    var i,j int

    for i = 0; i < 10; i++ {
        n[i] = i + 1
    }

    for j = 0; j < 10; j++ {
        fmt.Printf("Element[%d] = %d\n", j, n[j] )
    }
};
