package user

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
)

type service struct {
	asynq *asynq.Client
}

type Service interface {
	CreateUsersCsv(ctx context.Context, fileURL string) error
}

func NewService(asynq *asynq.Client) Service {
	return &service{asynq}
}

func (s *service) CreateUsersCsv(ctx context.Context, fileURL string) error {
	var request = map[string]interface{}{
		"fileURL": fileURL,
	}
	payload, err := json.Marshal(&request)
	if err != nil {
		return err
	}

	_, err = s.asynq.EnqueueContext(ctx, asynq.NewTask("create-users-csv", payload, asynq.MaxRetry(10)))
	return err
}
