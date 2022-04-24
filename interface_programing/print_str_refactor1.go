package interface_programing

import "fmt"

// WithName 引入WithName结构体
type WithName struct {
	Name string
}

type Country1 struct {
	WithName
}

type City1 struct {
	WithName
}

type Printable1 interface {
	PrintStr()
}

func (w WithName) PrintStr() {
	fmt.Println(w.Name)
}

func main() {
	c1 := Country1{WithName{"China"}}
	c2 := City1{WithName{"Beijing"}}
	c1.PrintStr()
	c2.PrintStr()
}
