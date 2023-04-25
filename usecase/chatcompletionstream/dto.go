package chatcompletion

import "github.com/rAndrade360/gpt-assistant/internal/domain/entity"

type ChatCompletionUseCaseStreamInputDto struct {
	ID      string
	UserID  string
	Message string
	Config  ChatConfig
}

type ChatConfig struct {
	InitialMessage   string
	Model            *entity.Model
	Temperature      float32
	TopP             float32
	N                int
	PresencePenalty  float32
	FrequencyPenalty float32
}

type ChatCompletionUseCaseStreamOutputDto struct {
	ChatID  string
	UserID  string
	Message string
	Role    string
}
