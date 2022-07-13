## TCP layer

![](https://img-blog.csdn.net/20181015142113271?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzQxNzI3MjE4/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

4位首部长度:限制了tcp头部不能超过64bit

6位标志位（即图中的保留6位）：标志位有如下几项

   + URG标志，表示紧急指针是否有效
   + ACK标志，表示确认号是否有效。称携带ACK标志的tcp报文段位确认报文段
   +  **PSH标志，提示接收端应用程序应该立即从tcp接受缓冲区中读走数据，为接受后续数据腾出空间（如果应用程序不将接收的数据读走，它们就会一直停留在tcp缓冲区中）**
   + RST标志，表示要求对方重新建立连接。携带RST标志的tcp报文段为复位报文段。
   + SYN标志，表示请求建立一个连接。携带SYN标志的tcp报文段为同步报文段。
   + FIN标志，表示通知对方本端要关闭连接了。携带FIN标志的tcp报文段为结束报文段。
  
16位窗口大小：是tcp流量控制的一个手段。这里说的窗口，指的是接收通告窗口。它告诉对方本端的tcp接收缓冲区还能容纳多少字节的数据，这样对方就可以控制发送数据的速度。


