# TCP中的拥塞控制和滑动窗口

####Question
tcp中checksum来防止是否存在单byte的错误，但是在数据链路层中CRC校验不是已经保证了数据无比特错误，为什么还要多此一举呢？
滑动窗口？

## 对于TCP的一些理解
TCP的设计保证了数据传输的完整性和可靠性，怎么样保证数据的完整性和可靠性呢?在数据传输的过程中，由于网络的不稳定性，tcp必须保证数据传输过程失误的情况下能够及时弥补，来确保数据的可靠传输。tcp的机制是，每发送一个数据报，都需要得到对应的数据报ack信息，来确保服务器收到信息。这个和我们接下来要讲的拥塞控制和滑动窗口有什么关系呢？我把这两个设计理解为尽可能的防止传输中的错误，对于滑动窗口来说，他的设计保证了两端都能够有足够的缓冲区来接受数据，避免了因为接受的数据过多而导致数据的丢失这种情况，对于拥塞控制而言，首先我们要在保证传输效率的同时保证传输的有效性，我们在发送数据的时候需要考虑每一次发送数据量的多少，是发送一个请求，还是多个请求，这是我们值得考量的点，当我们需要将信道带宽的利用率充分发挥的时候，不可避免的是一次性发送多个报文，但因为网络的复杂性，信道可用带宽会随时发生变化，我们怎么样才能够尽可能地判断当前信道带宽支持我们发送地数据量，这就是拥塞控制的设计理由。拥塞控制根据PTT(往返时延)来动态修改拥塞窗口的大小，做到尽可能地利用信道带宽。

+ 滑动窗口

滑动窗口的设计是为了确保接收端能够有足够地缓冲区来接受数据，避免数据地丢失。滑动窗口只有在发送端接受到ack的确认信息的时候，才会移动滑动窗口

+ 拥塞控制
 
  + cwnd会根据mss的大小进行init
  > linux 3.0以前，内核默认的initcwnd比较小，MSS为1460时，初始的拥塞控制窗口为3。
 linux3.0以后，采取了Google的建议，把初始拥塞控制窗口调到了10。


ssthresh：慢开始阈值

   1、cwnd < ssthresh，使用慢开始算法

   2、cwnd > ssthresh，停止慢开始，改用拥塞控制算法

   3、cwnd = ssthresh，慢开始与拥塞控制都可以使用，二者取其一


关于拥塞控制的四种算法:

慢开始:呈现指数级增长，刚开始向网络中注入的报文数量很少，但是增加速度很快，服务器每一次发送cwnd窗口大小的数据包到网络中，当收到多少数量的ack确认信息，则cwnd窗口大小就会对应增加多少。

拥塞避免:当cwnd的大小到一定程度了，如果再指数级增加，会导致突然的网络崩溃，当cwnd达到ssthresh的时候，我们判定现在网络承载的带宽已经被充分利用，这时候，无论一次接受到多少ack确认信息，cwnd只会增加1，当出现有数据包RTO超时的时候，会判定网络堵塞，这个时候，ssthresh会调整为出现拥塞的时候cwnd的一半,然后将cwnd重新赋值为1,并重新执行慢开始算法。


这两个算法的不足之处在于，如果数据包在网络传输过程中丢失，不是网络堵塞的原因，造成了RTO超时，将cwnd的值赋为1，严重损害了数据传输的效率，所以才有了快速重传机制以及快速恢复机制。

快速重传机制:建立在tcp的累计确认机制之上，当数据包丢失的时候，之后发送成功的数据包会传给发送端冗余的ack的信息，表示还未收到对应的数据包，当收到冗余ack信息三次的时候,会判定数据包丢失触发重传机制。
![](https://img-blog.csdnimg.cn/20200514225454520.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzM3NDM3OTgz,size_16,color_FFFFFF,t_70)

快速恢复机制:发送端连续接收端三个确认报文后可以确认只是丢失了个别报文（确认报文可以传回，网络没有拥塞），此时不需要启动慢开始算法。此时将ssthresh和cwnd赋值为此时的cwnd的一半，然后执行拥塞避免算法。



[TCP拥塞控制简析]( https://blog.csdn.net/qq_37437983/article/details/106130645?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522165552466216780357291802%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=165552466216780357291802&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~rank_v31_ecpm-4-106130645-null-null.142^v17^pc_search_result_control_group,157^v15^new_3&utm_term=cwnd%E5%92%8Cssthresh%E6%98%AF%E4%BB%80%E4%B9%88&spm=1018.2226.3001.4187 )


[TCP的快速重传机制]( https://blog.csdn.net/whgtheone/article/details/80983882?ops_request_misc=&request_id=&biz_id=102&utm_term=tcp%E6%80%8E%E4%B9%88%E5%86%B3%E5%AE%9A%E9%87%8D%E4%BC%A0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~sobaiduweb~default-1-80983882.142^v17^pc_search_result_control_group,157^v15^new_3&spm=1018.2226.3001.4187 )





[拥塞控制相关文章]( https://en.wikipedia.org/wiki/TCP_congestion_control#:~:text=In%20TCP%2C%20the%20congestion%20window,overloaded%20with%20too%20much%20traffic.](https://en.wikipedia.org/wiki/TCP_congestion_control#:~:text=In%20TCP%2C%20the%20congestion%20window,overloaded%20with%20too%20much%20traffic. )

  