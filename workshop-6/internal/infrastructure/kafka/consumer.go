package kafka

import (
	"github.com/Shopify/sarama"
	"time"
)

type Consumer struct {
	brokers        []string
	SingleConsumer sarama.Consumer
}

func NewConsumer(brokers []string) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 5 * time.Second
	/*
		sarama.OffsetNewest - получаем только новые сообщений, те, которые уже были игнорируются
		sarama.OffsetOldest - читаем все с самого начала
	*/
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer(brokers, config)

	if err != nil {
		return nil, err
	}

	/*
		consumer.Topics() - список топиков
		consumer.Partitions("test_topic") - партиции топика
		consumer.ConsumePartition("test_topic", 1, 12) - чтение конкретного топика с 12 сдвига в первой партиции
		consumer.Pause() - останавливаем чтение определенных топиков
		consumer.Resume() - восстанавливаем чтение определенных топиков
		consumer.PauseAll() - останавливаем чтение всех топиков
		consumer.ResumeAll() - восстанавливаем чтение всех топиков
	*/

	return &Consumer{
		brokers:        brokers,
		SingleConsumer: consumer,
	}, err
}
