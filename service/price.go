package service

import (
	"math"
	"math/rand"
	"time"
)

type priceService struct {

}

func (p *priceService) GetPriceFor (itemCode string) (float64, error){
	time.Sleep(time.Second)
	var price = math.Round(rand.Float64() * 10 * 50) / 10
	return price, nil
}

func NewPriceService() *priceService {
	return &priceService{}
}