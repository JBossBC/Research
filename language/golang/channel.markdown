# Channel 源码分析


chan 结构定义
>
`qcount  uint  //队列中的全部数据`

>`dataqsiz uint //循环队列的大小`

>`buf unsafe.Pointer //points to an array of dataqsiz  elements`

>`elemsize uint16`

>`closed uint32`

>`elemtype *_type //元素类星星`

>`sendx uint //要发送的下标位置`

>`recvx uint //接收数据存放的下标`

>`recvq waitq // receive 等待者列表`

>`sendq waitq // send 等待者列表`

>`lock mutex`


**lock保护hchan中所有的字段，以及在这个通道中被阻塞的sudogs之中的几个字段.持有此锁时不要改变其他g的状态，因为这可能会导致死锁**golang注释中 的解释是 as this can deadlock with stack shrinking stack shrinking

    type waitq struct{
      first *sudog
      last  *sudog
    }


waitq的结构体简单的来说就是一个标准的链表节点

    type sudog struct {
    	// The following fields are protected by the hchan.lock of the
    	// channel this sudog is blocking on. shrinkstack depends on
    	// this for sudogs involved in channel ops.
    
    	g *g
    
    	next *sudog
    	prev *sudog
    	elem unsafe.Pointer // data element (may point to stack)
    
    	// The following fields are never accessed concurrently.
    	// For channels, waitlink is only accessed by g.
    	// For semaphores, all fields (including the ones above)
    	// are only accessed when holding a semaRoot lock.
    
    	acquiretime int64
    	releasetime int64
    	ticket  uint32
    
    	// isSelect indicates g is participating in a select, so
    	// g.selectDone must be CAS'd to win the wake-up race.
    	isSelect bool
    
    	// success indicates whether communication over channel c
    	// succeeded. It is true if the goroutine was awoken because a
    	// value was delivered over channel c, and false if awoken
    	// because c was closed.
    	success bool
    
    	parent   *sudog // semaRoot binary tree
    	waitlink *sudog // g.waiting list or semaRoot
    	waittail *sudog // semaRoot
    	c*hchan // channel
    }

hchan中锁保护等待或发送队列中的sudog对象的属性




![](https://img-blog.csdnimg.cn/84f98b811872422196eebcf2712c6169.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAeWV4aXM=,size_20,color_FFFFFF,t_70,g_se,x_16)

channel 通过通信来实现内存共享。从本质上来看，计算机上线程和协程同步信息其实都是通过共享内存来实现的，因为无论是哪种通信模型，线程或者协程最终都会从内存中获取数据。

## 抽象层级

发送消息和共享内存这两种方式其实是用来传递信息的不同方式，但是它们两者有着不同的抽象层级，发送消息是一种相对高级的抽象，但是不同语言在实现这一机制时也都会使用操作系统提供的锁机制来实现，共享内存这种最原始和最本质的信息传递方式就是使用锁这种并发机制实现的

我们可以这样理解:更为高级和抽象的信息传递方式其实也只是对低抽象级别接口的组合和封装，go语言中的channel就提供了goroutine之间用于传递消息的方式，它在内部实现时就广泛用到了共享内存和锁，通过对两者进行的组合提供了更高级的同步机制
![](https://img-blog.csdnimg.cn/76d03ded102342ce9bcafa4f57e6563b.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAeWV4aXM=,size_20,color_FFFFFF,t_70,g_se,x_16)

### 耦合

使用发送消息的方式代替共享内存也能帮助我们减少多个模块之间的耦合，假设我们使用共享内存的方式在多个Goroutine之间传递消息，**每个Goroutine都可能是资源的生产者和消费者，他们需要在读取或者写入数据时先获取保护该资源的互斥锁
![](https://img-blog.csdnimg.cn/f263ddfa92a34c4b9bf610f0d2595456.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAeWV4aXM=,size_20,color_FFFFFF,t_70,g_se,x_16)

然而偶们使用发送消息的方式却可以将多个线程或者协程解耦，以前需要依赖同一个片内存的多个线程，现在可以成为消息的生产者和消费者，多个线程也不需要自己手动处理资源的获取和释放，其中Go语言实现的CSP机制通过引入Channel来解耦Goroutine

Go语言实现的CSP模型其实于消息队列非常相似，我们引入Channel这一中间层让资源的生产者和消费者更加清晰，当我们需要增加新的生产者或者消费者的时候也只需要直接增加Channel的发送方和接收方


![](https://img-blog.csdnimg.cn/1fba29b84b1f47c7849e17c72b47edac.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAeWV4aXM=,size_20,color_FFFFFF,t_70,g_se,x_16)