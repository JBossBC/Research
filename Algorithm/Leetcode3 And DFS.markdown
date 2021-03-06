# [508. 出现次数最多的子树元素和]( https://leetcode.cn/problems/most-frequent-subtree-sum/ )

## Question
**DFS**

给你一个二叉树的根结点 root ，请返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。

一个结点的 「子树元素和」 定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。

 
## 解题思路

根据此题的理解，我们要做的是首先求出所有子树元素和，然后再比对返回出现次数最多的字数元素和。怎么求出所有子树元素和，我首先想到的是深度优先搜索，对于整棵树的根节点的子树元素和，他的组成是左子树的元素和加上右子树的元素和，同时左子树和右子树的元素和又同样按上面所述组成，由此分析，他们计算方式一样，只是计算规模减小，有分治法的味道。


## DFS(深度优先遍历)

深度优先遍历实现的基础是递归，那么递归函数是如何实现深度优先遍历的呢。首先我们来讲一下递归，对于算法来说，我们要用计算机的思维去解决问题，对于递归，简而言之就是一个函数在内部再调用这个函数，有此类特征的函数我们称为递归函数。对于计算机底层来说，就是一个函数结束后通过跳转指令返回到调用他的函数。那么对于一个函数的声明周期来说，他能向"外界"传递信息的方式就两种，一个是修改全局变量的值，一个是通过函数返回值。那么递归有什么特别的地方呢?根据程序的执行流程来说，我们将第几次递归调用的函数分别从小到大编号，那么