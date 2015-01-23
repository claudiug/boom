package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
)

func Add1(x int, y int) int {
	return x + y
}

func Swap(x string, y string) (string, string) {
	return y, x
}

type Vertex struct {
	X int
	Y int
}

func (v *Vertex) Abs() float64 {
	return math.Abs(v.X, v.Y)
}

func foo() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	fmt.Println("hello")
	fmt.Println(math.Pi)
	fmt.Println(rand.Intn(100))
	result := Add(3, 100)
	fmt.Println(result)
	fmt.Println(Swap("x", "y"))
	fmt.Printf("v is the type %T\n", i)
	const (
		DEMO = "demo"
		LIFE = 1
	)
	fmt.Println(DEMO)
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("this is mac")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("no")
	}
	j := 100
	cc := &j
	fmt.Println(cc)
	*cc = 21
	fmt.Println(j)
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	p := &v
	p.X = 110
	fmt.Println(v)
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		cl = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, v2, v3, cl)
	var arr [3]string
	arr[1] = "demo"
	fmt.Println(arr)
	sli := []int{2, 3, 5, 7, 11, 13}
	for i := 0; i < len(sli); i++ {
		fmt.Println(sli[i])
	}
	fmt.Println(sli[:4])
	a := make([]int, 10)
	var sss []int
	fmt.Println(cap(a))
	sss = append(a, 10)
	fmt.Println(sss)
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40, 30,
	}
	fmt.Println(m)

	ddd := func(x int, y int) bool {
		return x == y
	}
	fmt.Println(ddd(3, 3))

	fmt.Println(v.Abs())
}
