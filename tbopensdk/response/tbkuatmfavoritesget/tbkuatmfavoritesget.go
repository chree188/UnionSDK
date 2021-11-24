package tbkuatmfavoritesget

import (
	"encoding/json"

	"github.com/chree188/UnionSDK/tbopensdk/response"
)

// Response  taobao.tbk.uatm.favorites.get( 获取淘宝联盟选品库列表 )
type Response struct {
	response.TopResponse
}

//解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}
