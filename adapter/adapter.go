package adapter

import "fmt"

type ITargetFood interface{
	GetFoods()	[]string
}

type TargetFood struct {
}

func (t *TargetFood)GetFoods() []string{
	return []string{"Steak", "Ceviche", "Hamburger"}
}

type Adaptee struct {}


func (a *Adaptee)GetDrinks() []string{
	return []string{"Wine", "Lemonade", "Coca Cola"}
}

type ChefAdapter struct {
	adaptee *Adaptee
	foods *TargetFood
}

func NewAdapter() *ChefAdapter{
	return &ChefAdapter{new(Adaptee), new(TargetFood)}
}

func (c *ChefAdapter) AdaptFoodsAndDrinks(){
	drinks := c.adaptee.GetDrinks()
	foods := c.foods.GetFoods()

	for i := 0; i < len(foods) ; i++  {
		fmt.Printf("Serve %s with %s\n", foods[i], drinks[i])
	}
}


