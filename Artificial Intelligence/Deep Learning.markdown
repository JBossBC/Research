# Deep Learning



深度学习:一般是指通过训练多层网络结构对未知数据进行分类或回归

深度学习分类:

+ 有监督学习方法--深度前馈网络、卷积神经网络、循环神经网络等
+ 无监督学习方法--深度信念网、深度玻尔兹曼机、深度自编码器等


深度学习的思想:深度神经网络的基本思想是通过构建多层网络，对目标进行多层表示，以期待通过多层的高层次的特征来表示数据的抽象语义信息，获得更好的特征鲁棒性。



## 神经网络

1. 人工神经网络(简写为ANNS)是一种模仿动物神经网络行为特征，进行分布式并行信息处理的算法数学模型。这种网络依靠系统的复杂程度，通过调整内部大量节点之间相互连接的关系，从而达到信息处理的目的，并具有自学习和自适应的能力。神经网络类型众多，其中最为重要的是多层感知器。为了详细地描述神经网络，我们先从简单的神经网络说起。

### 感知机

感知机是1957年，由Rosenblatt提出会，是神经网络和支持向量机的基础。

感知机是有生物学上的一个启发，他的参照对象和理论依据可以参照下图：（我们的大脑可以认为是一个神经网络，是一个生物的神经网络，在这个生物的神经网络里边呢，他的最小单元我们可以认为是一个神经元，一个neuron，这些很多个神经元连接起来形成一个错综复杂的网络，我们把它称之为神经网络。当然我们现在所说的，在深度学习包括机器学习指的神经网络Neural Networks实际上指的是人工神经网络Artificial Neural Networks，简写为ANNs。我们只是简化了。我们人的神经网络是由这样一些神经元来构成的，那么这个神经元他的一些工作机制呢就是通过这样一个下面图的结构，首先接收到一些信号，这些信号通过这些树突(dendrite)组织，树突组织接收到这些信号送到细胞里边的细胞核(nucleus)，这些细胞核对接收到的这些信号，这些信号是以什么形式存在的呢？这些信号比如说眼睛接收到的光学啊，或者耳朵接收到的声音信号，到树突的时候会产生一些微弱的生物电，那么就形成这样的一些刺激，那么在细胞核里边对这些收集到的接收到的刺激进行综合的处理，当他的信号达到了一定的阈值之后，那么他就会被激活，就会产生一个刺激的输出，那么就会形成一个我们大脑接收到的进一步的信号，那么他是通过轴突这样的输出计算的，这就是我们人脑的一个神经元进行感知的时候大致的一个工作原理。）


## 前向传播和反向传播介绍

神经网络的计算主要有两种：前向传播（foward propagation, FP）作用于每一层的输入，通过逐层计算得到输出结果；反向传播（backward propagation, BP）作用于网络的输出，通过计算梯度由深到浅更新网络参数。

### 前向传播


假设上一层结点i , j , k , . . . i,j,k,...i,j,k,...等一些结点与本层的结点w ww有连接，那么结点w ww的值怎么算呢？就是通过上一层的i , j , k , . . . i,j,k,...i,j,k,...等结点以及对应的连接权值进行加权和运算，最终结果再加上一个偏置项（图中为了简单省略了），最后在通过一个非线性函数（即激活函数），如R e L u ReLuReLu，s i g m o i d sigmoidsigmoid等函数，最后得到的结果就是本层结点w ww的输出。

最终不断的通过这种方法一层层的运算，得到输出层结果。


### 反向传播


由于我们前向传播最终得到的结果，以分类为例，最终总是有误差的，那么怎么减少误差呢，当前应用广泛的一个算法就是梯度下降算法，但是求梯度就要求偏导数，下面以图中字母为例讲解一下：
                                                                                                                                                                                                              