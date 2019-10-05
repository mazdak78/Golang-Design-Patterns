package factory

import "Golang-Design-Patterns/strategy"

//factory + strategy

const (
	StringType1 string = "StringType1"
	StringType2 string = "StringType2"
)

func Factory(typeCampaign string, name string)(ctx *strategy.CampaignContext){
	switch typeCampaign{
		case StringType1:
			return strategy.NewCampaignContext( name, &strategy.CampaignType1{})
		case StringType2:
			return strategy.NewCampaignContext( name, &strategy.CampaignType2{})
	}

	return nil
}
