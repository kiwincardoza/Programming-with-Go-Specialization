package main

import (
    "fmt"
    "math"
    "sort"
    //"time"
)

func toSort(sli []int, ch chan []int) {
    sort.Ints(sli)
    fmt.Println(sli)
    ch<-sli
}

func main() {
    var sli []int
    fmt.Println("Enter the number of integers to be sorted:")
    var n int
    fmt.Scan(&n)
    fmt.Println("Enter the integers:")
    for i:=0;i<n;i++ {
        var inp int
        fmt.Scan(&inp)
        sli = append(sli, inp)
    }
    quartet := float64(float64(n)/4)
    fmt.Println(quartet)
    if quartet == 1.5 {
        quartet = 1.0
    }
    quartet1 := int(math.Round(quartet))
    fmt.Println(quartet1)
    ch := make(chan []int)
    var sorted_slice []int
    for j:=0;j<4;j++ {
        if j==3 {
            go toSort(sli[j*quartet1:n], ch)
        } else {
            go toSort(sli[j*quartet1:(j+1)*quartet1], ch)
        }
        temp_slice := <-ch
        sorted_slice = append(sorted_slice, temp_slice...)
    }
    //go toSort(sli[quartet*4:n], ch)
    //temp_slice := <-ch
    //sorted_slice = append(sorted_slice, temp_slice...)
    sort.Ints(sorted_slice)
    fmt.Println("Final sorted list: ", sorted_slice)
    //time.Sleep(30*time.Second)    
}
