# RTT and RTO

+ RTT

RTT(Round-Trip Time)往返时延，往返时延由三个部分决定:链路的传播时间，末端系统的处理时间，路由器的缓存中排队和处理时间，前面两个值相对不变,路由器的缓存中的排队和处理时间会随着整个网络拥塞程度
的变化而变化，所以RTT的变化在一定程度上反映了网络拥塞程度的变化.

RTT怎么获得?

第一种方式:通过tcp报文中的timestamp
RTT=当前时间-timestamp回显的时间
timestamp回显的时间代表数据包发送出去的时间

*第二种方式:重传队列中数据包的TCP控制块*


+ RTO

RTO(retransmission time out):超时重传时间，tcp能够保证数据的完整性和有效性，那么在不确定的网络环境中，tcp如何判断数据报已经丢失并且做出相应调整呢?第一种方式是:RTO和RTT共同作用的结果，RTT反应了实时状态下网络的拥塞程度，对于一台主机来说，我们可以理解为主机认为在RTT的时间内，一定能收到发送报文的ack，如果未在RTT内收到的话，那么可以在一定程度上认为数据报已经丢失，需要重发。但是RTT的波动性会造成判断不准确，一般来说RTO大于RTT,这给了动态的网络环境一些buffer time，如何根据RTT设置RTO的值，能够对网络性能产生很大的影响;第二种方式:tcp的快速重传机制，之前我们讲到拥塞控制，这不需要等待rto超时。tcp中有一个累计确认机制(当接收端收到比期望序号大的报文段时，便会重复发送最近一次确认的报文段的确认信号，我们称之为冗余ACK)，比如说服务器A一次发送五个数据包给服务器B，因为某些原因，数据包2在传输的过程中丢失了，服务器A在收到服务器B发送的有关数据包3、4、5的ack信息中，会包含期望收到的序列号信息也就是ack 2，服务器A在收到多个需要ACK2的时候，会判定数据包2已经丢失，然后对数据包2进行重传。一般来说，收到3次冗余的ack就会判定该数据包已经丢失。快速重传的机制用来防止tcp拥塞控制的时候的错误判断导致信息传输效率下降。