package main

import (
    "fmt"
)


type Animal interface {
    Eat()
    Move()
    Speak()
    Name() string
}

type Cow struct {
    name string
    food string
    locomotion string
    noise string
}

type Bird struct {
    name string
    food string
    locomotion string
    noise string
}

type Snake struct {
    name string
    food string
    locomotion string
    noise string
}

func (c Cow) Eat() {
    fmt.Println(c.food)
}
func (c Cow) Move() {
    fmt.Println(c.locomotion)
}
func (c Cow) Speak() {
    fmt.Println(c.noise)
}
func (c Cow) Name() string {
    return c.name
}


func (b Bird) Eat() {
    fmt.Println(b.food)
}
func (b Bird) Move() {
    fmt.Println(b.locomotion)
}
func (b Bird) Speak() {
    fmt.Println(b.noise)
}
func (b Bird) Name() string {
    return b.name
}

func (s Snake) Eat() {
    fmt.Println(s.food)
}
func (s Snake) Move() {
    fmt.Println(s.locomotion)
}
func (s Snake) Speak() {
    fmt.Println(s.noise)
}
func (s Snake) Name() string {
    return s.name
}


func FindAnimal(animal_sli []Animal, name string) Animal {
    for _, animal := range animal_sli {
        if animal.Name() == name {
            return animal
            break
        }
    } 
    return nil
}


func main() {
    var animal_sli []Animal
    for {
        fmt.Println(">")
        var query_type string
        var name_or_type string
        var type_or_action string
        fmt.Scanf("%s %s %s", &query_type, &name_or_type, &type_or_action)
        if query_type == "newanimal" {
            if type_or_action == "cow" {
                animal_sli = append(animal_sli, Cow{name_or_type, "grass", "walk",
                "moo"})
                fmt.Println("Created it!")
            } else if type_or_action == "bird" {
                animal_sli = append(animal_sli, Bird{name_or_type, "worms", "fly",
                "peep"})
                fmt.Println("Created it!")
            } else if type_or_action == "snake" {
                animal_sli = append(animal_sli, Snake{name_or_type, "mice",
                "slither", "hsss"})
                fmt.Println("Created it!")
            }
        } else if query_type == "query" {
            var animal Animal
            animal = FindAnimal(animal_sli, name_or_type)
            if animal != nil {
                if type_or_action == "eat" {
                    animal.Eat()
                } else if type_or_action == "move" {
                    animal.Move()
                } else if type_or_action == "speak" {
                    animal.Speak()
                }
            }
        }
        //fmt.Println(animal_sli)

    }
}
