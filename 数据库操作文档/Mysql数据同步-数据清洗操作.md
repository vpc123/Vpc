## MySQL数据库的清洗流程



### MySQL数据库经典50题



1.student(sid,sname,ssex,sage,stel,saddress)



![5d15f5da7dbf331877](https://i.loli.net/2019/06/28/5d15f5da7dbf331877.png)

2.course(cid,course,tname,tid)



![5d15f5f1ae70036698](https://i.loli.net/2019/06/28/5d15f5f1ae70036698.png)

3.sc(sid,sname,cid,score)

![5d15f5fd2b54952020](https://i.loli.net/2019/06/28/5d15f5fd2b54952020.png)





4.teacher(tid,tname)

![5d15f605bd0fa80646](https://i.loli.net/2019/06/28/5d15f605bd0fa80646.png)

##### 1.查询“2018001”课程比“2018002”课程成绩高的所有学生的学号，姓名。

    $select a.sid,a.sname from 
    (select sid,sname,score from sc where cid = "2018001") as a,
    (select sid,sname,score from sc where cid = "2018002") as b
    where a.sid = b.sid and a.score > b.score;

##### 2.查询平均分大于60分的同学的学号，姓名和平均分。

```
$select sid,sname,avg(score)  from sc group  by sid having  avg(score)>60;
```

2.1.avg 返回数值列的平均值  
*语法：*

```
$select avg(列名) from 表名
```

2.2.group by 用于结合“聚合函数”，根据一个或多个列对结果集进行分组  
“聚合函数”（例如：sum,avg）常常需要添加group by 语句  
*语法：*

```
$SELECT column_name, aggregate_function(column_name) FROM table_name WHERE column_name operator value GROUP BY column_name
```

2.3.having 在 SQL 中增加 having 子句原因是,where关键字无法与“聚合函数”一起使用  
*语法：*

```
$SELECT column_name, aggregate_function(column_name)
FROM table_name
WHERE column_name operator value
GROUP BY column_name
HAVING aggregate_function(column_name) operator value  
```

##### 3.查询所有同学的学号，姓名，选课数，总成绩。

```
$select student.sid,student.sname,count(sc.cid),sum(score)  from student left  join sc on student.sid = sc.sid group  by student.sid,sname;
```

3.1.left join 关键字会从 左表 那里返回所有的行，即使在 右表 中没有匹配的行。  
*语法：*

```
$select 列 from 左表 left join 右表 on 左表.列 = 右表.列 order by 列
```

若有“聚合函数”（例如:sum,avg,count等）

```
$select 列 sum(列) from 左表 left join 右表 on 左表.列 = 右表.列 group by 列
```

*注：在某些数据库中， left join 称为 left outer join*

3.2.order by 与group by 的区别
（1）order by 从英文里理解就是行的排序方式，默认的为升序。 order by 后面必须列出排序的字段名，可以是多个字段名。
（2）group by 从英文里理解就是分组。必须有“聚合函数”来配合才能使用，使用时至少需要一个分组标志字段。

##### 注意：聚合函数是—sum()、count()、avg()、max()、min()等都是“聚合函数”。

##### 4.查询姓“风”的老师的个数。

> $select  count(tname)  from teacher where tname like  "风%";

##### 5.查询学过“风子”老师所教的所有课的同学的学号，姓名。

    $select sid,sname from student
    where sid in (select sid from sc,course,teacher 
    where sc.cid = course.cid and teacher.tid = course.tid
    and teacher.tname = "风子" group by sid
    having count(sc.cid) = (select count(cid) from course,teacher
    where teacher.tid = course.tid and teacher.tname = "风子"));

5.1.in允许在where子句中规定多个值  
5.2.select……from……where……group by …… having……

需要注意having和where的用法区别：  
（1）having只能用在group by之后，对分组后的结果进行筛选(即使用having的前提条件是分组)。  
（2）where肯定在group by 之前。  
（3）where后的条件表达式里不允许使用聚合函数，而having可以

5.3.当一个查询语句同时出现了where……group by……having……order by……的时候，执行顺序和编写顺序是：
1.执行where xx对全表数据做筛选，返回第1个结果集。
2.针对第1个结果集使用group by分组，返回第2个结果集。
3.针对第2个结果集中的每1组数据执行select xx，有几组就执行几次，返回第3个结果集。
4.针对第3个结集执行having xx进行筛选，返回第4个结果集。

5.针对第4个结果集排序。

##### 6.查询没有学过“风子”老师课的同学的学号、姓名。

```
$select student.sid,student.sname from student
where sid not in (select distinct sc.sid from sc,course,teacher
where sc.cid = course.cid and teacher.tid = course.tid and teacher.tname = "风子");
```

1.distinct 用于返回唯一不同的值

##### 7.查询学过“2018002”也学过“2018003”的学生的学号，姓名。

    $select student.sid,student.sname from student,sc
    where student.sid = sc.sid and sc.cid = "2018002"
    and exists (select * from sc as sc2 where sc2.sid = sc.sid and sc2.cid = "2018003");

1.exists 用于检查子查询是否至少会返回一行数据，该子查询实际上并不返回任何数据，而是返回值True或False  
exists 指定一个子查询，检测 行 的存在。

##### 8.查询课程编号“2018002”的成绩比课程编号“2018003”课程低的所有同学的学号、姓名。

```
$select a.sid,a.sname from
    (select sid,sname,score from sc where cid = "2018002") as a,
    (select sid,sname,score from sc where cid = "2018003") as b
    where a.sid = b.sid and a.score < b.score;
```

##### 9.查询所有课程成绩都小于60的同学的学号、姓名。

```
$select sid,sname from student where sid not  in  (select student.sid from student,sc where student.sid = sc.sid and score >  60);
```

##### 10.查询没有学全所有课的同学的学号、姓名。

```
$select student.sid,student.sname from student,sc
where student.sid = sc.sid
group by student.sid,student.sname
having count(cid)<(select count(cid) from course);
```

##### 11.查询至少有一门课与学号为“70602”同学所学相同的同学的学号和姓名。

    $select student.sid,student.sname from student,sc
    where student.sid = sc.sid
    and cid in (select cid from sc where sid = "70602")
    group by student.sid;

##### 12.查询至少学过学号为“70602”同学所学一门课的其他同学学号和姓名。

    $select distinct student.sid,student.sname from student,sc
    where student.sid = sc.sid
    and cid in (select cid from sc where sid = "70602");

##### 13.把“SC”表中“风子”老师教的课的成绩都更改为此课程的平均成绩。

哎，这题解不出来很是忧伤，只能查出平均成绩，就是换不进去……

```
$select  avg(a.score)as avgscore from  (select score FROM sc where cid IN  (select  DISTINCT cid from sc,teacher where tname='风子'))a
```

13.1.update 修改表中的数据  
语法：

> update 表 set 列 = 新值 where 列 = 某值

##### 14.查询和“70603”号的同学学习的课程完全相同的其他同学学号和姓名。

```
$select sid,sname from sc where cid IN
    (select cid from sc where sid = "70603")
    group by sid having count(*) = (select count(*) from sc where sid = "70603");
```

*以上代码实现的功能不完整，如果“70603”号的同学学习的课程是其它同学学习的课程的子集也会筛选出来，因此完善为如下代码：*

    $select a.sid,a.sname from 
    (select sid,sname,count(distinct cid) as cnt1 from sc where cid in 
    (select cid from sc where sid = "70603") and sid <> "70603" group by sid
    having count(distinct cid) = 
    (select count(distinct cid) from sc where sid = "70603")) as a,
    (select sid,count(distinct cid) as cnt2 from sc sc group by sid) as b
    where a.sid = b.sid and a.cnt1 = b.cnt2

##### 15.删除学习“风子”老师课的SC表记录。

```
$delete sc from sc,course,teacher where course.cid = sc.cid and course.tid = teacher.tid and teacher.tname =  "风子";
```

15.1.delete 删除表中的行  
语法

> $delete from 表 where 列 = 值

删除所有行，但不删除表，保留表结构，索引和属性：

> delete from 表; 或者：delete * from 表;

##### 16.向SC表中插入一些记录，这些记录要求符合以下条件：没有上过编号“2018001”课程的同学学号、2018002号课的平均成绩。

```
$insert into sc (select sid,"002all","2018002",
(select avg(score) from sc where cid = "2018002") from student where sid not in (select sid from sc where cid = "2018003"));
```

16.1.错误：Column count doesn’t match value count at row 1  
是由于类似

> $insert into table_name(col_name1, col_name2, col_name3) VALUES('value1','value2');

语句中，前后列数不等造成的。

2.insert into 像表格中插入新的行。  
语法：

> insert into 表 values（值1，值2，值3，……）

或者指定需要插入的数据列：

> insert into 表（列1，列2，……）values （值1，值2，……）

##### 17.按平均成绩从高到低显示所有学生的“语文”、“数学”、“英语”三门的课程成绩，按如下形式显示：学生ID，语文，数学，英语，有效课程数，有效平均分。

这题又不会了……，只能做到这样子了，然而以下的查询也是有问题的，只有大家都学了的情况下才能查出来，假如有一门没有学的话也查不出来的，那这时候该怎么办呢，忧伤+2。

```
$select aa.学号,aa.语文,bb.数学,cc.英语 from
    (select sc.sid as 学号,score as 语文 from sc,student where sc.sid = student.sid and cid = 2018001) as aa,
    (select sc.sid as 学号,score as 数学 from sc,student where sc.sid = student.sid and cid = 2018002) as bb,
    (select sc.sid as 学号,score as 英语 from sc,student where sc.sid = student.sid and cid = 2018003) as cc
where aa.学号 = bb.学号 and bb.学号 = cc.学号;
```

##### 18.查询各科成绩最高和最低的分： 以如下的形式显示：课程ID，最高分，最低分。

```
$ select L.cid as 课程ID,L.score as 最高分,R.score as 最低分
    from sc L,sc R
    where L.cid = R.cid
    and L.score = (select max(IL.score) from 
    sc IL,student as IM where L.cid = IL.cid and IM.sid = IL.sid
    group by IL.cid)
    and R.score = (select min(IR.score) FROM
    sc as IR where R.cid = IR.cid group by IR.cid)
group by L.cid;
```

##### 19.按各科平均成绩从低到高和及格率的百分数从高到低顺序。

    select t.cid as 课程号,
    max(course.course) as 课程名,
    ifnull(avg(score),0) as 平均成绩,
    100 * sum(case when ifnull(score,0) >= 60 then 1 else 0 end)/count(*) as 及格百分数
    from sc t,course
    where t.cid = course.cid
    group by t.cid
    order by 100 * sum(case when ifnull(score,0) >= 60 then 1 else 0 end)/count(*) desc;

19.1.isnull 函数用于如何处理null值




