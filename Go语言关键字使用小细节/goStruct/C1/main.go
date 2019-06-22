package main

import "fmt"

/**
定义一个结构体，类似于其他语言的class
*/
type person struct {
	Name string
	Age  int
}

func main() {
	// struct 对属性操作使用符号 . ;这里赋值采用两种方式，既可以使用默认值，也可以在外部赋值
	a := person{
		Name : "liang",
		Age : 29,
	}
	//a.Age = 29
	fmt.Println("a修改前：", a)
	//第一次修改并打印
	a=modifyPerson(a)
	fmt.Println("a第一次修改后：", a)


	/**
	  假如有一种场景有很多需要修改person内容，那么每次传入都需要取地址符号，这样很麻烦，可以在赋值对时候直接取得对应对地址
	  这种方式是开发对推荐方式
	*/
	b := &person{
		Name : "xuli",
		Age : 27,
	}
	fmt.Println("b修改前：", b)
	b=modifyPersonPointer(b)
	fmt.Println("b修改后：", b)
}

/**
从打印结果可以看出这里传入对是值类型，修改person内容并不会修改person原始值
*/
func modifyPerson(per person) (Per person){
	per.Age = 18
	fmt.Println("修改时：", per)
	return per
}

func modifyPersonPointer(per *person) (Per *person){
	per.Age = 19
	fmt.Println("修改时：", per)
	return per
}