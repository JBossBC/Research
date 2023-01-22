# Architecture

## node

kubernetes通过将容器放入在节点上运行的Pod来执行你的工作负载。节点可以是一个虚拟机或者物理机器，取决于所在的集群配置。每个节点包含运行Pod所需的服务;这些节点由控制面负责管理。

通常集群里会有若干个节点;而在一个资源受限的环境中，你的集群中也可能只有一个节点。

节点上的组件包括kubelet、容器运行时以及kube-proxy


### 管理


向API服务器添加节点的方式主要有两种:

1. 节点上kubelet向控制面执行自注册
2. 手动添加一个Node对象

在创建了Node对象或者节点上的kubelet执行了自注册操作之后，控制面会检查新的Node对象是否合法。

kubernetes会在内部创建一个Node对象作为节点的表示。kubernetes检查kubelet向API服务器注册节点时使用的metadata.name字段是否匹配。如果节点是简况的(即所有必要的服务都在运行中),则该节点可以用来运行Pod。否则，直到该节点变为健康之前，所有的集群活动都会忽略该节点，

> kubernetes会一直保存着非法节点对应的对象，并持续检查该节点是否已经变得健康。 可以通过管理员或者某个控制器显示删除该Node对象以停止健康检查操作。

Node对象的名称必须是合法的DNS子域名

### 节点名称唯一性

节点的名称是用来标识Node对象。没有两个Node可以同时使用相同的名称。kubernetes还假定名字相同的资源是同一个对象。就Node而言，隐式假定使用相同名称的实例会具有相同的状态(例如网络配置、根磁盘内容)和类似节点标签这类属性。**这可能在节点被更改但其名称未变时导致系统状态不一致。**如果某个Node需要被替换或者大量更改，需要从API服务器移除现有的Node对象，之后在更新之后重新将其加入。


### 节点自注册

当kubelet标志 --register-node 为 true(默认时),它会尝试向API服务注册自己。这是首选模式。

对于自注册模式，kubelet使用下列参数启动:

+ --kubeconfig:用于向API服务器执行身份认证所用的凭据的路径
+ --cloud-provider:与某云驱动进行注册以读取与自身相关的元数据的方式
+ --register-node: 自动向API服务注册
+  --register-with-taints:使用所给的污点列表注册节点。当register-node为false时无效
+  --node-ip:节点IP地址
+  --node-labels:在集群中注册节点时要添加的标签
+  --node-status-update-frequency:指定kubelet向API服务器发送其节点状态的频率


当Node鉴权模式和NodeRestriction准入插件被启用后，仅授权kubelet创建/修改自己的Node资源。

> 当Node配置需要被更新时，根据节点名称唯一性，一种好的做法是重新向API服务器注册该节点。例如，如果 kubelet 重启时其 --node-labels 是新的值集，但同一个 Node 名称已经被使用，则所作变更不会起作用， 因为节点标签是在 Node 注册时完成的。

> 如果在 kubelet 重启期间 Node 配置发生了变化，已经被调度到某 Node 上的 Pod 可能会出现行为不正常或者出现其他问题，例如，已经运行的 Pod 可能通过污点机制设置了与 Node 上新设置的标签相排斥的规则，也有一些其他 Pod， 本来与此 Pod 之间存在不兼容的问题，也会因为新的标签设置而被调到同一节点。 节点重新注册操作可以确保节点上所有 Pod 都被排空并被正确地重新调度。


### 手动节点管理

你可以使用kubectl来创建和修改Node对象

如果你希望手动创建节点对象，请设置kubelet标志 --register-node=false

你可以修改Node对象(忽略 --register-node设置)。你可以修改节点上的标签并将其标记为不可调度等

你可以结合使用Node上的标签和Pod上的选择运算符来控制调度。例如，你可以限制某Pod只能在符合要求的节点子集上运行。

如果标记节点为不可调度，将阻止新Pod调度到该Node之上，但不会影响任何已经在其上的Pod。这是重启节点或者执行其他维护操作之前一个有用的准备步骤。

标记一个Node为不可调度

` kubectl cordon $NODENAME`

> 被DaemonSet控制器创建的Pod能够容忍节点的不可调度属性。DaemonSet通常提供节点本地的服务，即使节点上的负载应用已经被腾空，这些服务也仍需运行在节点之上。

### 节点状态

一个节点的状态包含以下信息:

+ 地址
+ 状况
+ 容量与可分配
+ 信息

可以使用kubectl来查看节点状态和其他细节信息:

`kubectl describe node <节点名称>`

对于输出的每个部分

**地址**

这些字段的用法取决于你的云服务商或者物理机配置

+ HostName:由节点的内核报告。可以通过kubelet的`--hostname-override`参数覆盖
+ ExternalIP:通常是节点的可外部路由(从集群外可访问)的IP地址
+ InternalIP:通常是节点的仅可在集群内部路由的IP地址


**状况**

conditions字段描述了所有Running节点的状况

|节点状况|描述|
|--|--|
|ready|节点是健康的并已经准备好接收Pod为True;节点不健康而且不能接受Pod为False;节点控制器在最近node-monitor-grace-period期间(默认40秒)没有收到节点的消息为unknown|
|diskPressure|节点存在磁盘空间压力，即磁盘可用量低为True，否则为false|
|memoryPreesure|节点存在内存压力，即节点内存可用量低为True,否则为False|
|PIDPressure|节点存在进程压力，即节点上进程过多为True,否则为False|
|NetworkUnavailable|节点网络配置不正确为True,否则为False|


