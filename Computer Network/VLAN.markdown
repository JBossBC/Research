# VLAN(virtual local AreaNetwork)

在计算机网络中，传统的交换机虽然能隔离冲突域，提高每一个端口的性能，但并不能隔离广播域，当网络中的机器足够多时会引发广播风暴。同时，不同部门、不同组织的机器连在同一个二层网络中也会造成安全问题

因为，在交换机中划分子网、隔离广播域的思路便形成了VLAN的概念，VLAN按照功能、部门等因素将网络中的机器进行划分，使之分属于不同的部分，每一个部分形成一个虚拟的局域网络，共享一个单独的广播域。这样就可以把一个大型交换网络划分为许多个独立的广播域，即VLAN。

## Question



## 如何区分不同VLAN的流量呢？

VLAN技术将一个二层网络中的机器隔离开来，那么如何区分不同VLAN的流量呢？IEEE 802.1q协议规定了VLAN的实现方法，即在传统的以太网帧中再添加一个VLAN tag字段,用于标识不同的VLAN。这样，**支持VLAN的交换机在转发帧时，不仅会关注MAV地址，还会考虑VLAN tag字段**。VLAN tag中包含了TPID、PDP、CFI、VID，其中VID(VLAN ID)部分用来具体指出帧属于哪个VLAN的。VID占12位，所以其取值范围为0到4095

![](https://img-blog.csdnimg.cn/img_convert/0a75b419995199664eabd3dd8432244c.png)

## 过程分析


![](https://img-blog.csdnimg.cn/1f9852ddc5754ea3903030dec045f349.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBATXl5U29waGlh,size_20,color_FFFFFF,t_70,g_se,x_16)


### 什么是交换机的access端口和trunk端口

图中,Port1、Port2、Port5、Port6、Port7、Port8为access端口，**每一个access端口都会分配一个VLAN ID，标识它所连接的设备属于哪一个VLAN**。当数据帧从外界通过access端口进入交换机时，数据帧原本是不带tag的，access端口给数据帧打上tag(VLAN ID即为access端口所分配的VLAN ID);当数据帧从交换机内部通过access端口发送时，数据帧的VLAN ID必须和access端口的VLAN ID一致，access端口才接收此帧，接着access端口将帧的tag信息去掉，再发送出去。

Port3、Port4为trunk端口，**trunk端口不属于某个特定的VLAN，而是交换机与交换机之间多个VLAN的通道。trunk端口声明了一组VLAN ID,表明只允许带有这些VLAN  ID的数据帧通过，从trunk端口进去和出去的数据帧都是带tag的**PC1和PC3属于VLAN100, PC2和PC4属于VLAN200，所以PC1和PC3处在同一个二层网络中，PC2和PC4处在同一个二层网络中。**尽管PC1和PC2连接在同一台交换机中，但它们之间的通信是需要经过路由器的。**

### VLAN tag如何发挥作用的？

当PC1向PC3发送数据时，PC1将IP包封装在以太帧中，帧的目的MAC地址为PC3的地址，此时帧并没有tag信息。当帧到达port1时，port1给帧打上tag(VID=100)，帧进入switch1，然后帧通过port3、port4到达switch2(port3、port4允许VLAN ID为100、200的帧通过)。在switch2中，port5所标记的VID和帧相同，MAC地址也相匹配，帧就发送到Port5上，Port5将帧的tag信息去掉，然后发给PC3.由于PC2、PC4和PC1的VLAN不同，因此收不到PC1发送的帧

本质上VLAN相关字段的添加能让交换机再次进行分发对象的筛选，原来的交换机对所有的ARP广播请求进行广播，但是现在如果ARP广播请求中VLAN ID字段不与交换机所有端口的VLAN字段其中一个匹配，那么交换机就不会接受该请求。