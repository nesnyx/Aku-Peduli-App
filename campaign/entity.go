package campaign

import (
	"akupeduli/user"
	"time"
)

type Campaign struct {
	ID               int
	UserId           int
	Name             string
	Description      string
	ShortDescription string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImages
	User             user.User
}

type CampaignImages struct {
	ID         int
	CampaignId int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
