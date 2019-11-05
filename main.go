package main

import (
	"Golang-Design-Patterns/adapter"
	"Golang-Design-Patterns/bridge"
	"Golang-Design-Patterns/builder"
	chain_of_responsibility "Golang-Design-Patterns/chain-of-responsibility"
	"Golang-Design-Patterns/command"
	command_factory "Golang-Design-Patterns/command-factory"
	"Golang-Design-Patterns/decorator"
	"Golang-Design-Patterns/facade"
	"Golang-Design-Patterns/factory"
	"Golang-Design-Patterns/flyweight"
	"Golang-Design-Patterns/observer"
	"Golang-Design-Patterns/prototype"
	"Golang-Design-Patterns/proxy"
	"Golang-Design-Patterns/singleton"
	"Golang-Design-Patterns/state"
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


	fmt.Printf("Command Pattern:\n")
	notification := &command.Notification{}
	job1 := command.NewJob1Command(notification)
	job2 := command.NewJob2Command(notification)

	jobs := command.NewJobs(job1, job2)
	jobs.RunJob1()
	jobs.RunJob2()

	fmt.Printf("---------------------------\n")


	fmt.Printf("Command-Factory Pattern:\n")
	factory :=  command_factory.NewCommandFactory()
	commandFunc := factory.GetCommand("SaleEmailCommandType1")
	commandFunc()

	fmt.Printf("---------------------------\n")

	fmt.Printf("Proxy Pattern:\n")
	proxy := proxy.ProxyAccounting{}
	proxy.DoAccounting(100)

	proxy.DoAccounting(0)

	fmt.Printf("---------------------------\n")


	fmt.Printf("Decorator Pattern:\n")
	allDecorators := []decorator.URLDecorator{ decorator.RemoveHTTP, decorator.RemoveHTTPS, decorator.RemoveWWW}
	decoratedURL := decorator.URLDecorate("https://www.mywebsite.com", allDecorators...)
	fmt.Println(decoratedURL)

	fmt.Printf("---------------------------\n")


	fmt.Printf("Observer Pattern:\n")

	subject := observer.NewSubject()
	concreteObserverA := observer.NewConcreteNotifierA()
	concreteObserverA1 := observer.NewConcreteNotifierA()
	concreteObserverB := observer.NewConcreteNotifierB()
	concreteObserverB1 := observer.NewConcreteNotifierB()

	subject.Attach(concreteObserverA)
	subject.Attach(concreteObserverA1)
	subject.Attach(concreteObserverB)
	subject.Attach(concreteObserverB1)

	fmt.Printf("First Scan\n")
	subject.ScanData()

	subject.DeAttach(concreteObserverB)
	subject.DeAttach(concreteObserverA)

	fmt.Printf("Second Scan\n")
	subject.ScanData()

	fmt.Printf("---------------------------\n")


	fmt.Printf("Chain of Responsibility Pattern:\n")
	headEditor := chain_of_responsibility.HeadEditor{}
	headAccounter := chain_of_responsibility.HeadAccounter{}
	headManager := chain_of_responsibility.HeadManager{}

	request := chain_of_responsibility.CustomerBalanceRequest{
		CustomerName: "John",
		Balance:      500,
	}

	//create the chain here
	headEditor.Next = &headAccounter
	headAccounter.Next = &headManager

	headEditor.Handle(request)

	request.Balance = 3000
	request.CustomerName = "Jared"

	headEditor.Handle(request)

	request.Balance = 10000
	request.CustomerName = "Maz"

	headEditor.Handle(request)

	fmt.Printf("---------------------------\n")

	fmt.Printf("Prototype Pattern:\n")
	customer := prototype.NewCustomer("John", "john@gmail.com")
	fmt.Printf("First time, Customer name: %v, Email: %v\n", customer.Name, customer.Email)

	customer2 := customer.WithName("Jared")
	fmt.Printf("Cloned Customer, Customer name: %v, Email: %v\n", customer2.Name, customer2.Email)

	fmt.Printf("Second time, Customer name: %v, Email: %v\n", customer.Name, customer.Email)

	fmt.Printf("---------------------------\n")

	fmt.Printf("Flyweight Pattern:\n")
	shapeFactory := flyweight.NewShapeFactory()
	shape1 := shapeFactory.GetShapeToDisplay(flyweight.CircleType)
	shape1.DrawShape()

	shape2 := shapeFactory.GetShapeToDisplay(flyweight.RectangleType)
	shape2.DrawShape()

	shape3 := shapeFactory.GetShapeToDisplay(flyweight.CircleType)
	shape3.DrawShape()

	shape4 := shapeFactory.GetShapeToDisplay(flyweight.RectangleType)
	shape4.DrawShape()

	fmt.Printf("Number of created shapes: %v\n", shapeFactory.ObjectCount)

	fmt.Printf("---------------------------\n")

	fmt.Printf("Bridge Pattern:\n")

	app := bridge.App{
		CurrentUser: 20,
	}
	saleApp := bridge.SaleApp{}
	app.AppContext = &saleApp

	app.GetTasks()

	adminApp := bridge.AdminApp{}

	app.CurrentUser = 30
	app.AppContext = &adminApp

	app.GetAccount()

	fmt.Printf("---------------------------\n")

	fmt.Printf("State Pattern:\n")

	newLead := state.NewLead(&state.ConcreteState1{})
	newLead.GetStateName()
	newLead.GoToNextState()

	newLead.GetStateName()

	newLead.GoToPreviousState()

	newLead.GetStateName()

	newLead.GoToNextState()

	newLead.GetStateName()

	newLead.GoToNextState()

	fmt.Printf("---------------------------\n")
}

