package main

import "fmt"

type CoffeeMachine struct {
	beanAmount   float32
	grinderLevel int
	waterTemp    int
	waterAmt     float32
	milkAmount   float32
	addFoam      bool
}

func (c *CoffeeMachine) startCoffee(beanAmount float32, grind int) {
	c.beanAmount = beanAmount
	c.grinderLevel = grind
	fmt.Println("Starting the coffee order with beans:", beanAmount, "and grind level", grind)
}

func (c *CoffeeMachine) grindBeans() bool {
	fmt.Println("Grinding the beans:", c.beanAmount, "beans at", c.grinderLevel, "level")
	return true
}

func (c *CoffeeMachine) useMilk(amount float32) float32 {
	fmt.Println("Adding milk:", amount, "oz")
	c.milkAmount = amount
	return amount
}

func (c *CoffeeMachine) setWaterTemp(temp int) {
	fmt.Println("Setting water temp:", temp)
}

func (c *CoffeeMachine) useHotWater(amount float32) float32 {
	fmt.Println("Adding hot water:", amount)
	c.waterAmt = amount
	return amount
}

func (c *CoffeeMachine) endCoffee() {
	fmt.Println("Ending the coffee order")
}

func (c *CoffeeMachine) doFoam(foam bool) {
	fmt.Println("Foam settings:", foam)
	c.addFoam = foam
}
