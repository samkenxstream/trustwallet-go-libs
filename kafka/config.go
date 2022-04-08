package kafka

import (
	"github.com/Shopify/sarama"
)

func InitDefaultConfig() *sarama.Config {
	cfg := sarama.NewConfig()

	// Default Producer configurations
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Partitioner = sarama.NewRoundRobinPartitioner

	// Default Consumer configurations
	cfg.Consumer.Offsets.AutoCommit.Enable = false
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest

	return cfg
}
