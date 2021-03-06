## Mysql 

## Question

linux文件系统，如何读取文件，文件的存储?
MyISAM为什么不支持事务


### 为什么Mysql最终决定用B+树呢

+ Hash(哈希表)

哈希表是做数据快速检索的有效利器

hash算法，也叫散列算法，就是把任意值(key)通过哈希函数变换为固定长度的key地址，通过这个地址进行具体数据的数据结构。考虑到这个数据库表user，表中一共有7个数据，我们需要检索id=7的数据，sql语法是 `select * from user where id=7;`哈希算法首先计算存储id=7的数据的物理地址，通过该独立地址可以找到对应数据。这就是哈希算法快速检索数据的计算过程。

但是hash算法有个数据碰撞问题，也就是hash函数可能对不同的key会计算出同一个结果。解决碰撞问题的一个常见处理方式就是链地址法，即用链表把碰撞的数据连接起来。计算hash值之后，还需要检查该hash值是否存在碰撞数据链表，有则一直遍历到链表尾，直到找到真正的key对应的数据之后

从算法的时间复杂度来看，hash算法的时间复杂度为O(1),检索速度非常快，为什么Mysql并没有采取hash作为底层存储算法呢？

1. 首先我们要考虑算法的优越性，不光要从时间复杂度来看，还需要从空间复杂度来看。key值经过hash转化指向元素value，先不考虑碰撞，因为虚拟内存是连续的，如果存储的两个value过大，那么进程会占用很大的空间来保证数据存储正确。
2. 数据检索不光是单个数据的查找，还有范围查找，比如下面这个SQL语句:
`select * from user where id>3;`
针对以上这个语句，我们希望查出id>3的数据。如果我们使用hash算法实现的索引，思路就是一次性把所有数据加载到内存中，然后再去内存里筛选目标范围内的数据。这种查找方式没有一点效率

+ 二叉查找树(BST)

![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/ca87b379915f46bd89f204185a93995c~tplv-k3u1fbpfcp-zoom-in-crop-mark:3024:0:0:0.awebp)

二叉查找树的时间复杂度是o(logn)，比如针对上面这个二叉树结构，我们只需要计算比较3次就可以检索到id=7的数据，相对于直接遍历查询省了一半的时间，从检索效率上看来是能做到高速检索的。此外二叉搜索树的结构还能解决哈希索引不能提供的范围查找功能。二叉搜索树的叶子节点都是按序排列的，从左到右依次升序排列，如果我们需要找id>5的数据，我们只需要取出节点为6的节点以及其右子树的就可以了。

但是普通的二叉查找树有一个致命的缺点，极端环境下会退化成线性表，二分查找也会退化成遍历查找，时间复杂度退化为O(n),检索性能急剧下降。在数据库中，数据的自增是一种很常见的形式，比如一个表的主键是id，而主键一般默认都是自增的，如果采用二叉搜索树作为索引，那么下面不平衡状态导致的线性查找问题必然出现。

![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/ca87b379915f46bd89f204185a93995c~tplv-k3u1fbpfcp-zoom-in-crop-mark:3024:0:0:0.awebp)

+ AVL树和红黑树

二叉查找树存在不平衡问题，学者提出通过树节点的自动旋转和调整，让二叉树始终保持基本平衡的状态，这样就能保持二叉搜索树的最佳查找性能了。

红黑树是一颗会自动调整树形态的树结构，比如当二叉树处于一个不平衡状态时，红黑树就会自动左旋右旋节点以及节点变色来调整树的形态，使其保持基本的平衡状态(时间复杂度为O(logn))，也就保证了查找效率不会明显降低.（红黑树牺牲了一定时间来构造一颗稳定的二叉搜索树来提高查找效率）

但是红黑树也存在相同的问题，只是相比于二叉查找树要好一点，但是整体还是呈现二叉查找树的极端形式。
![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/6ae838e436ef4213a5da301a8db5bca1~tplv-k3u1fbpfcp-zoom-in-crop-mark:3024:0:0:0.awebp)

现在考虑一种更为严格的自平衡二叉树AVL树。因为AVL树是个绝对平衡的二叉树，因此他在调整二叉树的形态上会消耗更多性能。从树的形态上面来看AVL树不存在红黑树"右倾"问题，也就是说，大量的顺序插入不会导致查询性能的降低。这从根本上解决了二叉搜索树和红黑树的问题

