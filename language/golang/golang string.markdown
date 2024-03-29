
# golang String

在编程语言中，字符串是一种很重要的数据结构,通常由一系列字符组成。字符串一般由两种类型，一种在编译时指定长度，不能修改。一种具有动态的长度，可以修改。但是在Go语言中，字符串不能被修改，只能被访问,不能采取方式对字符串进行修改。

字符串的终止有两种方式，一种是C语言的隐式申明，以字符"\0"作为终止符。一种是GO的显示声明。Go语言运行时字符串string的表示结构如下

`
   type StringHeader struct{
     Data uintptr
     len int
  }
`

其中,Data指向底层的字符数组，len代表字符串的长度。**字符串在本质上是一串字符数组，每个字符在存储时都对应了一个或多个整数，这涉及到字符集的编码方式** 

Go语言中所有的文件都采用UTF-8的编码方式，同时字符常量使用UTF-8的编码字符集。UTF-8是一种长度可变的编码方式，可包含世界上大部分字符。

## 符文类型

Go语言的设计者认为,用字符表示字符串的组成元素可能会产生歧义,因为有些字符非常相似，这些相似的字符真正的区别在于其编码后的整数是不相同的，因此Go语言中使用符文类型来表示和区分字符串中的字符

## 字符串底层解析

字符串在词法解析的时候也有特殊表示。字符串常量在词法解析阶段最终会被标记为StringLit类型的Token并被传递到编译的下一个阶段。在语法分析阶段,采取递归下降的方式读取uft-8字符,单撇号或双引号是字符串的表示。分析的逻辑位于syntax/scanner.go文件中

如果在代码中识别到单撇号,则调用rawString函数;识别到双引号，则调用stdString函数,两者的处理略有不同。对于单撇号的处理比较简单,**一直循环向后读取,直到寻找到配对的单撇号**。双引号调用stdString函数，**如果出现另一个双引号则直接退出,如果出现了\\则对后面的字符进行转义**

在抽象语法树阶段，无论是import语句中包的路径、结构体中的字段标签还是字符串常量，都会调用strconv.Unquote(s)去掉字符串两边的引号等干扰，还原其本来的面目。

## 字符串拼接

当加号操作符两边是字符串时,编译时抽象语法树阶段具体操作的OP会被解析为OADDSTR。对两个字符串常量进行拼接时会将语法分析阶段调用noder.sum函数。noder.sum函数先将所有的字符串常量放到字符串数组中,然后调用strings.join函数完成对字符串常量数组的拼接。