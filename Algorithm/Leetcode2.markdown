# [剑指 Offer II 029. 排序的循环链表](https://leetcode.cn/problems/4ueAj6/)

给定循环单调非递减列表中的一个点，写一个函数向这个列表中插入一个新元素 insertVal ，使这个列表仍然是循环升序的。

给定的可以是这个列表中任意一个顶点的指针，并不一定是这个列表中最小元素的指针。

如果有多个满足条件的插入位置，可以选择任意一个位置插入新的值，插入后整个列表仍然保持有序。

如果列表为空（给定的节点是 null），需要创建一个循环有序列表并返回这个节点。否则。请返回原先给定的节点。

 

**解题思路:**

解决此题的首要问题是要理解什么是循环单调非递减列表
>输入：head = [3,4,1], insertVal = 2

>输出：[3,4,1,2]

>解释：在上图中，有一个包含三个元素的循环有序列表，你获得值为 3 的节点的指针，我们需要向表中插入元素 2 。新插入的节点应该在 1 和 3 之间，插入之后，整个列表如上图所示，最后返回节点 3 。

此题给出了一个示例，所谓单调非递减指的就是递增的一个序列，同时加加上循环这个限制，我们可以将这个列表看作一个循环列表来理解如何成立循环单调非递减列表。此题的本质就是在一个有序列表中插入一个节点，但此题的一些边界情况值得我们去研究，首先列表的生成有两种能够满足条件，第一种，从头节点开始到尾节点结束都是严格单调的，这种情况非常好说明，第二种情况，从头节点到某一个中间节点开始是单调的，这个中间节点的下一个节点到头节点同样是单调的，也就是循环二字的含义。那么题中有说道有一种情况能够满足多个插入条件的位置，这种情况就是插入元素比尾节点大，比首节点小的元素，它能够插到首节点的前面，也能插到尾节点的最后。我们尝试去用一种方式去均衡这两种情况来达到统一的解题方案。

**代码实现:**
    
    class Solution {
    public Node insert(Node head, int insertVal) {
    if(head==null){
    Node temp= new Node(insertVal);
    temp.next=temp;
    return temp;
    }
    Node FirstNode=head;
    while(FirstNode.next!=null){
    if((FirstNode.val<=insertVal&&FirstNode.next.val>=insertVal)||(FirstNode.val>FirstNode.next.val&&(FirstNode.next.val>insertVal||FirstNode.val<insertVal))){
       Node temp=new Node(insertVal,FirstNode.next);
       FirstNode.next=temp;
       break;
    }
    if(FirstNode.next==head){
    FirstNode.next=new Node(insertVal,head);
    break;
    }
    FirstNode=FirstNode.next;
    }
    return head;
    }
    }


此代码的时间复杂度为O(n),空间复杂度为O(1)