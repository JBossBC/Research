# Thread

## Question

## 简介

在早期的操作系统中，进程就是操作系统用来管理运行程序的最小单位。但是，随着硬件技术的发展，计算机拥有了更多的CPU核心，程序的可并行度提高，进程这一抽象开始显得过于笨重。第一、创建进程的开销较大，需要完成创建独立的地址空间、载入数据和代码段、初始化堆等等。即使使用fork接口创建进程，也需要对父进程的状态进行大量拷贝。第二、由于进程拥有独立的虚拟地址空间，在进程间进行数据共享和同步比较麻烦，一般只能基于共享虚拟内存页或者基于进程间通信。因此，操作系统的设计人员提出在进程内部添加可独立执行的单元，他们共享进程的地址空间，但又各自保存运行时所需的状态，这就是线程。

## 多线程的地址空间布局

多线程的地址空间主要有两个重要特征
+ 分离的内核栈与用户栈:由于每个线程的执行相对独立，进程为每个线程都准备了不同的栈，供他们存放临时数据。在内核中，每个线程也有对应的内核栈。当线程切换到内核中执行时，他的栈指针就会切换到对应的内核栈
+ 共享的其他区域:进程除栈以外的其他区域都由该进程的所有线程共享，包括堆、数据段、代码段等。当同一个进程的多个线程需要动态分配更多内存时，他们的内存分配操作都是在同一个堆上完成的。