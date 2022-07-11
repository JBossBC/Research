# Mysql sql



## 设置事务隔离级别

set session transaction isolation level 事务级别


## 开启事务 

start transaction

## 查看当前mysql事务隔离级别


## show 与select的区别

select针对的是表中具体项的查询

show 针对的是数据库，表，以及全局变量等全局因素的查询

## 查看连接的客户端详情

show processlist;

## 查询连接的客户端数量

show status like "Threads%"；

## limit and offset

limit y 表示:读取前y条数据
limit x，y 表示:跳过x条数据，读取y条数据
limit y offset x 表示:跳过x条数据，读取y条数据(读取x+1,x+1+y)


