# Factory Method


## 思想

工厂方法模式是一种创建型设计模式，其在父类中提供一个创建对象的方法，允许子类决定实例化对象的类型

## 问题

假如你正在开发一款物流管理应用。最终版本只能处理卡车运输，因此大部分代码都位于名为卡车的类中。一段时间后，应第三方需要，你需要扩展支持海上物流功能。

但对于目前状况来说，大部分代码都与卡车相关。在程序中添加轮船类可能需要修改全部代码。更糟糕的是，如果以后需要在程序中支持另一种运输方式，很可能需要再次对这些代码进行大幅度修改。


## 解决方案

工厂方法模式建议使用特殊的工厂方法代替对于对象构造函数的直接调用(即使用new运算符)。对象在工厂内仍然通过new运算符创建。工厂方法返回的对象通常被称作"产品".


## go代码实现

    
    type Transport interface {
    	method()
    	getName() string
    }
    type bus struct {
    	name string
    }
    
    func (Bus *bus) getName() string {
    	return Bus.name
    }
    
    func (*bus) method() {
    	fmt.Println("this is bus")
    }
    
    type car struct {
    	name string
    }
    
    func (Car *car) getName() string {
    	return Car.name
    }
    
    func (*car) method() {
    	fmt.Println("this is car")
    }
    
    func FactoryMode(name string) Transport {
    	if name == "car" {
    		return &car{name: "劳斯莱斯"}
    	}
    	if name == "bus" {
    		return &bus{name: "公交车"}
    	}
    	return nil
    }
    


## 总结    