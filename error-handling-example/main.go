package main

import "fmt"

type Stringer interface {
	String() string
}

func ToStringer(v interface{}) (Stringer, error) {
	if i, ok := v.(Stringer); ok {
		return i, nil
	}
	return nil, MyError("CastError")
}

type MyError string

func (e MyError) Error() string {
	return string(e)
}

type I int

func (i I) String() string {
	return "I"
}

func main() {
	i := I(100)
	s, err := ToStringer(i)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
