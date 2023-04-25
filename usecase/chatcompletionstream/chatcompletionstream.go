package chatcompletion

import (
	"context"
	"errors"
	"io"

	"github.com/rAndrade360/gpt-assistant/internal/domain/entity"
	"github.com/rAndrade360/gpt-assistant/internal/domain/gateway"
	openai "github.com/sashabaranov/go-openai"
)

type ChatCompletionUseCaseStream interface {
	Execute(ctx context.Context, input ChatCompletionUseCaseStreamInputDto) error
}

type chatCompletionUseCaseStream struct {
	gateway  gateway.ChatGateway
	aiclient openai.Client
	stream   chan ChatCompletionUseCaseStreamOutputDto
}

func NewChatCompletionUseCase(gateway gateway.ChatGateway, aiclient openai.Client, stream chan ChatCompletionUseCaseStreamOutputDto) ChatCompletionUseCaseStream {
	return &chatCompletionUseCaseStream{
		gateway:  gateway,
		aiclient: aiclient,
		stream:   stream,
	}
}

func (uc *chatCompletionUseCaseStream) Execute(ctx context.Context, input ChatCompletionUseCaseStreamInputDto) error {
	chat, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil && !errors.Is(errors.New("item not found"), err) {
		return err
	}

	if chat == nil {
		chat, err = entity.NewChat(input.UserID, input.Config.InitialMessage, *input.Config.Model)
		if err != nil {
			return err
		}
	}

	usrMsg, err := entity.NewMessage(input.Message, *input.Config.Model, entity.USER)
	if err != nil {
		return err
	}

	chat.AddMessage(usrMsg)

	var aimessages = make([]openai.ChatCompletionMessage, len(chat.Messages))

	for i := range chat.Messages {
		aimessages[i] = openai.ChatCompletionMessage{
			Role:    chat.Messages[i].Role.String(),
			Content: chat.Messages[i].Data,
		}
	}

	res, err := uc.aiclient.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Model:            input.Config.Model.Name,
		Messages:         aimessages,
		MaxTokens:        input.Config.Model.MaxTokens,
		Temperature:      input.Config.Temperature,
		TopP:             input.Config.TopP,
		N:                input.Config.N,
		Stream:           true,
		PresencePenalty:  input.Config.PresencePenalty,
		FrequencyPenalty: input.Config.FrequencyPenalty,
	})

	if err != nil {
		return err
	}

	for {
		response, err := res.Recv()
		if err == io.EOF {
			break
		}

		msg, err := entity.NewMessage(response.Choices[0].Delta.Content, *input.Config.Model, entity.SYSTEM)
		if err != nil {
			return err
		}

		err = chat.AddMessage(msg)
		if err != nil {
			return err
		}

		out := ChatCompletionUseCaseStreamOutputDto{
			ChatID:  chat.ID,
			UserID:  chat.UserID,
			Message: response.Choices[0].Delta.Content,
			Role:    entity.SYSTEM.String(),
		}

		uc.stream <- out
	}

	err = uc.gateway.Update(ctx, chat)

	return err
}
