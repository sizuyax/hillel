package services

import (
	"project-auction/internal/adapters/postgres/repository"
	"project-auction/internal/domain/entity"
)

type CommentService interface {
	CreateComment(entity.Comment) (entity.Comment, error)
}

type commentService struct {
	CommentRepository repository.PGCommentRepository
}

func NewCommentService(commentRepository repository.PGCommentRepository) CommentService {
	return &commentService{
		CommentRepository: commentRepository,
	}
}

func (ss commentService) CreateComment(inputComment entity.Comment) (entity.Comment, error) {
	expectedComment, err := ss.CommentRepository.InsertComment(inputComment)
	if err != nil {
		return entity.Comment{}, err
	}

	return expectedComment, nil
}
