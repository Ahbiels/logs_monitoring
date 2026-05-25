package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	logging "cloud.google.com/go/logging/apiv2"
	loggingpb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/IBM/sarama"
)

var dat map[string]any
var date_key string

func main() {
	var (
		ch     = make(chan *loggingpb.LogEntry)
		wgLogs sync.WaitGroup
	)
	inicio := time.Now()
	ctx := context.Background()
	c, err := logging.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lastHour := time.Now().Add(-5 * time.Minute)
		filter := fmt.Sprintf("timestamp >= %q", lastHour.Format(time.RFC3339))
		req := &loggingpb.ListLogEntriesRequest{
			ResourceNames: []string{
				"projects/cloud-cost-tam-dev/locations/global/logScopes/tam-log-monitoring",
			},
			Filter:  filter,
			OrderBy: "timestamp desc",
		}

		// it := c.ListLogEntries(ctx, req)
		for resp, err := range c.ListLogEntries(ctx, req).All() {
			wgLogs.Add(1)
			wgLogs.Go(func() {
				handler_error(err)
				handler_error(err)
				ch <- resp
				wgLogs.Done()
			})
		}

		go func() {
			wgLogs.Wait()
			close(ch)
		}()

		for logs := range ch {
			Producer(logs)
		}

		duracao := time.Since(inicio)
		result := fmt.Sprintf("A execução levou: %s\n", duracao)

		w.Write([]byte(result))
	})

	fmt.Println("Servidor Iniciado")
	log.Fatal(http.ListenAndServe(":7777", nil))

}

func Producer(logs *loggingpb.LogEntry) {
	date_key = logs.Timestamp.String()
	producer, err := sarama.NewSyncProducer([]string{"broker:9092"}, nil)
	handler_error(err)
	defer producer.Close()

	log_serialized, err := json.Marshal(logs)

	message := &sarama.ProducerMessage{
		Topic: "topic_log_monitoring",
		Value: sarama.ByteEncoder(log_serialized),
		Key:   sarama.StringEncoder(date_key),
	}
	partition, offset, err := producer.SendMessage(message)
	handler_error(err)
	log.Printf("Message sent! Partition=%d Offset=%d\n", partition, offset)
}

func handler_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
