package feishu

import (
	"encoding/json"
	"strings"
)

type Bot struct {
	webhook string
	sdk     *Sdk
}

func NewBot(webhook string) *Bot {
	return &Bot{
		webhook: webhook,
		sdk:     NewSDK("", ""),
	}
}

type TextMessageRequest struct {
	MsgType string             `json:"msg_type"`
	Content TextMessageContent `json:"content"`
}

type TextMessageContent struct {
	Text string `json:"text"`
}

// SendText https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN#756b882f
func (b Bot) SendText(text string) error {
	bs, err := json.Marshal(TextMessageRequest{
		MsgType: "text",
		Content: TextMessageContent{
			Text: text,
		},
	})
	if err != nil {
		return err
	}

	return b.sdk.WebhookV2(b.webhook, strings.NewReader(string(bs)))
}
