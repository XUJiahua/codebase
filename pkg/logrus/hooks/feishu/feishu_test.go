package feishu

import (
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestNewFeishuBotHook(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	hook := NewFeishuBotHook("https://open.feishu.cn/open-apis/bot/v2/hook/eb4caa67-8b03-4e8e-84ba-fc90f7900308", []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
	})
	logrus.AddHook(hook)

	logrus.Info("info")
	logrus.Error("error")
	_, err := time.Parse(time.RFC3339, "")
	if err != nil {
		logrus.Error(err)
		return
	}
}
