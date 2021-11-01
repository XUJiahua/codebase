package feishu

import (
	"bytes"
	"text/template"
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

const textMessage = `{
    "msg_type": "text",
    "content": {
        "text": "{{.}}"
    }
}`

var textTemplate = template.Must(template.New("text").Parse(textMessage))

// SendText https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN#756b882f
func (b Bot) SendText(text string) error {
	var buf bytes.Buffer
	err := textTemplate.Execute(&buf, text)
	if err != nil {
		return err
	}
	return b.sdk.WebhookV2(b.webhook, &buf)
}
