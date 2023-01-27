# volume

container中的文件在磁盘上存放是临时的，这给container中运行的较为重要的应用程序带来一点问题。问题之一是当容器崩溃时文件丢失。kubelet会重新启动容器，但容器会以干净的状态重启。第二个问题会在同一pod中运行多个容器并共享文件时出现。

## 背景

docker卷是磁盘上或者另外一个容器内的一个目录。docker提供卷驱动程序，但是其功能非常有限。

kubernetes支持很多类型的库。pod可以同时使用任意数量的卷类型。临时卷类型的生命周期与pod相同，但持久卷可以比pod的存活期长。当pod不再存在时，kubernetes也会销毁临时卷;不过kubernetes不会销毁持久卷对于给定pod中任何类型的卷，在容器重启期间数据都不会丢失。

卷的核心是一个目录，其中可能存在数据，pod中的容器可以访问该目录中的数据。所采用的特定的卷类型将决定该目录如何形成的、使用何种介质保存数据以及目录中存放的内容。

使用卷时，在.spec.volumes字段中设置为pod提供的卷，并在.spec.containers[*].volumeMounts字段中声明卷在容器中的挂载位置。容器中的进程看到的文件系统视图是由它们的容器镜像的初始内容以及挂载在容器中的卷所组成的。其中根文件系统同容器镜像的内容相温和。任何在该文件系统下的写入操作，如果被允许的话，都会影响接下来容器中进程访问文件系统所看到的内容。

卷挂载在镜像中的指定路径下。 Pod 配置中的每个容器必须独立指定各个卷的挂载位置。

卷不能挂载到其他卷之上（不过存在一种使用 subPath 的相关机制），也不能与其他卷有硬链接。

## 卷类型

kubernetes支持下列类型的卷:

+ cephfs

cephfs卷允许你将现存的cephfs卷挂载到pod中。不像emptyDir那样会在pod被删除的同时也会被删除,cephfs卷的内容在pod被删除时会被保留，只是卷被卸载了。这意味着cephfs卷可以被预先填充数据，且这些数据可以在pod之间共享。同一cephfs卷可同时被多个写着挂载。

> 在使用 Ceph 卷之前，你的 Ceph 服务器必须已经运行并将要使用的 share 导出（exported）。

+ configMap

configMap卷提供了向pod注入配置数据的办法。configmap对象中存储的数据可以被configmap类型的卷引用，然后被pod中运行的容器化应用使用

引用configmap对象时，你可以在卷中通过它的名称来引用。你可以自定义configmap中特定条目所要使用的路径。

+ downwardAPI

downwardAPI卷用于为应用提供downward API数据。在这类卷中,所公开的数据以纯文本格式的只读文件形式存在。

> 容器以subPath卷挂载方式使用downward API时，在字段值更改时将不能接收到它的更新

+ emptyDir

当pod分配到某个节点上时，emptyDir卷会被创建，并且pod在该节点上运行期间，卷一直存在，就像其名称表示的那样，卷最初是空的。尽管pod中的容器挂载emptyDir卷的路径可能相同也可能不相同，这些容器都可以读写emptyDir卷中相同的文件。当pod因为某些原因被从节点删除时，emptyDir也会被永久删除