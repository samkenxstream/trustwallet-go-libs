package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type SyncProducer struct {
	conn sarama.SyncProducer
}

func InitSyncProducer(config *sarama.Config, brokers []string) (*SyncProducer, error) {
	sarama.Logger = logrus.New()

	// Sarama requires to set this flag `true` to use sync producer.
	// Producers have 2 channels: successes and errors, if we turn off this flag in config - it will be locked.
	config.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize kafka producer: %w", err)
	}

	return &SyncProducer{conn: syncProducer}, nil
}

func (sp *SyncProducer) Close() error {
	return sp.conn.Close()
}

func (sp *SyncProducer) PublishByteMessage(message []byte, topic string) (int32, int64, error) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	partition, offset, err := sp.conn.SendMessage(msg)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to send message to kafka topic (%s): %w", topic, err)
	}

	return partition, offset, nil
}
