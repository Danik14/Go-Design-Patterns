package main

import "fmt"

type Coffee interface {
	getPrice() int
}

type CoffeeBoom struct {
}

func (c *CoffeeBoom) getPrice() int {
	return 15
}

type WithSugar struct {
	coffee Coffee
}

// decorator function
func (c *WithSugar) getPrice() int {
	coffeePrice := c.coffee.getPrice()
	return coffeePrice + 2
}

type WithMilk struct {
	coffee Coffee
}

// decorator function
func (c *WithMilk) getPrice() int {
	coffee := c.coffee.getPrice()
	return coffee + 4
}

func main() {

	coffee := &CoffeeBoom{}

	//Add sugar to our coffee
	pizzaWithCheese := &WithSugar{
		coffee: coffee,
	}

	//Add milk to our coffee
	pizzaWithCheeseAndTomato := &WithMilk{
		coffee: pizzaWithCheese,
	}

	//as you can see the standart price of our coffee has increased
	//due to wrapping it with decorators
	fmt.Printf("Price of CoffeeBoom with sugar and milk is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
