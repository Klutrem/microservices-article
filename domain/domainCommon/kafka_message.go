package domainCommon

import (
	"github.com/IBM/sarama"
)

type KafkaMessage struct {
	sarama.ConsumerMessage
	ReplyTopic string
}
