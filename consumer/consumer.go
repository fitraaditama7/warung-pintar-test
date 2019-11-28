package consumer

import (
	"os"
	"time"
	"warung-pintar-test/cmd/consumer/config"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	Consumer sarama.Consumer
}

/*
 * @desc Function consume data from kafka
 *
 * @param  {[]string} topics - topics kafka
 * @param  {channels of os.Signal} signals - unix signal
 *
 * @return {*sarama.Config}
 */
func (c *KafkaConsumer) Consume(topics []string, signals chan os.Signal) {
	chanMessage := make(chan *sarama.ConsumerMessage, 1024)
	for _, topic := range topics {
		partitionList, err := c.Consumer.Partitions(topic)
		if err != nil {
			logrus.Errorf("unable to get partition got error: %v", err.Error())
			continue
		}
		for _, partition := range partitionList {
			go consumeMessage(c.Consumer, topic, partition, chanMessage)
		}
	}

	logrus.Infof("Kafka is consuming")
	done := time.After(1 * time.Second)
	config.LISTMESSAGE = make([]string, 0)

ConsumerLoop:
	for {
		select {
		case <-done:
			break ConsumerLoop

		case msg := <-chanMessage:
			logrus.Infof("New message from kafka, message: %v", string(msg.Value))
			message := string(msg.Value)
			config.LISTMESSAGE = append(config.LISTMESSAGE, message)

		case sig := <-signals:
			if sig == os.Interrupt {
				break ConsumerLoop
			}
		}
	}
}

/*
 * @desc Function consume message from kafka by partition
 *
 * @param  {sarama.Consumer} consumer - Consumer manages PartitionConsumers which process Kafka messages from brokers.
 * @param  {string} topic - topic kafka
 * @param  {int32} partition - partition kafka
 * @param  {channel of *sarama.ConsumerMessage} c - channel for ConsumerMessage encapsulates a Kafka message returned by the consumer.
 *
 * @return {*sarama.Config}
 */
func consumeMessage(consumer sarama.Consumer, topic string, partition int32, c chan *sarama.ConsumerMessage) {
	msg, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		logrus.Errorf("Unable to consume partition %v got error %v", partition, err)
		return
	}

	defer func() {
		if err := msg.Close(); err != nil {
			logrus.Errorf("unable to close partition %v: %v", partition, err)
		}
	}()

	for {
		msg := <-msg.Messages()
		c <- msg
	}
}
