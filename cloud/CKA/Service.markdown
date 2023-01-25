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

> 端口名称只能包含小写字母数字字符和-。端口名称还必须以字母数字字符开头和结尾。


## 选择自己的IP地址

在Service创建的请求中，可以通过设置spec.clusterIp字段来指定自己的集群IP地址。比如，希望替换一个已经已存在的DNS条目，或者遗留系统已经配置了一个固定的IP且很难重新配置。

用户选择的IP地址必须合法，并且这个IP地址在`service-cluster-ip-range`CIDR范围内，这对API服务器来说是通过一个标识来指定的。如果IP地址不合法，API服务器会返回HTTP状态码422，表示值不合法


## 服务发现

kubernetes支持两种基本的服务发现模式 --- 环境变量和DNS

### 环境变量

当Pod运行在Node上，kubelet会为每个活跃的Service添加一组环境变量。kubelet为Pod添加环境变量{SVCNAME}_SERVICE_HOST和{SVCNAME}_SERVICE_PORT.这里service的名称需要大写，横线被转换成下划线。


### DNS


支持集群的DNS服务器监视kubernetes API中的新服务，并为每个服务创建一组DNS记录。如果在整个集群中都启用了DNS，则所有Pod都应该能通过其DNS名称自动解析服务。

例如，如果你在kubernetes命名空间my-ns中有一个名为my-service的服务，则控制平面和DNS服务共同为my-service.my-ns创建DNS记录。my-ns命名空间中的Pod应该能够通过按名检索my-service来找到服务。

其他命名空间中的 Pod 必须将名称限定为 my-service.my-ns。 这些名称将解析为为服务分配的集群 IP。

Kubernetes 还支持命名端口的 DNS SRV（服务）记录。 如果 my-service.my-ns 服务具有名为 http　的端口，且协议设置为 TCP， 则可以对 _http._tcp.my-service.my-ns 执行 DNS SRV 查询以发现该端口号、"http" 以及 IP 地址。


## 发布服务


kubernetes serviceTypes允许指定你所需要的service类型。

Type的取值以及行为如下:

+ ClusterIP:通过集群的内部 IP 暴露服务，选择该值时服务只能够在集群内部访问。 这也是你没有为服务显式指定 type 时使用的默认值。
+ NodePort: 通过每个节点上的IP和静态端口暴露服务。 为了让节点端口可用，Kubernetes 设置了集群 IP 地址，这等同于你请求 type: ClusterIP 的服务。
+ LoadBalancer:使用云提供商的负载均衡器向外部暴露服务。外部负载均衡器可以将流量路由到自动创建的NodePort服务和ClusterIP服务上。
+ ExternalName:通过返回CNAME记录和对应值，可以将服务映射到externalName字段的内容。无需创建任何类型代理