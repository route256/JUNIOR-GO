package answer

import "fmt"

type Sender interface {
	sendAsyncMessage(message paymentMessage) error
	sendMessage(message paymentMessage) error
	sendMessages(messages []paymentMessage) error
}

type Repository interface {
	getAnswer(ID int) *Answer
	verifyAnswer(answerID int, success bool) error
}

type Service struct {
	repo   Repository
	sender Sender
}

func NewService(repo Repository, sender Sender) *Service {
	return &Service{
		repo:   repo,
		sender: sender,
	}
}

func (s Service) Verify(answerID int, success bool, sync bool) {
	answer := s.repo.getAnswer(answerID)
	err := s.repo.verifyAnswer(answerID, success)
	if err != nil {
		fmt.Println("Error in verify", err)
	}

	if sync {
		err = s.sender.sendMessage(
			paymentMessage{
				answer.ID,
				answer.userID,
				answer.sum,
				success,
			},
		)

		if err != nil {
			fmt.Println("Send sync message error: ", err)
		}

		return
	}

	err = s.sender.sendAsyncMessage(
		paymentMessage{
			answer.ID,
			answer.userID,
			answer.sum,
			success,
		},
	)

	if err != nil {
		fmt.Println("Send async message error: ", err)
	}
}

func (s Service) VerifyBatch(answerIDs []int, success bool) {
	var messages []paymentMessage

	for _, ID := range answerIDs {
		answer := s.repo.getAnswer(ID)
		err := s.repo.verifyAnswer(ID, success)
		if err != nil {
			fmt.Println("Error in verify", err)
		}

		messages = append(messages, paymentMessage{
			answer.ID,
			answer.userID,
			answer.sum,
			success,
		})
	}

	err := s.sender.sendMessages(messages)

	if err != nil {
		fmt.Println("Send message error: ", err)
	}
}
