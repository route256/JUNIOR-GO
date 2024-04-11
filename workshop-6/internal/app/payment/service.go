package payment

import (
	"fmt"
)

type Receiver interface {
	Subscribe(topic string) error
}

type Service struct {
	receiver Receiver
}

func NewService(receiver Receiver) *Service {
	return &Service{
		receiver: receiver,
	}
}

func (s *Service) StartConsume(topic string) {
	err := s.receiver.Subscribe(topic)

	if err != nil {
		fmt.Println("Subscribe error ", err)
	}
}
