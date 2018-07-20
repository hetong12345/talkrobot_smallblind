package SDK

import (

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"bytes"
	"io/ioutil"
	"fmt"
)

const CHUNK_SIZE = 6400

func isBase64(str string) bool {

	return true

}
func Nlpresolve(name string, text string) data{
	urlAll := API_URL_PATH + "nlp/" + name
	Mapper["text"] = string(Utf8ToGbk([]byte(text)))
	Mapper["sign"] = getReqSign()

	return doHttpPost(urlAll, Mapper)
}

func GbkToutf8(gbk []byte) string {
	reader := transform.NewReader(
		bytes.NewReader(gbk), simplifiedchinese.GBK.NewDecoder())
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
		return "err"
	}
	return string(d)
}
func Utf8ToGbk(utf8 []byte) ([]byte) {//已经ok
	reader := transform.NewReader(
		bytes.NewReader(utf8), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil
	}
	return d
}