package SDK

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"encoding/json"
)

var httpCode int //获取上一次Http请求的状态码
//var resp http.Response
type res struct {
	Ret  int    `json:"ret"`  //返回码； 0表示成功，非0表示出错
	Msg  string `json:"msg"`  //返回信息；ret非0时表示出错时错误原因
	Data data `json:"data"` //返回数据；ret为0时有意义
}
type data struct {
	Text        string  `json:"text"`
	Mix_tokens  []token `json:"mix_tokens"`
	Base_tokens []token `json:"base_tokens"`
}
type token struct {
	Word   string `json:"word"`   //分词
	Offset int    `json:"offset"` //分词所在文本偏移量（字节）
	Length int    `json:"length"` //分词长度（字节）
}
var LastData data
func GetHttpCode() int  {
	return httpCode
}
func resolveResp(resp http.Response)bool {
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil ReadAll failed :", err.Error())
		return false
	}
	r := &res{}
	json.Unmarshal([]byte(GbkToutf8([]byte(response))), r)
	LastData=r.Data
	return true
}
func doHttpPost(urlAll string, Mapper map[string]string) data {
	val := url.Values{}
	for k, v := range Mapper {
		val.Add(k, v)
	}
	resp, err := http.PostForm(urlAll, val)
	if err != nil {
		errMsg = "发送post请求错误"
		fmt.Println(err)
		return data{}
	}

	httpCode = resp.StatusCode
	resolveResp(*resp)
	return LastData
}
