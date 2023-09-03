package repository

import (
	"hbcase/internal/domain/entity"

	"gorm.io/gorm"
)

type CampaignRepositoryI interface {
	SaveCampaign(campaign *entity.Campaign) error
	GetCampaignInfo(campaingName string) (*entity.Campaign, error)
	GetCampaigns() ([]entity.Campaign, error)
}

type campaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) CampaignRepositoryI {
	return &campaignRepository{
		db: db,
	}
}

func (cr *campaignRepository) SaveCampaign(campaign *entity.Campaign) error {
	return cr.db.Save(campaign).Error
}

func (cr *campaignRepository) GetCampaignInfo(campaignName string) (*entity.Campaign, error) {
	var campaign entity.Campaign
	err := cr.db.Where("name =?", campaignName).Order("-id").First(&campaign).Error
	return &campaign, err
}

func (cr *campaignRepository) GetCampaigns() ([]entity.Campaign, error) {
	var campaigns []entity.Campaign
	err := cr.db.Find(&campaigns).Error
	return campaigns, err
}
