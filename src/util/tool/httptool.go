package tool

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type ResMsgS struct {
	Msg string `json:"msg"`
}

func ResMsg(res http.ResponseWriter, code int, msg string) {
	resMsg_ := ResMsgS{Msg: msg}
	result := ""
	if code != 200 {
		b, err := json.Marshal(resMsg_)
		if err != nil {
			result = `{"msg":"内部错误"}`
			code = 500
		} else {
			result = string(b)
		}
	} else {
		result = msg
	}
	res.WriteHeader(code)
	res.Write([]byte(result))
}

func ResInvalidRequestBody(res http.ResponseWriter) {
	ResMsg(res, 400, "请求报文格式错误")
}

func ParseResMsg(r string) string {
	data := ResMsgS{}
	reqContent := []byte(r)
	err := json.Unmarshal(reqContent, &data)
	if err != nil {
		return "JSON格式错误"
	}
	return data.Msg
}

func ResSuccessMsg(res http.ResponseWriter, code int, msg string) {
	resMsg_ := ResMsgS{Msg: msg}
	result := ""
	b, err := json.Marshal(resMsg_)
	fmt.Println(b)
	if err != nil {
		result = `{"msg":"内部错误"}`
		code = 500
	} else {
		result = string(b)
		fmt.Println(result)
		//result = b
	}
	//res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(code)
	res.Write([]byte(result))
}

func ParseJsonStr(str string, body interface{}) error {
	data := []byte(str)
	err := json.Unmarshal(data, body)
	if err != nil {
		return err
	} else {
		return nil
	}
}