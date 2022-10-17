package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	var (
		ctx     = context.Background()
		brokers = []string{
			"redpanda-0.redpanda.default.svc.cluster.local.:9093",
			"redpanda-1.redpanda.default.svc.cluster.local.:9093",
			"redpanda-2.redpanda.default.svc.cluster.local.:9093",
		}
		topic = "onTestConsumer"
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
		kgo.WithLogger(kgo.BasicLogger(log.Writer(), kgo.LogLevelInfo, log.Prefix)),
	)
	if err != nil {
		log.Println("[ERROR] error initialize redpanda client")
		return
	}

	defer client.Close()

	reqID := 0

	for {
		res := client.ProduceSync(ctx, &kgo.Record{
			Key:   []byte("testproduce"),
			Value: []byte("{\"foo\":\"bar-" + strconv.Itoa(reqID) + "\"}"),
			Topic: topic,
		})

		if res.FirstErr() != nil {
			log.Println("[ERROR] error produce message ", res.FirstErr().Error())
			continue
		}

		time.Sleep(1 * time.Second)
		reqID++
	}
}
