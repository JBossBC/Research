# netpoll

Multi-Reactor模型中根据角色的，可以将Reactor分为两类:mainRactor、subReactor。一般mainReactor是一个，而subReactor会有多个。Multi-Reactor会有多个。

Multi-Reactor模型的原理如下:

1. mainReactor主要负责接收客户端的连接请求，建立新连接，接受完连接后mainReactor就会按照一定的负载均衡策略分发给其中一个subReactor进行管理。
2. subReactor会将新的客户端连接进行管理，负责后续该客户端的请求处理。
3. 通常Reactor线程主要负责IO的操作(数据读写)、而业务逻辑的处理会由专门的工作线程来执行。

>此处所指的Reactor,以epoll为例可以简单理解为一个Reactor对应一个epoll对象,由一个线程进行及处理，Reactor线程又被称为IO线程。

## netpoll server端内部结构

+ Listener: 主要用来初始化Listener,内部调用标准库的net.Listen(),然后再封装了一层。具体实现则是调用socket()、bind()、listen()等系统第哦啊用。
+ EventLoop: 框架对外提供的结构，对外暴露serve()方法来创建server端程序
+ Poll:是抽象出的一套接口，屏蔽底层不同操作系统平台接口的差异,linux下采用epoll来实现、bsd平台下则采用kqueue来实现。
+ pollmanager:Poll的管理器，可以理解为一个poll池,也就是一组epoll或者kqueue集合。
+ loadbalance:负载均衡封装，主要用来从pollmanager按照一定的策略(随机、轮询、最小连接等)选择出来一个poll实例,一般在客户端初始化完成后，server会调用该接口拿到一个poll实例,并将新建立的客户端连接加入到poll管理。