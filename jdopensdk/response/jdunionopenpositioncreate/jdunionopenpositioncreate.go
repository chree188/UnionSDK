package jdunionopenpositioncreate

import (
	"encoding/json"

	"github.com/chree188/UnionSDK/jdopensdk/response"
)

// Response jd.union.open.position.create 创建推广位
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_position_create_responce"`
}

// WrapResult 解析输出结果
func (t *Response) WrapResult(result string) {
	err := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if err != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Message = err.Error()
	} else {
		//解析queryResult
		if t.Responce.QueryResultStr == "" {
			t.ErrorResponse.Code = -2
			t.ErrorResponse.Message = "empty queryResult"
		} else {
			if err = json.Unmarshal([]byte(t.Responce.QueryResultStr), &t.Responce.QueryResult); err != nil {
				t.ErrorResponse.Code = -1
				t.ErrorResponse.Message = err.Error()
			} else {
				t.ErrorResponse.Code = t.Responce.QueryResult.Code
				t.ErrorResponse.Message = t.Responce.QueryResult.Message
			}
		}
	}
	t.Responce.QueryResultStr = ""
}

// Responce 响应结果
type Responce struct {
	Code           string `json:"code"`
	QueryResultStr string `json:"createResult"`
	QueryResult    QueryResult
}

// QueryResult 具体内容
type QueryResult struct {
	Code    int64  `json:"code"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

type Data struct {
	SesultList string `json:"resultList"`
	SiteId     int64  `json:"siteId"`
	Type       int64  `json:"type"`
	UnionId    int64  `json:"unionId"`
	Pid        string `json:"pid"` //pid：仅uniontype传4时，展示pid；可用于内容平台推广
}

type SesultList struct {
	Result string `json:"result"`
}

type PidList struct {
	Pid string `json:"pid"` //pid结果，仅uniontype传4时，展示pid；可用于内容平台推广
}
