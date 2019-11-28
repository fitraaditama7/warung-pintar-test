package utils

import (
	"time"

	"github.com/Shopify/sarama"
)

/*
 * @desc Function for setting configuration kafka
 *
 * @param  {string} username - username for kafka
 * @param  {string} password - password for kafka
 *
 * @return {*sarama.Config}
 */
func GetKafkaConfig(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}
