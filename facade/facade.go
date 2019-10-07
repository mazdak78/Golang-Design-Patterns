package facade

import "fmt"

type Customer struct{}
type Campaign struct {}
type Post struct {}

func (*Customer) CreateCustomer(name string){
	fmt.Printf("Customer created: %s \n", name)
}

func (*Campaign) CreateCampaign(name string){
	fmt.Printf("Campaign created: %s \n", name)
}

func (*Post) CreatePost(title string){
	fmt.Printf("Post created: %s \n", title)
}

type Facade struct {
	_customer *Customer
	_campaign *Campaign
	_post *Post
}

func NewFacade() *Facade{
	return &Facade{new(Customer), new(Campaign), new(Post)}
}

func(f *Facade) CompleteProcess(customerName string, campaignName string, postTitle string){
	f._customer.CreateCustomer(customerName)
	f._campaign.CreateCampaign(campaignName)
	f._post.CreatePost(postTitle)
}
