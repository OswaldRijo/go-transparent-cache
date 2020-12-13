package sample1

import (
	"sync"
	"time"
	"fmt"
)

// Search: interface with privates methods to improve calls to PriceService
type Search interface {
	//Makes synchronization between all threads built for searching prices
	syncSearch( code string)

	//Sets the parallel search
	parallelizeSearch( itemCodes *[]string, results *[]float64, e *error)
}

// PriceService is a service that we can use to get prices for the items
// Calls to this service are expensive (they take time)
type PriceService interface {
	GetPriceFor(itemCode string) (float64, error)
}

// TransparentCache is a cache that wraps the actual service
// The cache will remember prices we ask for, so that we don't have to wait on every call
// Cache should only return a price if it is not older than "maxAge", so that we don't get stale prices
type TransparentCache struct {
	actualPriceService PriceService
	maxAge             time.Duration
	prices             sync.Map
	priceChannel		chan float64
}

func (c *TransparentCache) parallelizeSearch( itemCodes *[]string, results *[]float64, e *error) {
	c.priceChannel = make(chan float64)
	var i = 0
	var price float64
	var qtyProcess = len(*itemCodes)

	for _, itemCode := range *itemCodes {
		go c.syncSearch(itemCode)
	}
	for i < qtyProcess{
		select {
		case price = <-c.priceChannel:
			*results = append(*results, price)
			i++
		}
	}
	close(c.priceChannel)
}

func (c *TransparentCache) syncSearch( code string) {
	price, _ := c.GetPriceFor(code)

	c.priceChannel <- price
}


func NewTransparentCache(actualPriceService PriceService, maxAge time.Duration) *TransparentCache {
	return &TransparentCache{
		actualPriceService: actualPriceService,
		maxAge:             maxAge,
		prices:             sync.Map{},
	}
}

// GetPriceFor gets the price for the item, either from the cache or the actual service if it was not cached or too old
func (c *TransparentCache) GetPriceFor(itemCode string) (float64, error) {
	price, ok := c.prices.Load(itemCode)
	if ok {
		// TODO: check that the price was retrieved less than "maxAge" ago!
		return (price).(float64), nil
	}
	price, err := c.actualPriceService.GetPriceFor(itemCode)
	if err != nil {
		return 0, fmt.Errorf("getting price from service : %v", err.Error())
	}
	c.prices.Store(itemCode, price)
	return (price).(float64), nil
}

// GetPricesFor gets the prices for several items at once, some might be found in the cache, others might not
// If any of the operations returns an error, it should return an error as well
func (c *TransparentCache) GetPricesFor(itemCodes ...string) ([]float64, error) {
	results := &[]float64{}
	var err error
	c.parallelizeSearch(&itemCodes, results, &err)

	return *results, err
}
