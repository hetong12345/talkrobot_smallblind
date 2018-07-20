package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"encoding/json"
)

var GJsonEncodeReplacer *strings.Replacer = strings.NewReplacer(
	"\\", "\\\\",
	"\"", "\\\"",
	"\b", "\\b",
	"\f", "\\f",
	"\n", "\\n",
	"\r", "\\r",
	"\t", "\\t")

func openFile() *os.File {
	f, err := os.OpenFile("./a.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("OpenFile, err:%s", err.Error()))
	}
	return f
}

func Dump(data []byte, f *os.File) {
	f.Write(data)
	f.Write([]byte{'\n'})
}

func UTF8ToGBK(utf8 []byte) (gbk []byte, err error) {
	reader := transform.NewReader(
		bytes.NewReader(utf8), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

type GBKString []byte

func (s *GBKString) MarshalJSON() ([]byte, error) {
	return []byte("\"" + GJsonEncodeReplacer.Replace(string(*s)) + "\""), nil
}

type Raw struct {
	Data string `json:"Raw"`
}

type GBKWrap struct {
	Data GBKString `json:"Raw"`
}

func main() {
	// 1. utf8 转码成 gbk
	s := "所有命运赠送的礼物,都早已在暗中标好了价格。出处:出自《断头女王》"
	gbk, err := UTF8ToGBK([]byte(s))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(gbk)
	fmt.Println(string(gbk))

	b:=GBKString{}
	b=[]byte(gbk)
	c,_:=b.MarshalJSON()
	fmt.Println(c)
	fmt.Println(string(c))

	f := openFile()
	Dump(gbk, f)

	// 2. gbk格式的string用json直接Marshal
	var raw Raw
	raw.Data = string(gbk)
	var tmpBytes []byte
	tmpBytes, _ = json.Marshal(&raw)
	Dump(tmpBytes, f)
	fmt.Println(string(tmpBytes),f)

	// 3. gbk格式的string的正确使用Marshal的方法
	var gbkWarp GBKWrap

	gbkWarp.Data = gbk

	tmpBytes, _ = json.Marshal(&gbkWarp)
	Dump(tmpBytes, f)
	fmt.Println(string(tmpBytes),f)

	c:=strings.NewReplacer(
		"a","b",
		"c","d",
	)
	str:="aaaaqweqweaaaccccqweccccqweccc"
	println(c.Replace(str))
}
