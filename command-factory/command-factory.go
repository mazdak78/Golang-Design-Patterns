package command_factory

import (
	"fmt"
)

type CommandFactory struct{
	Commands map[string]func()
}

//Informer - Receiver
type Informer struct{}

// SendSaleEmailType1 - Concrete command
func (l *Informer) SendSaleEmailType1() {
	fmt.Println("Inform customers type 1")
}

//SaleEmailCommand - ConcreteCommand
type SaleEmailCommand struct {
	Informer *Informer
}

func NewCommandFactory() CommandFactory{
	_factory := CommandFactory{}
	seCommand := SaleEmailCommand{}
	_factory.Commands = map[string]func(){ "SaleEmailCommandType1" : seCommand.Informer.SendSaleEmailType1}
	return _factory
}

func (cf *CommandFactory) GetCommand(commandName  string) func(){
	return cf.Commands[commandName]
}

