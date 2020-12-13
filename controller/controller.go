package controller

import(
	cache "Golang-challenge/pkg"
	"Golang-challenge/service"
	"time"
)
var cacheInstance = cache.NewTransparentCache(service.NewPriceService(), time.Minute)
type Controller struct {

}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) GetPrice(itemCode string) (float64, error)  {
	return cacheInstance.GetPriceFor(itemCode)
}

func (c *Controller) GetPrices(itemCodes ...string) ([]float64, error)  {
	return cacheInstance.GetPricesFor(itemCodes...)
}