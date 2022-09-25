# Raft

raft是一个用于管理日志一致性的协议。它将分布式一致性分解为多个子问题:Leader选举、日志复制、安全性、日志压缩等。同时，raft算法使用了更强的假设来减少了需要考虑的状态，使之变的易于理解和实现。raft将系统中的角色分为领导者(leader)、跟从者(follower)、候选者(Candidate)

+ leader:接受客户端请求，并向Follower同步请求日志，当日志同步到大多数节点后告诉Follower提交日志
+ Follower:接受并持久化Leader同步的日志,在Leader告之日志可以提交之后，提交日志。
+ Candidate:leader选举过程中的临时角色

   raft要求系统在任意时刻最多只有一个Leader,正常工作期间只有leader和followers。raft算法将时间分为一个个的任期，每一个term的开始都是Leader选举。在成功选举Leader之后，Leader会在整个term内管理整个集群。如果Leader选举失败，该term就会因为没有Leader而结束。

## Term

raft算法将时间划分为任意不同长期的任期。任期用连续的数学进行表示。每一个任期的开始都是一次选举，一个或多个候选人会尝试成为领导人。如果一个候选人赢得了选举,它就会在该任期的剩余时间担任领导人。在某些情况下,选票会被瓜分，有可能没有选出领导人,那么，将会开始另一个任期,并且立刻开始下一次选举。raft算法保证在给定的一个任期最多只有一个领导人。

## RPC

raft算法中服务器节点之间通信使用远程过程调用(RPC),并且基本的一致性算法只需要两种类型的RPC,为了在服务器之间运输快照增加了第三种RPC

+ requestvote RPC:候选人在选举期间发起。
+ AppendEntries RPC:领导人发起的一种心跳机制，复制日志也在该命令中完成。
+ InstallSnapshot RPC:领导者使用该RPC来发送快照给落后的追随者

## Leader选举

raft使用心跳来触发Leader选举。当服务器启动时,初始化为Follower。Leader向所有Followers周期性发送heartbeat。如果Follower在选举超时时间内没有收到leader的heartbeat,就会等待一段随机事件后发起一次Leader选举

每一个follower都有一个时钟,是一个随机的值，表示的是follower等待成为leader的时间,谁的时钟先跑完,则发起leader选举

follower将其当前term加一然后转化为candidate,它首先给自己投票并且给集群中的其他服务器发送requestVote rpc。结果有以下三种情况

+ 赢得了多数的选票,成功选举为Leader
+ 收到了leader的消息，表示有其他服务器已经抢先当选了leader
+ 没有服务器赢得多数的选票，leader选举失败，等待选举时间超时后发起下一次选举


## leader选举的限制

**在raft协议中，所有的日志条目都只会从leader节点往follower节点写入，且leader节点上的日志只会增加，绝对不会删除或者覆盖。**


这意味着leader节点必须包含所有已经提交的日志，即能被选举为leader的节点一定需要包含所有的已经提交的日志。因为日志只会从leader向follower传输，所以如果被选举出的leader缺少已经commit的日志，那么这些已经提交的日志就会丢失，显然这是不符合要求的。
