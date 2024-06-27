package services

import (
	"log/slog"
	"project-auction/internal/adapters/postgres/repository"
	"project-auction/internal/domain/entity"
)

type CommentService interface {
	CreateComment(entity.Comment) (entity.Comment, error)
}

type commentService struct {
	log               *slog.Logger
	CommentRepository repository.PGCommentRepository
}

func NewCommentService(log *slog.Logger, commentRepository repository.PGCommentRepository) CommentService {
	return &commentService{
		log:               log,
		CommentRepository: commentRepository,
	}
}

func (ss commentService) CreateComment(inputComment entity.Comment) (entity.Comment, error) {
	expectComment, err := ss.CommentRepository.InsertComment(inputComment)
	if err != nil {
		ss.log.Error("failed to insert comment", slog.String("error", err.Error()))
		return entity.Comment{}, err
	}

	return expectComment, nil
}
