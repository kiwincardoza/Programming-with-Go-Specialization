package main

import (
    "fmt"
)


func GenDisplaceFn(acc float64, vel float64, dis float64) func (float64) float64 {
    fn := func (time_inp float64) float64 {
        return (0.5*acc*time_inp*time_inp) + (vel*time_inp) + (dis)
    }
    return fn
} 


func main() {
    fmt.Println("Please enter acceleration value: ")
    var acc float64
    fmt.Scan(&acc)
    fmt.Println("Please enter initial velocity value: ")
    var vel float64
    fmt.Scan(&vel)
    fmt.Println("Please enter initial displacement value: ")
    var dis float64
    fmt.Scan(&dis)
    fmt.Println("Enter the time: ")
    var time_inp float64
    fmt.Scan(&time_inp)
    fn := GenDisplaceFn(acc, vel, dis)
    fmt.Println(fn(time_inp))
}
