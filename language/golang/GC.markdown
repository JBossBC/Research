# Golang GC

垃圾回收器一直是被诟病最多，也是整个运行时中改进最努力的部分。所有变化都是为了缩短STW时间，提高程序实时性。


## 什么是垃圾回收机制

垃圾回收可以有效的防止内存泄露，有效的使用空闲的内存。

内存泄露是指该内存空间使用完毕之后未回收，内存泄漏过多导致内存溢出，导致应用程序所占内存超出系统限制，最终被系统杀掉。

## 什么是垃圾

这里的垃圾是指无用的对象或其他数据等已经不被需要，但却无法被GC所释放。

GC所使用的两种判断垃圾的算法

-  引用计数法

     引用计数是垃圾收集器中的早期策略。在这种方法中，堆中每个对象实例都有一个引用计数。当一个对象被创建时，且将该对象实例分配给一个变量，该变量计数设置为1.当任何其它变量被赋值为这个对象的引用时，计数加1，但当一个对象实例的某个引用超过了声明周期或者被设置为一个新值时，对象实例到的引用计数器减1.任何引用计数器为0的对象实例可以被当作垃圾收集。当一个对象实例被垃圾收集时，它引用的任何对象实例的引用计数器减1.

    引用计数法无法检测出循环引用，两个对象互相引用的情况下无法被检测出

- 可达性算法(根搜索算法)
    
    从根引用节点开始检索这个节点引用的节点，当所有的被引用节点都被找到后，剩余的节点则被认为是没有被引用的节点，需要被回收


GC所使用的几种回收算法

1. 标记清除算法 从根集合扫描，将存活的对象进行标记，标记完毕后再扫描一遍，对未进行标记的对象进行回收。标识清除算法不需要进行对象的移动，仅对不存活的对象进行处理，容易造成内存碎片。适用于存活率高的情况。
2. 复制算法 将可用内存空间分为两个部分，开始只使用其中的一半，在这一半用完后再将存活的对象复制到另外一块上。适用于存活率低的情况，避免了内存碎片的情况。
3. 标记整理算法 同标记清除算法一样进行标记，但在清除时将后面的对象在内存空间上向左移动，相比标记清楚来说成本更高，但没内存碎片的问题
4. 分代收集算法 对对象进行分代，不同代的对象采取不同的回收算法，提高效率。年轻代，新生成的对象都归为年轻代，年轻代在存放满和其他情况会触发minor GC。 老年代，大对象直接进入老年代，年轻代经历几次垃圾回收后仍然存活的对象也进入老年代。老年代存放满后触发Full GC



Golang的基本特征是"非分代、非紧缩、写屏障、并发标记处理"

与之前版本在STW状态上完成标记不同，并发标记和用户代码同时执行让一切都处于不稳定状态。用户代码随时可能修改已经被扫描过的区域，在标记过程中还会不断分配新对象，这让垃圾回收变得很麻烦。

究竟什么时候启动垃圾回收？过早会严重浪费CPU资源，影响用户代码执行性能。而太晚，会导致堆内存恶意膨胀。如何正确平衡这些问题就是个巨大的挑战。

所有问题的核心:抑制堆增长，充分利用CPU资源。


### 三色标记和写屏障

**golang gc中的root node**
+ 全局变量，生命周期贯穿程序的变量
+ 执行栈，groutine执行栈的变量及指向堆内存的指针
+ 寄存器，参与计算的指针或变量 

**这是让标记和用户代码并发的基本保障**,基本原理:

+ 起初所有对象都是白色。
+ 扫描找出所有可达对象，标记为灰色，放到待处理队列。
+ 从队列中提取灰色对象，将其引用对象标记为灰色放入队列，自身标记为黑色。
+ 写屏障监视对象内存修改，重新标色或放回队列。

白色:没有找到对象被引用位置

灰色:已找到对象被引用位置，还未分析对象引用的其他对象

黑色:已找到对象被引用位置，已经分析对象引用的其他对象

当完全全部扫描和标记工作后，剩余不是白色就是黑色，分别代表要待回收和活跃对象，清理操作只需将白色对象内存回收即可。

### 后台并发标记

为了避免GC长时间STW影响正常任务执行，golang将主要标记工作放在单独的g上执行。这些g可以和其他任务g一起被调度，并发执行，而不需要STW

由于执行并发标记的g有多个，这些g需要并发的从全局GC队列中获取对象。为了减少这些g的并发冲突，golang在每个g维护了两段buffer，用于批量缓存扫描对象，这样每个g可以从gc队列中批量获取对象，减少冲突。？？？

### 后台并发清除

被清除的对象已经不可能再被访问，因此没有必要在STW中处理。golang使用一个专门的g来清除回收内存，这个g可以和其他g并发调度执行。

### 写屏障

写屏障是后台并发标记的重要补充。在非STW状态下进行标记，必然会出现已经扫描过的对象又引用了新对象的情况，这时新对象没有被标记。写屏障解决了这个问题。写屏障打开后，对指针的修改操作会将指针的新旧值全部加入gc工作队列。并在并发标记完成后再进入STW最后检查一次GC工作队列，保证所有对象在最新状态扫描标记过。

### 辅助GC

辅助GC是mallocgc中执行一部分gc标记的机制。每分配一定数量的内存，就会在mallocgc中做出一些gc标记工作。如果mallocgc分配了过多内存却没有完成足够多的标记工作，就会被挂起，直到其他gc工作线程完成了足够多的工作或gc结束时才会被唤醒。这个机制的目的是防止GC过程中mallocgc执行过快分配过多新内存，导致gc持续时间过长或无法完成。

