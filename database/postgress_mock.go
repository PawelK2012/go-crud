package database

import (
	"context"

	"github.com/PawelK2012/go-crud/models"
)

type PostgressClientMock struct {
	mockDatabase    map[string]map[string]string
	simulateFailure bool
}

func NewPostgressClientMock(mockDatabase map[string]map[string]string, simulateFailure bool) (ClientInterface, error) {
	return &PostgressClientMock{
		mockDatabase:    mockDatabase,
		simulateFailure: simulateFailure,
	}, nil
}

func (client *PostgressClientMock) Init(ctx context.Context) error {
	return nil
}

func (client *PostgressClientMock) GetNoteById(ctx context.Context, id int) (*models.Note, error) {
	return nil, nil
}

func (client *PostgressClientMock) Create(ctx context.Context, note *models.Note) (int64, error) {
	return 1, nil
}

func (client *PostgressClientMock) GetAll(ctx context.Context) ([]models.Note, error) {
	return nil, nil
}
