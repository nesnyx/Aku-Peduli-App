package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(userId int) ([]Campaign, error)
	FindById(id int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	CreateImage(campaign CampaignImages) (CampaignImages, error)
	MarkAllImagesAsNonPrimary(campaignId int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Preload("User").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil

}

func (r *repository) FindByUserId(userId int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", userId).Preload("User").Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *repository) FindById(id int) (Campaign, error) {
	var campaign Campaign
	err := r.db.Preload("CampaignImages").Preload("User").Where("id = ?", id).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {

	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) CreateImage(campaignImage CampaignImages) (CampaignImages, error) {
	err := r.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}
	return campaignImage, nil
}

func (r *repository) MarkAllImagesAsNonPrimary(campaignId int) (bool, error) {
	// update campaign_images set is_primary = false where campaignId = 1
	err := r.db.Model(&CampaignImages{}).Where("campaign_id = ?", campaignId).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
