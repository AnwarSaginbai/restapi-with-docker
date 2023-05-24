package comment

import (
	"context"
	"errors"
	"fmt"
	"log"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment - реализация комментарий для нашего проекта
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - это структура, которая хранит всю нашу логику

type Service struct {
	Store Store
}

// NewService  - возвращает указатель на новый Service

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	log.Println("retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
