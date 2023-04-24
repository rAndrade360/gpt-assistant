package entity

import (
	"time"

	"github.com/google/uuid"
	tiktoken_go "github.com/pkoukk/tiktoken-go"
)

type Role int8

const (
	USER Role = iota + 1
	SYSTEM
)

type Message struct {
	ID          string
	Role        Role
	Data        string
	Model       Model
	TotalTokens int
	CreatedAt   time.Time
}

func NewMessage(data string, model Model, role Role) (*Message, error) {
	tk, err := tiktoken_go.EncodingForModel(model.Name)
	if err != nil {
		return nil, err
	}

	tokens := tk.Encode(data, nil, nil)

	return &Message{
		ID:          uuid.NewString(),
		Role:        role,
		Data:        data,
		Model:       model,
		TotalTokens: len(tokens),
		CreatedAt:   time.Now(),
	}, nil
}
