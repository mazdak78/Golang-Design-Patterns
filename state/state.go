package state

import "fmt"

type ILeadState interface{
	Next(l *Lead)
	Prev(l *Lead)
	String(l *Lead) string
}


// Lead is our Context
type Lead struct{
	leadState ILeadState
}

func NewLead(l ILeadState) *Lead{
	return &Lead{leadState: l}
}

func (l *Lead) GoToPreviousState(){
	l.leadState.Prev(l)
}

func (l *Lead) GoToNextState(){
	l.leadState.Next(l)
}

func (l *Lead) GetStateName() string{
	return l.leadState.String(l)
}

// Concrete States 1:
type ConcreteState1 struct{

}

func (c1 *ConcreteState1) Next(l *Lead){
	fmt.Print("Check conditions. Going to ConcreteState2.\n")
	l.leadState = &ConcreteState2{}

}

func (c1 *ConcreteState1) Prev(l *Lead){
	fmt.Print("The Lead is in the root state. There is no prev state.\n")
}

func (c1 *ConcreteState1) String(l *Lead) string{
	fmt.Print("ConcreteState1\n")
	return "ConcreteState1"
}

// Concrete States 2:
type ConcreteState2 struct{

}

func (c2 *ConcreteState2) Next(l *Lead){
	fmt.Print("There is no next State.\n")
}

func (c2 *ConcreteState2) Prev(l *Lead){
	fmt.Print("Check conditions. Going to ConcreteState1.\n")
	l.leadState = &ConcreteState1{}

}

func (c2 *ConcreteState2) String(l *Lead) string{
	fmt.Print("ConcreteState2\n")
	return "ConcreteState2"
}
