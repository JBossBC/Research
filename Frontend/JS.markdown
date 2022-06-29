# JS

## Question

后端已经用viewUtil.writeJson封装成了json数组，为什么前端使用ajax的时候如果不指定dataType,默认success返回的是字符串呢？

undefined和null的区别

var trobj=$(\<tr>\</tr>)为什么成立?

## 为什么要学习JavaScript

1. HTML定义了网页的内容
2. CSS描述了网页的布局
3. JavaScript控制了网页的行为

JavaScript是一种轻量级的编程语言
JavaScript是可插入HTML页面的编程语言
JavaScript插入HTML页面后，可由所有的现代浏览器执行


DOM(document object model)(文档对象模型)是用于访问HTML元素的正式W3C标准。


JavaScript可以通过不同的方式来输出数据

+ 使用window.alert()弹出警告框
+ 使用document.write()方法将内容写到HTML文档中
+ 使用innerHTML写到HTML元素
+ 使用console.log()写到浏览器控制台



javascript数据类型
值类型:字符串、数字、布尔、空、未定义(Undefined)、symbol
引用数据类型:对象、数组、函数、正则和日期

对象的两种访问方式

1. 对象名.属性名
2.  对象名["属性名"]






构造函数

    function Persion(name,age){
     this.name=name;
     this.age=age;
     this.sayName=function(){console.log(this.name)};
    }


构造函数其实就是一个普通的函数，不同的是构造函数习惯上首字母大写，构造函数和普通函数的还有一个区别就是调用方式的不同，普通函数是直接调用，而构造函数需要使用new关键字进行调用。

构造函数执行创建对象的过程:

1. 调用构造函数，他会立刻创建一个新的对象
2. 将新建的对象设置为函数中this,在构造函数中可以使用this来引用新建的对象。
3. 逐行执行函数中的代码
4. 将新建的对象作为返回值返回



## this:

   + 当以函数的形式调用时，this是window
   + 当以方法的形式调用时，谁调用方法,this就是谁
   + 当以构造函数的形式调用，this就是新创建的那个对象




js切割字符串
   字符串.slice(begin,end)

