package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
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

func NewChat(userID, initialMessage string, model Model) (*Chat, error) {
	m, err := NewMessage(initialMessage, model, USER)
	if err != nil {
		return nil, err
	}

	c := Chat{
		ID:        uuid.New().String(),
		UserID:    userID,
		Status:    ACTIVE,
		CreatedAt: time.Now(),
	}

	c.AddMessage(m)

	return &c, nil
}

func (c *Chat) AddMessage(m *Message) error {
	if c.Status == ENDED {
		return errors.New("the chat was ended")
	}

	c.Messages = append(c.Messages, *m)

	msgTokens := 0
	for i := len(c.Messages) - 1; i >= 0; i-- {
		msgTokens += c.Messages[i].TotalTokens
		if msgTokens > c.Messages[i].Model.MaxTokens {
			c.Messages = c.Messages[i+1:]
			c.Offset += (i + 1)
			break
		}
	}

	return nil
}
