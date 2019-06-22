### Mysql 使用完全解析



#### 1 Mysql 基础使用

1.1   创建数据库

> CREATE DATABASE  database-name

1.2   删除数据库

> drop  database dbname



> show databases;                # 查看当前Mysql都有那些数据，根目录都有那些文件夹

> create database 数据库名;   # 创建数据库文件夹

> use 数据库名;               # 使用选中数据库，进入目录

> show tables;                # 查看当前数据库下都有那些表，

> create table 表名(nid int,name varchar(20), pwd varchar(64)); # 创建数据库表

> drop table 表名                  # 删除整个表
> delete from 表名               # 删除表中的数据，保留表结构，可恢复  
> truncae table 表名          # 以更快的方式删除表中的数据，不可恢复。

> desc 标明                     #查看描述

> select * from 表名;       # 查看表中的所有数据

> insert into 表名(nid,name,pwd) values(1,'alex','123');  # 插入数据

> alter table user_info add constraint fk_u_p foreign key user_info(part_nid) references part(nid);  # 创建外键关联，下面有另一种方式（子表创建时就关联外键）

 

#### 连表操作:

> select * form a,b where a.x = b.o  join,

     a. left join  
         select * from a LEFT OUTER JOIN b ON a.x = b.o  
    
     b. inner join ,永远不会出现Null  
         select * from a inner JOIN b ON a.x = b.o

> select part_nid as a, count(nid) as b from userinfo group by part_nid;



#### 2    用户管理

创建用户  
            > create user '用户名'@'IP地址' identified by '密码';  
删除用户  
             >drop user '用户名'@'IP地址';  
修改用户  
             >rename user '用户名'@'IP地址'; to '新用户名'@'IP地址';;  
修改密码  
             >set password for '用户名'@'IP地址' = Password('新密码')  

> PS：用户权限相关数据保存在mysql数据库的user表中，所以也可以直接对其进行操作（不建议）
>  select User,Host,password_expired from user;



#### 3. 授权管理

`show grantsfor'用户'@'IP地址'-- 查看权限`

`grant  权限 on 数据库.表 to  '用户'@'IP地址'-- 授权`

`*GRANT SELECT ON `test`.`t1` TO 'lyy'@'%'*`

``revoke 权限 on 数据库.表 from`'用户'``@``'IP地址'``-- 取消权限`



> flush privileges，将数据读取到内存中，从而立即生效。



### 三、数据表级别常用SQL语句

#### 1 创建建新表

```
create table tb5(
    nid int not null auto_increment primary key,
   name varchar(16),
   age int default 19
   )engine=innodb default charset=utf8;

```

`注释：`

`primary-key //创建主键，为了加速查找。主键在表中唯一且不可重复、不能为null。`

`auto-increment//设为自增列，一个表中只能有一个自增列`

`default  //是设置默认值   update userinfo set part_nid=2;`

`engine=innodb //设置存储引擎`



**外键的作用有二：**

1. 软件工程的模块化思想，一张表不宜过大；

2. 可以用主表（这里也叫外表，主键表）来约束从表（外键表）；比如主表设置一个学号，子表的外键的学号字段就只能在这个范围内输入；

**外键说明：**

对于从表来说，外键不一定需要作为从表的主键，外键也不一定是外表的主键，外表的唯一键就可以作 为从表的外键。



### 2.  表内容操作

**1. 增**

`insert``into``表 (列名,列名...)``values``(值,值,值...)`

`insert``into``表 (列名,列名...)``values``(值,值,值...),(值,值,值...)`

`insert``into``表 (列名,列名...)``select``(列名,列名...)``from``表`



**2. 删**

`delete``from``表`

`delete``from``表``where``id＝1``and``name``＝``'jack’`



`**drop table表名 //整表删除**`

truncae table 表名 //整表删除



**3. 改**

`update``表``set``name``＝``'jack'``where``id>1  （修改表内容）`



