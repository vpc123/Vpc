### MySQL调用项目封装接口



作者：流火夏梦                时间：2019-06-22

前情提要：go语言项目封装好了一些mysql调用关系的接口，接口调用关系举例说明如下

```
//2019-06-22-vpc:定义返回类型的结构体（StoreID,UserID,排名顺序）  
rsp = new(RspHubStatisticsDateInfo)  
//2019-06-21-vpc:选择进行操作的数据库表  
//分类：需要进行分类条件进行聚合判断才可以-示例（）  
//1 select sum(duration),user_id,store_id from wd_call_flow where  store_id=0 and start_date<= '2019-06-21' and start_date>= '2019-06-20'   group by user_id,store_id order by  sum(duration) desc ;  
//2 select sum(duration),user_id,store_id from wd_call_flow where start_date<= '2019-06-21' and start_date>= '2019-06-20' group by user_id,store_id order by  sum(duration) desc ;  
fmt.Println(req.StartDate)  
query := d.NewQueryHelper("wd_call_flow")  
query.Select("sum(duration)","user_id","store_id")  
//查询字段区间(pbdal.SignOp_GTE:大于，pbdal.SignOp_LTE：小于)  
query.And("start_date", pbdal.SignOp_GTE, req.StartDate+" 00:00:00.000000")  
query.And("start_date", pbdal.SignOp_LTE, req.EndDate+" 23:59:59.959280")  
//query.And("start_date", pbdal.SignOp_GTE, "2014-06-20"+" 00:00:00.000000")  
//query.And("start_date", pbdal.SignOp_LTE, "2020-05-20"+" 23:59:59.959280")  
query.OrderBy("sum(duration)",pbdal.OrderType_Desc)  
query.GroupBy("user_id","store_id")  
//query.GroupBy("user_id")  
//query.And("user_id", pbdal.SignOp_Equal, req.UserID)  
//查询字段区间(pbdal.SignOp_Equal：等于)  
//query.And("start_date", pbdal.SignOp_Equal, "2018-06-20 14:49:41")  
//按照用户id进行降序排列(pbdal.OrderType_Desc:降序，pbdal.OrderType_Asc)  
//query.OrderBy("user_id", pbdal.OrderType_Asc)  
sql, _ := query.BuildSelect()
```

以上代码片段是在一个go语言项目的框架下进行的项目落地开发工作，针对项目开发语言的设计对查询或者插入语句使用框架封装的方式进行实现和代入。

示例1：

如下查询语句：（通过分组，条件，排序进行的语句查询实现）

> select sum(duration),user_id,store_id from wd_call_flow where start_date<= '2019-06-21' and start_date>= '2019-06-20' group by user_id,store_id order by sum(duration) desc ;

如何分解上述代码到软件开发架构中呢！（这里介绍的是一个go语言开发项目的框架）



第一步：

            定义好go的结构体准备接收函数字段信息，这个需要结合项目的返回值进行字段的设计实现。譬如我们这里定义好的项目字段信息如下：

```
type RspHubStatisticsDateInfo struct {  
   MaxRecord   int64                        //一共有多少的排名个数  
  RecordList HubStatisticsDateInfoList1  
}
type ReqHubStatisticsDateInfo struct {
	StartDate   string                     //查询起始时间日期
	EndDate     string                     //查询结束时间日期

	RecordIndex int64                      //索引字段
	RecordSize  int64                      //记录条数
	ReqHub ReqHubStatisticsDateInfo1       //嵌套结构体类型

}
```

第二步：

            编辑落库查询的mysql语言，譬如我们这里使用的语句如下

```
>select sum(duration),user_id,store_id from wd_call_flow where start_date<= '2019-06-21' and start_date>= '2019-06-20' group by user_id,store_id order by sum(duration) desc ;
```

说明：我们来分析下mysql语句实现的逻辑

1   选择字段sum(duration),user_id,store_id

2   条件start_date<= '2019-06-21' and start_date>= '2019-06-20'

3   根据用户id和门店id进行分类

4   按照sum(duration)的数值进行降序排列



第三步：

            按照拆分的逻辑关系进行封装语句的查询实现

1   新建针对数据库表的连接(这里我们打开数据表wd_call_flow的连接)

> query := d.NewQueryHelper("wd_call_flow")

2    确定查询字段的信息(这里我们确定查询的字段为"sum(duration)","user_id","store_id")

> query.Select("sum(duration)","user_id","store_id")

3    明确查询字段的区间代码(这里我们查询日期区间的数据信息)

> query.And("start_date", pbdal.SignOp_GTE, req.StartDate+" 00:00:00.000000")
> query.And("start_date", pbdal.SignOp_LTE, req.EndDate+" 23:59:59.959280")

4    进行分组查询的语句设置("user_id","store_id")

> query.GroupBy("user_id","store_id")

5     根据查询结果进行升序降序排列操作

> query.OrderBy("sum(duration)",pbdal.OrderType_Desc)

6     通过查询系统封装接口实现数据查询

> sql, _ := query.BuildSelect()



总结：针对大型的业务代码系统设计而言，针对数据库的系统查询接口封装实现是非常必要的。这也是大厂业务员的代码开发过程中针对业务的实现，目前入职的公司采用了consul的为服务架构设计理念，虽然入职第一周已经熟悉了系统项目业务的开发设计流程，但是整个系统的构建和设计还有很多不是很明白的地方。接下来的计划和安排就是熟悉甚至自己写一个微服务框架进行项目的开发设计。
