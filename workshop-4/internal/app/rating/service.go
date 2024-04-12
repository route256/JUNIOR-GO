package rating

import "context"

type Type string

const (
	SimpleType   Type = "simple"
	MultiplyType Type = "multiply"
	AdditionType Type = "addition"
)

type Repository interface {
	get(ctx context.Context, userID uint32) (uint32, error)
	save(ctx context.Context, ratingRow ratingRow) error
}

type CalcVariant interface {
	calc(currentValue uint32) uint32
}

type Service struct {
	repository   Repository
	calcVariants map[Type]CalcVariant
}

func NewService(repository Repository, calcVariants map[Type]CalcVariant) *Service {
	return &Service{
		repository:   repository,
		calcVariants: calcVariants,
	}
}

func (s *Service) Calc(ctx context.Context, userID uint32) error {
	currentValue, err := s.repository.get(ctx, userID)
	if err != nil {
		return err
	}

	ratingRow := ratingRow{}
	ratingRow.UserID = userID
	ratingRow.SimpleValue = s.calcVariants[SimpleType].calc(currentValue)
	ratingRow.MultiplyValue = s.calcVariants[MultiplyType].calc(currentValue)
	ratingRow.AdditionValue = s.calcVariants[AdditionType].calc(currentValue)

	return s.repository.save(ctx, ratingRow)
}
