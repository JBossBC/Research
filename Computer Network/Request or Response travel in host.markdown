# Request or Response travel in host


当通过诸如http.Get这样的API向服务器发送请求时，底层实现了以下几个过程

+ 通过DNS协议将域名解析为IP地址
+ 通过操作系统提供的系统调用创建一个socket连接，这实际上完成了TCP的三次握手过程
+ 通过socket连接以文本形式向服务端发送请求，在代码层面实际上是在向一个socket文件描述符写入数据，写入的数据就是一个HTTP请求.






Socket 通过 <源 IP、源 Port、目的 IP、目的 Port> 的四元组来区分 (实际上还有协议，TCP 或 UDP)，只要有一处不同，就是不同的 socket。因此，尽管 TCP 支持的端口号最多为 65535 个，但是每台机器理论上可以建立无数个 socket 连接。比如 HTTP 服务器只消耗一个 80 端口号，但可以和不同 IP:Port 的客户端建立连接，实际受限于操作系统的内存大小。
转载请附上原文出处链接及本声明。
