# seata



## 分布式事务


一个分布式事务代表一个由一批分支事务组成的全局事务，通常分支事务只是本地事务。


## seata中的三个角色

+ Transaction Coordinator(TC):维持着全局事务和分支事务的状态，驱动全局的提交或回滚
+ Transaction Manager(TM):定义一个全局事务的范围:全局事务的开始，提交和回滚
+ Resource Manager(RM):管理分支事务工作所需的资源，向Transaction Coordinator注册分支事务并且报告分支事务的状态，驱动分支事务的提交和回滚


![](https://camo.githubusercontent.com/05a283bea3d9313f03d63bf7e917c016249f14721e49796a3a386e61df12c5a4/68747470733a2f2f63646e2e6e6c61726b2e636f6d2f6c61726b2f302f323031382f706e672f31383836322f313534353031333931353238362d34613930663064662d356664612d343165312d393165302d3261613364333331633033352e706e67)


## seata 对一个分布式事务的基本过程

+ Transaction manager 开启一个新的全局事务.Transaction Coordinator生成一个XID 代表这个全局事务
+ XID通过微服务的调用链传播
+ Resource Manager 将本地事务注册到Transaction Manager中XID响应的全局事务的分支
+ Transaction Manager请求Transaction Coordinator提交或者回滚相应的全局事务的XID
+ Transaction coordinator 驱动对应全局事务下的所有分支事务以完成分支提交或回滚。