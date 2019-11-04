package bridge

import (
	"encoding/json"
	"fmt"
)

type IApp interface{
	GetTasks(user int) ([]byte, error)

	GetAccount(user int)([]byte, error)
}

type App struct{
	AppContext  IApp
	CurrentUser int
}


func (s *App) GetTasks() ([]byte, error){
	return s.AppContext.GetTasks(s.CurrentUser)
}

func (s *App) GetAccount() ([]byte, error){
	return s.AppContext.GetAccount(s.CurrentUser)
}

type SaleApp struct{
}

func (s *SaleApp) GetTasks(user int) ([]byte, error){
	fmt.Printf("Current User: %v\n", user)
	fmt.Printf("Get Current Tasks for Sale App\n")
	return json.Marshal("Return a json")
}

func (s *SaleApp) GetAccount(user int) ([]byte, error){
	fmt.Printf("Current User: %v\n", user)
	fmt.Printf("Get Current  Account of Sale App\n")
	return json.Marshal("Return a json")
}

type AdminApp struct{
}

func (a *AdminApp) GetTasks(user int) ([]byte, error){
	fmt.Printf("Current User: %v\n", user)
	fmt.Printf("Get Current Tasks for Admin App\n")
	return json.Marshal("Return a json")
}

func (a *AdminApp) GetAccount(user int) ([]byte, error){
	fmt.Printf("Current User: %v\n", user)
	fmt.Printf("Get Current Account of Admin App\n")
	return json.Marshal("Return a json")
}

