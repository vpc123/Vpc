package main

import (
	"testing"
)

func Test_Division_1(t *testing.T) {
	a := person{
		Name: "liang",
		Age:  29,
	}
	a= modifyPerson(a)
	if a.Age !=29 && a.Name !="liang"{
		t.Errorf("没有通过函数的单元测试，请确保函数的正确性！")
	}

}