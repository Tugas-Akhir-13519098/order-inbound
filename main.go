package main

import (
	"log"
    "net/http"
    "github.com/segmentio/kafka-go"
    "encoding/json"
)

type Message struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}

func handlePublishMessage(w http.ResponseWriter, r *http.Request) {
	config := kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test",
	}
	writer := kafka.NewWriter(config)

    // Parse the JSON payload from the request body
    var message Message
    err := json.NewDecoder(r.Body).Decode(&message)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Publish the message to Kafka
    err = writer.WriteMessages(r.Context(), kafka.Message{
        Key:   []byte(message.Key),
        Value: []byte(message.Value),
    })
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Return a success response to the client
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Message published to Kafka"))
}

func main() {
	log.Printf("Application is running")
    
	http.HandleFunc("/publish", handlePublishMessage)
    http.ListenAndServe(":8080", nil)
}