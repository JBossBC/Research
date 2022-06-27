# [剑指 Offer II 091. 粉刷房子 ]( https://leetcode.cn/problems/JEj789/submissions/ )

假如有一排房子，共 n 个，每个房子可以被粉刷成红色、蓝色或者绿色这三种颜色中的一种，你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。

当然，因为市场上不同颜色油漆的价格不同，所以房子粉刷成不同颜色的花费成本也是不同的。每个房子粉刷成不同颜色的花费是以一个 n x 3 的正整数矩阵 costs 来表示的。

例如，costs[0][0] 表示第 0 号房子粉刷成红色的成本花费；costs[1][2] 表示第 1 号房子粉刷成绿色的花费，以此类推。

请计算出粉刷完所有房子最少的花费成本。







## Greedy algorithm
    class Solution {
    public int minCost(int[][] costs) {
    int [] canUseColor=new int[]{0,0,0,0};
    int tempSelect=0;
    int preColor=3;
    int selectColor=0;
    int sumCost=0;
    //进入循环
    for(int i=0;i<costs.length;i++){
    tempSelect=0;
    while(tempSelect<3){
    if(costs[i][tempSelect]<costs[i][selectColor]&&canUseColor[tempSelect]==0){
    selectColor=tempSelect;
    }
    tempSelect++;
    }
    sumCost+=costs[i][selectColor];
    canUseColor[preColor]=0;
    canUseColor[selectColor]=1;
    preColor=selectColor;
    }
    return sumCost;
    }
    }