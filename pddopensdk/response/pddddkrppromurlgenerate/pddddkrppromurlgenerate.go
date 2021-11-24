package pddddkrppromurlgenerate

import (
	"encoding/json"

	response2 "github.com/chree188/UnionSDK/pddopensdk/response"
)

//pdd.ddk.rp.prom.url.generate生成营销工具推广链接
type Response struct {
	response2.TopResponse
	RpPromotionURLGenerateResponse RpPromotionURLGenerateResponse `json:"rp_promotion_url_generate_response"`
}

//解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.ErrorCode = -1
		t.ErrorResponse.ErrorMsg = unmarshal.Error()
	}
}

type RpPromotionURLGenerateResponse struct {
	URLList      []URLList      `json:"url_list"`
	RequestID    string         `json:"request_id"`
	ResourceList []ResourceList `json:"resource_list"`
}

type ResourceList struct {
	Desc string `json:"desc"`
	Url  string `json:"url"`
}

type URLList struct {
	QqAppInfo                *AppInfo `json:"qq_app_info"`
	MultiGroupMobileShortURL string   `json:"multi_group_mobile_short_url"`
	MultiGroupURL            string   `json:"multi_group_url"`
	MobileURL                string   `json:"mobile_url"`
	MultiGroupShortURL       string   `json:"multi_group_short_url"`
	SchemaURL                string   `json:"schema_url"`
	WeAppInfo                *AppInfo `json:"we_app_info"`
	MobileShortURL           string   `json:"mobile_short_url"`
	MultiGroupMobileURL      string   `json:"multi_group_mobile_url"`
	URL                      string   `json:"url"`
	ShortURL                 string   `json:"short_url"`
}

type AppInfo struct {
	UserName          string  `json:"user_name"`
	PagePath          string  `json:"page_path"`
	QqAppIconURL      *string `json:"qq_app_icon_url,omitempty"`
	SourceDisplayName string  `json:"source_display_name"`
	Title             string  `json:"title"`
	AppID             string  `json:"app_id"`
	Desc              string  `json:"desc"`
	WeAppIconURL      *string `json:"we_app_icon_url,omitempty"`
}
