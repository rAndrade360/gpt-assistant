package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	tiktoken_go "github.com/j178/tiktoken-go"
)

type Status int8

const (
	ACTIVE Status = iota + 1
	ENDED
)

type Chat struct {
	ID        string
	Messages  []Message
	UserID    string
	Status    Status
	Offset    int
	CreatedAt time.Time
}

func NewChat(userID, initialMessage, model string) Chat {
	m := NewMessage(initialMessage, model, USER)
	c := Chat{
		ID:        uuid.New().String(),
		Status:    ACTIVE,
		CreatedAt: time.Now(),
	}

	c.AddMessage(m)

	return c
}

func (c *Chat) AddMessage(m *Message) error {
	if c.Status == ENDED {
		return errors.New("the chat was ended")
	}

	c.Messages = append(c.Messages, *m)

	msgTokens := 0
	for i := len(c.Messages) - 1; i >= 0; i-- {
		s := tiktoken_go.GetContextSize(c.Messages[i].Model)

		if c.Messages[i].TotalTokens+msgTokens > s {
			c.Messages = c.Messages[i+1:]
			c.Offset += (i + 1)
			break
		} else {
			msgTokens += c.Messages[i].TotalTokens
		}
	}

	return nil
}
