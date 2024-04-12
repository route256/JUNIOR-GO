package tasks

type Status uint8

func (s Status) String() string {
	switch s {
	case ToDo:
		return "TO DO"
	case InProgress:
		return "IN PROGRESS"
	case Done:
		return "DONE"
	}

	return "UNKNOWN"
}

const (
	ToDo       Status = iota // Ожидает начала
	InProgress               // В работе
	Done                     // Завершено
)
