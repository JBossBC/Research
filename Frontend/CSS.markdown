# CSS

css是一种用来结构化文档添加样式的计算机语言，CSS文件扩展名为.css。


## CSS盒子模型

所有HTML元素可以看作盒子,在CSS中,"box model"这一术语是用来设计和布局时使用。

css盒模型本质上是一个盒子，封装周围的HTML元素,它包括:边框,边距,填充,和实际内容。

盒模型允许我们在其他元素和周围元素边框之间的空间放置元素

+ Margin(外边距)-清楚边框外的区域，外边距是透明的
+ Border(边框)-围绕在内边距和内容外的边框
+ Padding(内边距)-清楚内容周围的区域，内边距是透明的
+ Content(内容)-盒子的内容，显示文本和图像。


> 当你指定一个CSS元素的宽度和高度属性时,只是设置内容区域的宽度和高度。

总元素的宽度=宽度+左填充+右填充+左边框+右边框+左边距+右边距



## CSS语法

css规则由两个主要的部分构成:选择器，以及一条或多条声明

选择器通常是您需要改变样式的HTML元素

每条声明由一个属性和一个值组成

属性是您希望设置的样式属性。每个属性有一个值。属性和值被冒号分开

## CSS id和class选择器

如果你要在HTML元素中设置CSS样式，你需要在元素中设置"id"和"class"选择器。

### id选择器

id选择器可以为标有特定id的HTML元素指定特定的样式。

HTML元素以id属性来设置id选择器，CSS中id选择器以"#"来定义

ID属性不要以数字开头，数字开头的ID在mozilla/firefox浏览器中不起作用


### class选择器

class选择器用于描述一组元素的样式，class选择器有别于id选择器，class可以在多个元素中使用。class选择器在HTML中以class属性表示，在CSS中，类选择器以一个点.号显示。

类名的第一个字符不能使用数字!它无法在mozilla或firefox中起作用


## CSS创建

插入样式表的方法有三种:

+ 外部样式表
+ 内部样式表
+ 内联样式


+ 外部样式表

当样式需要应用于很多页面时，外部样式表将会是理想的选择。在使用外部样式表的情况下，你可以通过改变一个文件来改变整个网页的外观。每个页面使用\<link>标签链接到样式表。\<link>标签在文档的头部:

> 不要在属性值与单位之间留有空格


+ 内部样式表

你可以使用<style>标签在文档头部定义内部样式表


+ 内联样式

由于要将表现与内容混杂在一起，内联样式会损失掉样式表许多优势

## CSS背景属性

background-attachment:背景图像是否固定或者随着页面的其余部分滚动

background-color:设置元素的背景颜色

background-image:把图像设置为背景

background-position:设置背景图像的起始位置

background-repeat:设置背景图像是否及如何重复


## text

文本颜色被用来设置文字的颜色:color

文本的对齐方式:文本排列属性是用来设置文本的水平对齐方式,当text-align设置为justify,每一行被展开为宽度相等，左右外边距是对其的

文本修饰:text-decoration(从设计的角度来看text-decoration属性主要是用来删除链接的下划线。)


文本转换:text-transform(文本转换属性用来指定在一个文本中的大写和小写字母。可用于所有字句变成大写或小写字母,或每个单词的首字母大写。)

文本缩进:text-indent(文本缩进属性是用来指定文本的第一行的缩进。)


## 链接

 + a:link-正常,未访问过的链接
 + a:visited-用户已访问过的链接
 + a:hover-当用户鼠标放在链接上时
 + a:active-链接被点击的那一刻

## display和visibility

display属性设置一个元素应如何显示，visibility属性指定一个元素应可见还是隐藏

### display:none或visibility:hidden

隐藏一个元素可以通过把display设置为none或者把visibility设置为hidden,两者的区别在于，visibility隐藏后的元素仍然会占用空间，而display隐藏后的元素不会占用任何空间

## position(定位)

position属性指定了元素的定位类型

+ static
+ relative
+ fixed
+ absolute
+ sticky


> 元素可以使用的顶部，底部，左侧和右侧属性定位。然而，这些属性无法工作，除非是先设定position属性。他们也有不同的工作方式

### static

html元素的默认值，即没有定位，遵循正常的文档流对象。静态定位的元素不会受到top,bottom,left,right影响

### fixed

元素的位置相对于浏览器窗口是固定位置,即使窗口是滚动的它也不会移动

> **fixed定位使元素的位置与文档流无关，因此不占据空间,fixed定位的元素和其他元素重叠**

### relative定位

相对定位元素的定位是相对其正常位置。

### absolute定位

绝对定位的元素的位置相对于最近的已定位的父元素,如果元素没有已定位的父元素，那么它的位置相对于<html>

> absolute定位使元素的位置与文档流无关，因此不占据空间

### sticky定位

sticky称为粘性定位,基于用户的滚动位置来定位,在position:relative与position:fixed定位之间切换。

他的行为就像position:relative,而当页面滚动超出目标区域时,它的表现就像position:fixedd;它会固定在目标位置

元素定位表现为在跨越特定阈值前为相对定位，之后为固定定位

这个特定阈值指的是top,right,bottom或left之一,换言之,指定top.right,bottom或left四个阈值其中之一,才可使粘性定位生效,否则其行为与相对定位相同。