> 如果使用命令行工具来打印已保护(cordoned)节点的细节。其中的condition字段可能包括schedulingdisabled。schedulingdisabled不是kubernetes API中定义的Condition,被保护起来的节点在其规约中被标记为不可调度(unschedulable) 

如果Ready状况的status处于unknown或者false状态的时间超过了Pod-eviction-timeout值(一个传递给kube-controller-manager的参数)，节点控制器会对节点上的所有Pod出发API发起的驱逐。默认的逐出超时时长为5分钟。

某些情况下，当节点不可达时，API服务器不能和其上的kubelet通信。删除Pod的决定不能传达给kubelet,直到它重新建立和API服务器的连接为止。与此同时，被计划删除的Pod可能会继续在游离的节点上运行。


节点控制器在确认Pod在集群中已经停止运行前，不会强制删除它们。你可以看到在这些无法访问的节点上运行的Pod处于Terminating或者unknown状态。如果kubernetes不能基于下层基础设施推断出节点是否已经永久离开了集群，集群管理员可能需要手动删除该节点对象。从kubernetes删除节点对象将导致API服务器删除节点上所有运行的Pod对象并释放它们的名字。

当节点上出现问题时，kubernetes控制面会自动创建与影响节点的状况对应的污点。调度器在将pod指派到某Node时会考虑Node上的污点设置。Pod也可以设置容忍度，以便能够在设置了特定污点的Node上运行。

### 容量(Capacity)与可分配(Allocatable)

这两个值描述节点上的可用资源:CPU、内存和可以调度到该节点上的Pod的个数上限。

capacity块中的字段标示节点拥有的资源总量。allocatable块指示节点上可供普通Pod消耗的资源量。

### 信息(info)

info指的是节点的一般信息,如内核版本，kubernetes版本，容器运行时详细信息，以及节点使用的操作系统。kubelet从节点收集这些信息并将其发布到kubernetes API

### 心跳

kubernetes节点发送的心跳帮助你的集群确定每个节点的可用性，并在检测故障时采取行动。

对于节点，有两种形式的心跳。

+ 更新节点的.status
+ kube-node-lease名称空间中lease对象。每个节点都有一个关联的lease对象。

与Node的.status更新相比，lease是一种轻量级资源。使用lease来表达心跳在大型集群钟可以减少这些更新对性能的影响

kubelet负责创建的更新节点的.status,以及更新它们对应的lease。

+ 当节点状态发生变化时，或者在配置的时间间隔内没有更新事件时，kubelet会更新.status。.status更新的默认间隔为5分钟(比节点不可达事件的40秒默认超时时间要长很多).
+ kubelet会创建并每10秒(默认更新间隔时间)更新lease对象。lease的更新独立于Node的.status更新而发生。如果lease的更新操作失败，kubelet会采用指数回退机制，从200毫秒开始重试，最长重试间隔为7秒钟。


### 节点控制器

节点控制器是kubernetes控制面组件，管理节点的方方面面。

节点控制器在节点的生命周期中扮演多个角色。第一个是当节点注册时为它分配一个CIDR区段(如果启用了CIDR分配)

第二个是保持节点控制器内的节点列表与服务商所提供的可用机器列表同步。如果在云环境下运行，只要是某个节点不健康，节点控制器就会询问云服务是否节点的虚拟机仍可用。如果不可用，节点控制器会将该节点从它的节点列表删除。


第三个是监控节点的健康状况。节点控制器负责

+ 在节点不可达的情况下，在Node的.status中更新Ready状况。在这种情况下，节点控制器将NodeReady更新为Unknown
+ 如果节点仍然无法访问:对于不可达节点上的所有Pod出发API发起的逐出操作。默认情况下，节点控制器将节点标记为unknown后等待5分钟后提交第一个驱逐请求。

默认情况下，节点控制器每5秒检查一次节点状态，可以使用kube-controller-manager组件上的`--node-monitor-period`参数来配置周期


### 逐出速率限制

大部分情况下，节点控制器把逐出速率限制在每秒`--node-eviction-rate`个(默认为0.1)。这表示它每10秒钟内至多从一个节点驱逐pod。

当一个可用区域(Availability Zone)中的节点变为不健康时，节点的驱逐行为将会发生改变。节点控制器会同时检查可用区域中不健康(Ready状况为Unknown或False)的节点的百分比:

+ 如果不健康节点的比例超过`unhealthy-zone-threshold`(默认为0.55),驱逐速率会降低。
+ 如果集群较小(--large-cluster-size-threshold个节点-默认为50)
,驱逐操作将会停止
否则驱逐效率将会降为每秒 `--secondary-node-eviction-rate`个(默认为0.01)

在逐个可用区域中实施这些策略的原因是，当一个可用区域可能从控制面脱离时其他可用区域可能仍然保持连接。如果你的集群没有跨越云服务商的多个可用区域，那整个集群就只有一个可用区域。

跨多个可用区域部署你的节点的一个关键原因是当某个可用区域整体出现故障时，工作负载可以转移到健康的可用区域。因此，如果一个可用区域中的所有节点都不健康时，节点控制器将会以正常的速率`--node-eviction-rate`进行驱逐操作。在所有的可用区域都不健康(也即是集群中没有健康节点)的极端情况下，节点控制器将假设控制面与节点间的连接除了某些问题，它将停止所有驱逐动作(如果故障后部分节点重新连接，节点控制器会从剩下不健康或者不可达节点中驱逐Pod

节点控制器还负责驱逐运行在拥有NoExecute污点的节点上的Pod，除非这些Pod能够容忍此污点。节点控制器还负责根据节点故障(例如节点不可访问或者没有就绪)为其添加污点。这意味着调度器不会将Pod调度到不健康的节点上。

