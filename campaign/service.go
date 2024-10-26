package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(id GetCampaignDetailInput, input CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampignImageInput, filLocation string) (CampaignImages, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userId int) ([]Campaign, error) {
	if userId != 0 {
		campaigns, err := s.repository.FindByUserId(userId)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignById(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindById(input.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {

	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.Description
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.User.ID = input.User.ID

	slugMake := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugMake)
	// generated slug

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}

func (s *service) UpdateCampaign(inputId GetCampaignDetailInput, input CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindById(inputId.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserId != input.User.ID {
		return campaign, errors.New("not an owner of the campaign")
	}

	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount

	updateCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updateCampaign, err
	}
	return updateCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampignImageInput, fileLocation string) (CampaignImages, error) {
	isPrimary := 0
	campaign, err := s.repository.FindById(input.CampaignID)
	if err != nil {
		return CampaignImages{}, err
	}
	if campaign.UserId != input.User.ID {
		return CampaignImages{}, errors.New("not an owner of the campaign")
	}

	if input.IsPrimary {
		isPrimary = 1
		_, err := s.repository.MarkAllImagesAsNonPrimary(input.CampaignID)
		if err != nil {
			return CampaignImages{}, err
		}
	}

	campaignImages := CampaignImages{}
	campaignImages.ID = input.CampaignID

	campaignImages.IsPrimary = isPrimary
	campaignImages.FileName = fileLocation

	newCampaignImage, err := s.repository.CreateImage(campaignImages)
	if err != nil {
		return newCampaignImage, err
	}
	return newCampaignImage, nil
}
