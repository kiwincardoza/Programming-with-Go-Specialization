package main


import (
    "fmt"

)


func Swap(s1 []int, ind1 int) {
    temp := s1[ind1]
    s1[ind1] = s1[ind1+1]
    s1[ind1+1] = temp
}


func BubbleSort(s1 []int) {
    for i:=0; i<len(s1); i++ {
        for j:=0; j<len(s1)-i-1; j++ {
            if (s1[j] > s1[j+1]) {
                Swap(s1, j)
                //temp := s1[j]
                //s1[j] = s1[j+1]
                //s1[j+1] = temp
            }
        }
    }
}


func main() {
    s1 := make([]int, 0)
    var inp int
    for i:=0; i<10; i++ {
        fmt.Println("Enter an integer: ")
        _, err := fmt.Scanf("%d", &inp)
        if err != nil {
            fmt.Println("Error in reading the integer-", err)
        }
        s1 = append(s1, inp)
    }
    BubbleSort(s1)
    fmt.Println(s1)
}
