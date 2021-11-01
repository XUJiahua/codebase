package feishu

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBot_SendText(t *testing.T) {
	err := NewBot("https://open.feishu.cn/open-apis/bot/v2/hook/eb4caa67-8b03-4e8e-84ba-fc90f7900308").SendText("hello")
	require.NoError(t, err)
}
