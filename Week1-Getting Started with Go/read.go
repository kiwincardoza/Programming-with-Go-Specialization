package main

import (
    "fmt"
    //"io/ioutil"
    "os"
    "bufio"
    "strings"
)

type Person struct {
    fname string
    lname string
}

func main() {
    fmt.Println("Enter the name of the file to read: ")
    var file_name string
    fmt.Scan(&file_name)

    file, err := os.Open(file_name)
    if err != nil {
        fmt.Println("Error in opening the file")
    }
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var persons []Person

    for scanner.Scan() {
        //fmt.Println(scanner.Text())
        full_name := scanner.Text()
        full_name_l := strings.Split(full_name, " ")
        p := Person{fname:full_name_l[0], lname:full_name_l[1]}
        persons = append(persons, p)
    }
    for _, person := range persons {
        fmt.Println("First Name: ", person.fname, "; Last Name: ", person.lname)
    }
    
    file.Close()
}

