package broker

import (
	"context"
	"errors"
	"github.com/Shopify/sarama"
	"route256/libs/broker/kafka"
	"sync"
)

type HandleFunc func(ctx context.Context, message *sarama.ConsumerMessage)

type KafkaReceiver struct {
	consumer *kafka.Consumer
	handlers map[string]HandleFunc
}

func NewReceiver(consumer *kafka.Consumer) *KafkaReceiver {
	return &KafkaReceiver{
		consumer: consumer,
		handlers: map[string]HandleFunc{},
	}
}

func (r *KafkaReceiver) RegisterHandler(topic string, handler HandleFunc) {
	r.handlers[topic] = handler
}

func (r *KafkaReceiver) Subscribe(ctx context.Context, topic string) error {
	handler, ok := r.handlers[topic]

	if !ok {
		return errors.New("can not find handler")
	}

	// получаем все партиции топика
	partitionList, err := r.consumer.SingleConsumer.Partitions(topic)

	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	for _, partition := range partitionList {

		pc, err := r.consumer.SingleConsumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			return err
		}
		wg.Add(1)
		go func(pc sarama.PartitionConsumer, partition int32) {
			defer wg.Done()

			for {
				select {
				case message := <-pc.Messages():
					handler(ctx, message)
				case <-ctx.Done():
					return
				}
			}
		}(pc, partition)
	}

	wg.Wait()

	return nil
}
