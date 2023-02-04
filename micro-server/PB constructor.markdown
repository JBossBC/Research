# Protocol Buffers


 Protocol Buffers(后面简称protobuf)是google团队开发的一种语言中立，平台无关，可扩展的数据压缩编码方式（序列化），其很适合做数据存储或RPC数据交换格式。可用于通讯协议、数据存储等领域的语言无关、平台无关、可扩展的序列化结构数据格式。

## protocol buffer的作用

Protobuf的存在有两个原因，一个是为了从数据存储的角度来加速站点间数据传输速度，另一个是为了解决站点间数据传输的协议不规范问题。
传输数据的大小无疑是影响传输速度的关键因素，当前流行的数据传输协议（json、xml等）会携带一些“结构化”的数据（如标签等），另外，它们对数据的压缩也没有很极致，所以需要有一个对数据存储很“紧凑”，对数据压缩很高效的工具来进行革新。

最初的数据传输协议是request/response形式的，没有 protocol buffers 之前，google 已经存在了一种 request/response 格式，用于手动处理 request/response 的编码和反编码。但是这种协议往往没有很明确的格式，所以开发人员经常会遇到新旧版协议不兼容的问题。因此，急需一个协议不需要了解所有业务字段还能灵活地应对各种改动的需求。


## protocol buffers原理

protobuf对传输的数据采取一种最简单的key-value形式的存储方式（但其中有一种类型的数据不是k-v形式，后面会讲），这钟存储方式极大的节省了空间。除此之外protobuf还采取了varint(变长编码)形式来压缩数据，对体积较小的字段分配较少的空间，由此使得压缩后的文件非常“紧凑”。


protobuf定义了一种后缀名为“.proto”的描述型文件为待传输的结构化数据作为数据协议，待传输的数据必须符合“.proto”文件中的相关定义。“.proto”文件在简洁易读的同时很好地保留了原数据的结构信息，并且还给出了一些实用的关键字，灵活地让开发者对数据中的字段做选取，给了开发人员很大的发挥空间，基本解决了其他协议出现的被需求牵着走的局面。“.proto”文件可以看作一种数据传输的协议，需要开发人员按照语法编写，其还可以通过protobuf工具编译成C++/Java/Python对象。

protobuf数据是连续的key-value组成，每个key-value对表示一个字段，value可以是基础类型，也可以是一个子消息。其中，**key表示了该字段数据类型，字段id信息，而value则是该字段的原始数据**。若字段类型是string/bytes/子message（长度不固定），则在key和value之间还有一个值表示该字段的数据长度。


### varint编码(极致压缩int)

在实际场景中，一般我们不会使用到很大的数，因此使用4/8字节表示数值会有冗余。为了进一步减小序列化后的数据大小，protobuf引入了varint编码，解决数值表示过程中的冗余问题。

在protobuf中，key值/length/int32/int64/uint32/uint64使用了varint编码表示。

varint编码原理如下图所示，每个字节首位表示后面是否还有字节，而后7位表示数据。

### zigzag编码(解决负数不能极致压缩的问题)

由于负数在计算机中用补码表示，首位永远是1，无法使用varint编码进行压缩。为了解决此问题，protobuf引入了zigzag编码，目前，sint32和sint64类型的数据使用zigzag编码。

zigzag编码原理是，首先按照下表中的规律，将数据全部转化成正数，然后再用varint进行编码。






## protobuf存储原理

protobuf对不同数据类型进行分类，分别对其选择不同存储格式，不过多数数据对存储格式都是键值对的key-value形式

|type|meaning|used for|
|--|--|--|
|0|varint|int32,int64,uint32,uint64,sint32,sint64,bool,enum|
|1|64-bit|fixed64,sfixed64,double|
|2|Length-delimited|string,bytes,embedded message,packed repeated fields|
|5|32-bit|fixed32,sfixed32,float|

对于message这种结构体数据，protobuf把message中所有字段的数据表示成k-v/k-l-v形式后拼接在一起，减少分隔符所占用的空间。这样一来有个问题出现了，那就是如何在一长串的拼接起来的二进制数据中找到对应的field？，并且还要确定该field是k-v存储还是k-l-v存储的呢？

这个问题的答案在存储field的“key”中，“key”由两个值决定，分别是field id和该field的类型编号，其计算方法是key = （field_id << 3） | type，其中|代表两个二进制类型的拼接，示例如下图所示：

要注意的是，key在某些时地方也会被称作tag。


### varint类型存储(type=0)

在type=0的数据类型中，除了正数整型还存在负数整型,一个负数整数一般会被表示为一个很大的证书，因为计算机定义负数的符号位为数字的最高位。如果采用Varints这种编码方式表示一个负数，那么一定需要 10 个byte长度。

然而protobuf定义了一种新的sint32/64类型，其采用zigzag编码方式，首先将所有整数映射为无符号证书,然后再采用varint编码方式编码。这样一来绝对值小的整数，编码后也有一个较小的varint编码值,无缝对接。

### 32/64-bit类型存储(type=1、type=5)



### length-delimited类型存储(type=2)


## 压缩原理


protobuf采取的压缩算法是Varints，Varint是一种紧凑的、不定长的表示数字的算法，它用一个或多个字节来表示一个数字，其中值越小的数字使用越少的字节数，这样可以节省表示较小数字的空间。


