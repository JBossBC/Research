# File Descriptor

## Question

文件系统和文件描述符

文件描述符在形式上是一个非负整数(centos /dev/fd下面)。实际上，它是一个索引值,指向内核为每一个进程所维护的该进程打开文件的记录表。当程序打开一个现有文件或者创建一个新文件时，内核向进程返回一个文件描述符。

+ 概念

linux系统中，一切都是文件，当进程打开现有文件或创建新文件时，内核向进程返回一个文件描述符,文件描述符就是内核为了高效管理已被打开的文件所创建的索引,用来指向被打开的文件，执行所有I/O操作的系统调用都会通过文件描述符。

+ 文件描述符、文件、进程间的关系

      + 每个文件描述符会与一个打开的文件相对应
      + 不同的文件描述符也可能指向同一个文件
      + 相同的文件可以被不同的进程打开，也可以在同一个进程被多次打开

+系统为维护文件描述符,建立了三个表
       + 进程级的文件描述符表
       + 系统级的文件描述符表
       + 文件系统的i-node表

   ![](https://imgconvert.csdnimg.cn/aHR0cHM6Ly9zZWdtZW50ZmF1bHQuY29tL2ltZy9iVk9YM2w_dz02MTkmaD0zMTI)