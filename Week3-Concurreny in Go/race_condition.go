package main

import (
    "fmt"
    "time"
)


/*
The below main() calls adder and subtractor functions as goroutines. Adder
increments the variable num by 1 and subtractor decrements num by 1. reader()
function prints the num variable value. Since both adder and subtractor alters
the num variable, it leads to a race condition where we cannot say for sure
after adder() is completed, subtractor() will run. This may lead to
unpredictable values for the variable 'num'
*/


func adder(a *int) {
    for i:=0;i<10000;i++ {
        time.Sleep(1 * time.Second)
        (*a)++
    }
}

func subtractor(a *int) {
    for i:=0;i<1000;i++ {
        time.Sleep(1 * time.Second)
        (*a)--
    }
}


func reader(a *int) {
    fmt.Println(*a)
}


func main() {
    var num int 
    num = 0
    go adder(&num)
    go subtractor(&num)
    //reader(&num)
    for {
        time.Sleep(2 * time.Second)
        reader(&num)
    }
}
