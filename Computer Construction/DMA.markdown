# DMA(Direct Memory Access)

DMA传输将数据从一个地址空间复制到另一个地址空间，提供在外设和存储器之间或者存储器之间的高速数据传输。

DMA作用就是解决大量数据转移过度消耗CPU资源的问题。有了DMA使CPU更专注于更加实用的操作-计算、控制等。


DMA的作用就是实现数据的直接传输，而去掉了传统数据传输需要CPU寄存器参与的环节，主要涉及四种情况的数据传输，但本质上是一样的，都是从内存的某一区域传输到内存的另一区域（外设的数据寄存器本质上就是内存的一个存储单元）。四种情况的数据传输如下：

    外设到内存
    内存到外设
    内存到内存
    外设到外设

DMA的主要特征

每个通道都直接连接专用的硬件DMA请求，每个通道都同样支持软件触发。这些功能通过软件来配置；

    在同一个DMA模块上，多个请求间的优先权可以通过软件编程设置（共有四级：很高、高、中等和低），优先权设置相等时由硬件决定（请求0优先于请求1，依此类推）；
    独立数据源和目标数据区的传输宽度（字节、半字、全字），模拟打包和拆包的过程。源和目标地址必须按数据传输宽度对齐；
    支持循环的缓冲器管理；
    每个通道都有3个事件标志（DMA半传输、DMA传输完成和DMA传输出错），这3个事件标志逻辑或成为一个单独的中断请求；
    存储器和存储器间的传输、外设和存储器、存储器和外设之间的传输；
    闪存、SRAM、外设的SRAM、APB1、APB2和AHB外设均可作为访问的源和目标；
    可编程的数据传输数目：最大为65535。







