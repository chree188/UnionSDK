package request

import (
	"net/url"

	utils2 "github.com/chree188/UnionSDK/utils"
)

//pdd.ddk.theme.prom.url.generate（多多进宝主题推广链接生成
//https://open.pinduoduo.com/#/apidocument/port?id=pdd.ddk.resource.url.gen
type PddDdkThemePromUrlGenerateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkThemePromUrlGenerateRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("pid"), "pid")
	utils2.CheckNotNull(tk.Parameters.Get("theme_id_list"), "theme_id_list")

}

//添加请求参数
func (tk *PddDdkThemePromUrlGenerateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkThemePromUrlGenerateRequest) GetApiName() string {
	return "pdd.ddk.theme.prom.url.generate"
}

//返回请求参数
func (tk *PddDdkThemePromUrlGenerateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
