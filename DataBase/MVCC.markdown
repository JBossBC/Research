# MVCC

## Question


乐观锁+版本号

----


MVCC:多版本并发控制，通过这种机制可以做到读写不阻塞，且避免了类似脏读这样的问题，主要通过undo日志链来实现

read commit:语句级快照

repeatable read:事务级快照



