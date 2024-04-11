package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

type Producer struct {
	brokers       []string
	syncProducer  sarama.SyncProducer
	asyncProducer sarama.AsyncProducer
}

func newAsyncProducer(brokers []string) (sarama.AsyncProducer, error) {
	asyncProducerConfig := sarama.NewConfig()

	asyncProducerConfig.Producer.Partitioner = sarama.NewHashPartitioner
	asyncProducerConfig.Producer.RequiredAcks = sarama.WaitForAll

	asyncProducerConfig.Producer.Return.Successes = true
	asyncProducerConfig.Producer.Return.Errors = true

	asyncProducer, err := sarama.NewAsyncProducer(brokers, asyncProducerConfig)

	if err != nil {
		return nil, errors.Wrap(err, "error with async kafka-producer")
	}

	go func() {
		// Error и Retry топики можно использовать при получении ошибки
		for e := range asyncProducer.Errors() {
			fmt.Println(e.Error())
		}
	}()

	go func() {
		for m := range asyncProducer.Successes() {
			fmt.Println("Async success with key", m.Key)
		}
	}()

	return asyncProducer, nil
}

func newSyncProducer(brokers []string) (sarama.SyncProducer, error) {
	syncProducerConfig := sarama.NewConfig()

	// случайная партиция
	// syncProducerConfig.Producer.Partitioner = sarama.NewRandomPartitioner

	// по кругу
	// syncProducerConfig.Producer.Partitioner = sarama.NewRoundRobinPartitioner

	// по ключу
	// syncProducerConfig.Producer.Partitioner = sarama.NewHashPartitioner
	/**
	Кейсы:
		- одинаковые ключи в одной партиции
		- при cleanup.policy = compact останется только последнее сообщение по этому ключу
	*/
	syncProducerConfig.Producer.Partitioner = sarama.NewRandomPartitioner

	syncProducerConfig.Producer.RequiredAcks = sarama.WaitForAll

	/*
	  Если хотим exactly once, то выставляем в true

	  У продюсера есть счетчик (count)
	  Каждое успешно отправленное сообщение учеличивает счетчик (count++)
	  Если продюсер не смог отправить сообщение, то счетчик не меняется и отправляется в таком виде в другом сообщение
	  Кафка это видит и начинает сравнивать (в том числе Key) сообщения с одниковыми счетчиками
	  Далее не дает отправить дубль, если Idempotent = true
	*/
	syncProducerConfig.Producer.Idempotent = true
	syncProducerConfig.Net.MaxOpenRequests = 1

	// Если хотим сжимать, то задаем нужный уровень кодировщику
	syncProducerConfig.Producer.CompressionLevel = sarama.CompressionLevelDefault

	syncProducerConfig.Producer.Return.Successes = true
	syncProducerConfig.Producer.Return.Errors = true

	// И сам кодировщик
	syncProducerConfig.Producer.Compression = sarama.CompressionGZIP

	syncProducer, err := sarama.NewSyncProducer(brokers, syncProducerConfig)

	if err != nil {
		return nil, errors.Wrap(err, "error with sync kafka-producer")
	}

	return syncProducer, nil
}

func NewProducer(brokers []string) (*Producer, error) {
	syncProducer, err := newSyncProducer(brokers)
	if err != nil {
		return nil, errors.Wrap(err, "error with sync kafka-producer")
	}

	asyncProducer, err := newAsyncProducer(brokers)
	if err != nil {
		return nil, errors.Wrap(err, "error with async kafka-producer")
	}

	producer := &Producer{
		brokers:       brokers,
		syncProducer:  syncProducer,
		asyncProducer: asyncProducer,
	}

	return producer, nil
}

func (k *Producer) SendSyncMessage(message *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	return k.syncProducer.SendMessage(message)
}

func (k *Producer) SendSyncMessages(messages []*sarama.ProducerMessage) error {
	err := k.syncProducer.SendMessages(messages)
	if err != nil {
		fmt.Println("kafka.Connector.SendMessages error", err)
	}

	return err
}

func (k *Producer) SendAsyncMessage(message *sarama.ProducerMessage) {
	k.asyncProducer.Input() <- message
}

func (k *Producer) Close() error {
	err := k.syncProducer.Close()
	if err != nil {
		return errors.Wrap(err, "kafka.Connector.Close")
	}

	// k.asyncProducer.AsyncClose() - не использовать в общем случае
	err = k.asyncProducer.Close()
	if err != nil {
		return errors.Wrap(err, "kafka.Connector.Close")
	}

	return nil
}
