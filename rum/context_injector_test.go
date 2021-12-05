package rum

import (
	"fmt"
	"testing"
)

type MyString string

func TestStringify(t *testing.T) {

	s := &Person{
		Name: "zhangsan",
		Age:  10,
	}

	r := stringify(s)
	fmt.Printf("testset-%s", r)
}

type Person struct {
	Name string
	Age  int
}
