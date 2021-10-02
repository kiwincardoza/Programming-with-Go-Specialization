package main

import "fmt"
import "strings"


func main () {
    var str1 string
    fmt.Println("Enter a string: ")
    fmt.Scan(&str1)
    if ((string(str1[0])=="i") || (string(str1[0])=="I")) && ((string(str1[len(str1)-1])=="n") || (string(str1[len(str1)-1])=="N")) && ((strings.Contains(str1, "a")) || (strings.Contains(str1, "A")))  {
        fmt.Println("Found!")
    } else {
        fmt.Println("Not Found!")
    }

}
