## Qos 类
kubernetes创建pod时就给它指定了下列一种Qos类

+ Guaranteed
+ Burstable
+ BestEffort


### 创建一个Qos类为Guaranteed的pod

对于Qos类为Guaranteed的pod:

+ pod中的每个容器都必须指定内存限制和内存请求
+ 对于pod中的每个容器，内存限制必须等于内存请求
+ pod中的每个容器都必须指定cpu限制和cpu请求
+ 对于pod中的每个容器，cpu限制必须等于cpu请求


## 创建一个Qos类为Burstable的pod

如果满足以下条件，将会指定pod的qos类为burstable

+ pod不符合guaranteed qos类的标准
+ pod中至少一个容器具有内存或cpu的请求和限制。


## 创建一个Qos类为BestEffort的Pod

对于Qos类为BestEffort的pod，pod中的容器必须没有设置内存和cpu限制或请求。