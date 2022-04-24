package interface_programing

import "fmt"

type Country2 struct {
	Name string
}

type City2 struct {
	Name string
}

type Stringable interface {
	ToString() string
}

func (c Country2) ToString() string {
	return "Country = " + c.Name
}

func (c City2) ToString() string {
	return "City = " + c.Name
}

// PrintStr
/**
我们使用了一个叫 Stringable 的接口，
我们用这个接口把 “业务类型” Country 和 City 和 “控制逻辑” Print() 给解耦了。
于是，只要实现了 Stringable 接口，都可以传给 PrintStr() 来使用
*/
func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
}

func main() {
	c1 := Country2{"USA"}
	c2 := City2{"Los Angeles"}
	PrintStr(c1)
	PrintStr(c2)
}
