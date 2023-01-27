# memory
一个kubernetes集群可以被划分为多个命名空间。如果你在具有默认内存限制的名称空间内尝试创建一个pod，并且这个pod中的容器没有声明自己的内存资源限制，那么控制面会为该容器设定默认的内存限制


> limitrange不会检查它应用的默认值的一致性。这意味着limitrange设置的limit的默认值可能小于客户端提交给api服务器的声明中为容器指定的request值。发生这种情况，最终会导致pod无法调度。

## 设置默认内存限制和请求的动机

如果你的命名空间设置了内存资源配额，那么为内存限制设置一个默认值会很有帮助。以下是内存资源配额对命名空间施加的三条限制:

+ 命名空间中运行的每个pod中的容器都必须有内存限制。(如果为pod中的每个容器声明了内存限制，kubernetes可以通过将其容器的内存限制相加推断出pod级别的内存限制)
+ 内存限制用来在pod被调度到的节点上执行资源预留。预留给命名空间中所有pod使用的内存总量不能超过规定的限制
+ 命名空间中所有pod实际使用的内存总量也不能超过规定的限制。

当你添加limitrange时

如果该命名空间内的任何pod的容器未指定内存限制，控制面将默认内存限制应用于该容器，这样pod可以受到内存resourcequota限制的命名空间中运行。

# cpu

一个 Kubernetes 集群可被划分为多个命名空间。 如果你在具有默认 CPU限制 的命名空间内创建一个 Pod，并且这个 Pod 中任何容器都没有声明自己的 CPU 限制， 那么控制面会为容器设定默认的 CPU 限制。

## 默认CPU限制和请求的动机

如果你的命名空间设置了 CPU 资源配额， 为 CPU 限制设置一个默认值会很有帮助。 以下是 CPU 资源配额对命名空间的施加的两条限制:

+ 命名空间中运行的每个 Pod 中的容器都必须有 CPU 限制。
+  CPU 限制用来在 Pod 被调度到的节点上执行资源预留

预留给命名空间中所有 Pod 使用的 CPU 总量不能超过规定的限制。

当你添加 LimitRange 时：

如果该命名空间中的任何 Pod 的容器未指定 CPU 限制， 控制面将默认 CPU 限制应用于该容器， 这样 Pod 可以在受到 CPU ResourceQuota 限制的命名空间中运行。

# 配置命名空间的最小和最大内存约束

# 配置命名空间的最小和最大cpu约束

# 为命名空间配置内存和cpu配额

resourcequota在命名空间中设置了如下要求:
+ 在该命名空间中每个pod的所有容器都必须要有内存请求和限制，以及cpu请求和限制
+ 在该命名空间中所有pod的内存请求总和不能超过规定大小
+ 在该命名空间中所有pod的内存限制总和不能超过规定大小
+ 在该命名空间中所有pod的cpu请求总和不能超过规定大小
+ 在该命名空间中所有pod的cpu限制总和不能超过规定大小

> 这里面的request 和limit 看起来比较矛盾，也许可能会问请求的总和一旦满足那么这个限制也必然满足?
> 这里是对request和limit的理解有误，对于request来说比较好理解，我能够规定一个容器能够申请的大小,以便于方便容器的管理。对于这个limit的理解，我理解为弹性的标准，request可以理解为这个容器被创建的时候最高的标准，但是一个容器的启动肯定是为了外界服务的，在为外界服务的过程中，因为计算和IO肯定会消耗计算机资源，那么肯定会导致cpu和memory的增长，我们如何保证这个增长的上限，这就是limit的作用。request保证了容器的初始值的上限，而容器的初始值与limit的差值的绝对值，即是这个容器能够在处理过程中消耗的计算机资源的规范，这就是弹性的过程，我允许容器的运行时消耗的计算机资源处在一定范围内，但不会无限的容忍他扩展，为了保证自身物理机或者说是虚拟机的安全，以及其他服务的可用，同时也方便了物理资源的更有效利用。

**一旦为某个namespace申请了quota，那么所有创建的pod必须被指定request和limit**