package main

import (
	"fmt"
)
/**
结构体嵌套，使用对就是匿名结构体
*/
type person struct {
	UserName string
	UserAge int
	Constact struct{
		Phone, City string
	}
}

func main() {
	/**
	  匿名结构体对应用案例
	*/
	st := &struct {
		Name string
		Age int
	}{
		Name : "liang",
		Age : 29,
	}
	fmt.Println(st)

	/**
	  机构体嵌套打印
	*/
	per := person{UserName: "liangyongxing", UserAge: 29}
	per.Constact.Phone = "15701183662"
	per.Constact.City = "北京"
	fmt.Println(per)
}