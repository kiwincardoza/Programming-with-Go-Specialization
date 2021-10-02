package main

import (
    "fmt"
    //"strings"
    "sort"
)


func main() {
    s1 := make([]int, 0)
    var inp int
    for {
        _, err := fmt.Scanf("%d", &inp)
        if err != nil {
            //fmt.Println("Error in reading integer: ", err)
            if err.Error() == "expected integer" {
                break
            }
        }
        //fmt.Println(inp)
        //break
        s1 = append(s1, inp)
        sort.Slice(s1, func(i, j int) bool {
                return s1[i] < s1[j]
        })
        fmt.Println(s1)
    }
}
