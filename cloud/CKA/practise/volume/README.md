## emptyDir

当 Pod 分派到某个节点上时，emptyDir 卷会被创建，并且在 Pod 在该节点上运行期间，卷一直存在。 就像其名称表示的那样，卷最初是空的。 尽管 Pod 中的容器挂载 emptyDir 卷的路径可能相同也可能不同，这些容器都可以读写 emptyDir 卷中相同的文件。 当 Pod 因为某些原因被从节点上删除时，emptyDir 卷中的数据也会被永久删除。

> 容器崩溃并不会导致pod被从节点移除，因此容器崩溃期间emptyDir卷中的数据是安全的
emptyDir的一些用途:

+ 缓存空间，例如基于磁盘的归并排序
+ 为耗时较长的计算任务提供检查点,以便任务能方便地从崩溃前状态恢复执行
+ 在web服务器容器服务数据时,保存内容管理器容器获取内容

emptyDir.medium字段用来控制emptyDir卷的存储位置。默认情况下，emptyDir卷存储在该节点所使用的介质上;此处介质可以是磁盘、SSD或网络存储，这取决于你的环境。你可以将emptyDir.medium字段设置为memory,以告诉kubernetes为你挂载tmpfs(基于RAM的文件系统)。虽然tmpfs速度非常快，但是他与磁盘不同:tmpfs在节点重启时会被清除,并且你所写入的所有文件都会计入容器的内存消耗，受容器内存限制约束

你可以通过为默认介质指定大小限制，来限制emptyDir卷的存储容量。此存储时从节点临时存储中分配的。如果来自其他来源的数据占满了存储，emptyDir可能会在达到此限制之前发生存储容量不足的问题。

## 通过文件将pod信息呈现给容器(downward volume)
在kubernetes里面，有两种方式可以将pod和容器字段呈现给pod中运行的容器。downwardAPI卷可以呈现pod和容器字段

在kubernetes中，有两种方式可以将pod和容器字段呈现给运行中的容器:

+ 环境变量
+ downward volume

这两种呈现pod和容器字段的方式都称为downward API