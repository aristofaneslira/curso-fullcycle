package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type Producer struct {
	ConfigMap *ckafka.ConfigMap
}

func NewKafkaProducer(configMap *ckafka.ConfigMap) *Producer {
	return &Producer{
		ConfigMap: configMap,
	}
}

func (p *Producer) Publish(msg interface{}, key []byte, topic string) error {
	producer, err := ckafka.NewProducer(p.ConfigMap)
	if err != nil {
		return err
	}
}
