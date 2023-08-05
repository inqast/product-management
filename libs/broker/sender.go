package broker

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"route256/libs/broker/kafka"
	log "route256/libs/logger"
)

type message interface {
	GetKey() string
}

type KafkaSender struct {
	producer *kafka.Producer
	topic    string
}

func NewKafkaSender(
	producer *kafka.Producer,
	topic string,
) *KafkaSender {
	return &KafkaSender{
		producer: producer,
		topic:    topic,
	}
}

func (s *KafkaSender) SendMessage(message message) error {
	kafkaMsg, err := s.buildMessage(message)
	if err != nil {
		log.Error("Send message marshal error", err)
		return err
	}

	partition, offset, err := s.producer.SendSyncMessage(kafkaMsg)

	if err != nil {
		log.Error("Send message connector error", err)
		return err
	}

	log.Info("Partition: ", partition, " Offset: ", offset, " AnswerID:", message.GetKey())
	return nil
}

func (s *KafkaSender) SendMessages(messages []message) error {
	var kafkaMsg []*sarama.ProducerMessage
	var message *sarama.ProducerMessage
	var err error

	for _, m := range messages {
		message, err = s.buildMessage(m)
		kafkaMsg = append(kafkaMsg, message)

		if err != nil {
			log.Error("Send message marshal error", err)
			return err
		}
	}

	err = s.producer.SendSyncMessages(kafkaMsg)

	if err != nil {
		log.Error("Send message connector error", err)
		return err
	}

	log.Info("Send messages count:", len(messages))
	return nil
}

func (s *KafkaSender) buildMessage(message message) (*sarama.ProducerMessage, error) {
	msg, err := json.Marshal(message)

	if err != nil {
		log.Error("Send message marshal error", err)
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic:     s.topic,
		Value:     sarama.ByteEncoder(msg),
		Partition: -1,
		Key:       sarama.StringEncoder(message.GetKey()),
	}, nil
}
