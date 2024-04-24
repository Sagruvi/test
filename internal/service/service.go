package service

import (
	"context"
	"testproject/internal/entity"
	"testproject/internal/repo"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.3 --name=Service --output=../mocks/service
type Service interface {
	ProcessEvent(ctx context.Context, event entity.Event) error
	GetEvents(ctx context.Context) ([]entity.Event, error)
	Search(request entity.Request) ([]entity.Event, error)
	Insert() error
}

type service struct {
	repo.Repository
}

func NewService(repo repo.Repository) Service {
	return &service{
		Repository: repo,
	}
}

func (s *service) ProcessEvent(ctx context.Context, event entity.Event) error {
	err := s.Repository.CreateEvent(ctx, event)
	if err != nil {
		return err
	}
	return nil
}
func (s *service) GetEvents(ctx context.Context) ([]entity.Event, error) {
	return s.Repository.GetEvent(ctx)
}
func (s *service) Search(request entity.Request) ([]entity.Event, error) {
	return s.Repository.LogData(request)
}
func (s *service) Insert() error {
	return s.Repository.InsertData()
}
