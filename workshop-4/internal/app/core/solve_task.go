package core

import "context"

func (s *Service) SolveTask(ctx context.Context, ID uint32, userID uint32, answer string) error {
	err := s.taskService.Solve(ctx, ID, userID, answer)
	if err != nil {
		return err
	}

	return s.ratingService.Calc(ctx, userID)
}
