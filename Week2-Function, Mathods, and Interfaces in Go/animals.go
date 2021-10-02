package main 

import (
    "fmt"
)

type Animal struct {
    food string
    locomotion string
    noise string
}


func (a Animal) Eat() {
    fmt.Println(a.food)
}

func (a Animal) Move() {
    fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
    fmt.Println(a.noise)
}


func main () {
    cow := Animal{"grass", "walk", "moo"}
    bird := Animal{"worms", "fly", "peep"}
    snake := Animal{"mice", "slither", "hsss"}
    for {
        fmt.Println(">")
        var name string
        var action string
        fmt.Scanf("%s %s", &name, &action)
        //input_slice := strings.Split(input, " ")
        //name := input_slice[0]
        //action := input_slice[1] 
        
        if name == "cow" {
            if action == "eat" {
                cow.Eat()
            }
            if action == "move" {
                cow.Move()
            }
            if action == "speak" {
                cow.Speak()
            }
        }
        if name == "bird" {
            if action == "eat" {
                bird.Eat()
            }
            if action == "move" {
                bird.Move()
            }
            if action == "speak" {
                bird.Speak()
            }
        }
        if name == "snake" {
            if action == "eat" {
                snake.Eat()
            }
            if action == "move" {
                snake.Move()
            }
            if action == "speak" {
                snake.Speak()
            }
        }

    }
}
