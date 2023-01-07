# go-zero rpc源码分析


在go-zero rpc 一个服务的main中:核心的方法有

+ conf.mustload(加载配置文件)
+ server.NewOrderServer(初始化服务的配置信息,加载context<-来自配置文件)
+ zrpc.MustNewServer(初始化rpc服务)****
+ s.start(启动rpc服务)



前两个基本上就很简单，加载etc/***.yaml文件,第二个注册一个服务，并且将配置信息注入到context里面。


## zrpc.MustNewServer

### NewServer

这里面传入了rpc服务对应的配置以及注册函数,返回了一个rpcServer.函数在处理过程中除了给rpcServer中的register赋值为注册函数之外，注册函数没有被其他地方使用，所以这一块我们就不管，重点在于对配置的处理。

+ 首先判断了rpc服务是否需要auth鉴权
+ 然后初始化了一个服务器
+ 向这个服务器里面注册metrics进行数据记录
+ 然后注册直连rpc服务还是向etcd注册rpc服务
+ 如果需要向etcd注册rpc服务,会写一个注册函数到RPCserver中的server里面
+ 最后通过setup开启prometheus、linktracing(根据你的etc下面的服务配置决定)、以及日志


## s.start


其实总的来说，除了rpc一些go-zero自带的中间件还没有注册之外，rpc注册函数以及向etcd里面注册的函数已经封装到rpcserver里面了，我们在这一步其实主要关注这两个函数的调用。

+ 使用适配器进入interal.rpcserver.start方法
+ 在处理过程中注册一些拦截器(prometheus、自动熔断、链路追踪)，同时，在流传输过程中注册一些拦截器


**这里面其实有一些步骤刚开始看起来感觉很怪,但其本质一点是因为作者封装grpc为RPCServer，但是这仅仅是封装而已，对于rpc服务启动最终还是grpc这一套，所以之前做的rpcServer配置，在这里的最重要的环节就是通过grpc.NewServer进行转化,将所有的配置信息(options、streamInterceptors、unaryInterceptors)全部传递给grpc.Server中的opts**

+ 调用register(server)->这里的server为grpc.NewServer(options...)生成的原生grpc的server

register中的pb.Register服务名()填充grpc.server中的services字段(根据上面传递的grpc.opts),reflection.register()将生成的grpcserver再向上封装一层为GRPCServer，这里封装了grpc server 同时还提供了服务端的XDS功能(详细的我也不太了解)

> Envoy通过文件系统或通过查询一个或多个管理服务器来发现其各种动态资源。这些发现服务及其相应的API统称为xDS。 通过订阅，指定要监视的文件系统路径，启动gRPC流或轮询REST-JSON URL来请求资源。后两种方法涉及使用DiscoveryRequest proto 载荷发送请求。在所有方法中资源以DiscoveryResponse proto 负载的形式发送。我们在下面讨论每种类型的订阅。


+ 下面的我也没太看懂，需要了解一下grpc的底层实现再来
