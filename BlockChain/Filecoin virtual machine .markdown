# FVM Introduction

## 0.Question
actor具体是指什么?
filecoin node 四种接口有什么不同?

## 1.what is filecoin virtual machine?
fvm充当以参与者形式为部署在链上的代码(智能合约)提供了一个执行的环境，

filecoin node implementations:

+ Lotus(使用go语言进行编写)，自己的参与者实施
+ Forest(使用rust语言进行编写),自己参与者实施
+  Venus(使用go语言进行编写，以前的Go filecoin实现),通过Go导入重用Lotus actors
+  Fuhon(用c++进行编写)，通过FFI重用Lotus actors

fvm