package producer

import (
	"fmt"
	"testing"

	"github.com/Shopify/sarama/mocks"
)

func TestSendMessage(t *testing.T) {
	t.Run("Send message success", func(t *testing.T) {
		mockerProducer := mocks.NewSyncProducer(t, nil)
		mockerProducer.ExpectSendMessageAndSucceed()
		kafka := KafkaProducer{
			Producer: mockerProducer,
		}

		msg := "Test ngirim message"
		err := kafka.SendMessage("test_topic", msg)
		if err != nil {
			t.Errorf("Send message should not be error but have %v", err)
		}
	})

	t.Run("Send message failed", func(t *testing.T) {
		mockerProducer := mocks.NewSyncProducer(t, nil)
		mockerProducer.ExpectSendMessageAndFail(fmt.Errorf("Error"))
		kafka := KafkaProducer{
			Producer: mockerProducer,
		}

		msg := "Test ngirim message"

		err := kafka.SendMessage("test_topic", msg)
		if err == nil {
			t.Error("this should be error")
		}
	})
}
