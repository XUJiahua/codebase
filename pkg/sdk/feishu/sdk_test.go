package feishu

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSdk_AccessToken(t *testing.T) {
	sdk := NewSDK("", "")
	token, err := sdk.AccessToken("")
	require.NoError(t, err)
	spew.Dump(token)
}
