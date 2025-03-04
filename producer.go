/*
307832097 Roy Kosary
311319313 Nir Agam
*/
package main

// Order represents a single food order.
type Order struct {
	Restaurant  string
	FoodType    string
	OrderNumber int
}

// StartProducer creates orders and sends them out.
// The final message is a "DONE" signal with OrderNumber = -1.
func StartProducer(cfg ProducerConfig, producerChan chan<- Order) {
	defer close(producerChan)

	for i := 0; i < cfg.Orders; i++ {
		order := Order{
			Restaurant:  cfg.Restaurant,
			FoodType:    cfg.FoodType,
			OrderNumber: i,
		}
		producerChan <- order
	}

	done := Order{
		Restaurant:  cfg.Restaurant,
		FoodType:    "",
		OrderNumber: -1,
	}
	producerChan <- done
}
