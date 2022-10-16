package main

import (
	"context"
	"log"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	var (
		ctx     = context.Background()
		brokers = []string{"redpanda.default.svc.cluster.local:9092"}
		topic   = "onTestConsumer"
	)

	log.Println("[INFO] start producer")
	log.Println("[INFO] brokers ", brokers)
	log.Println("[INFO] topic ", topic)

	produceMessage(ctx, brokers, topic)
}

func produceMessage(ctx context.Context, brokers []string, topic string) {
	log.Println("[INFO] initializing produce message")
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)
	if err != nil {
		log.Println("[ERROR] error initialize redpanda client")
		return
	}

	defer client.Close()

	for {
		res := client.ProduceSync(ctx, &kgo.Record{
			Key:   []byte("testproduce"),
			Value: []byte("{\"foo\":\"bar\"}"),
			Topic: topic,
		})

		if res.FirstErr() != nil {
			log.Println("[ERROR] error produce message ", res.FirstErr().Error())
			continue
		}

		time.Sleep(1 * time.Second)
	}
}
