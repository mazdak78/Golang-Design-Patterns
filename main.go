package main

import(
	"Golang-Design-Patterns/factory"
	s "Golang-Design-Patterns/strategy"
	"fmt"
)

func main(){

	fmt.Printf("Strategy Pattern\n")
	ctx := s.NewCampaignContext( "Karl Marx", &s.CampaignType1{})
	ctx.CreateCampaign()

	ctx = s.NewCampaignContext( "Hannah Arendt", &s.CampaignType2{})
	ctx.CreateCampaign()

	fmt.Printf("---------------------------\n")

	fmt.Printf("Factory+Strategy Pattern\n")
	factory.Factory(factory.StringType1, "Capital").CreateCampaign()
	factory.Factory(factory.StringType2, "The Origins of Totalitarianism").CreateCampaign()

	fmt.Printf("---------------------------\n")

}