AVL树的优点

1. 不错的查找性能O(logn),**不存在极端的低效查找的情况**
2. 可以实现范围查找，数据排序。


AVL作为数据查找的数据结构确实很不错，但是AVL树并不适合做Mysql数据库的索引数据结构，因为数据库查询数据的瓶颈在于磁盘IO，如果使用的是AVL树，我们每一个树节点只存储了一个数据，我们一次磁盘IO只能取一个节点上的数据加载到内存里，那比如查找id=7这个数据我们就要进行磁盘IO三次，非常的消耗时间。当我们设计数据库索引的时候需要首先考虑怎么尽可能减少磁盘IO的次数

磁盘IO有个特点，就是从磁盘上面读取1B数据和1KB数据所消耗的时间是基本一样的，我们就可以根据这个思路，尽可能在一个树节点上尽可能存储多的数据，一次磁盘IO就能多加载数据到内存，这就是B树和B+树的设计原理了.**不管是什么数据结构，在底层不是通过指针连接就是顺序存储，我们必须尽可能将我们理论上需要的数据尽可能连续存储，保证一次磁盘IO就能够读取到多个想要的数据来保证数据库的查找效率**



+ B树

因为磁盘IO读一个数据和读取100个数据的时间基本一致，我们的优化思路为:尽可能在一次磁盘IO种多读一点数据到内存，在树结构的体现就是**每个节点能存储的key可以适当增加**
![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/d7ce7195cf2a418e870dc75cb2c273e5~tplv-k3u1fbpfcp-zoom-in-crop-mark:3024:0:0:0.awebp)

对于数据库索引数据结构而言，B树是一个不错的选择，总体来说,B树有以下优点

1. 很好的检索速度，时间复杂度:B树的查找性能为O(h*logn)，其中h为树高，n为每个节点关键词的个数
2. 尽可能少的磁盘IO，加快了检索速度
3. 可以支持范围查找


b树和b+树有什么不同呢

1. B树一个节点里面存储的是数据，而B+树存储的是索引(地址)，所以B树里一个节点存不了很多数据，但是B+树一个节点能存很多索引，B+树叶子节点存所有的数据
2. B+树的叶子节点是数据阶段用了一个链表串联起来，便于范围查找找


通过B树和B+树的对比我们可以看出，B+树存储的是索引，在单个节点存储容量有限的情况下，单节点也能存储大量索引，使得整个B+树高度降低，减少磁盘IO。其次，B+树的叶子节点是真正数据存储的地方，叶子节点用了链表连接起来，这个链表本身就是有序的，在数据范围查找时，更具备效率。所以Mysql的索引用的就是B+树

