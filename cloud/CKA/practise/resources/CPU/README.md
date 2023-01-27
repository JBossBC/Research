## cpu 单位
cpu 单位以cpu单位度量。kubernetes中的一个CPU等同于:

+ 1个AWSvCPU
+ 1个GCP核心
+ 1个Azure vCore
+ 裸机上具有超线程能力的英特尔处理器上的一个超线程

小数值是可以使用的。一个请求0.5CPU的容器保证会获得请求1个CPU的容器的CPU的一半。可以使用后缀m表示毫。例如 100m CPU、100 milliCPU 和 0.1 CPU 都相同。 精度不能超过 1m。

CPU 请求只能使用绝对数量，而不是相对数量。0.1 在单核、双核或 48 核计算机上的 CPU 数量值是一样的。

## 设置超过节点能力的CPU请求

CPU请求和限制都与容器相关。Pod对CPU用量的请求等于Pod中所有容器的请求数量之和。同样，pod的CPU资源限制等于pod中所有容器CPU资源限制数之和。

pod调度是基于资源请求值来进行的。仅在某节点具有足够的cpu资源来满足pod cpu请求时，pod就会在对应的节点上运行。

## 如果不指定cpu限制

如果没有为容器指定CPU限制，则会发生以下情况之一:

+ 容器在可以使用的CPU资源上没有上限。因而可以使用所在节点上所有的可用CPU资源。
+ 容器在具有默认CPU限制的名字空间中运行，系统会自动为容器设置默认限制。集群管理员可以使用limitrange指定cpu限制的默认数。

## 未设置cpu请求但设置了cpu限制

如果你为容器指定了 CPU 限制值但未为其设置 CPU 请求，Kubernetes 会自动为其 设置与 CPU 限制相同的 CPU 请求值。类似的，如果容器设置了内存限制值但未设置 内存请求值，Kubernetes 也会为其设置与内存限制值相同的内存请求。

## CPU 请求和限制的初衷

通过配置你的集群中运行的容器的 CPU 请求和限制，你可以有效利用集群上可用的 CPU 资源。 通过将 Pod CPU 请求保持在较低水平，可以使 Pod 更有机会被调度。 通过使 CPU 限制大于 CPU 请求，你可以完成两件事：

+  Pod 可能会有突发性的活动，它可以利用碰巧可用的 CPU 资源。
+  Pod 在突发负载期间可以使用的 CPU 资源数量仍被限制为合理的数量。