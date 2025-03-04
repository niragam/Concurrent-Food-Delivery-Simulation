/*
307832097 Roy Kosary
311319313 Nir Agam
*/
package main

import (
	"sync"
)

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		return
	}

	producerChans := make([]<-chan Order, 0, len(config.Producers))
	for _, pCfg := range config.Producers {
		ch := make(chan Order, pCfg.QueueSize)
		producerChans = append(producerChans, ch)
		go StartProducer(pCfg, ch)
	}

	mergedChan := fanIn(producerChans)

	zones := make(map[string]chan Order)
	for _, z := range config.Zones {
		zones[z.Name] = make(chan Order, config.ZoneQueueSize)
	}

	var wgZones sync.WaitGroup
	wgZones.Add(len(config.Zones))
	for _, z := range config.Zones {
		tokens := make(chan struct{}, z.Workers)
		go StartZoneManager(z.Name, zones[z.Name], tokens, &wgZones, len(config.Zones))
	}

	go StartDispatcher(mergedChan, zones, len(config.Producers))

	var wgHTTP sync.WaitGroup
	wgHTTP.Add(1)
	go StartHTTPServer(config.HTTPServerPort, &wgHTTP)

	wgZones.Wait()

	wgHTTP.Wait()
}
