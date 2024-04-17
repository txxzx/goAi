package zhipu

import (
	"testing"

	"github.com/swxctx/xlog"
)

/**
    @date: 2024/4/16
**/
func TestAuth(t *testing.T) {
	reloadClient()
	accessToken, expireIn := GetAccessToken()
	t.Logf("GetAccessToken: token-> %s, expireIn-> %d", accessToken, expireIn)
}

func reloadClient() {
	if err := NewClient(
		"bcb3c9fe3e1b147252c6a81a9dc6ab70.bCaavKN64xeiqjDU",
		true); err != nil {
		xlog.Errorf("NewClient: err-> %v", err)
	}
}

func TestChatStream(t *testing.T) {
	reloadClient()
	err := Chat(
		&RequestArgs{
			Model: "glm-3-turbo",
			Messages: []Messages{
				{
					Role:    "user",
					Content: "你好",
				},
			},
		})
	if err != nil {
		t.Errorf("Chat: err-> %v", err)
		return
	}
}
