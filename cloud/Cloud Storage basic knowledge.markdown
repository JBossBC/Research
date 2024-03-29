# Cloud Storage

## Question

块存储和文件存储的区别?


## 三种存储形态

+ 块存储

块存储简单的理解就是一块一块的硬盘，直接挂载在主机上，在主机上我们能够看到的就是一块块的硬盘以及硬盘分区。从存储架构的角度而言，块存储又分为DAS存储(Direct-Attached Storage,直连式存储)和SAN存储(Storage Area Network,存储区域网络)

+ 文件存储

文件存储，指的是在文件系统上的存储，也就是主机操作系统中的文件系统。在文件系统中有分区，有文件夹，子文件夹，形成一个自上而下的文件结构；文件系统下的文件，用户可以通过操作系统中的应用程序进行打开、修改等操作。从架构上来说，文件存储提供一种NAS(Network Attached Storage,网络附属存储)架构，使得主机的文件系统不仅限于本地的文件系统，还可以连接基于局域网的共享文件系统。

+ 对象存储

对象存储是面向对象/文件的、海量的互联网存储，它也可以直接被称为"云存储"。对象尽管是文件，它是已被封装的文件(编程中的对象就有封装性的特点),也就是说，在对象存储系统中，你不能直接打开/修改文件，但可以像ftp一样上传文件，下载文件等。另外**对象存储没有像文件系统那样有一个很多层级的文件结构，而是只有一个"桶"的概念(也就是存储空间)，"桶"里面全部都是对象，是一种非常扁平化的存储方式。其最大的特点就是它的对象名称就是一个域名地址，一旦对象被设置为"公开",所有网名都可以访问到它；