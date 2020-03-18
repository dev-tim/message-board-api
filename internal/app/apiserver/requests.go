package apiserver

import (
	"github.com/dev-tim/message-board-api/internal/app/model"
	"time"
)

type CreateMessageBodyV1ClientRequest struct {
	Id    string
	Name  string
	Email string
	Text  string
}

func (r *CreateMessageBodyV1ClientRequest) ToMessage() *model.Message {
	now := time.Now()
	return &model.Message{
		Id:           r.Id,
		Name:         r.Name,
		Email:        r.Email,
		Text:         r.Text,
		CreationTime: &now,
		CreatedAt:    nil,
		UpdatedAt:    nil,
	}
}

type PatchMessageBodyV1ClientRequest struct {
	Text string
}
