# Golang reflect

在计算机科学中，反射是程序在运行时检查、修改自身结构和行为的能力。最早的计算机以原始的汇编语言进行编程，**汇编语言具有固有的反射性，因为它可以通过将指令定义为数据并修改这些指令数据对原始体系结构进行修改。**但随后出现的高级语言(Algol、Pascal、C)导致反射的能力在很大程度上消失了。


反射为Go语言提供了复杂、意想不到的处理能力以及灵活性。例如，我们没有办法在运行时获取结构体变量内部的方法名和属性名。对于函数或方法，我们没有办法动态地检查参数的个数和返回值的个数，更不能在运行时通过函数名动态调用函数，这些都可以由反射做到。

反射在Go程序中使用的不会特别多，一般会作为框架或者是基础服务的一部分(使用json标准库序列化时就用到了反射)

## 反射的基本使用方法

### 反射的两种类型

`func ValueOf(i interface{})Value`

`func TypeOf(i interface{})Value`

这两个函数的参数都是空接口interface{},内部存储了即将被反射的变量。因此，反射与接口之间存在很强的联系。

可以将reflect.Value看作反射的值,reflect.Type看作反射的实际类型。其中，reflect.Type是一个接口，包含和类型有关的许多方法签名。

reflect.Value是一个结构体，其内部包含了很多方法。可以简单地使用fmt打印reflect.TypeOf与reflect.ValueOf函数生成的结果。reflect.ValueOf将打印出反射内部的值,reflect.TypeOf会打印出反射的类型。

reflect.Value类型中的Type方法可以获取当前反射的类型

**因此reflect.Value可以转换为reflect.Type(加了包名)。reflect.Value与reflect.Type都具有Kind方法，可以获取标识类型的Kind,其底层是uint。Go语言的内置类型都可以用唯一的整数进行标识**

通过Kind类型可以方便地验证反射地类型是否相等


### 反射转化为接口

reflect.Value中的Interface方法以空接口的形式返回reflect.Value中的值。如果要进一步获取空接口的真实值，可以通过接口的断言语法对接口进行转换
