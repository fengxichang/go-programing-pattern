package interface_programing

import "fmt"

type Shape interface {
	Sides() int
	Area() int
}

type Squara struct {
	len int
}

/**
声明一个 _ 变量（没人用），其会把一个 nil 的空指针，从 Square 转成 Shape,
这样，如果没有实现完相关的接口方法，编译器就会报错
*/
var _ Shape = (*Squara)(nil)

func (s Squara) Sides() int {
	return s.len * 2
}

func (s Squara) Area() int {
	return s.len * s.len
}

/**
我们可以看到 Square 并没有实现 Shape 接口的所有方法，
程序虽然可以跑通，但是这样编程的方式并不严谨
*/
func main() {
	s := Squara{len: 5}
	fmt.Printf("%d\n", s.Sides())
}
