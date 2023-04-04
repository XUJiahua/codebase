package feishu

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xujiahua/codebase/pkg/sdk/feishu"
)

type FeishuBotHook struct {
	bot       *feishu.Bot
	logLevels []logrus.Level
}

func NewFeishuBotHook(webhook string, logLevels []logrus.Level) *FeishuBotHook {
	bot := feishu.NewBot(webhook)
	return &FeishuBotHook{
		bot:       bot,
		logLevels: logLevels,
	}
}

func (f FeishuBotHook) Levels() []logrus.Level {
	return f.logLevels
}

func (f FeishuBotHook) Fire(entry *logrus.Entry) error {
	text, err := entry.String()
	if err != nil {
		return err
	}

	go func() {
		err := f.bot.SendText(text)
		if err != nil {
			fmt.Printf("send feishu bot message failed: %v\n", err)
		}
	}()

	return nil
}
