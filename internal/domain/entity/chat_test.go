package entity

import (
	"testing"
	"time"
)

func TestNewChat(t *testing.T) {
	type args struct {
		userID         string
		initialMessage string
		model          Model
	}
	tests := []struct {
		name string
		args args
		want Chat
	}{
		{
			name: "should be able to return a chat",
			args: args{
				userID:         "123456789",
				initialMessage: "You are the best!",
				model: Model{
					Name:      "gpt-4",
					MaxTokens: 4092,
				},
			},
			want: Chat{
				ID: "1234567890",
				Messages: []Message{
					{
						Role: USER,
						Data: "You are the best!",
						Model: Model{
							Name:      "gpt-4",
							MaxTokens: 4092,
						},
						TotalTokens: 5,
					},
				},
				UserID: "123456789",
				Status: ACTIVE,
				Offset: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := NewChat(tt.args.userID, tt.args.initialMessage, tt.args.model)
			if got.Offset != tt.want.Offset {
				t.Errorf("Offset error. want: %d, got: %d", tt.want.Offset, got.Offset)
			}
			if got.UserID != tt.want.UserID {
				t.Errorf("UserID error. want: %s, got: %s", tt.want.UserID, got.UserID)
			}
			if got.Status != tt.want.Status {
				t.Errorf("Status error. want: %d, got: %d", tt.want.Status, got.Status)
			}
			for i := range got.Messages {
				if got.Messages[i].Role != tt.want.Messages[i].Role {
					t.Errorf("Message Role error. want: %d, got: %d", tt.want.Messages[i].Role, got.Messages[i].Role)
				}
				if got.Messages[i].Data != tt.want.Messages[i].Data {
					t.Errorf("Message Data error. want: %s, got: %s", tt.want.Messages[i].Data, got.Messages[i].Data)
				}
				if got.Messages[i].Model != tt.want.Messages[i].Model {
					t.Errorf("Message Model error. want: %v, got: %v", tt.want.Messages[i].Model, got.Messages[i].Model)
				}
				if got.Messages[i].TotalTokens != tt.want.Messages[i].TotalTokens {
					t.Errorf("Message TotalTokens error. want: %d, got: %d", tt.want.Messages[i].TotalTokens, got.Messages[i].TotalTokens)
				}
			}
		})
	}
}

func TestChat_AddMessage(t *testing.T) {
	type fields struct {
		ID        string
		Messages  []Message
		UserID    string
		Status    Status
		Offset    int
		CreatedAt time.Time
	}
	type args struct {
		m *Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chat{
				ID:        tt.fields.ID,
				Messages:  tt.fields.Messages,
				UserID:    tt.fields.UserID,
				Status:    tt.fields.Status,
				Offset:    tt.fields.Offset,
				CreatedAt: tt.fields.CreatedAt,
			}
			if err := c.AddMessage(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Chat.AddMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
