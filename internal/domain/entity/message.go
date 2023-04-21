package entity

import (
	"time"

	"github.com/google/uuid"
	tiktoken_go "github.com/j178/tiktoken-go"
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
	Model       string
	TotalTokens int
	CreatedAt   time.Time
}

func NewMessage(data, model string, role Role) *Message {
	tokens := tiktoken_go.CountTokens(model, data)
	return &Message{
		ID:          uuid.NewString(),
		Role:        role,
		Data:        data,
		Model:       model,
		TotalTokens: tokens,
		CreatedAt:   time.Now(),
	}
}
