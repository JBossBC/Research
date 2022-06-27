# Rest style



Rest是一种设计风格，他明确了六种结构约束，构建成了一种组织Web服务的架构，目的是为了创建具有良好扩展性的分布式系统。

rest强调在url中只使用名词，对于url的增删改查仅用http协议中的动词来实现。

1. Uniform interface
 
一个REST系统需要使用一个统一的接口来完成子系统之间以及服务与用户之间的交互。这使得REST系统中的各个子系统可以独自完成演化。

2. Client–server

客户与服务器之间通过一个统一的接口来互相通讯

3. Stateless

在一个REST系统中，服务端并不会保存有关用户的任何状态。也就是说，客户端自身负责用户状态的维持，并在每次发送请求时都需要提供足够的信息。(这可能会导致服务器处理的数据量变大，但对于分布式系统来说，每一个请求可能是不同服务器去解决，REST服务器在这里满足了分布式系统)

4. Cacheable

REST系统需要能够恰当地缓存请求，以尽量减少服务端和客户端之间的信息传输，以提高性能。

5. Layered system

在一个REST系统中，客户端并不会固定地与一个服务器打交道。

6. Code on demand (optional)


对于REST风格来说，其设计理念契合了如今分布式系统的需求。但从目前我的理解上面来看，他通过冗余了客户的状态，对于传统的Web服务，我们通常使用Session来保存客户端的相关数据，客户端下次访问服务器数据时，不用再走繁琐的验证流程，但对于REST服务，客户端每一次都要发送用户的用户名和密码，服务器也要每一次进行验证，这重新回到了http无状态的时代，很多人给出了以下解决方案来对客户端的权限进行验证
![](https://img-blog.csdnimg.cn/20201223172931903.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L1NlbmlvclNoZW4=,size_16,color_FFFFFF,t_70#pic_center)