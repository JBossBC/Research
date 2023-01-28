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

+ hostpath

> hostpath卷存在许多安全风险，最佳做法是尽可能避免使用hostpath。当必须使用hostpath卷时,它的范围应仅限于所需的文件或目录，并以只读方式挂载
> 如果通过AdmissionPolicy限制hostpath对特定目录的访问，则必须要求volumeMounts使用readOnly挂载以使策略生效

hostpath卷能将主机节点文件系统上的文件或目录挂载到pod中。虽然这不是大多数pod需要的，但是它为一些应用程序提供了强大的逃生舱。

hostPath 的一些用法有：

+ 运行一个需要访问 Docker 内部机制的容器；可使用 hostPath 挂载 /var/lib/docker 路径。
+ 在容器中运行 cAdvisor 时，以 hostPath 方式挂载 /sys。
+ 允许 Pod 指定给定的 hostPath 在运行 Pod 之前是否应该存在，是否应该创建以及应该以什么方式存在。

除了必须的path属性之外，你可以选择性地为hostpath卷指定type

支持的type值如下:

|取值|行为|
|--|--|
||空字符串（默认）用于向后兼容，这意味着在安装 hostPath 卷之前不会执行任何检查。|
|DirectoryOrCreate|如果在给定路径上什么都不存在，那么将根据需要创建空目录，权限设置为 0755，具有与 kubelet 相同的组和属主信息。|
|Directory|在给定路径上必须存在的目录。|
|FileOrCreate|如果在给定路径上什么都不存在，那么将在那里根据需要创建空文件，权限设置为 0644，具有与 kubelet 相同的组和所有权。|
|File|在给定路径上必须存在的文件。|
|Socket|在给定路径上必须存在的 UNIX 套接字。|
|CharDevice|在给定路径上必须存在的字符设备。|
|BlockDevice|在给定路径上必须存在的块设备。|


当使用这种类型的卷时要小心，因为：
+  HostPath 卷可能会暴露特权系统凭据（例如 Kubelet）或特权 API（例如容器运行时套接字），可用于容器逃逸或攻击集群的其他部分。
+  具有相同配置（例如基于同一 PodTemplate 创建）的多个 Pod 会由于节点上文件的不同而在不同节点上有不同的行为。
+  下层主机上创建的文件或目录只能由 root 用户写入。 你需要在特权容器中以 root 身份运行进程，或者修改主机上的文件权限以便容器能够写入 hostPath 卷



## ISCSI

iscsi卷能够将ISCSI(基于IP的SCSI)卷挂载到你的Pod中。不像emptyDir那样会在删除pod的同时也会被删除,iscsi卷的内容在删除pod时会被保留，卷只是被卸载。这意味着iSCSI卷可以被预先填充数据，并且这些数据可以在pod之间共享。

> 在使用iSCSI卷之前，你必须拥有自己的iSCSI服务器，并在上面创建卷

iSCSI的一个特点是它可以同时被多个用户以只读方式挂载。这意味着你可以用数据集预先填充卷，然后根据需要在尽可能多的Pod上使用它。不幸的时，iSCSI卷只能由单个使用者以读写模式挂载。不允许同时写入。


## local

local卷所代表的是某个被挂载的本地存储设备，例如磁盘、分区或者目录

local卷只能用作静态创建的持久卷。不支持动态配置。

与hostpath卷相比，local卷能够以持久和可移植的方式使用，而无需手动将pod调度到节点。系统通过查看presistentvolume的**节点亲和性**配置,就能了解卷的节点约束。

然而，local卷仍然取决于底层节点的可用性，并不适合所有应用程序。如果节点变得不健康，那么local卷也将变得不可被pod访问。使用它的pod将不能运行。使用local卷的应用程序必须能够容忍这种可用性的降低，以及因底层磁盘的耐用性特征而带来的潜在的数据丢失风险。

## NFS

nfs卷能将NFS(网络文件系统)挂载到你的pod中。不像emptyDir那样会在删除pod的同时也会删除，nfs卷的内容在删除pod时会被保存，卷只是被卸载。这意味着nfs卷也可以被预先填充数据，并且这些数据可以在pod之间共享。

## persistentVolumeClaim

persistentVolumeClaim卷用来将持久卷挂载到Pod中。持久卷申领使用户在不知道特定云环境细节的情况下"申领"持久存储的一种方法


## secret

secret卷用来给pod传递敏感信息,例如密码。你可以将secret存储在kubernetes API服务器上，然后以文件的形式挂载到Pod中，无需直接与kubernetes耦合。secret卷由tmpfs(基于RAM的文件系统)提供存储，因此它们永远不会被写入持久化的存储器