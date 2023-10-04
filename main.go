package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Trevor",
		LastName:    "Sawler",
		PhoneNumber: "1-555-111-2222",
	}

	log.Println(user.FirstName, user.LastName, user.PhoneNumber)

}

/*
This program demonstrates the following main features of Go:

    Structs: Structs are a way to define custom data types in Go. They are composed of named fields, which can be of any type.
    Functions: Functions are first-class citizens in Go, meaning that they can be passed around and assigned to variables.
    Methods: Methods are functions that are associated with a struct type. They can be used to operate on instances of that struct type.
    Goroutines and channels: Goroutines are lightweight threads of execution in Go. Channels are a way to communicate between goroutines.
    Interfaces: Interfaces are a way to define contracts between types. They can be used to implement polymorphism in Go.
    Pointers: Pointers are a way to reference values in memory. They are useful for passing values to functions by reference and for modifying values that are stored in other variables.
    Slices: Slices are a dynamic array in Go. They can be resized as needed and they are efficient for storing and accessing data.
    Maps: Maps are a key-value store in Go. They can be used to store any type of data, both as keys and values.
    Error handling: Go provides a built-in error handling mechanism that allows you to handle errors gracefully.
    Higher-order functions: Higher-order functions are functions that take other functions as arguments or return functions as results. Go supports higher-order functions, which makes it a powerful language for functional programming.
    Closures: Closures are functions that have access to the variables in the scope in which they were created, even after the function that created them has returned. Go supports closures, which makes it a powerful language for writing code that is concise and easy to read.
    Defer: Defer allows you to execute code after a function has returned. This can be useful for cleaning up resources or ensuring that certain operations are always performed.
    Waitgroups: Waitgroups allow you to coordinate the execution of multiple goroutines. This can be useful for ensuring that all goroutines have finished executing before your program exits.

*/
