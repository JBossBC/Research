# Address

## Question
为什么需要页内偏移

## MMU

 实际上完成虚拟地址转换为物理地址转换的硬件是 CPU 中含有一个被称为 内存管理单元（Memory Management Unit, MMU） 的硬件。

## 物理地址

这里说的物理地址是内存中的内存单元实际地址，不是外部总线连接的其他电子元件的地址

物理地址就是内存中每个内存单元的编号，这个编号是顺序排好的，物理地址的大小决定了内存中有多少个内存单元，物理地址的大小由地址总线的位宽决定。

## 虚拟地址

虚拟地址是CPU保护模式下的一个概念，保护模式是80286系列和之后的X86兼容CPU操作模式，在CPU引导操作系统内核后，操作系统内核会进入一种CPU保护模式，也叫虚拟内存管理，在这之后的程序在运行时都处于虚拟内存当中，虚拟内存里的所有地址都是不直接的，所以你有时候可以看到一个虚拟地址对应不同的物理地址，比如A进程里的call函数入口虚拟地址是0x001，而b也是，但是他俩对应的物理地址却是不同的，操作系统利用这种内存管理方法

1. 防止程序对物理地址写数据造成一些不可必要的问题，比如知道了A进程的物理地址，那么向这个地址写入数据就会造成A进程出现问题，在虚拟进程中运行程序永远不知道自己处于内存中哪一段的物理地址上!现在的操作系统运行在保护模式下即便知道其他进程的物理地址也不允许向其写入!但是可以通过操作系统留下的后门函数获取该进程上的虚拟地址空间所有控制权限并写入指定数据
2. 虚拟内存管理采用一种拆东墙补西墙的形式，所以虚拟内存的内存要比物理内存大许多。在进入虚拟模式之前CPU以及Bootloader(Bootloader是在操作系统内核运行之前运行。可以初始化硬件设备、建立内存空间映射图，从而将系统的软硬件环境带到一个合适状态，以便为最终调用操作系统内核准备好正确的环境),操作系统内核均运行在实模式下，直接对物理地址进行操作。
3. 当不同的进程使用同样的代码时，比如库文件的代码，物理内存可以储存一份这样的代码，不同的进程只需要把自己虚拟内存映射过去就行了，节省内存，
4. 程序可以使用一系列相邻的虚拟地址来访问物理内存中不相邻的大内存缓冲区。

虚拟内存中也有分页管理，这种管理方式是为了确保内存中不会出现内存碎片，当操作系统内核初始化完毕内存中的分页表后CPU的分页标志位会被设置，这个分页标志位是给MMU看的!

### 分页管理

内存分页其实就是我们所说的4G空间，内存的所有内存被操作系统内核以4G为每页划分开，当我们程序运行时会加载到内存中的4G空间里，其实说是有4G其实并没有真正在的4G空间，4G空间中有一小部分被映射到了物理内存中，或者被映射到了硬盘的文件上，或者没有被映射，还有一小部分在内存当中就会被划分栈、堆，其中有大片大片的内存是没有被映射的，同样物理内存也是被分页了用来与虚拟内存产生映射关系!

其实真正情况下只有3G用户空间，假如你的内存是4G的那么其中有1G是给操作系统内核使用的，所谓的4G空间只是操作系统基于虚拟内存这种拆东墙补西墙的形式给你一种感觉每个进程都有4G的可用空间。
这里来说一下拆东墙补西墙，当我们程序被加载进4G空间时其实根本用不了所谓的4G空间，其中有大片内存被闲置，那么这个时候，其他程序被加载进来的时候发现内存不够了，就把其他程序里的4G空间里闲置部分拿出来给这个进程用，换之这个进程内存不够时就会把其他进程里闲置的空间拿过来给该进程使用。
当我们要对物理地址做操作时比如if语句要根据CPU的状态标志寄存器来做不同的跳转，那么这个时候就要对CPU状态寄存器做操作了，必须要知道它的物理地址，内存中有一个电子元件叫MMU负责从操作系统已经初始化好的内存映射表里查询与虚拟地址对应的物理地址并转换，比如mov 0x4h8这个是虚拟地址，当我们要对这个虚拟地址里写数据时那么MMU会先判断CPU的分页状态寄存器里的标志状态是否被设定，如果被设定那么MMU就会捕获这个虚拟地址并在操作系统内核初始化好的内存映射表里查询与之对应的物理地址，并将其转换为真正的实际物理地址，然后在对这个实际的物理地址给CPU，再由CPU去执行对应的命令，相反CPU往内存里读数据时比如A进程要读取内存中某个虚拟地址的数据，A进程里的指令给的是虚拟地址，MMU首先会检查CPU的分页状态寄存器标志位是否被设置，如果被设置MMU会捕获这个虚拟地址并将其转换成相应的物理地址然后提交给CPU，再由CPU到内存中去取数据。

## 虚拟地址和物理地址的关系

分页就是把整个虚拟和物理内存切成一段段固定大小的空间,连续且尺寸固定的内存空间叫页，linux下每一页大小4kb。
虚拟地址和物理内存之间通过页表来映射；虚拟地址分为:页号和页内偏移
![](https://img-blog.csdnimg.cn/115aa212b37e47b7b293ffef6b7fe6a2.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA54mb54mbY29kaW5n,size_20,color_FFFFFF,t_70,g_se,x_16#pic_center)

页号作为页表的索引，页表包含物理页每页所在的物理内存的基地址，这个基地址和页内偏移的组合就形成了物理内存地址。

简单分页的缺点:空间上的缺陷:32位，单进程一个页4kb，虚拟内存4GB，一个页表项4字节，结果是4GB空间映射需要4mb的存储页表。

**解决方法:采用多级页表或段页式存储**
![](https://img-blog.csdnimg.cn/8d116aefda3c4bda802f3fc7d377fe9a.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA54mb54mbY29kaW5n,size_20,color_FFFFFF,t_70,g_se,x_16#pic_center)
地址结构就由段号、段内页号和页内位移三部分组成。

+ 第一次访问段表，得到页表起始地址
+ 第二次访问页表，得到物理页号
+ 第三次将物理页号与页内位移组合，得到物理地址。
+ 
内存碎片:内部碎片和外部碎片

### 内部碎片:内存中已经被分配出去的内存，但是进程不使用这一块内存，进程却一直占用着导致操作系统无法回收给其他进程使用，为了有效的防止这种空间上的浪费现象所以使用了内存分页管理机制。

操作系统在内存中会维护一个内存信息分页表用于标示某段到某段为个页面。

![](https://img-blog.csdn.net/20171111212410986)

比如在内存中分配了这样的一个地址，当ID为1的内存不用了，但是该进程一直占用着这段0-2的内存，如果此时分配一个长度为2字节的内存空间，ID1的内存刚好足够分配但是这段地址一直被该进程所占用着，所以无法分配。

后面Intel工程师为了防止这种情况的出现用页为单位的内存管理方式，有效的防止了这种内存碎片的情况发生。

这样的话页ID为1的地方为单独的一个页，当进程不使用时操作系统可以将该页内存分配给其他进程所使用。但是这种分页内存往往也会出现一些内存碎片，比如分页分到最后剩下一部分不足以分配给其他进程所使用的内存页面也称为内部碎片，只不过这种浪费比原本的浪费要节约许多。






