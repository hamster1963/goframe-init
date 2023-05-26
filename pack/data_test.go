package packed

import (
	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"testing"
)

var (
	CryptoKey = []byte("x76cgqt36i9c863bzmotuf8626dxiwu0")
)

// 测试方法注释
func Test09_44_37(t *testing.T) {
	binContent, err := gres.Pack("/Users/laixin/Developer/chat-im/core-demo")
	if err != nil {
		panic(err)
	}
	binContent, err = gaes.Encrypt(binContent, CryptoKey)
	if err != nil {
		panic(err)
	}
	if err := gfile.PutBytes("data.bin", binContent); err != nil {
		panic(err)
	}
}
