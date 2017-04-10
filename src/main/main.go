package main

import (
	"log"
	"net/http"
	"time"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

func main() {
	log.Println("Starting server")
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":9090", nil)
}

func handleRequest(response http.ResponseWriter, request *http.Request) {
	log.Println("Handling request")
	start := time.Now()

	words, numHelpers := decodeRequest(request)

	/* Fanning Out */

	chProducer := getProducerChannel(words) // producer

	helperChannels := FanOut(chProducer, numHelpers) // consumer

	/* Fanning In */

	chMerged := FanIn(helperChannels)

	var data []WordAndDefinition
	for tuple := range chMerged {
		data = append(data, tuple)
	}

	/* Send the information back to the user */

	mainBody := JsonResponseBody{Pairs: data, Time: fmt.Sprintf("%s", time.Since(start))}
	bytes, _ := json.Marshal(mainBody)
	response.Write(bytes)

	log.Println("Request complete")
}

/* Generically-named method to simulate a goroutine that is a ~producer~ */
func getProducerChannel(data []string) <- chan string {
	channel := make(chan string)
	go func(){
		log.Println("Starting producer")
		for _, value := range data {
			channel <- value
		}
		close(channel)
		log.Println("Producer complete")
	}()
	return channel
}

func decodeRequest(request *http.Request) ([]string, int) {
	var decodedRequest Request

	rawData, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(rawData, &decodedRequest)
	if err != nil {
		panic(err)
	}

	defer request.Body.Close()

	return decodedRequest.Words, decodedRequest.NumHelpers
}