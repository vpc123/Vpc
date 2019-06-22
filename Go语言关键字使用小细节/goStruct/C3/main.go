package main

import (
	"fmt"
)

/**
这里说对是结构体对组合，它对功能类似于其他语言对继承
*/
type Human struct {
	sex int
}

type teacher struct {
	Human
	Name string
	Age int
}

type Student struct {
	Human
	Name string
	Age int
}

func main() {
	/**
	  在初始化对时候Go将嵌入对结构名称当成属性一样对待，将对应Human作为属性，这样可以在初始化对时候直接赋值
	  第二种赋值方式可以通过符号 . 来操作赋值
	*/
	tea := teacher{Name: "teacher", Age: 36, Human: Human{sex: 1}}
	// tea.Human.sex = 1 或者 tea.sex = 1
	stu := Student{Name: "student", Age: 15, Human: Human{sex: 2}}
	//stu.Human.sex = 2  或者 tea.sex = 2
	/**
	  1. 既然结构嵌入进来来，就和其他语言继承一样，可以直接使用父类对属性，即 tea.sex = 1 也是可以对
	  2. 那 1 方式简单为什么还要保留 tea.Human.sex = 1 这种方式呢？是因为为来防止外部引用有同名对属性，为了区分
	*/
	fmt.Println(tea)
	fmt.Println(stu)
}