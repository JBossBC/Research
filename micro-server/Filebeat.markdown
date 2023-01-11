# Filebeat


## 工作原理

当你开启filebeat的时候，他会启动一个或者多个输入去监控你指定的日志数据存在的位置。对于每一个filebeat查看的日志，filebeat开启一个harvest.每个harvest从最新的内容中读取单条的日志并且发送这个新的日志数据到libbeat，它聚合事件并将聚合的数据发送到filebeat配置的输出。


## harvest

一个采集器负责读取单个文件的内容.采集器读取文件时，以行为单位，并且发送内容到output.采集器将对文件的开启和关闭负责，这意味着当这个收集器正在运行的时候文件描述符会一直保持打开。当采集器正在收集的时候，如果一个文件被删除或者重命名，采集器将继续读取该文件。这样有一个副作用，即磁盘上的空间被保留，直到收割机关闭。默认情况下，filebeat保持文件打开，直到达到close_inactive

关闭一个采集器有以下的结果

+ 文件处理程序关闭，如果在采集器仍在读取文件时删除了文件，则释放了底层资源。
+ 只有在scan_frequency结束后，才能再次开启获取文件
+ 当采集器被关闭时，如果一个文件被移动或者删除，这个文件的采集工作不会继续。


为了方便控制当一个采集器被关闭时的状态，使用close_* 配置选项进行配置

## input

一个input负责管理harvests并且找到所有的sources file 然后 读取他们

如果input的类型是log，input会找到对应paths配置下的所有符合的文件并且对每个符合的file开启一个harvester。每个input起一个单独的go routine进行工作。

filebeat 目前支持serveral input types.每个input类型能够被多次定义.log input检查每个file去判断是否需要开启采集器，是否已经有一个已经在running,或者是否这个文件被忽略掉(ignore_older). New lines are only picked up if the size of the file has changed since the harvester was closed.(实现可以看看)


## principle

### filebeat怎样保持文件的状态

filebeat保持每个文件的状态同时经常刷新状态到磁盘的注册文件中。这个状态被用来记录最后一次harvester读取的偏移量确保所有的日志行都会被发送。如果the output，例如ES或者Logstash是不可达的，filebeat将会保持跟踪发送的最后一行同时持续读取文件，等待output再一次变为可达。当filebeat正在running的时候，对于每个input这个状态信息将会保存在内存中。当filebeat被重启时，来自registry 文件的数据将会被用来重新构建状态，同时filebeat继续在已知的最后位置继续采集。


对于每个input,filebeat保持它找到的每个文件的状态。这个文件名和路径不能够充足说明这个文件的身份，因为这个文件可能被重命名或者移动。因此对于每个文件，filebeat存储了唯一身份标识去识别这个文件是否曾经被采集过。

如果你使用的时候需要每天创建大量的新文件，那么注册表文件将会变得很大，对于这种情况，参考配置文件解决。


### filebeat怎样确保至少一次传递

filebeat保证这些事件将会被传递到配置的output中至少一次，同时没有数据丢失。filebeat能够达到这个目的的原因在于它存储了每次事件的传递状态到注册表文件中。

对于定义的output被阻塞或者不确认所有的事件，filebeat将会保持尝试去发送事件直到the output承认他已经收到了这些事件。


如果filebeat在他还在发送事件的过程中关闭，它不会等待the output在关闭之前承认所有的事件。在filebeat关闭前，任何一个发送到output但是没有确认的事件，都将会在filebeat重启后再次发送。这个确保了每次event都至少被发送一次，但是最终可能会将重复的事件发送到输出(无法保证幂等性)，你可以配置filebeat来在关闭前等待特定的时间。
