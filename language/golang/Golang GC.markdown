# Golang GC



## 三色标记法

三色标记法将对象的颜色分为黑、灰、白三种颜色

+ 黑色:该对象已经被标记过了，且该对象下的属性也全部被标记过了(程序所需要的对象)
+ 灰色:该对象已经被标记过了，但该对象下的属性没有全被标记完(GC需要从此对象中去寻找垃圾)
+ 白色:该对象没有被标记过(垃圾对象)


在GC开始工作时，从GC Roots开始进行遍历访问，访问步骤可以分为下面几步:

1. GC Roots根对象会被标记为灰色
2. 然后从灰色集合中获取对象，将其标记为黑色，将该对象引用到的对象标记为灰色
3. 重复步骤二，直到没有灰色集合可以标记为止
4. 结束后，剩下的白色对象即为GC Roots不可达，可以进行回收。

### 存在的问题

+ 如果标记过程启用STW,则没有问题，但是会影响程序可用性
+ 如果标记过程中没有启用STW,则会出现问题

    + 产生浮动垃圾:例如，当灰色对象B将白色对象A标记为灰色的瞬间,灰色对象B取消了灰色对象A的引用，从此时来讲，灰色对象A是不可达的应该被作为垃圾回收，但是此轮GC会认为A可达。(这本质上是清理垃圾不够干净，但是从程序可用的角度上面来看是可以接受度的)
    + 产生错误垃圾:当一个灰色对象A即将引用白色对象B的时候，灰色对象A断开与白色对象B的连接，转而被黑色对象C所引用，从此刻来讲，白色对象B应该被作为可用对象被标记，因为他被一个黑色对象C所引用，但是从三色标记法的步骤来看，B对象最终会被回收，这会导致程序错误(不能忍受)


从设计的角度来看，如果使用STW去用三色标记法清除垃圾的话，对于性能和程序可用性来讲是不能接受的,如果不使用STW去用三色标记法清除垃圾的话，程序会出错。

## 内存屏障

为了解决上面的产生错误垃圾的问题，同时尽量避免STW的时间,通常引入屏障技术来保障数据的一致性

> 屏障技术在很多方面都有使用，本质上可以理解为一个Hooks函数去"过滤"操作.例如在DTM中,使用屏障去保证程序的幂等特性,本质上也就是通过Hooks去加锁保证此函数只被成功执行一次,在编译过程中，使用内存屏障去保证不出现指令重排....



内存屏障，是一种屏障指令，它能使CPU或编译器对在该屏障指令之前和之后发出的内存操作强制执行排序约束，在内存屏障前执行的操作一定会先于内存屏障后执行的操作。

在上述的描述过程中,在GC时开启STW可以保证垃圾的正常回收，但是因为STW会导致程序在一定时间内不可用，所以这里加入内存屏障的思路我们可以理解为优化性能----为了缩短STW的时间。

GC的过程无疑是两种,标记、清除,缩短STW的时间无疑是将标记或者清除或者标记和清除同时并发地进行,所以绕来绕去，本质上我们需要去解决并发中三色标记法出现的程序错误 

那么为了在标记算法中保证正确性，那么我们需要达成下面任意一个条件：

> 首先我们来分析一下上面三色标记法在并发过程中会出现得问题:本来被灰色对象所引用的白色对象突然间被黑色对象引用之后，按照三色标记法来说,白色对象会被当作垃圾处理。但是对于程序来说，此白色对象被黑色对象引用，说明此白色对象应该是可达对象，应该转换为黑色对象。

此处就引出了两种解决方式

+ 强三色不变性:黑色对象不会指向白色对象，只会指向灰色对象或者黑色对象。(强三色不变性强制性让并发过程中理应出现得问题不出现，是直接的解决方式的问题)
+ 弱三色不变性:即便黑色对象指向白色对象，那么从灰色对象出发，总存在一条可以找到该白色对象的路径(弱三色不变性允许黑色对象指向白色对象，但也有前提存在，也就是该白色对象最终能被变成黑色的前提下(这也是对这句话"总存在一条可以找到该白色对象的路径"最直接的转换))

根据操作类型的不同，我们可以将内存屏障分为Read barrier和write barrier,在Golang中都使用write barrier,原因是因为对于一个不需要对象拷贝的垃圾回收器来说，read barrier代价是非常搞得，因为对于这类垃圾回收器来说是不需要保存读操作的版本指针问题。相对于write

 barrier来说代码量更大，因为堆中的写操作远远大于堆中的读操作。


## dijkstra write barrier


go1.7之前是使用的 dijkstra write barrier

