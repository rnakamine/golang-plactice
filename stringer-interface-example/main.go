package main

import "fmt"

type Stringer interface {
	String() string
}

type I int

func (i I) String() string {
	return "INT"
}

type F float64

func (f F) String() string {
	return "FLOAT"
}

type B bool

func (b B) String() string {
	return "BOOL"
}

func checkType(s Stringer) {
	switch v := s.(type) {
	case I:
		fmt.Printf("%T %d\n", v, v)
	case F:
		fmt.Printf("%T %g\n", v, v)
	case B:
		fmt.Printf("%T %t\n", v, v)
	}
}

func main() {
	i := I(10)
	checkType(i)
	f := F(10.5)
	checkType(f)
	b := B(true)
	checkType(b)
}