[为什么选择B+树](https://juejin.cn/post/7109421951025709093)


### Mysql的存储引擎

Mysqll底层数据引擎以插件形式设计，最常见的是Innodb引擎和Myisam引擎，用户根据个人需求选择不同的引擎作为mysql数据库的底层引擎。我们知道了B+树作为mysql的索引数据结构非常合适，但是数据和索引到底怎么组织起来也是需要设计的，设计理念的不同导致了Innodb和Myisam的出现，各自呈现独特的性能

MyISAM虽然数据查找性能极佳，但是不支持事务处理。Innodb最大的特色就是支持ACID兼容的事务功能，而且他支持行级锁。mysql建立表的时候就可以指定引擎。


+ Innodb 创建表后生成的文件有：

frm:创建表的语句
idb:表里面的数据+索引文件

+ Myisam 创建表后生成的文件有

frm:创建表的语句
MYD:表里面的数据文件（myisam data）
MYI:表里面的索引文件（myisam index）

从生成的文件来看，这两个引擎底层数据结构和索引的组织方式并不一样，mylsam引擎把数据和索引分开了，一人一个文件，这叫做非聚集索引方式；innodb引擎把数据和索引放在同一个文件里了，这叫做聚集索引方式。


1. MyISAM引擎的底层实现(非聚集索引方式)
     
       MyISAM用的是非聚集索引方式，即数据和索引落在不同的两个文件上。MyISAM在建表时以主键作为KEY来建立主索引B+树，树的叶子节点存储的是对应数据的物理地址，我们拿到这个物理地址后，就可以到MyISAM数据文件中直接定位到具体的数据记录了。
       ![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/608959a02a1a4ba9b0604a135d7107e8~tplv-k3u1fbpfcp-zoom-in-crop-mark:3024:0:0:0.awebp)
       当我们为某个字段添加索引时，我们同样也会生成对应字段的索引树，该字段的索引树的叶子节点同样时对应了数据的物理地址，然后也是拿着这个物理地址去数据文件中定位到具体的数据记录

2. Innodb引擎的底层实现(聚集索引方式) 
  
       InnoDB时聚集索引方式，因此索引和数据都存储在同一个文件里面。首先InnoDB会根据主键ID作为KEY 建立B+树，B+树叶子节点存储的是主键ID对应的数据。
       
       建表的时候InnoDB就会自动建立好主键ID索引树，这也是为什么Mysql在建表的时候要求指定主键的原因。当我们为表里某个字段添加索引时InnoDB会怎么建立索引树呢？比如我们要给user_name这个字段加索引，那么InnoDB就会建立user_name索引B+树，节点里面存的是user_name这个KEY，叶子节点存储的是主键KEY。**注意，叶子节点存储的是主键KEY!**。拿到主键KEY后，InnoDB才会去主键索引树里根据刚在user_name索引树找到的主键key查找对应的数据。InnoDB这样做的原因是为了节省存储空间。一个表里可能会有许多个索引，InnoDB都会给每个加了索引的字段生成索引树，如果每个字段索引树都存储了具体数据，那么这个表的索引数据文件就变成非常巨大。

+ MyISAM直接找到物理地址后就可以直接定位到数据记录，但是InnoDB查询到叶子节点后，还需要再查询一次主键索引树，才能定位到具体数据。等于MyISAM一步就查到了数据，但是InnoDB需要两步



### Mysql的锁

Mysql的锁分为表锁和行锁

表锁:开销小，不会产生死锁，发生锁冲突概率高，并且并发度低。

行锁:开销大，会产生死锁，发生锁冲突的概率低，并发度高。

+ MyISAM的锁机制

读锁:当某一线程对某张表进行读操作时，其他进程也可以读，但是不能写

写锁:当某一线程对某种表某张表进行写操作时，其他线程不能写也不能读。

因此MyISAM的读操作和写操作，以及写操作之间是串行的。MyISAM在执行读写操作的时候会自动给表加上相应的锁

在MyISAM中当一个进程请求某张表的读锁，而另一个进程同时也请求写锁，Mysql会先让后者获得写锁。即使读请求比写请求先到达锁等待队列，写锁也会插入到读锁之前。
因为Mysql总是认为写请求一般比读请求重要，这也就是MyISAM不太适合有大量的读写操作的应用的原因，因为大量的写请求会让查询操作很难获取到读锁，有可能永远阻塞。


### InnoDB锁模式

InnoDB实现了两种类型的行锁

共享锁(S):允许一个事务去读一行，阻止其他事务获得相同数据集的排他锁

排他锁(X):允许获得排他锁的事务更新数据，但是组织其他事务获得相同数据集的共享锁和排他锁

可以这么理解:

共享锁就是我读的时候，你可以读，但是不能写。排他锁就是我写的时候，你不能读也不能写。其实就是MyISAM的读锁和写锁，但是针对的对象不同了。

InnoDB还有两个类型的表锁

意向共享锁(IS):表示事务准备给数据行加入共享锁，也就是说一个数据行加共享锁前必须先取得该表的IS锁。

意向排他锁(IX),类似上面，表示事务准备给数据行加入排他锁，说明事务在一个数据行加排他锁前必须先取得该表的IX锁

![](https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/gold-user-assets/2019/8/24/16cc2d1a98a8b003~tplv-t2oaga2asx-zoom-in-crop-mark:3024:0:0:0.awebp)


注意:

当一个事务请求的锁模式与当前的锁模式兼容，InnoDB就将请求的锁授予该事务，反之如果请求不兼容，则该事务就等待锁释放。

意向锁是InnoDB自动加的，不需要用户干预

 

