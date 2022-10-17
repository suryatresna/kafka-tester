package main

import (
	"context"
	"fmt"
	"log"

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
		group = "consumerapp"
		topic = "onTestConsumer"
	)

	log.Println("[INFO] start consumerapp")
	log.Println("[INFO] brokers ", brokers)
	log.Println("[INFO] group ", group)
	log.Println("[INFO] topic ", topic)

	startConsuming(ctx, brokers, group, topic)
}

func startConsuming(ctx context.Context, brokers []string, group, topic string) {

	log.Println("[INFO] consumer intializing")
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.FetchIsolationLevel(kgo.ReadCommitted()), // only read messages that have been written as part of committed transactions
		kgo.ConsumerGroup(group),
		kgo.ConsumeTopics(topic),
		kgo.WithLogger(kgo.BasicLogger(log.Writer(), kgo.LogLevelInfo, log.Prefix)),
	)
	if err != nil {
		fmt.Printf("error initializing Kafka consumer: %v\n", err)
		return
	}
	// The default blocking commit on leave will not run because the only
	// way to kill this program is to interrupt it, but, usually you will
	// close the client and wait for it to close before quitting. If you
	// want to perform an action on commit errors, you can use the
	// AutoCommitCallback option.
	defer client.Close()

	log.Println("[INFO] entering consume message")

	for {
		fetches := client.PollFetches(ctx)
		iter := fetches.RecordIter()

		for _, fetchErr := range fetches.Errors() {
			log.Printf("[ERROR] error consuming from topic: topic=%s, partition=%d, err=%v\n",
				fetchErr.Topic, fetchErr.Partition, fetchErr.Err)
			continue
		}

		for !iter.Done() {
			record := iter.Next()
			log.Printf("[INFO] consumed record from partition %d with message: %v", record.Partition, string(record.Value))
		}
	}
}
