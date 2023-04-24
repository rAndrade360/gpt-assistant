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
		name       string
		fields     fields
		args       args
		wantErr    bool
		wantOffset int
	}{
		{
			name: "should be able to add message with offset 1",
			fields: fields{
				ID: "",
				Messages: []Message{
					{
						ID:   "12345678",
						Role: USER,
						Data: "You are the Best. My friend!",
						Model: Model{
							Name:      "gpt-4",
							MaxTokens: 8,
						},
						TotalTokens: 7,
					},
				},
				UserID: "12345789",
				Status: ACTIVE,
				Offset: 0,
			},
			args: args{
				m: func() *Message {
					m, _ := NewMessage("I kow. I kow!", Model{Name: "gpt-4", MaxTokens: 8}, SYSTEM)
					return m
				}(),
			},
			wantErr:    false,
			wantOffset: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewChat(tt.fields.UserID, tt.fields.Messages[0].Data, tt.fields.Messages[0].Model)
			if err != nil {
				t.Fatalf("Err to create chat: %s", err.Error())
			}

			if err := c.AddMessage(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Chat.AddMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
			if c.Offset != tt.wantOffset {
				t.Errorf("Offset error, want: %d, got: %d", tt.wantOffset, c.Offset)
			}
		})
	}
}
