# Context

在GO程序中可能同时存在许多协程，这些协程被动态地创建和销毁。例如，在典型的http服务器中，每个新建立的连接都可能新建一个协程。当请求完成后，协程也随之被销毁。但是，**请求可能临时终止也可能超时，这个时候我们希望安全并及时地停止协程**，而不必一直占用系统的资源。因此，需要一种能够优雅控制协程退出的手段，

## Question

context与协程的关系?
父进程结束后子进程也会结束,但是

## 为什么需要context

有一句关于go的名言--如果你不知道如何退出一个协程，那么就不要创建这个协程。在context之前，管理协程退出需要协助通道close的机制，该机制会唤醒所有监听该通道的协程，并触发响应的退出逻辑。


## context的使用方式

context是使用频率非常高的包,context一般作为接口的第一个参数传递超时信息。在Go源码中，net/http,net,sql包的使用如下

```
 
//net/http

func(r *Request)WithContext(ctx context.Context)*Reqeust

//sql

func(db *DB)BeginTx(ctx context.Context,opts *TxOptions)(*Tx,error)

//net

func (d *Dialer)DialContext(ctx context.Context,network,address string)(Conn,error) 

```

context.Context其实是一个接口

```
type Context interface{

  Deadline()(deadline time.Time,ok,bool)

  Done() <-chan structp{}

  Err() error

  Value(key interface{})interface{}
}

```

Deadline 方法的第一个返回值表示还有多久到期，第二个返回值表示是否到期。Done是使用最频繁的方法,其返回一个通道，一般的做法是监听该通道的信号，如果收到信号则表示通道已经关闭，需要执行退出。如果通道已经关闭,则Err()返回退出的原因。Value方法返回指定key对应的Value,这是context携带的值。

context中携带值是非常少见的,其一般在跨程序的API中使用，并且该值的作用域在结束时终结。key必须时访问安全的，因为可能有多个协程同时访问它。一种常见的策略是在context中存储授权相应的值,这些鉴权不会影响到程序的核心逻辑。

Value主要用于安全凭证、分布式跟踪ID、操作优先级、退出信号与到期时间等场景。尽管如此，在使用value方法时也需要谨慎，如果参数与函数核心处理逻辑有关，那么仍然建议显示地传递参数。

### context退出与传递

context是一个接口，这意味着有相应地实现。用户可以按照接口中定义地方法，严格实现其语义。一般用的最多的还是go标准库地简单实现。调用context.Background函数或context.TODO函数会返回最简单地context实现。context.Background函数一般作为根对象存在，其不可以退出，也不可以携带值。要具体地使用context的功能，需要派生出新的context,配套的使用函数如下,其中前三个函数用于处理退出。

```
func WithCancel(parent Context)(ctx Context,cancel CancelFunc)

func WithTimeout(parent Context,timeout time.Duration)(Context,CancelFunc)

func WithDeadline(parent Context,d time.Time)(Context,CancelFunc)

func WithValue(parent Context,key,value interface{}) Context

```


WithCancel函数返回一个子context并且有cancel退出方法。子context在两种情况下会退出，一种是调用cancel,另一种情况是当参数中的父context退出时,该context及其关联的子context都会退出。

WithTimeout函数指定超时时间，当超时发生时，子context将退出。因此子context的退出有3种时机，一种是父context退出;一种是超时退出;一种是主动调用cancel函数退出。WithDeadline和WithTimeout函数的处理方法相似，不过其参数指定的是最后到期的时间。WithValue函数返回带key-value的子context。
