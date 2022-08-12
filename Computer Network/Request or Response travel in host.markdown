# Request or Response travel in host


当通过诸如http.Get这样的API向服务器发送请求时，底层实现了以下几个过程

+ 通过DNS协议将域名解析为IP地址
+ 通过操作系统提供的系统调用创建一个socket连接，这实际上完成了TCP的三次握手过程
+ 通过socket连接以文本形式向服务端发送请求，在代码层面实际上是在向一个socket文件描述符写入数据，写入的数据就是一个HTTP请求.
+ 