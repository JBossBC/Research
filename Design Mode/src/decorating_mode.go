package main

import (
	"fmt"
)

// decorating moded 能够无限套娃，这个基类的成员变量的作用非常关键***
func main() {
	pizza := &VeggieMania{}
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}
	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}

//零件接口
type IPizza interface {
	getPrice() int
}

// 具体零件
type VeggieMania struct {
}

func (p *VeggieMania) getPrice() int {
	return 15
}

// 具体装饰
type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}
