# Service


将运行在一组Pods上的应用程序公开为网络服务的抽象方法。
kubernetes为pod提供自己的IP地址，并为一组Pod提供相同的DNS名，并且可以在它们之间进行负载均衡


## 动机

创建和销毁kubernetes Pod以匹配集群的期望状态。pod是非永久性资源。如果使用Deployment来运行应用程序，则它可以动态创建和销毁pod。

每个pod都有自己的ip地址，但是在deployment中，在同一时刻运行的pod集合可能与稍后运行该应用程序的pod集合不同。

这导致一个问题:如果一组pod为集群内的其他pod提供功能，那么前端如何找出并跟踪要连接的IP地址，以便前端可以使用提供工作负载的后端部分


## service资源

kubernetes service定义了这样一种抽象:逻辑上的一组Pod,一种可以访问它们的策略--通常称为微服务。service所针对的pod集合通常是通过**选择算符**来确定的。

service定义的抽象能够解耦服务发现所带来的关联

## 云原生服务发现

如果想要在应用程序中使用kubernetes API进行服务发现，则可以查询API服务器用于匹配EndpointSlices。只要服务中的pod集合发生改变，kubernetes就会为服务更新EndpointSlices。

## 定义service

service在kubernetes中是一个REST对象，和pod类似。

kubernetes为该service服务分配一个IP地址，该IP地址由服务代理使用。

即使 Service 中使用同一配置名称混合使用多个 Pod，各 Pod 通过不同的端口号支持相同的网络协议， 此功能也可以使用。这为 Service 的部署和演化提供了很大的灵活性。 例如，你可以在新版本中更改 Pod 中后端软件公开的端口号，而不会破坏客户端。


## 没有选择算符的service

由于没有选择算符的存在，服务最常见的用法是为kubernetes pod 的访问提供抽象，但是当**与相应的EndpointSlices对象一起使用且没有选择算符时，服务也可以为其他类型的后端提供抽象，包括在集群外运行的后端。**


## 自定义EndpointSlices

当为服务创建EndpointSlice对象时，可以为EndpointSlice使用任何名称。命名空间的每个EndpointSlice必须有一个唯一的名称。通过在EndpointSlice上设置kubernetes.io/service-name label可以将EndpointSlice链接到服务