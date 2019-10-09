package main

import(
	"Golang-Design-Patterns/adapter"
	"Golang-Design-Patterns/facade"
	"Golang-Design-Patterns/factory"
	"Golang-Design-Patterns/builder"
	"Golang-Design-Patterns/singleton"
	s "Golang-Design-Patterns/strategy"
	"Golang-Design-Patterns/template"
	"fmt"
)

func main(){

	fmt.Printf("Strategy Pattern:\n")
	ctx := s.NewCampaignContext( "Karl Marx", &s.CampaignType1{})
	ctx.CreateCampaign()

	ctx = s.NewCampaignContext( "Hannah Arendt", &s.CampaignType2{})
	ctx.CreateCampaign()

	fmt.Printf("---------------------------\n")

	fmt.Printf("Factory+Strategy Pattern:\n")
	factory.Factory(factory.StringType1, "Capital").CreateCampaign()
	factory.Factory(factory.StringType2, "The Origins of Totalitarianism").CreateCampaign()

	fmt.Printf("---------------------------\n")

	fmt.Printf("Template Pattern:\n")
	c := template.StripeCustomer()
	c.Create()

	c = template.PaypalCustomer()
	c.Create()

	fmt.Printf("---------------------------\n")

	fmt.Printf("Facade Pattern:\n")
	f := facade.NewFacade()
	f.CompleteProcess("Microsoft", "Microsoft Azul", "10 tips abouts Microsoft Azul")

	fmt.Printf("---------------------------\n")

	fmt.Printf("Adapter Pattern:\n")
	a := adapter.NewAdapter()
	a.AdaptFoodsAndDrinks()

	fmt.Printf("---------------------------\n")
	fmt.Printf("Singleton Pattern:\n")
	config1, _ := singleton.ConfigRepository().Get("Config 1")
	config2, _ := singleton.ConfigRepository().Get("Config 2")
	fmt.Printf("Config 1: %s \n", config1)
	fmt.Printf("Config 2: %s \n", config2)

	fmt.Printf("---------------------------\n")
	fmt.Printf("Builder Pattern:\n")

	chef := builder.ChefDirector{}

	_steakBuilder := builder.SteakBuilder{}
	food := chef.MixMaterials(_steakBuilder)
	food.Cook()

	_pastaBuilder := builder.PastaBuilder{}
	food = chef.MixMaterials(_pastaBuilder)
	food.Cook()


	fmt.Printf("---------------------------\n")
}

