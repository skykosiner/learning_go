package main;

import (
    "net/http"
    sayHello "example.com/hello_func"
    //"fmt"
);

func main() {
    http.HandleFunc("/", sayHello);
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    };
};
