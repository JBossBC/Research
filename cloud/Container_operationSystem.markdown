# container 


## AUFS

AUFS 全称 Union File System 又叫做Another UnionFS。所谓的Union FS就是吧不同物理位置的目录合并mount到同一个目录，而docker就是通过这个特性实现了镜像层的重叠，容器层的存储和显示层的展示。

我们再来看一下这张图，我们通过docker history + 镜像ID来查看镜像的历史。
当镜像启动的时候，一个新的可写层会加载到镜像的顶部，这一层我们一般称为容器层，之下是镜像层。
容器层可以读写，容器所有发生文件变更都发生在这一层，而镜像层是read-only只读。

根据aUFS的定义，容器的文件系统就是由下面的15个只读镜像层和1个可写的容器层通过aUFS mount出来的。

到这里，就能和前面的aUFS联系起来了，X就是容器层，可修改，可记录，Y就是镜像层，不可更改，只读，而Z就是我们进入联合起来的视图层。

## LXC

LXC又名Linux container,是一种虚拟化解决方案，这种是内核级的虚拟化(主流的解决方案Xen,KVM,LXC)

通过namespace进行资源的隔离，Guset1下的进程与Guset2下的进程是独立的，可以看作运行在两台物理机上一样。Container管理工具就是对Guest进行管理的(创建、销毁)。


## namespace

linux namespaces机制提供一种资源隔离方案。PID,IPC,Network等系统资源不再是全局性的(在linux2.6内核以前是全局的),而是属于特定的namespace。每个namespace里面的资源对其他namespace都是透明的。namespace是container中使用到的重要技术之一，是对系统资源的操作上的隔离。

## Linux3.8以上的内核实现的namespace

+ Mount namespace
+ UTS namespace
+ IPC namespace
+ Net namespace
+ Pid namespace
+ User namespace

### Mount

mount namespace是对挂载的文件系统布局进行隔离

### IPC

处于同一namespace下的进程才可以进行进程间的通信

### NET

NET namespace实现网络协议栈上的隔离，在自己的namespace中对网络的设置只能在本namespace中生效。

### PID

我们通过fork来创建进程时可以为每个进程指定命名空间。linux下的进程关系时一棵树，所以有了父命名空间和子命名空间之分。

### USER

User namespace中使用到了map转换，由于container并不是真正的虚拟化,所以在Guest-OS中创建的root用户会被映射到Host-OS中的普通用户中去。
