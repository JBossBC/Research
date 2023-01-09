# zero-copy


零拷贝是一种高效的数据传输机制，在追求低延迟的传输场景中十分常用。


## 传统的数据传输方式


    JVM向OS发出read()系统调用，触发上下文切换，从用户态切换到内核态。
    从外部存储（如硬盘）读取文件内容，通过直接内存访问（DMA）存入内核地址空间的缓冲区。
    将数据从内核缓冲区拷贝到用户空间缓冲区，read()系统调用返回，并从内核态切换回用户态。
    JVM向OS发出write()系统调用，触发上下文切换，从用户态切换到内核态。
    将数据从用户缓冲区拷贝到内核中与目的地Socket关联的缓冲区。
    数据最终经由Socket通过DMA传送到硬件（如网卡）缓冲区，write()系统调用返回，并从内核态切换回用户态。


**内核缓冲区的意义(当一个用户进程要从磁盘读取数据时，内核一般不直接读磁盘，而是将内核缓冲区中的数据复制到进程缓冲区中。但若是内核缓冲区中没有数据，内核会把对数据块的请求，加入到请求队列，然后把进程挂起，为其它进程提供服务。等到数据已经读取到内核缓冲区时，把内核缓冲区中的数据读取到用户进程中，才会通知进程)


## 零拷贝的数据传输方法


“基础的”零拷贝机制

通过分析可以看出，第2、3次拷贝（也就是从内核空间到用户空间的来回复制）是没有意义的，数据应该可以直接从内核缓冲区直接送入Socket缓冲区。零拷贝机制就实现了这一点。不过零拷贝需要由操作系统直接支持，不同OS有不同的实现方法。大多数Unix-like系统都是提供了一个名为sendfile()的系统调用，在其man page中，就有这样的描述：

>sendfile() copies data between one file descriptor and another.


>Because this copying is done within the kernel, sendfile() is more efficient than the combination of read(2) and write(2), which would require transferring data to and from user space.


可见确实是消除了从内核空间到用户空间的来回复制，因此“zero-copy”这个词实际上是站在内核的角度来说的，并不是完全不会发生任何拷贝。

在Java NIO包中提供了零拷贝机制对应的API，即FileChannel.transferTo()方法。不过FileChannel类是抽象类，transferTo()也是一个抽象方法，因此还要依赖于具体实现。FileChannel的实现类并不在JDK本身，而位于sun.nio.ch.FileChannelImpl类中，零拷贝的具体实现自然也都是native方法.


