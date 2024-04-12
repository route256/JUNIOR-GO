package answer

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"route-kafka/internal/infrastructure/kafka"
)

type paymentMessage struct {
	AnswerID int
	UserID   int
	Sum      int
	Success  bool
}

type KafkaSender struct {
	producer *kafka.Producer
	topic    string
}

func NewKafkaSender(producer *kafka.Producer, topic string) *KafkaSender {
	return &KafkaSender{
		producer,
		topic,
	}
}

func (s *KafkaSender) sendAsyncMessage(message paymentMessage) error {
	kafkaMsg, err := s.buildMessage(message)
	if err != nil {
		fmt.Println("Send message marshal error", err)
		return err
	}

	s.producer.SendAsyncMessage(kafkaMsg)

	fmt.Println("Send async message with key:", kafkaMsg.Key)
	return nil
}

func (s *KafkaSender) sendMessage(message paymentMessage) error {
	kafkaMsg, err := s.buildMessage(message)
	if err != nil {
		fmt.Println("Send message marshal error", err)
		return err
	}

	partition, offset, err := s.producer.SendSyncMessage(kafkaMsg)

	if err != nil {
		fmt.Println("Send message connector error", err)
		return err
	}

	fmt.Println("Partition: ", partition, " Offset: ", offset, " AnswerID:", message.AnswerID)
	return nil
}

func (s *KafkaSender) sendMessages(messages []paymentMessage) error {
	var kafkaMsg []*sarama.ProducerMessage
	var message *sarama.ProducerMessage
	var err error

	for _, m := range messages {
		message, err = s.buildMessage(m)
		kafkaMsg = append(kafkaMsg, message)

		if err != nil {
			fmt.Println("Send message marshal error", err)
			return err
		}
	}

	err = s.producer.SendSyncMessages(kafkaMsg)

	if err != nil {
		fmt.Println("Send message connector error", err)
		return err
	}

	fmt.Println("Send messages count:", len(messages))
	return nil
}

func (s *KafkaSender) buildMessage(message paymentMessage) (*sarama.ProducerMessage, error) {
	msg, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Send message marshal error", err)
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic:     s.topic,
		Value:     sarama.ByteEncoder(msg),
		Partition: -1,
		Key:       sarama.StringEncoder(fmt.Sprint(message.AnswerID)),
		Headers: []sarama.RecordHeader{ // например, в хедер можно записать версию релиза
			{
				Key:   []byte("test-header"),
				Value: []byte("test-value"),
			},
		},
	}, nil
}
