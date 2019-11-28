package libs

import (
	"warung-pintar-test/cmd/utils"
	"warung-pintar-test/producer"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

/*
 * @desc Function for sending message to kafka
 *
 * @param  {string} msg - message for kafka
 *
 * @return {error}
 */
func SendMessage(msg string) error {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)

	kafkaConfig := utils.GetKafkaConfig("", "")

	producers, err := sarama.NewSyncProducer([]string{"kafka: 9092"}, kafkaConfig)
	if err != nil {
		logrus.Errorf("Unable to create kafka producer got error: %v", err)
		return err
	}

	defer func() {
		if err := producers.Close(); err != nil {
			logrus.Errorf("unable to stop kafka producer got error: %v", err)
			panic(err)
		}
	}()

	logrus.Infof("success create kafka sync producer")
	kafka := producer.KafkaProducer{
		Producer: producers,
	}

	err = kafka.SendMessage("test_topic", msg)
	if err != nil {
		logrus.Errorf("Unable to send message to kafka producer got error: %v", err)
		return err
	}
	return nil
}
