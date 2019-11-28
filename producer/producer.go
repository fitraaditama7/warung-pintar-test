package producer

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type KafkaProducer struct {
	Producer sarama.SyncProducer
}

/*
 * @desc Function sending message to kafka
 *
 * @param  {string} topic - topic kafka
 * @param  {string} msg - message for send to kafka
 *
 * @return {*sarama.Config}
 */
func (p *KafkaProducer) SendMessage(topic, msg string) error {
	kafkaMessage := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	partition, offset, err := p.Producer.SendMessage(kafkaMessage)
	if err != nil {
		logrus.Errorf("Send message error: %v", err)
		return err
	}

	logrus.Infof("Send message success, Topic %v, Partition %v, Offset %d", topic, partition, offset)
	return nil
}
