package gateway

import (
	"context"

	"github.com/rAndrade360/gpt-assistant/internal/domain/entity"
)

type ChatGateway interface {
	FindByID(ctx context.Context, id string) (*entity.Chat, error)
	Create(ctx context.Context, chat *entity.Chat) error
	Update(ctx context.Context, chat *entity.Chat) error
}
