### Go语言 Strcut用法梳理

作者:流火夏梦                                  时间:2019-06-22

#### 1  Go语言结构体基础

```
　　1.Go中的struct与C中的struct非常相似，并且Go没有class

　　2.使用 type <Name> struct{} 定义结构，名称遵循可见性规则(即首字母大写对外可见)。 type person struct{}

　　3.支持指向自身的指针类型成员，支持匿名结构，可用作成员或定义成员变量

　　4.匿名结构也可以用于struct的值，可以使用字面值对结构进行初始化

　　5.允许直接通过指针来读写结构成员

　　6.相同类型的成员可进行直接拷贝赋值

　　7.支持 == 与 != 比较运算符，但不支持 > 或 <

　　8.支持匿名字段，本质上是定义了以某个类型名为名称的字段

　　9.嵌入结构作为匿名字段看起来像继承，但不是继承
　　　　
　　10. 可以使用匿名字段指针

```

一、结构体的基本使用案例



说明：下面的案例是将struct的值传递和指针传递作为样式进行验证，使用struct的形式仿照class的构建方法和模式进行验证。

相关代码在文档同级目录下的(go-struct(go语言结构体函数代码)目录中)

说明:

```
#定义结构体:
type person struct {
    Name string Age int 
}
```

1   值传递进行修改

> modifyPerson(a)

2    指针传递进行修改

> modifyPersonPointer(&a)

通过代码我们可以看出值传递无法对参数进行修改，但是指针传递却可以做到。

相关的代码在目录文件中main函数中有说明实现，并在main_test.go中进行函数方法测试。



**二、匿名结构体以及结构体内嵌案例**

匿名结构体:

```
type person struct {
    UserName string
    UserAge int
    Constact struct{
        Phone, City string
         }
}
```



调用匿名结构体:



    st := &struct {
        Name string
        Age int
    }{
        Name : "liang",
        Age : 29,
    }
    fmt.Println(st)

**三、结构体的内嵌组合模拟继承案例**



嵌套结构体模拟继承：



```
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

```



嵌套结构体的调用：



```
tea := teacher{Name: "teacher", Age: 36, Human: Human{sex: 1}}
    // tea.Human.sex = 1 或者 tea.sex = 1
stu := Student{Name: "student", Age: 15, Human: Human{sex: 2}}

```



总结：结构体的使用形式及用法有如下几种方式

1   一般定义使用

2   匿名使用方式

3   嵌套模拟继承

总的来说，Go语言的Struct类型对go语言的开发设计是很关键的，struct虽然没有那么多的方法，但是本身的简洁性 和易用性就是go语言的一个很大的利器。深入理解并掌握struct的使用方法对之后的开发设计是十分必要的，我们同样可以从struct的设计实现理解go语言的大道至简的设计模式。





















