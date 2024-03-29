# What is database

## Question

什么是红黑树？
什么是AVL树?
什么是B树
什么是B+树

## 索引

+ 聚集索引:索引文件和数据文件在一起的(Innodb)
+ 非聚集索引:索引文件和数据文件分开存储的(MyISAM) 

## 数据库的发展历史

### 数据库1.0----文件系统

当我们需要存储图书馆的所有图书信息的时候，一开始我们将所有图书信息放在csv文件中

Book.csv ( title , author , year )
> "Gone with the Wind","Margaret Mitchell",1936
> 
> "Hamlet","William Shakespeare",1602
> 
> "活着","余华",1993
> 
> "三体","刘慈欣",2006

这种存储方式，实现起来简单，但当我们需要查询某一个图书内容的时候，这种存储模式，我们只能通过遍历的方式去查询，这是一种非常糟糕的查询方式。

面对庞大的数据量，数据被存放在多个文件里面。每次查询，我们需要打开多个文件，打开后还要遍历里面的数据，磁盘IO和时间复杂度都很高。

问题的根源在于:我们存储数据的时候是没有规律的

一旦存储数据的时候没有规律，在我们查找数据的时候，就只能"抹黑"查找，这极大降低了存储本身的意义。

所以，让数据规律存储，是优化这个文件系统的第一步。

### 数据库2.0---规律存储

让数据有规律的存储，一旦数据有规律，我们就可以使用各种算法去高效地查找他们。

让数据按照字典排序升序存储，于是我们可以进行二分查找，时间复杂度从o(n)->o(log2n)，缺点是每次插入都要排序

让书籍按照hash表的结构进行排序，于是我们可以进行hash查找，用空间换时间，时间复杂度是o(1)

让书籍按照二叉树的结果进行存储，于是我们可以进行二叉查找，时间复杂度为o(log2n)

二叉树极端的情况下会退化成o(n)，于是有了平衡二叉树

平衡二叉树终究还是二叉，只有两个节点，一次从磁盘load的数据太少了，于是有了可以有多于2个子节点的b树

b树找出来的数据是无序的，如果你要求数据排好序返回，还要在内存手动排一次序，于是有了叶子节点是一个双向链表的b+树






