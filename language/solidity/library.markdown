# library

库与合约类似，库的目的是只需要在特定的地址部署一次，而他们的代码可以通过EVM的DELEGATECALL特性进行宠用。

这意味着如果库函数被调用，它的代码在调用合约的上下文中执行，即this指向调用合约，特别注意，他访问的是调用合约存储的状态。因为每个库都是一段独立的代码，所以它仅能访问调用合约明确提供的状态变量（否则它就无法通过名字访问这些变量）。

因为我们假定库是无状态的，所以如果它们不修改状态（如果它们是 view 或者 pure 函数），库函数仅能通过直接调用来使用（即不使用 DELEGATECALL 关键字）， 特别是，任何库不可能被销毁。

使用库就显示是使用基类合约方法类似。虽然它们在继承关系中不会显式可见，但调用库函数与调用显式的基类合约十分类似（可以使用 L.f() 调用）。

当然，使用内部调用约定来调用库的内部函数，这意味着所有的 internal 类型，和 保存在内存类型 都是通过引用而不是复制来传递。

EVM 为了实现这些，合约所调用的内部库函数的代码及内部调用的所有函数都在编译阶段被包含到调用合约中，然后使用一个 JUMP 指令调用来代替 DELEGATECALL。