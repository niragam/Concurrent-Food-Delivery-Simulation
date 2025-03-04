/*
307832097 Roy Kosary
311319313 Nir Agam
*/
package main

import (
	"sync"
)

// fanIn merges all producer channels into a single channel.
func fanIn(chs []<-chan Order) <-chan Order {
	out := make(chan Order)
	var wg sync.WaitGroup
	wg.Add(len(chs))

	for _, c := range chs {
		go func(c <-chan Order) {
			defer wg.Done()
			for o := range c {
				out <- o
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// StartDispatcher routes orders to the correct zone channel based on food type.
// Once all producers signal DONE, it sends a "DONE" message to each zone.
func StartDispatcher(in <-chan Order, zones map[string]chan Order, producerCount int) {
	doneCount := 0

	for order := range in {
		if order.OrderNumber == -1 {
			doneCount++
			if doneCount == producerCount {
				for _, zch := range zones {
					zch <- Order{FoodType: "DONE", OrderNumber: -1}
					close(zch)
				}
				return
			}
			continue
		}
		zones[mapFoodTypeToZone(order.FoodType)] <- order
	}
}

func mapFoodTypeToZone(foodType string) string {
	switch foodType {
	case "Pizza":
		return "PizzaZone"
	case "Burger":
		return "BurgerZone"
	case "Sushi":
		return "SushiZone"
	default:
		return "UnknownZone"
	}
}
