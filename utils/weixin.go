package utils

import (
	"fmt"

	. "github.com/silenceper/wechat/oauth"
	"github.com/silenceper/wechat"
)


func GetUserInfoByCode(code string) (result UserInfo) {
	//配置微信参数 TODO viper配置
	config := &wechat.Config{
		AppID:          "your app id",
		AppSecret:      "your app secret",
		Token:          "your token",
		EncodingAESKey: "your encoding aes key",
	}
	wc := wechat.NewWechat(config)
	oauth := wc.GetOauth()
	resToken, err := oauth.GetUserAccessToken(code)
	if err != nil {
		fmt.Println(err)
		return
	}
	//getUserInfo
	result, errs := oauth.GetUserInfo(resToken.AccessToken, resToken.OpenID)
	if errs != nil {
		fmt.Println(errs)
		return
	}
	fmt.Println(result)
	return result
}

type Config struct {
	AppID     string `json:"app_id"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	Signature string `json:"signature"`
}

