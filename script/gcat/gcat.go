package main

import (
    "fmt"
    "io/ioutil"
    "flag"
)

func main() {

    flag.Parse()
    var filename = flag.Args()[0]

    if contents,err := ioutil.ReadFile(filename);err != nil{
        fmt.Println(err)
    } else {
        fmt.Printf("%s\n",contents)
    }   
}