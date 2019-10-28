package observer

import (
	"container/list"
	"fmt"
	"math/rand"
)

//interface that receive updates
type Observer interface{
	Update(s *Subject)
}

// interface that manage notifying
type Subject struct{
	observers *list.List
	status int
}

func NewSubject() *Subject {
	//this can be categorized from database etc
	return &Subject{observers:new(list.List)}
}

func (s *Subject) Attach(o Observer){
	s.observers.PushBack(o)
}

func (s *Subject) DeAttach(o Observer){
	for obs := s.observers.Front(); obs != nil; obs = obs.Next() {
		if obs.Value.(Observer) == o {
			s.observers.Remove(obs)
		}
	}
}

func (s *Subject) notify(){
	for obs := s.observers.Front(); obs != nil; obs = obs.Next() {
		observer := obs.Value.(Observer)
		observer.Update(s)
	}
}

// some business logic function to notify
func (s *Subject) ScanData() {

	// just generate random number for status
	// this can be a business logic in real world to calculate or fetch status for a specific source or logic
	s.status = rand.Intn(1000)
	s.notify()
}

//Concrete class A for Observer
type ConcreteNotifierA struct{
}

func NewConcreteNotifierA() *ConcreteNotifierA {
	return &ConcreteNotifierA{
	}
}

func (a *ConcreteNotifierA) Update(s *Subject) {
	if s.status > 500 {
		fmt.Printf("Status is bigger than 500: %v \n", s.status)
	}
}


//Concrete class B for Observer
type ConcreteNotifierB struct{
}

func NewConcreteNotifierB() *ConcreteNotifierB {
	return &ConcreteNotifierB{
	}
}

func (b *ConcreteNotifierB) Update(s *Subject) {
	if s.status < 501 {
		fmt.Printf("Status is less than 500: %v \n", s.status)
	}
}

