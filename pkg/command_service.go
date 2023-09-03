package pkg

import (
	"hbcase/internal/domain/entity"
	"hbcase/internal/domain/services"
	"hbcase/utils"
	"log"
	"strconv"
	"strings"
)

type CommandService struct {
	ps services.ProductService
	os services.OrderService
	cs services.CampaignServiceI
}

func NewCommandService(ps services.ProductService, os services.OrderService, cs services.CampaignServiceI) CommandService {
	return CommandService{ps: ps, os: os, cs: cs}
}

type Command struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}

func GenerateCommands(list []string) []Command {
	commands := make([]Command, 0)
	for i := range list {
		params := strings.Split(list[i], " ")
		cmd := Command{
			Name: params[0],
			Args: params[1:],
		}
		commands = append(commands, cmd)
	}

	return commands
}

func ExecuteCommands(commands []Command, cmdService CommandService) []string {
	var finalStr string
	outputs := make([]string, 0)
	for i := range commands {
		cmd := commands[i]
		switch cmd.Name {
		case "create_product":
			finalStr = commands[i].CreateProduct(cmdService.ps)
		case "get_product_info":
			finalStr = commands[i].GetProductInfo(cmdService.ps)
		case "create_order":
			finalStr = commands[i].CreateOrder(cmdService)
		case "create_campaign":
			finalStr = commands[i].CreateCampaign(cmdService.cs)
		case "get_campaign_info":
			finalStr = commands[i].GetCampaign(cmdService)
		case "increase_time":
			finalStr = commands[i].IncreaseTime(cmdService)
		default:
			log.Printf("Unknown command: %s", cmd.Name)
		}
		outputs = append(outputs, finalStr)
	}
	return outputs
}

func (c Command) CreateProduct(ps services.ProductService) string {
	args := c.Args
	code := args[0]
	price, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		log.Println("Price needs to be a number")
		return err.Error()
	}

	stock, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		log.Println("Stock needs to be a number")
		return err.Error()
	}

	product := entity.NewProduct(code, price, stock)
	errs := product.ValidateProduct()
	if errs.HasError() {
		return errs.GetErrorsString()
	}

	if err := ps.SaveProduct(product); err != nil {
		log.Println("Error occured while saving product")
	}

	return product.String(c.Name)
}

func (c Command) GetProductInfo(ps services.ProductService) string {
	product, err := ps.GetProductInfo(c.Args[0])
	if err != nil {
		log.Println("Product not found")
		return err.Error()
	}

	return product.String(c.Name)
}

func (c Command) CreateOrder(cmdService CommandService) string {
	args := c.Args
	productCode := args[0]
	quantity, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		log.Println("Quantity needs to be a number")
		return err.Error()
	}

	product, err := cmdService.ps.GetProductInfo(productCode)
	if err != nil {
		log.Println("Product not found")
		return err.Error()
	}

	order := entity.NewOrder(productCode, product.Price, quantity)
	errs := order.ValidateOrder()
	if errs.HasError() {
		return errs.GetErrorsString()
	}

	if err := cmdService.os.Create(order); err != nil {
		log.Println("Error occured while saving order")
	}

	product.DecreaseStock(quantity)
	if err := cmdService.ps.SaveProduct(product); err != nil {
		log.Println("Error occured while saving product")
	}

	return order.String()
}

func (c Command) CreateCampaign(cs services.CampaignServiceI) string {
	name := c.Args[0]
	pCode := c.Args[1]
	duration, err := strconv.ParseInt(c.Args[2], 10, 64)
	if err != nil {
		log.Println("Duration needs to be a number")
		return err.Error()
	}

	limit, err := strconv.ParseInt(c.Args[3], 10, 64)
	if err != nil {
		log.Println("Price manipulation limit needs to be a number")
		return err.Error()
	}

	salesCount, err := strconv.ParseInt(c.Args[4], 10, 64)
	if err != nil {
		log.Println("Target sales count needs to be a number")
		return err.Error()
	}

	campaign := entity.NewCampaign(name, pCode, int(duration), limit, salesCount)

	errs := campaign.ValidateCampaign()
	if errs.HasError() {
		return errs.GetErrorsString()
	}

	if err := cs.SaveCampaign(campaign); err != nil {
		return "Error occured while saving campaign"
	}

	return campaign.ToString(c.Name, nil)
}

func (c Command) GetCampaign(cmdService CommandService) string {
	campaign, err := cmdService.cs.GetCampaignInfo(c.Args[0])
	if err != nil {
		log.Println("Campaign not found")
		return err.Error()
	}

	orders, err := cmdService.os.GetOrders(campaign.ProductCode)
	if err != nil {
		return err.Error()
	}

	data := entity.CalculateOrdersData(orders, campaign.TargetSalesCount)

	return campaign.ToString(c.Name, data)
}

func (c Command) IncreaseTime(cmdService CommandService) string {
	incTime, _ := strconv.ParseInt(c.Args[0], 10, 64)

	if err := utils.IncreaseTime(int(incTime)); err != nil {
		return err.Error()
	}

	if campaigns, err := cmdService.cs.GetCampaigns(); err == nil {
		for i := range campaigns {
			campaign := &campaigns[i]

			product, err := cmdService.ps.GetProductInfo(campaigns[i].ProductCode)
			if err != nil {
				log.Println("Product couldn't find")
				continue
			}

			if errs := campaign.ValidateCampaign(); errs.HasError() {
				campaign.SetCampaignStatus(false)
				if err := cmdService.cs.SaveCampaign(campaign); err != nil {
					log.Println("Error occured while updating campaign")
					continue
				}

				product.SetPriceatFirst() // if campaign is already done, price should return first price
			} else {
				// according to time, product's price iterates
				product.ChangeProductPrice(int64(campaign.Duration), campaign.PriceManipulationLimit)
			}

			if err := cmdService.ps.SaveProduct(product); err != nil {
				log.Println("Error occured while updating product")
			}
		}
	}

	return utils.GetSystemTime()
}
