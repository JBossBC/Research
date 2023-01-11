# jaeger


jaeger是一个分布式追踪系统,它被用于监视基于微服务的分布式系统并对其进行故障排除

+ 分布式上下文传播
+ 分布式事务监控
+ 根本原因分析
+ 服务依赖分析
+ 性能延迟优化


## 特点

+ OpenTracing启发的数据模型
+ 对每个服务/endpoint的概率使用一致的前期抽样
+ 支持多样的存储方式:ES,Cassandra,in-memory
+ 社区通过grpc插件支持外部存储后端:postgresql、clickhouse
+ 系统拓扑图
+ 自适应采样
+ 收集后的数据进行管道化处理
+ 服务性能监控


## terminology

### span

一个span代表一个有操作名称以及开始和结束时间的逻辑工作单元。span互相之间可以嵌套以及排序来模拟一个因果关系


### trace

一个trace代表这个数据在这个系统的执行路径，他被认为是一个由span组成的有向无环图。


### baggage(行李)

baggage是用户定义的一个元数据(键值对)，它被定义在distributed context中，能够存在于一个trace的整个生命周期


## component

jaeger能够被部署为一个一体式二进制文件，也就是说所有的jaeger后端组件都将作为单个进程来运行，jaeger也能作为一个可伸缩的分布式系统来部署。

下面由两个主要的部署选项

+ 收集器被直接写入存储中
+ 收集器被写入kafka作为一个初步缓冲区


### jaeger client libraries


jaeger客户端是一个语言特有的接口，遵循OpenTracing规范。这意味只需要遵循OpenTracing这套接口，就可以和jaeger服务器无缝对接。

一个被检测的服务在接收新的请求时创建span并且附加context信息(trace id，spanid,and baggage).只有id和baggage会随着请求进行传播，其他的信息，比如说操作时间，操作名称，tags还有日志并不会随着请求进行传播。相反，他们会在后台异步发送到jaeger backend。


jaeger客户端在生产中始终开启，为了最大限度地减少开销，jaeger客户端采用了各种各样的策略去优化，当跟踪被采样时，分析跨度数据被捕获并传输到jaeger后端。如果未对跟踪进行采样，则根本不会收集分析数据，并且对OpenTracing API的调用进行抽样检测，以产生最小的开销。默认情况，jaeger客户端对0.1%的跟踪进行采样，并且能够从jaeger后端检索采样策略。


### agent

jaeger agent 是一个网络守护进程，它监听通过UDP发送的span,这些span被批处理然后发送到collector。它被设计为一个基础组件被部署到每个主机上面。agent从客户端抽象出收集器的路由和发现。

agent是能够被忽略的组件，也就是说你的应用程序可以通过直接将数据发送到collector的方式来省略agent

### collector

jaeger collector 接收 来自SDK或者jaeger agent的trace，通过处理管道运行它们以进行验证和清理、补充，并且将他们发送到存储点。

jaeger内置了对多个存储后端的支持，以及用于实现自定义存储插件的可扩展插件框架

### Query

jaeger query是一项服务，公开API，用于从存储中检索trace并且为web UI提供搜索，分析trace的功能


### ingester

jaeger ingester是一项读取kafka的trace并将其写入存储后端的服务，实际上，他是以jaeger collector的分离分离版本，它支持kafka作为唯一的输入协议。

笼统来说就是jaeger collector从直接到存储的模式分离，编程jaeger collector->kafka->存储后端。






