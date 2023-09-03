package entity

import (
	"fmt"
	"hbcase/utils"
)

type Campaign struct {
	ID                     int64  `json:"id"`
	Name                   string `json:"name"`
	ProductCode            string `json:"product_code"`
	Duration               int    `json:"duration"`
	PriceManipulationLimit int64  `json:"price_manipulation_limit"`
	TargetSalesCount       int64  `json:"target_sales_count"`
	Status                 bool   `json:"status"`
	EndTime                string `json:"end_time"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
}

func NewCampaign(name, productCode string, duration int, priceManLimit, tSalesCount int64) *Campaign {
	return &Campaign{
		Name:                   name,
		ProductCode:            productCode,
		Duration:               duration,
		PriceManipulationLimit: priceManLimit,
		TargetSalesCount:       tSalesCount,
		EndTime:                utils.CalculateEndTime(duration),
		Status:                 true,
	}
}

func (c *Campaign) ValidateCampaign() *utils.Errors {
	errs := utils.NewErrors()
	if utils.TimeDifference(c.EndTime) <= 0 {
		errs.Add("campaign status", "campaign is ended")
	}

	return errs
}

func (c *Campaign) SetCampaignStatus(status bool) {
	c.Status = status
}

func (c *Campaign) ToString(funcName string, exprs map[string]interface{}) string {
	var str string

	switch funcName {
	case "create_campaign":
		str = fmt.Sprintf("Campaign created; name %s, product %s, duration %d, limit %d, target sales count %d", c.Name, c.ProductCode, c.Duration, c.PriceManipulationLimit, c.TargetSalesCount)
	case "get_campaign_info":
		str = fmt.Sprintf("Campaign %s info; Status %t, Target Sales %d, Total Sales %d, Turnover %d, Average Item Price %d", c.Name, c.Status, c.TargetSalesCount, exprs["totalSales"], exprs["turnover"], exprs["averageItemPrice"])
	default:
		str = "Unknown function name"
	}

	return str
}
