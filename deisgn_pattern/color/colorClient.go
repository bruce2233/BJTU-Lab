package color

import (
	"design/factory"
	"flag"
	"fmt"
)

type Person struct {
	age int
}

func main() {

	// bob := Person{22}
	// chvalue(&bob)
	// print(bob.age)
	// fmt.Printf("%T",f.ColorStrFactory)
	var f = flag.Int("f", 8, "设置字体颜色")
	var b = flag.Int("b", 8, "设置背景颜色")
	var s = flag.Int("s", 1, "设置特效:    ")
	help:= "请使用-str参数输入字符串\n"
	help += "使用-f -b 修改字体和背景颜色, 0黑 1红  2绿  3黄  4蓝  5紫红  6青蓝  7白 8终端默认颜色\n"
	help += "使用-s 修改文字特效  0终端默认 1高亮显示  4使用下划线  5闪烁  7反白显示  8不可见"
	var originStr = flag.String("str",help, "原始字符串")

	flag.Parse()
	// fmt.Println(*f)
	fmt.Printf(factory.ColorStrFactory(*originStr, 30+*f, 40+*b, *s))
}

func chvalue(person *Person) {
	(*person).age++
}
