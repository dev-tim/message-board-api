package store

import "github.com/dev-tim/message-board-api/internal/app/model"

type IMessagesRepository interface {
	Create(m *model.Message) (*model.Message, error)
	FindById(id string) (*model.Message, error)
	FindLatest(limit, offset int) ([]*model.Message, error)
	Update(id, text string) (*int64, error)
}
