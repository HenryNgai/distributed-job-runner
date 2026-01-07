package main

import (
	"encoding/json"
	"log"
)

type Job struct {
	/*
		Note
		Struct tag used for JSON serialization. This is basically saying when
		this struct is converted to JSON, the field ID should be represented
		with the key "id".
	*/

	ID   int    `json:"id"`
	Data string `json:"data"`
}

func createDummyJobs() []Job {
	jobs := []Job{
		{ID: 1, Data: "First Job"},
		{ID: 2, Data: "Second Job"},
		{ID: 3, Data: "Third Job"},
	}
	return jobs
}

func pushToQueue(job Job) {
	// Placeholder for pushing job to a message queue
	log.Printf("Pushing job to queue: %+v\n", job)
}

func main() {
	log.Println("Producer started")
	firstJob := Job{ID: 1, Data: "Sample Data"}
	log.Printf("Created job: %+v\n", firstJob)
	jsonByteSlice, err := json.Marshal(firstJob)
	if err != nil {
		log.Fatalf("Failed to marshal job: %v", err)
	}
	log.Printf("JSON data: %s\n", jsonByteSlice)
}
