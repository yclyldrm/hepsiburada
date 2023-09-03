package services

import (
	"hbcase/internal/domain/entity"
	"hbcase/internal/domain/repository"
)

type CampaignServiceI interface {
	SaveCampaign(*entity.Campaign) error
	GetCampaignInfo(string) (*entity.Campaign, error)
	GetCampaigns() ([]entity.Campaign, error)
}

type campaignService struct {
	rp repository.CampaignRepositoryI
}

func NewCampaignService(rp repository.CampaignRepositoryI) CampaignServiceI {
	return &campaignService{
		rp: rp,
	}
}

func (s *campaignService) SaveCampaign(campaign *entity.Campaign) error {
	return s.rp.SaveCampaign(campaign)
}

func (s *campaignService) GetCampaignInfo(code string) (*entity.Campaign, error) {
	return s.rp.GetCampaignInfo(code)
}

func (s *campaignService) GetCampaigns() ([]entity.Campaign, error) {
	return s.rp.GetCampaigns()
}
