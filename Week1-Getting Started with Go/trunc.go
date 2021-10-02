package main

import (
    "fmt"
    "math"
)

func main() {
   var fl float64
   fmt.Println("Enter a floating point number:")
   fmt.Scan(&fl)
   res := math.Trunc(fl)
   fmt.Println("Truncated value:", res)
}
