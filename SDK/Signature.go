package SDK

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
	"net/url"
)

// getReqSign ：根据 接口请求参数 和 应用密钥 计算 请求签名
// 参数说明
//   - $params：接口请求参数（特别注意：不同的接口，参数对一般不一样，请以具体接口要求为准）
// 返回数据
//   - 签名结果
func getRandomString(lens int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func getReqSign() string {
	// 0. 补全基本参数
	Mapper["app_id"] = getAppId()
	if Mapper["nonce_str"] == "" {
		Mapper["nonce_str"] = getRandomString(32)
	}

	if Mapper["time_stamp"] == "" {
		Mapper["time_stamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	}
	//fmt.Println(Mapper["time_stamp"])
	// 1. 字典升序排序
	var SortString []string
	for k := range Mapper {
		SortString = append(SortString, k)
	}
	sort.Strings(SortString) //会根据字母的顺序进行排序。

	var str string
	for _, v := range SortString {
		if v != "" {
			a,_:=url.Parse(Mapper[v])
			b:=a.String()
			str = str + v + "=" + b + "&"
		}
	}
	//str = str + "app_key=" + strings.ToUpper(getAppKey())
	str = str + "app_key=" + getAppKey()
	//fmt.Println(str)
	sign := md5.Sum([]byte(str))
	return strings.ToUpper(fmt.Sprintf("%x", sign))
}
