package main

import (
	"sync"
	"log"
)

/* Merges the slice of channels into one channel */
func FanIn(helperChannels [] <- chan WordAndDefinition) <- chan WordAndDefinition {
	chMerged := make(chan WordAndDefinition)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(helperChannels))

	log.Println("Starting fan-in process")
	for _, channel := range helperChannels {
		go func(chInner <- chan WordAndDefinition) {
			for data := range chInner {
				chMerged <- data
			}
			waitGroup.Done()
		}(channel)
	}

	go func() {
		waitGroup.Wait()
		close(chMerged)
		log.Println("Fan-in complete")
	}()

	return chMerged
}
