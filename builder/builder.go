package builder

import "fmt"

type foodBuilder interface {
	GetFood() *Food
}

type ChefDirector struct {
}

type Food struct{
	Name string
	TimeToCook int
	Materials string
}

func (f *Food) Cook(){
	fmt.Printf("%s in %v minutes will turn into %s \n", f.Materials, f.TimeToCook,f.Name)
}

//MixMaterials
func (d *ChefDirector) MixMaterials(builder foodBuilder) *Food{
	return builder.GetFood()
}

type SteakBuilder struct {
}

func (b SteakBuilder) GetFood() *Food{
	return &Food{Name : "Steak", TimeToCook: 20, Materials: "Meat, Onion, Mushroom"}
}

type PastaBuilder struct {
}

func (b PastaBuilder) GetFood() *Food{
	return &Food{Name : "Pasta", TimeToCook: 30, Materials: "Tomato, Pasta, Meat, Onion"}
}


