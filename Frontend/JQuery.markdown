# JQuery


## Question

HTML加载顺序?
什么时候用post，什么时候用get
monitor_create monitor

JQuery是一个优秀的JavaScript库，极大地简化了JavaScript开发人员遍历HTML文档，操作DOM，执行动画和开发Ajax操作


JQuery能做以下事情:

  + HTML元素选取
  + HTML元素操作
  + CSS操作
  + HTML事件函数
  + JavaScript特效和动画
  + HTML DOM遍历和修改
  + AJAX
  + Utilities


JQuery库只提供了一个叫JQuery的函数，该函数中以及该函数的原型中定义了大量的方法。JQuery函数具有四种参数:

+ 选择器(字符串):jquery函数通过该选择器获取对应的DOM，然后将这些DOM封装到一个jQuery对象中并返回
+ DOM对象(即Node实例):jQuery函数将该DOM封装成jQuery对象并返回
+ HTML文本字符串：jQuery函数根据传入的文本创建好HTML元素并封装成jQuery对象并返回。$("<div class="one">one</div>");
+ 一个匿名函数$(function(){})当文档加载完毕之后jquery函数调用匿名函数。


### jQuery选择器

+ 基本选择器:

 >      所有选择器 \*
>     
>      标签选择器 标签名
>     
>      ID选择器 #id
>     
>      类选择器 .className
>     
>      群组选择器 .one,.two 多个选择器使用逗号分隔，取并集
>     
>      复合选择器 .one.two 多个选择器组合使用，取交集


+ 层次选择器
   
>    后代选择器 .one .two(两个选择器使用空格隔开，表示可以获取当前元素的子代以及孙子代等等)
> 
>    子代选择器   .one>.two(两个选择器使用>隔开,表示只能获取当前选中元素的子代元素)
> 

+ 兄弟选择器:
    
>   下一个兄弟选择器 .one+.two(两个选择器使用+隔开，表示可以获取当前元素的下一个兄弟元素，下一个兄弟元素要符合.two)
>    
>   之后所有兄弟选择器 .one~.two(两个选择器使用~隔开，表示可以获取当前元素之后的所有兄弟元素，只要所有兄弟元素能符合.two)


### jquery过滤器

jQuery过滤器必须用在jQuery选择器后，表示对通过前面的jQuery选择器选择到的内容进行过滤。语法: selector:过滤器

+ 基本过滤器

 >  selector:first 获取所有已选择到的元素中第一个元素
> 
>   selector:last 获取所有已选择到的元素中的最后一个元素
> 
>   selector:even 获取所有已选择到的元素中索引为偶数的元素
> 
>   selector:odd 获取所有已选择到的元素中索引为奇数的元素
> 
>   selector:eq(index) 获取所有已选择到的元素中索引为index的元素
> 
>   selector:lt(num) 获取所有已选择到的元素中索引小于num的元素
> 
>   selector:gt(num)  获取所有已选择到的元素中索引大于num的元素
> 
>   selector:not(selector2) 获取所有已选择到的元素除了selector2的元素
> 
>   selector:header 获取所有已选择到的元素中的标题元素(h1~h6)

+ 内容过滤器

  > selector:contains(text)获取所有已选择到的元素中包含text的元素
> 
>   selector:empty 获取所有已选择到的元素中的空元素
> 
>   selector:parent 获取所有已选择到的元素中非空的元素(有子节点)
> 
>   selector1:has(selector2)获取所有已选择到的元素中包含selector2的元素
>   


+ 




match()正则表达式匹配

$.Ajax或$.get或$.post，必须要指定返回类型，否则会出现解析错误



datatable params:

datasrc:对返回的数据进行处理，比如返回一系列json数据，但只对其中一个数据进行做表

ajax默认返回字符串，只有当确定"dataType"的时候才返回指定类型
var trobj=$("<tr></tr>")


JS中for(x in y)
x代表下标，而不是值



$("#searchTemple option:selected").text()
获取下拉框的值

$().empty()
删除元素所有的子节点


children()函数用于选取每个匹配元素的子元素，并以jQuery对象返回。你还可以使用选择器进一步缩小筛选范围，筛选出符合指定选择器的元素。

children()函数的返回值为jQuery类型，返回一个新的jQuery对象，该对象封装了当前jQuery对象匹配元素的所有符合指定选择器的子元素。

如果没有匹配的元素，则返回空的jQuery对象。

children()函数只在当前jQuery对象匹配元素的所有子元素中查找，不会查找"孙子"以及更后代的元素。



## JQuery对象与js的dom对象的区别

在很多场景中，我们需要jQuery与DOM能够相互的转换，它们都是可以操作的DOM元素，**jQuery是一个类数组对象，而DOM对象就是一个单独的DOM元素。**

相比较jQuery转化成DOM，开发中更多的情况是把一个dom对象加工成jQuery对象。$(参数)是一个多功能的方法，通过传递不同的参数而产生不同的作用。

   如果传递给$(DOM)函数的参数是一个DOM对象，jQuery方法会把这个DOM对象给包装成一个新的jQuery对象

   通过$(dom)方法将普通的dom对象加工成jQuery对象之后，我们就可以调用jQuery的方法了
