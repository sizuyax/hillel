package services

import (
	"github.com/stretchr/testify/assert"
	"project-auction/internal/adapters/postgres/repository/mocks"
	"project-auction/internal/domain/entity"
	"testing"
)

func TestInsertComment(t *testing.T) {
	mockRepo := mocks.NewPGCommentRepository(t)
	svc := &commentService{
		CommentRepository: mockRepo,
	}

	inputComment := entity.Comment{
		Body: "test",
	}

	expectedComment := entity.Comment{
		ID:      1,
		ItemID:  1,
		OwnerID: 1,
		Body:    inputComment.Body,
	}

	mockRepo.On("InsertComment", inputComment).Return(expectedComment, nil)

	createComment, err := svc.CreateComment(inputComment)

	assert.NoError(t, err)
	assert.Equal(t, expectedComment, createComment)

	mockRepo.AssertExpectations(t)
}
