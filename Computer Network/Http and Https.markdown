# Http

## what is http?

http:超文本传输协议，是一个基于请求与响应，无状态，应用层的协议。设计HTTP的初衷是为了提供一种发布和接收HTML页面的方法

 | 版本 | 时间 |  内容 | 发展现状 |
 | :-:|:-:|:-: | :-: |
 |   HTTP/0.9  |1991年|不涉及数据包传输，规定客户端和服务器之间通信格式，只能GET请求|没有作为标准
 |HTTP/1.0|1996年|传输内容格式不限制，增加PUT、PATCH、HEAED、OPTIONS、DELETE命令|正式作为标准|
|HTTP/1.1|1997年|持久连接(长连接)、节约带宽、HOST域、管道机制、分块传输编码|2015年前广泛使用
|HTTP/2|2015年|多路复用、服务器推送、头信息压缩、二进制协议等|逐渐覆盖市场|
![](https://img-blog.csdn.net/20180723105652242?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3hpYW9taW5nMTAwMDAx/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)
![](https://img-blog.csdn.net/20180723105652242?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3hpYW9taW5nMTAwMDAx/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

多路复用:通过单一的HTTP/2连接请求发起多重的请求-响应消息，多个请求stream共享一个tcp连接，实现多留并行而不是依赖建立多个TCP连接。

### HTTP特点:

1. 无状态:协议对客户端没有状态存储，对事务处理没有记忆能力，比如访问一个网站需要反复进行登录操作
2. 无连接:HTTP/1.1之前，因为无状态特点，每次请求需要通过tcp三次握手四次挥手，和服务器重新建立连接。比如某个客户机在短时间多次请求同一资源，服务器并不能区别是否已经响应过用户的请求，所以每次需要重新相应请求，需要消耗不必要的时间和流量
3. 基于请求和相应:由客户端发起请求，服务端响应
4. 简单快速、灵活
5. 通信使用明文、请求和响应不会对通信方确认，无法保护数据的完整性

针对HTTP无状态的解决策略

+ cookie/session技术
+ HTTP/1.1持久连接(HTTP keep-alive)方法，只要任意一端没有明确提出断开连接，则保持TCP连接状态，在请求首部字段中的Connection:keep-alive即为表明使用了持久连接

### HTTPS特点:

+ 基于HTTP协议，通过SSL或TLS提供加密处理数据、验证对方身份以及保护数据完整性

       + 内容加密:采用混合加密技术，中间者无法直接查看明文内容
       + 验证身份，通过证书认证客户端访问的是自己的服务器
       + 保护数据完整性:防止传输的内容被中间人冒充或者篡改

> 混合加密:结合非对称加密和对称加密技术。客户端使用对称加密生成密钥对传输数据进行加密，然后使用非对称加密的公钥再对密钥进行加密，所以网络上传输的数据是被密钥加密的密文和用公钥加密后的秘密密钥，因此即便被黑客截取，因为没有私钥，所以无法获取到加密明文的密钥，也就无法获取到明文数据
>
>数字摘要:通过单项hash函数对原文进行哈希，将需加密的明文hash成一串固定长度的消息摘要，不同明文的消息摘要结果不相同。

![](https://img-blog.csdn.net/20180719103559793?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3hpYW9taW5nMTAwMDAx/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

![](https://img-blog.csdnimg.cn/20190803111825690.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3hpYW9taW5nMTAwMDAx,size_16,color_FFFFFF,t_70)