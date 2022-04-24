package interface_programing

import "fmt"

type Country struct {
	Name string
}

type City struct {
	Name string
}

// Printable 面向接口编程
type Printable interface {
	PrintStr()
}

func (c Country) PrintStr() {
	fmt.Println(c.Name)
}

func (c City) PrintStr() {
	fmt.Println(c.Name)
}

func main() {
	c1 := Country{"China"}
	c2 := City{"ShangHai"}

	c1.PrintStr()
	c2.PrintStr()
}
