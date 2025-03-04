/*
307832097 Roy Kosary
311319313 Nir Agam
*/
package main

import (
	"sync"
	"time"
)

var (
	completedOrders   []Order
	completedOrdersMu sync.Mutex

	zoneManagersDone   int
	zoneManagersDoneMu sync.Mutex
	allDone            bool
)

// StartZoneManager processes orders using a token pool to limit concurrency.
func StartZoneManager(zoneName string, zoneChan <-chan Order, tokens chan struct{}, wg *sync.WaitGroup, totalZones int) {
	defer wg.Done()

	var localWG sync.WaitGroup

	for order := range zoneChan {
		if order.FoodType == "DONE" && order.OrderNumber == -1 {
			break
		}

		localWG.Add(1)
		tokens <- struct{}{} // Acquire a token

		go func(o Order) {
			defer localWG.Done()
			defer func() { <-tokens }()

			time.Sleep(100 * time.Millisecond)
			completedOrdersMu.Lock()
			completedOrders = append(completedOrders, o)
			completedOrdersMu.Unlock()
		}(order)
	}

	localWG.Wait()

	zoneManagersDoneMu.Lock()
	zoneManagersDone++
	if zoneManagersDone == totalZones {
		allDone = true
	}
	zoneManagersDoneMu.Unlock()
}
