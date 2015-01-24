package main

import (
	"fmt"
)

const (
	FileClosedSaved = 1
	FileNotClosed   = 2
	FileNotFound    = 3
)

type Stringer interface {
	String() string
}

type fakeString struct {
	content string
}

func (s *fakeString) String() string {
	return s.content
}

func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())
	}
}

type Artist struct {
	Name  string
	Songs int
}

func newRelease(artist Artist) int {
	artist.Songs++
	return artist.Songs
}
func newRelease1(artist *Artist) int {
	artist.Songs++
	return artist.Songs
}

func main() {
	foo := func(s string) int {
		return len(s)
	}
	fmt.Println("le")

	result := foo("claudiu")
	fmt.Println(result)
	fmt.Println(FileNotFound)
	city, loc := Location("Berlin")
	fmt.Println(city)
	fmt.Println(loc)
	fmt.Sprintf("The region for %s and the city is %s ", city, loc)

	me := Artist{Name: "Raluca", Songs: 43}
	me1 := &Artist{Name: "Claudiu", Songs: 43}
	fmt.Printf("%s was release they song %d\n", me.Name, me.Songs)
	fmt.Printf("%s was release they song %d\n", me1.Name, me1.Songs)
	newRelease(me)
	newRelease1(me1)
	fmt.Println(me.Songs)
	fmt.Println(me1.Songs)
	s := &fakeString{"Ceci n'est pas un string"}
	printString(s)
	printString("Hello, Gophers")
}

func Location(city string) (string, string) {
	var region string
	var country string
	switch city {
	case "Berlin":
		region, country = "Berlin", "Berlin"
	case "Munchen":
		region, country = "Bavaria", "Munchen"
	default:
		region, country = "no", "no"
	}
	return region, country
}
