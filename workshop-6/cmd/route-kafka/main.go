package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"route-kafka/internal/app/answer"
	"route-kafka/internal/app/payment"
	"route-kafka/internal/infrastructure/kafka"
	"sync"
	"syscall"
	"time"
)

type paymentMessage struct {
	AnswerID int
	UserID   int
	Sum      int
	Success  bool
}

func main() {
	// получаем из конфига
	var brokers = []string{
		"127.0.0.1:9091",
		"127.0.0.1:9092",
		"127.0.0.1:9093",
	}

	fmt.Println(brokers)

	// producerExample(brokers)

	// consumerExample(brokers)

	consumerGroupExample(brokers)
}

func producerExample(brokers []string) {
	kafkaProducer, err := kafka.NewProducer(brokers)
	if err != nil {
		fmt.Println(err)
	}

	answerService := answer.NewService(
		answer.NewRepository("some connector"),
		answer.NewKafkaSender(kafkaProducer, "payments"),
	)

	answerService.Verify(1, true, false)
	answerService.Verify(2, false, false)

	answerService.Verify(3, false, true)
	answerService.Verify(4, true, true)

	answerService.VerifyBatch([]int{5, 6, 7, 8}, true)

	err = kafkaProducer.Close()
	if err != nil {
		fmt.Println("Close producers error ", err)
	}
}

func consumerExample(brokers []string) {
	kafkaConsumer, err := kafka.NewConsumer(brokers)
	if err != nil {
		fmt.Println(err)
	}

	// обработчики по каждому из топиков
	handlers := map[string]payment.HandleFunc{
		"payments": func(message *sarama.ConsumerMessage) {
			pm := paymentMessage{}
			err = json.Unmarshal(message.Value, &pm)
			if err != nil {
				fmt.Println("Consumer error", err)
			}

			fmt.Println("Received Key: ", string(message.Key), " Value: ", pm)
		},
	}

	payments := payment.NewService(
		payment.NewReceiver(kafkaConsumer, handlers),
	)

	// При условии одного инстанса подходит идеально
	payments.StartConsume("payments")

	<-context.TODO().Done()
}

func consumerGroupExample(brokers []string) {
	keepRunning := true
	log.Println("Starting a new Sarama consumer")

	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion

	/*
		sarama.OffsetNewest - получаем только новые сообщений, те, которые уже были игнорируются
		sarama.OffsetOldest - читаем все с самого начала
	*/
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Используется, если ваш offset "уехал" далеко и нужно пропустить невалидные сдвиги
	config.Consumer.Group.ResetInvalidOffsets = true

	// Сердцебиение консьюмера
	config.Consumer.Group.Heartbeat.Interval = 3 * time.Second

	// Таймаут сессии
	config.Consumer.Group.Session.Timeout = 60 * time.Second

	// Таймаут ребалансировки
	config.Consumer.Group.Rebalance.Timeout = 60 * time.Second

	const BalanceStrategy = "roundrobin"
	switch BalanceStrategy {
	case "sticky":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategySticky}
	case "roundrobin":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRoundRobin}
	case "range":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRange}
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", BalanceStrategy)
	}

	/**
	 * Setup a new Sarama consumer group
	 */
	consumer := kafka.NewConsumerGroup()
	group := "route-example"

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	consumptionIsPaused := false
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, []string{"payments"}, &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
		}
	}()

	<-consumer.Ready() // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
		case <-sigusr1:
			toggleConsumptionFlow(client, &consumptionIsPaused)
		}
	}

	cancel()
	wg.Wait()

	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		log.Println("Resuming consumption")
	} else {
		client.PauseAll()
		log.Println("Pausing consumption")
	}

	*isPaused = !*isPaused
}
