package main_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"metacasts.tv/shared_actors"
)

func Test_AskForName(t *testing.T) {
	setup()

	a := assert.New(t)
	b := []byte("Mark\n")

	r := bytes.NewBuffer(b)

	main.AskForName(r)

	a.Equal(len(main.ActorNames), 1)
	a.Equal(main.ActorNames[0], "Mark")
}

func Test_AskForNames(t *testing.T) {
	setup()

	a := assert.New(t)
	b := []byte("Mark\nBates\nn\n")

	r := bytes.NewBuffer(b)

	main.AskForNames(r)

	a.Equal(len(main.ActorNames), 2)
	a.Equal(main.ActorNames[0], "Mark")
	a.Equal(main.ActorNames[1], "Bates")
}

func Test_AskForNames_FourNames(t *testing.T) {
	setup()

	a := assert.New(t)
	b := []byte("Mark\nBates\ny\nMary\ny\nSmith\nn\n")

	r := bytes.NewBuffer(b)

	main.AskForNames(r)

	a.Equal(len(main.ActorNames), 4)
	a.Equal(main.ActorNames[0], "Mark")
	a.Equal(main.ActorNames[1], "Bates")
	a.Equal(main.ActorNames[2], "Mary")
	a.Equal(main.ActorNames[3], "Smith")
}
func setup() {
	main.ActorNames = []string{}
}