表本身的修改：

`添加列：alter table 表名 add 列名 类型`

`删除列：alter table 表名 drop column 列名`

`修改列：`

`alter table 表名 modify column 列名 类型;  -- 类型`

`alter table 表名 change 原列名 新列名 类型; -- 列名，类型`

`添加主键：`

`alter table 表名 add primary key(列名);`

`删除主键：`

`alter table 表名 drop primary key;`

`alter table 表名  modify  列名 int, drop primary key;`

`添加外键：alter table 从表 add constraint 外键名称（形如：FK_从表_主表） foreign key 从表(外键字段) references 主表(主键字段);`

`删除外键：alter table 表名 drop foreign key 外键名称`

`修改默认值：ALTER TABLE testalter_tbl ALTER i SET DEFAULT 1000;`

`删除默认值：ALTER TABLE testalter_tbl ALTER i DROP DEFAULT;`



**4 查**

`select``*``from``表`

`select``*``from``表``where``id > 1`

`select``nid as num,``name``,gender``as``gg``from``表``where``id > 1`



更多查询操作（更新中）



`a、条件`

`select``*``from``表``where``id > 1``and``name``!=``'alex'``and``num = 12;`

`select``*``from``表``where``id``between``5``and``16;`

`select``*``from``表``where``id``in``(11,22,33)`

`select``*``from``表``where``id``not``in``(11,22,33)`

`select``*``from``表``where``id``in``(``select``nid``from``表)`

`b、通配符`

`select``*``from``表``where``name``like``'ale%'``- ale开头的所有（多个字符串）`

`select``*``from``表``where``name``like``'ale_'``- ale开头的所有（一个字符）`

`c、限制`

`select``*``from``表 limit 5;            - 前5行`

`select``*``from``表 limit 4,5;          - 从第4行开始的5行`

`select``*``from``表 limit 5 offset 4    - 从第4行开始的5行`

`d、排序`

`select``*``from``表``order``by``列``asc``- 根据 “列” 从小到大排列`

`select``*``from``表``order``by``列``desc``- 根据 “列” 从大到小排列`

`select``*``from``表``order``by``列1``desc``,列2``asc``- 根据 “列1” 从大到小排列，如果相同则按列2从小到大排序`

`e、分组`

`select``num``from``表``group``by``num`

`select``num,nid``from``表``group``by``num,nid`

`select``num,nid``from``表 ``where``nid > 10``group``by``num,nid``order``nid``desc`

`select``num,nid,``count``(*),``sum``(score),``max``(score),``min``(score)``from``表``group``by``num,nid`

`select``num``from``表``group``by``num``having``max``(id) > 10`

`特别的：``group``by``必须在``where``之后，``order``by``之前`

`f、连表`

`无对应关系则不显示`

`select``A.num, A.``name``, B.``name`

`from``A,B`

`Where``A.nid = B.nid`

`无对应关系则不显示`

`select``A.num, A.``name``, B.``name`

`from``A``inner``join``B`

`on``A.nid = B.nid`

`A表所有显示，如果B中无对应关系，则值为``null`

`select``A.num, A.``name``, B.``name`

`from``A``left``join``B`

`on``A.nid = B.nid`

`B表所有显示，如果B中无对应关系，则值为``null`

`select``A.num, A.``name``, B.``name`

`from``A``right``join``B`

`on``A.nid = B.nid`

`g、组合`

`组合，自动处理重合`

`select``nickname`

`from``A`

`union`

`select``name`

`from``B`

`组合，不处理重合`

`select``nickname`

`from``A`

`union``all`

`select``name`

`from``B`

`select part_nid as a, count(nid) as b from userinfo group by part_nid;`





总结：由于项目需要，实际开发的任务中经常需要操作mysql数据库，简单的增删改查，实际项目中多会进行操作接口的封装。所以mysql的一般性操作要熟知并掌握核心要领。




