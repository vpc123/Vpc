### Go语言依赖的gcc环境配置



##### 1 使用go语言进行开发时出现没有gcc环境

##### 2 或者遇到如下情况时

```
github.com/wendal/go-oci8

cc1.exe: sorry, unimplemented: 64-bit mode not compiled in

```



解决办法：
（1）解压统计目录下的压缩包文件x86_64-4.8.2-release-posix-seh-rt_v3-rev2.7z

（2）拷贝目录文件夹到除C盘以外的任意位置

![5cf91a97a49ba19599](https://i.loli.net/2019/06/06/5cf91a97a49ba19599.png)



（3）将bin目录加入到环境变量中并重启电脑

例如：我的环境变量路径是

```
D:\标配软件\mingw64\bin
```

(4)  检查是否有效

![5cf91b698695835452](https://i.loli.net/2019/06/06/5cf91b698695835452.png)





总结：

go语言进行编译开发需要依赖gcc的环境，但是windows并不是天生支持这样的开发环境，所以一开始就需要我们进行配置gcc环境从而支持go语言的开发设计任务。这是开发的一个小坑，千万需要留心。
