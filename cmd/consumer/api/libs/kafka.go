package libs

import (
	"os"
	"warung-pintar-test/cmd/utils"
	"warung-pintar-test/consumer"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

/*
 * @desc Function for consume all message in kafka
 *
 * @return {error}
 */
func GetAllMessage() error {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)

	kafkaConfig := utils.GetKafkaConfig("", "")

	consumers, err := sarama.NewConsumer([]string{"kafka:9092"}, kafkaConfig)
	if err != nil {
		logrus.Errorf("Unable to create consumer got error: %v", err)
		return err
	}

	defer func() {
		if err := consumers.Close(); err != nil {
			logrus.Fatal(err)
			panic(err)
		}
	}()

	kafkaConsumer := consumer.KafkaConsumer{
		Consumer: consumers,
	}

	signals := make(chan os.Signal, 1)
	kafkaConsumer.Consume([]string{"test_topic"}, signals)
	return nil
}
