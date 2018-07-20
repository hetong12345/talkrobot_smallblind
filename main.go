package main

import (
	"./SDK"
	"fmt"

)

const (
	ID="1106980079"
	KEY="q95QUrpLGwB41INk"
)
func main() {
	SDK.SetAppInfo(ID,KEY)  //初始化

	text1 := "我是个科学家"
	res1:=SDK.Nlpresolve("nlp_wordseg", text1)
	res2:=SDK.Nlpresolve("nlp_wordpos", text1)
	res3:=SDK.Nlpresolve("nlp_wordner", text1)
	res4:=SDK.Nlpresolve("nlp_wordsyn", text1)
	fmt.Println(res1,res2,res3,res4)
}


