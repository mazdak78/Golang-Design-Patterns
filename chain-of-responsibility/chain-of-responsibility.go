package chain_of_responsibility

import (
	"fmt"
	"reflect"
)

//Request - will be handled by different handlers
type CustomerBalanceRequest struct{
	CustomerName string
	Balance int
}

type BalanceRequest interface {
	Handle(request CustomerBalanceRequest)
}

type HeadEditor struct{
	Next BalanceRequest
}

func (h *HeadEditor) Handle(b CustomerBalanceRequest){
	if b.Balance < 1000 {
		fmt.Printf("%v approved balance for %v request. Balance: %v\n", reflect.TypeOf(h).Elem().Name(), b.CustomerName, b.Balance)
	} else{
		h.Next.Handle(b)
	}
}

type HeadAccounter struct{
	Next BalanceRequest
}

func (h *HeadAccounter) Handle(b CustomerBalanceRequest){
	if b.Balance <= 5000 {
		fmt.Printf("%v approved balance for %v request. Balance: %v\n", reflect.TypeOf(h).Elem().Name(), b.CustomerName, b.Balance)
	} else{
		h.Next.Handle(b)
	}
}

type HeadManager struct{
	Next BalanceRequest
}

func (h *HeadManager) Handle(b CustomerBalanceRequest){
	if b.Balance > 5000 {
		fmt.Printf("%v approved balance for %v request. Balance: %v\n", h, b.CustomerName, b.Balance)
	} else{
		h.Next.Handle(b)
	}
}

// this is just for test and practicing, it has nothing to do with Cor pattern
func (h *HeadManager) String() string {
	return "HeadManager"
}


