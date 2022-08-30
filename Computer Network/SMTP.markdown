# SMTP

SMTP即简单邮件传输协议,它是一组用于由源地址到目的地址传送邮件的规则，由它来控制信件的中转方式。SMTP协议属于TCP/IP协议簇，它帮助每台计算机在发送或中转信件时找到下一个目的地。通过SMTP协议所指定的服务器，就可以把E-mail寄到收信人的服务器上了，整个过程只要几分钟。SMTP服务器则是遵守SMTP协议的发送邮件服务器，用来发送或中转发出的电子邮件。SMTP是一种TCP协议支持的提供可靠且有效电子邮件传输的应用层协议。


## 过程

首先，运行在发送端邮件服务器主机上的SMTP客户，发起建立一个到运行在接收端邮件服务器主机上的SMTP服务器端口号25之间的TCP连接。如果接收邮件服务器当前不在工作，SMTP客户就等待一段时间后再尝试建立该连接

![](https://img-blog.csdn.net/20170403164254171?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzU2NDQyMzQ=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

SMTP使用一些命令和应答，在MTA客户和MTA服务器之间进行传输报文。

![](https://img-blog.csdn.net/20170403164317197?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzU2NDQyMzQ=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


## 邮件报文的封装和报文格式

SMTP协议可以将互联网邮件报文封装在邮件对象中。SMTP协议的邮件对象由两个部分组成:信封和内容。

   + 信封实际上是一种SMTP命令
   + 邮件报文是邮件对象中的内容，它又有首部和主体两个部分

![](https://img-blog.csdn.net/20170403164927322?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzU2NDQyMzQ=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


RFC文档对报文格式的定义:

```
 + 所有报文都是由ASCII码组成
 + 报文由报文行组成，各行之间用回车(CR)、换行(LF)符分割
 + 报文的长度不能超过998个字符
 + 报文行的长度<=48个字符之内(不包括回车换行符)
 + 报文中柯包括多个首部字段和首部内容
 + 报文可包括一个主体,主体必须用一个空行与其首部分割
 + 除非需要使用回车与换行符，否则报文中不使用回车与换行符

```


## 邮件报文的发送过程

+ 建立连接

   + 从客户端使用熟知的端口号25建立与服务器的TCP连接，SMTP服务器向该客户送回应答码220，并且还为客户端提供了服务器的域名
   + 客户端收到应答码后，发送HELO命令，启动客户端和服务器之间的SMTP会话。该客户端发送的HELO用来向服务器提供客户端的标识信息
   + 服务器端回应应答码250,通知客户端:请求建立邮件服务会话已经实现

+ 报文发送
   + 客户用"Mail FROM"向服务器报告发信人的邮箱与域名
   + 服务器向客户回应应答码"250",代表请求命令完成
   + 客户用"RCPT TO"命令向服务器报告收信人的邮箱与域名
   + 服务器向客户回应应答码"250"，代表请求命令完成
   + 客户用"DATA"命令对报文的传送进行初始化
   + 服务器回应"354",表示可以进行邮件输入了
   + 客户用连续的行向服务器传送报文的内容，每行以两字符的行结束表示(CR与LF)终止
   + 服务器向客户回应响应码"250",代表请求命令完成

+ 连接终止
  
   + 客户端发送"QUIT"命令
   + 服务器收到后，回应应答码"221",并结束会话