package proxy

import "fmt"

// To use proxy and to object they must implement same methods
type IAccounting interface {
	DoAccounting()
}

// Accounting represents real objects which proxy will delegate data
type Accounting struct {
}

// DoAccounting implements IAccounting interface and handel's all logic
func (a *Accounting) DoAccounting() {
	fmt.Printf("Do Accounting.\n",)
}

// ProxyObject represents proxy object
type ProxyAccounting struct {
	Accounting *Accounting
}

// ObjDo are implemented IObject and intercept action before send in real Object
func (p *ProxyAccounting) DoAccounting(balance int) {

	if balance <= 0{
		fmt.Printf("Not enough balance.\n")
	} else{
		p.Accounting.DoAccounting()
	}
}



