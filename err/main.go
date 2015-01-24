package main

import (
	"errors"
	"log"
)

func main() {
	err := Boom()
	if err != nil {
		log.Println(err)
	}
}

func Boom() error {
	return errors.New("Boom!")
}