### GC触发条件

gc的初始函数是runtime.gcstart，但调用gcstart不一定会开启一轮gc。gostart有一个gctrigger类型的参数，需要验证trigger.test()为true才会真正触发GC。而trigger有三种类型，分别对应3种触发GC的条件:

1. gcTriggerHeap。这种类型的触发条件是memstats.heap_live>=memstats.gc_trigger，即当前使用的内存量超过GC触发阈值。而触发阈值memstats.gc_trigger，这个gc_trigger并不是根据gcpercent计算的下一次GC内存大小next_gc,而是基于一个持续估算更新的triggerRatio计算的值。triggerRatio的更新算法很复杂，主要在gcController.endCycle()中，但triggerRatio的值一定在gcpercent/100的0.6~0.95倍之间，因此gc_trigger一定小于next_gc
2. gcTriggerTime。这类的触发条件是当前时间与上次GC事件memstats.last_gc_nanotime间的时间间隔已经超过forcegcperiod(120s)
3. gcTriggerCycle。这种类型的trigger中会指定一个GC执行轮数，触发条件是指定的轮数n大于已执行过的GC轮数work.cycles。即如果已执行的GC次数没有达到指定次数，则触发一次GC

而与上面的三类trigger对应，调用gcstart开始一次GC的位置有3处:

1. 内存分配函数runtime.mallocgc中。mallocgc是goruntime中唯一的堆分配函数，在函数的末尾，会根据此次调用是否真正申请了新的堆内存来决定是否需要触发GC。这里使用的trigger类型是gcTriggerHeap。如前文所述，这里触发GC的条件是堆内存的用量超过一个阈值，这个阈值是由上一次gc结束后堆内存的大小乘以一个系数计算而来的。这个系数事实上是由更复杂的算法得出的，会略小于1+gcpercent/100。gcpercent的初始值来自于GOGC环境变量，默认为100.也就是说，如果堆内存用量达到上一次GC结束时用量的两倍，就会触发下一次GC.gcpercent可以debug.setGCpercent接口动态调整。通过这种方式，go runtime达到了两个目的 ①、内存使用量控制在一个比较稳定的范围 ②、gc触发不会太过频繁
2. forcegchelper定时触发。forcegchelper是在runtime/proc.go的init函数中创建的一个goroutine(G)，专门用于定时触发GC。这个G由sysmon定时触发执行，触发间隔为120秒，这里使用的trigger类型是gcTriggerTime。
3. runtime.GC()接口调用触发。在程序中调用这个接口会强制触发一次GC，而这个调用需要等到下一次GC完成后才会返回，而不是它出发的这次GC完成。原因是设计者希望显示调用GC的效果可以在函数返回后立刻通过heap profile看到，而下一次GC完成后这次GC的结果才会体现在heap profile中。这里使用的trigger类型是gcTriggerCycle。

### GC过程

为了避免长时间停止任务运行，golang将gc过程分为多个步骤，其中一些步骤是可以和任务goroutine并发执行的。目的是尽可能减少需要STW的时间。通过debug.gcstoptheworld可以控制gc的方式。如果debug.gcstoptheworld=1则标记阶段会完全停止任务运行，如果debug.gcstoptheworld=2则标记（mark）和清除（sweep）阶段都会完全停止任务运行。默认情况下，debug.gcstoptheworld=0。


### gcStart:初始化与启动

gcStart负责GC的舒适化，并启动后台标记工作，返回时会将GC阶段修改为_GCmark状态。流程如下:

1. 调用gcBgMarkStartWorkers为每个p启动一个gcBgMarkWorker的G，用于标记全局和每个p上分配的内存。gcBgMarkWorker启动之后是可以长期存在的，因此不会每次执行GC都去创建，但在一些条件下部分gcBgMarkWorker会退出，这时gcStart会重新启动一个。gcStart会等待这些G全部创建完毕再进入下一步。但这些G不会立刻开始执行标记，而是要等到标记阶段后被gcController.findRunnableGCWorker唤醒执行。gcController.findRunnableGCWorker在runtime.schedule中被调用，如果正处于GC状态且worker数量没有达到阈值，就会运行gcBgMarkWorker。
2. stopTheWorldWithSema,停止所有任务
3. gcController.startCycle()初始化一些GC执行参数，其中主要包括gcControllerState.dedicatedMarkWorkNeeded和gcControllerState.fractionalUtilizationGoal.这两个参数用于计算和设置每个P上的gcBgMarkWorker是独占P运行还是分时间片运行的。设置的原则和目的是控制GC标记所占用的CPU比例，这个比例通常由全局变量const gcBackgroundUtilization=0.25指定不超过25%
4. setGCPhase(_GCmark),将GC阶段修改为_GCmark状态。setGCphase除了修改状态外，更重要的是**启动写屏障(write barrier)**如果状态为_GCmark/_GCmarktermination，写屏障就会打开。
5. gcMarkRootPrepare，收集根节点信息，计算出根节点标记的任务量，记录在work.markrootJobs中。
6. gcMarkTinyAllocs,完成每个p的小型对象内存标记。**gcMarkTinyAllocs会将每个p使用本地缓存mcache分配的小型数据结构地址标记为灰色并加入到p.gcw队列中**