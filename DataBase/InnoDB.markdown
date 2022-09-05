# InnoDB

## Question

session产生的sort和join？？？

逻辑日志？？？


![](https://img-blog.csdnimg.cn/20210317103404374.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzMxOTYwNjIz,size_16,color_FFFFFF,t_70)


InnoDB存储引擎的具体架构如图。上半部分是实例层(计算层),位于内存中，下半部分是物理层，位于文件系统中


## 实例层

实例层分为线程和内存。InnoDB重要的线程有Master Thread，Master Thread是InnoDB的主线程，负责调度其他各线程。


Master Thread的优先级最高，其内部包含几个循环:主循环(loop)、后台循环(background loop)、刷新循环(flush loop)、暂停循环(suspend loop)。Master Thread会根据其内部运行的相关状态在各循环间进行切换??????????

大部分操作在主循环(loop)中完成，其包含1s和10s两种操作。

### 1s操作主要包括如下

 + 日志缓冲刷新到磁盘(这个操作总是被执行，即使事务还没有提交)
 + 最多可能刷100个新脏页到磁盘。
 + 执行并改变缓冲的操作。
 + 若当前没有用户活动，可能切换到后台循环等

### 10s操作主要包括如下

  + 最多可能刷新100个脏页到磁盘
  + 合并至多5个被改变的缓冲(总是)
  + 日志缓冲刷新到磁盘(总是)
  + 删除无用的undo页(总是)
  + 刷新100个或者10个脏页到磁盘(总是)产生一个检查点(总是)等。
  + buf_dump_thread负责将buffer pool中的内容dump到物理文件中，以便再次启动mysql时，可以快速加热数据。(redo log防止buffer pool没有加载到物理文件中的时候发生宕机导致数据丢失)
  + page_cleaner_thread负责将buffer pool中的脏页刷新到磁盘，刷新操作都是由主线程完成的，所以在刷新脏页时会非常影响mysql的处理能力，在5.7版本之后可以通过参数设置开启多个page_cleaner_thread。
  + purge_thread负责将不再使用的undo日志进行回收。
  + read_thread处理用户的读请求，并负责将数据页从磁盘上读取出来，可以通过参数设置线程数量。
  + write_thread负责将数据页从缓冲区写入磁盘，也可以通过参数设置线程数量，page_cleaner线程发起刷脏页操作后write_thread就可以开始工作了。
  + redo_log_thread负责把日志缓冲中的内容刷新到Redo log 文件中
  + insert_buffer_thread负责把insert buffer中的内容刷新到磁盘，实力层的内存部分主要包括innodb buffer pool，这里包含innodb最重要的缓存内容。数据和索引页、undo页、insert buffer页、自适应Hash索引页、数据字典页和锁信息等。additional memory pool后续已经不再使用。Redo buffer里存储数据修改所产生的Redo log。double write buffer 是double write所需要的buffer，主要解决由于宕机引起的物理写入操作中断，数据页不完成的问题。


## 物理层


**物理层在逻辑上分为系统表空间、用户表空间、redo日志和undo日志等**

系统表空间里有ibdata文件和一些undo，ibdata文件里有insert buffer段、double write段、回滚段、索引段、数据字典段和undo信息段


用户表空间是指以.ibd为后缀的文件，文件中包含insert buffer的bitmap页、叶子页(这里存储真正的用户数据)、非叶子页。InnoDB是索引组织表，采用B+树组织存储，数据都存储在叶子节点中，分支节点(即非叶子页)存储索引分支查找的数据值。

Redo日志中包括多个Redo文件，这些文件循环使用，当达到一定存储阈值时会触发checkpoint刷脏页操作，同时也会在mysql实例异常宕机后重启，InnoDB表数据自动还原恢复过程中所使用。


![](https://img-blog.csdnimg.cn/20210317114758391.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzMxOTYwNjIz,size_16,color_FFFFFF,t_70#pic_center)



------------


## 缓冲池(Buffer Pool)

用户读取或者写入的最新数据都存储在Buffer Pool中，如果Buffer Pool中没有找到则会读取物理文件进行查找，之后存储到Buffer Pool中并返回给Mysql server。Buffer Pool采用LRU机制。

Buffer Pool决定了一个SQL执行得到速度快慢，如果查询结果页都在内存中则返回结果速度很快，否则会产生物理读(磁盘读)，返回结果时间变长，性能远不如存储在内存中。

但我们又不能将所有数据页都存储在Buffer Pool中，比如物理ibd文件有500DB，我们的机器不可能配置能容下500GB数据页的内存，因为我们这样做成本很高而且没有必要。

在单机单实例情况下，我们可以配置Buffer Pool为物理内存的60%~80%,剩余内存用于session产生的sort和join等，以及运维管理使用。

如果是单机多实例，所有实例的buffer pool总量也不要超过物理内存的80%。开始时我们可以根据经验设置一个Buffer Pool的经验值，比如16GB，之后业务在Mysql运行一段时间后可以根据 `show global status like '%buffer_pool wait%'`的值来看是否需要调整Buffer Pool的大小。


## 重做日志(Redo log)

> 确保事务的持久性。防止在发生故障的时间点，尚有脏页未写入磁盘，在重启mysql服务的时候，根据redo log 进行重做，从而达到事务的持久性这一特性。Redo log是一个循环复用的文件集，负责记录InnoDB中所有对Buffer Pool的物理修改日志.


### redo log执行流程

当Redo log文件空间中，检查点位置的LSN和最新写入的LSN差值达到Redo log文件总空间的75%后，InnoDB会进行异步刷新操作，直到降到75%以下，并释放redo log的空间

当checkpoint_age达到文件总量大小的90%后，会触发同步刷新，此时InnoDB处于挂起状态无法操作。

这样我们就看到Redo log的大小直接影响了数据库的处理能力，如果设置太小会导致强行checkpoint操作频繁刷新脏页，那我们就需要把Redo log设置的大一些，5.6版本之前Redo log总大小不能超过3.8GB，5.7版本之后就放开了这个限制，需要权衡考虑

事务提交时log buffer会刷新到redo log文件中，具体刷新机制由参数控制

若参数 innodb_file_per_table=ON，则表示用户建表时采用用户独立表空间，即一个表对应一组物理文件，.frm表定义文件和.idb表数据文件。????????????????????????

若这个参数设置为OFF，则表示用户建表存储在ibdata文件中，不建议采用共享表空间，这样会导致ibdata文件过大，而且当表删除后空间无法回收。独立表空间可以在用户删除大量数据后回收物理空间，执行一个DDL就可以将表的高水位降下来了。？？？？？？？？？？？？？？？？？？？？


在一个事务中的每一次sql操作之后都会写入一个redo log到buffer中，在最后commit的时候，必须先将该事务的所有日志写入到redo log file进行持久化(这里的写入是顺序写的)，待事务的commit操作完成才算完成。


![](https://img-blog.csdnimg.cn/20210317144803230.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzMxOTYwNjIz,size_16,color_FFFFFF,t_70#pic_center)


由于重做日志文件打开没有使用O_DIRECT选项，因此重做日志文件缓冲先写入文件系统缓存。为了确保重做日志写入磁盘，必须进行一次fsync操作。由于fsync的效率取决于磁盘的性能，因此磁盘的性能决定了事务提交的性能，也就是数据库的性能。由此我们可以得出在进行批量操作时，不要for循环里面嵌套事务。(磁盘IO一次操作一个数据和1w个连续数据的消耗时间是一样的)


### redo log记录形式

前面说过，redo log实际上记录数据页的变更，而这种变更记录是没必要全部保存，因此redo log实现采用了大小固定，循环写入的方式，当写到结尾时，会回到开头循环写日志

![](https://img-blog.csdnimg.cn/20210317113210313.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzMxOTYwNjIz,size_16,color_FFFFFF,t_70)


同时我们很容易得知，在innodb中，既有redo log需要刷盘，还有数据页也需要刷盘，redo log存在的意义主要就是降低对数据页刷盘的要求。在上图中，write pos表示redo log当前记录的LSN(逻辑序列号)位置，check point表示数据页更改记录刷盘后对应redo log所处的LSN(逻辑序列号)位置。write pos 到checkpoint之间的部分是redo log空着的部分，用于记录新的记录；check point到write pos之间是redo log待落盘的数据页更改记录。当write pos 追上check point时，会先推动check point往前移动，空出位置再记录新的日志。

启动innodb的时候，不管上次是正常关闭还是异常关闭，总是会进行恢复操作。因为redo log记录的是数据页的物理变化，因此恢复的啥时候速度要比逻辑日志要快很多。重启innodb时，首先会检查磁盘中数据页的lsn，如果数据页的lsn小于日志的lsn，则会从checkpoint开始恢复。还有一种情况，在宕机前正处于checkpoint的刷盘阶段，则数据页的刷盘进度超过了日志页的刷盘进度，此时会出现数据页中记录的lsn大于日志中的lsn，这时超出日志进度的部分将不会重做，因为这本身就已经做过的事情，无需在重做