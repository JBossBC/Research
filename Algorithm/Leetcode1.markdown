# [Leetcode 1089:复写零](https://leetcode.cn/problems/duplicate-zeros/)

给你一个长度固定的整数数组 arr，请你将该数组中出现的每个零都复写一遍，并将其余的元素向右平移。

注意：请不要在超过该数组长度的位置写入元素。

要求：请对输入的数组 就地 进行上述修改，不要从函数返回任何东西。


+ 解题思路

 + 此题本质上是在原来的数组上面添加元素，通用方法是遍历所有元素，当遇到0元素的时候，将0后面的元素统一向后移一位，然后再对零进行复写，但由此可知此题在最差的情况下时间复杂度θ(n)=x^2，空间复杂度为O(n)=1,时间复杂度相对较高.
 
JAVA代码实现

     public static void duplicateZeros(int[] arr) {
              int tempIndex=0;
              int temp=0;
              boolean flag=true;
              for(int i=0;i<arr.length;i++) {
                  tempIndex = i;
                     if (arr[tempIndex] == 0) {
                       while (tempIndex + 1 < arr.length) {
                          if (flag) {
                             temp = arr[tempIndex + 1];
                             arr[tempIndex + 1] = 0;
                             flag = false;
                       } else {
                        temp = temp ^ arr[tempIndex + 1];
                        arr[tempIndex + 1] = temp ^ arr[tempIndex + 1];
                         temp = temp ^ arr[tempIndex + 1];
                        }
                        tempIndex++;
                        }
                        if (!flag) {
                        i++;
                        }
                        flag = true;
                        }
                        }
                        }

**用此种代码实现可知效率为**
> 执行用时：36 ms, 在所有 Java 提交中击败了5.60% 的用户

> 内存消耗：41.3 MB, 在所有 Java 提交中击败了86.69% 的用户

