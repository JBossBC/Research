# cond


sync.cond字面意义就是同步条件变量，它实现的是一种监视器模式

对于cond而言，它实现一个条件变量，是goroutine间等待和通知的点。条件变量与共享的数据隔离，它可以同时阻塞多个goroutine，直到另外的goroutine更改了条件变量，并通知唤醒阻塞着的一个或多个goroutinue。



    type Cond struct {
       noCopy noCopy
       L Locker
       notify  notifyList
       checker copyChecker
    }
    

notifyList记录的是一个基于票号的通知列表

    type notifyList struct {
       wait   uint32 // 用于记录下一个等待者 waiter 的票号
       notify uint32 // 用于记录下一个应该被通知的 waiter 的票号
       lock   uintptr// 内部锁
       head   unsafe.Pointer // 指向等待者 waiter 的队列队头
       tail   unsafe.Pointer // 指向等待者 waiter 的队列队尾
    }


其中，head与tail是指向sudog结构体的指针，sudog是代表的处于等待列表的goroutinue,它本身就是双向链表。值得一提的是，在sudog中有一个字段ticket就是用于给当前goroutine记录票号使用的。

cond实现的核心模式为票务系统，每一个想要来买票的goroutine(调用cond.wait())，我们称之为waiter，票务系统会给每个waiter分配一个取票码，等供票方有该取票码时，就会唤醒waiter。卖票的goroutine有两种，一种时调用cond.signal()的，它会按照票号唤醒一个买票的waiter。第二种时调用Cond.Broadcast()的，它会通知唤醒所有的阻塞waiter。

cond字段中notifylist结构体是一个记录票号的通知列表。这里将notifylist比作排队取票买电影票，当g1通过wait来买票时，发现此时并没有票可买，因此他只能阻塞等待有票之后的通知，此时他手上已经取得勒专属取票码0.