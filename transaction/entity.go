package transaction

import (
	"akupeduli/campaign"
	"akupeduli/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignId int
	UserId     int
	User       user.User
	Amount     int
	Status     string
	Code       string
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
