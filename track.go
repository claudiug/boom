package main

import (
	"fmt"
)

func Greet(name string) {
	fmt.Println("Hello " + name)
}
func GreetNames(names []string, suffix string) {
	for _, x := range names {
		Greet(x + " " + suffix)
	}
}
func main() {
	comm := make(chan string)
	name := "claudiu"
	arr := []string{"claudiu", "raluca", "garba", "ion", "georgiana", "badoi"}
	fmt.Println(name)
	fmt.Println(arr)
	go func() {
		GreetNames(arr, "<C>")
		comm <- "finished greeting concurrently "
	}()
	GreetNames(arr, "<M>")
	fmt.Println(<-comm)
}
