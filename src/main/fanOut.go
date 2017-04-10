package main

import "log"

/* Kicks off our helpers -> return a slice of channels that these helpers feed data into */
func FanOut(chSource <- chan string, numGoroutines int) [] <- chan WordAndDefinition {
	helpers := make([] <- chan WordAndDefinition, numGoroutines)
	log.Println("Starting fan-out process")

	for i := range helpers {
		helpers[i] = definitionHelper(chSource)
	}

	log.Println("Fan-out complete")
	return helpers
}

func definitionHelper(words <- chan string) <- chan WordAndDefinition {
	channel := make(chan WordAndDefinition)
	go func() {
		log.Println("Starting helper")

		for word := range words {
			// this is the expensive operation
			definition := SendRequest(word)
			channel <- WordAndDefinition{word, definition}
		}

		close(channel)

		log.Println("Helper complete")
	}()
	return channel
}