package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type User struct {
	FirstName string `json:"fname"`
	LastName  string `josn:"last`
	Password  string `json:"-"`
}

func (u User) MarshalJSON() ([]byte, error) {
	m := map[string]string{
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"full_name":  fmt.Sprintf("%s %s", u.FirstName, u.LastName),
	}
	return json.Marshal(m)
}

func main() {
	u := User{FirstName: "foo", LastName: "bar", Password: "foobar"}
	//encode
	m, _ := json.Marshal(u)
	fmt.Println(m)
	fmt.Println(string(m))
	fmt.Println(json.NewEncoder(os.Stdout).Encode(u))
	//decode
	s := `{"fname":"foo","LastName":"bar"}`
	fmt.Println(s)
	uu := User{}
	json.NewDecoder(strings.NewReader(s)).Decode(&uu)
	fmt.Println(uu)
	fmt.Println(string(u.MarshalJSON()))
}
