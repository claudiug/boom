package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ActorsNames = []string{}

func AskForName() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please Enter an actor name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println(name)
}

func main() {
	AskForName()
}
