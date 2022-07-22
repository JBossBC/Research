# GOroutine

## 什么是协程

协程，又称微线程，纤程。对于进程、线程，都是有内核进行调度，有CPU时间片的概念，进行抢占式调度。协程的调用有点类似子程序，但是和子程序相比，协程有挂起的概念，协程可以挂起跳转执行其他协程，合适的机会再跳转回来。

goroutine的使用方式非常的简单，只需要使用go关键字就可以启动一个协程，并且它是处于异步方式运行，你不需要等他运行完以后再执行以后的代码

## 协程调度原理

![](https://img-blog.csdnimg.cn/20200708103112182.png)

    G:一个G代表一个goroutine
    
    M:内核级线程，一个M代表了一个内核线程，等同于系统线程
    
    P:处理器，用来管理和执行goroutine，一个p代表了m所需的上下文环境
  
    Sched:代表调度器，他维护有存储M和G的队列以及调度器的一些状态信息等


    G-M-P三者的关系和特点:

    P的个数取决于设置的GOMAXPROCS，go新版本使用最大内核数，比如你有8核处理器，那么p的数量就是8

    M的数量和P不一定匹配，可以设置很多M，M和P绑定后才可运行，多余的M处于休眠状态。

    P包含一个LRQ本地运行队列，这里面保存着P需要执行的协程G的队列，除了每个P自身保存G的队列外，调度器还拥有一个全局的G队列GRQ，这个队列存储的是所有未分配的协程G

    三者关系:G需要绑定在M上才可以运行，M需要绑定在P才能运行

    简单来说，一个G的执行需要M和P的支持，一个M在于一个P关联之后形成了一个有效的G运行环境(内核线程+上下文环境),每个P都包含一个可运行的G的队列

## 调度实现

![](https://img-blog.csdnimg.cn/20200709133709779.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZpbmdodGluZzMyMQ==,size_16,color_FFFFFF,t_70) 

从上图中看，有2个物理线程M，每一个M都拥有一个处理器P，每一个也都有一个正在运行的goroutine。P的数量可以通过GOMAXPROCS()来设置，它其实也就代表了真正的并发度，即有多少个goroutine可以同时运行。图中灰色的那些goroutine并没有运行，而是处于ready的就绪态，正在等待被调度。P维护着这个队列(称为runqueue),在GO语言里，启动一个goroutine很容器:go function就行了，所以每有一个go语句被执行，runqueue队列就在其末尾加入一个goroutine，在下一个调度点，就从runqueue中取出

### 调度器的有两大思想：

+ 复用线程：协程本身就是运行在一组线程之上，不需要频繁的创建、销毁线程，而是对线程的复用。在调度器中复用线程还有2个体现：1）work stealing，当本线程无可运行的G时，尝试从其他线程绑定的P偷取G，而不是销毁线程。2）hand off，当本线程因为G进行系统调用阻塞时，线程释放绑定的P，把P转移给其他空闲的线程执行。
+ 利用并行：GOMAXPROCS设置P的数量，当GOMAXPROCS大于1时，就最多有GOMAXPROCS个线程处于运行状态，这些线程可能分布在多个CPU核上同时运行，使得并发利用并行。另外，GOMAXPROCS也限制了并发的程度，比如GOMAXPROCS = 核数/2，则最多利用了一半的CPU核进行并行。


### 调度器的两小策略:

+ 抢占:在coroutine中要等待一个协程主动让出CPU才执行下一个协程，在GO中，一个goroutine最多占用CPU 10ms，防止其他goroutine被饿死，这就是goroutine不同于coroutine的一个地方。全局队列，在新的调度器中依然有全局队列，但功能已经被弱化，当M执行work stealing从其他P偷不到G时，它可以从全局G队列获取G
+ 全局G队列:在新的调度器中依然有全局G队列，但功能已经被弱化，当M执行work stealing从其他P偷不到G时，它可以从全局G队列获取G