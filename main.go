package main

import (
	"fmt"
	"log"
	"miner/runner"
	"sync"
	"time"
)

func main() {
	fmt.Printf("Starting Trademiner...")

	fetch_ticker := time.NewTicker(60 * 60 * 6 * time.Second)
	process_ticker := time.NewTicker(60 * 60 * 6 * time.Second)
	quit := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			select {
			case <-fetch_ticker.C:
				log.Println("Data Fetch Start!")
				runner.Run()
				log.Println("Data Fetch Complete!")
			case <-process_ticker.C:
				log.Println("Ai Stat Start!")
				runner.RunAiStats()
				log.Println("Ai Stat Complete!")
			case <-quit:
				fetch_ticker.Stop()
				process_ticker.Stop()
				wg.Done()
				return
			}
		}
	}()
	wg.Wait()
}
