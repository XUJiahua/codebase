package feishu

import (
	"github.com/sirupsen/logrus"
	"github.com/xujiahua/codebase/pkg/sdk/feishu"
	"strings"
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
	// trim \n
	text = strings.TrimSpace(text)
	// escape double quotes
	text = strings.Replace(text, `"`, `\"`, -1)

	return f.bot.SendText(text)
}
