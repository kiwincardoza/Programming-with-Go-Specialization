package main

import (
   "fmt"
   "encoding/json"
)


func main() {
    var name string
    var address string
    fmt.Println("Enter user name: ")
    fmt.Scan(&name)
    fmt.Println("Enter user address: ")
    fmt.Scan(&address)
    //fmt.Println(name, address)
    //var idM map[string]string
    idM := map[string]string {name:address}
    barr, _ := json.Marshal(idM)
    fmt.Println(string(barr))
}

