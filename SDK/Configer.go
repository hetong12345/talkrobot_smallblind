package SDK

import "fmt"

const API_URL_PATH = "https://api.ai.qq.com/fcgi-bin/"

type BaseApi struct {
	Mapper map[string]string
	ErrMsg string
}

var Mapper = make(map[string]string)

var (
	appId  string
	appKey string
	errMsg string
)

// setAppInfo ：设置AppID和AppKey
// 参数说明
//   - $app_id
//   - $app_key
// 返回数据
//   - 是否设置成功
func SetAppInfo(id, key string) bool {
	if id == "" || key == "" {
		errMsg = "id或key为空"
		fmt.Println(errMsg)
		return false
	}
	appId = id
	appKey = key
	return true
}

func GetErrMsg() string {
	return errMsg
}
func getAppId() string {
	return appId
}
func getAppKey() string {
	return appKey
}
