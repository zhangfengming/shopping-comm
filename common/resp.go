package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H struct {
	Code                   int
	Message                string
	TraceId                string
	Data                   interface{}
	Rows                   interface{}
	Total                  interface{}
	SkyWalkingDynamicField string
}

func Resp(w http.ResponseWriter, code int, data interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:    code,
		Data:    data,
		Message: message,
	}

	ret, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}

func RespOK(w http.ResponseWriter, data interface{}, message string) {
	Resp(w, 200, data, message)
}
func RespFail(w http.ResponseWriter, data interface{}, message string) {
	Resp(w, 500, data, message)
}
func RespCreated(w http.ResponseWriter, code int, data interface{}, message string) {
	Resp(w, 201, data, message)
}

func RespList(w http.ResponseWriter, code int, data interface{}, message string, rows interface{}, total interface{}, skyWalkingDynamicField string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:                   code,
		Data:                   data,
		Message:                message,
		Total:                  total,
		Rows:                   rows,
		SkyWalkingDynamicField: skyWalkingDynamicField,
	}

	ret, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}

func RespListOK(w http.ResponseWriter, code int, data interface{}, message string, rows interface{}, total interface{}, skyWalkingDynamicField string) {
	RespList(w, 200, data, message, rows, total, skyWalkingDynamicField)
}
func RespListFail(w http.ResponseWriter, code int, data interface{}, message string, rows interface{}, total interface{}, skyWalkingDynamicField string) {
	RespList(w, 500, data, message, rows, total, skyWalkingDynamicField)
}
