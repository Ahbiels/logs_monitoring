package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	logging "cloud.google.com/go/logging/apiv2"
	loggingpb "cloud.google.com/go/logging/apiv2/loggingpb"
)

var dat map[string]any

func main() {
	var (
		ch     = make(chan []byte)
		wgLogs sync.WaitGroup
	)
	inicio := time.Now()
	ctx := context.Background()
	c, err := logging.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

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
			handler_func(err)
			value, err := json.Marshal(resp)
			handler_func(err)
			ch <- value
			wgLogs.Done()
		})
	}

	go func() {
		wgLogs.Wait()
		close(ch)
	}()

	for log_serelialize := range ch{
		_ = log_serelialize
	}

	duracao := time.Since(inicio)
	fmt.Printf("A execução levou: %s\n", duracao)
}

func handler_func(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
