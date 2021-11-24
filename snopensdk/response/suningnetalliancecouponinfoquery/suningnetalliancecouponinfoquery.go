package suningnetalliancecouponinfoquery

import (
	"encoding/json"

	response2 "github.com/chree188/UnionSDK/snopensdk/response"
)

//suning.netalliance.couponinfo.query 查询券领用情况
type Response struct {
	response2.TopResponse
	SnResponseContent SnResponseContent `json:"sn_responseContent"`
}

//解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.ErrorCode = "-1"
		t.ErrorResponse.ErrorMsg = unmarshal.Error()
	} else {
		if t.SnResponseContent.SnError != nil {
			t.ErrorResponse.ErrorCode = t.SnResponseContent.SnError.ErrorCode
			t.ErrorResponse.ErrorMsg = t.SnResponseContent.SnError.ErrorMsg
		}
	}
}

type SnResponseContent struct {
	SnBody  SnBody   `json:"sn_body"`
	SnError *SnError `json:"sn_error"`
}
type SnError struct {
	ErrorMsg  string `json:"error_msg"`
	ErrorCode string `json:"error_code"`
}

type SnBody struct {
	QueryCouponinfo QueryCouponinfo `json:"queryCouponinfo"`
}

type QueryCouponinfo struct {
	CouponRemainingAmount string              `json:"couponRemainingAmount"`
	CouponCount           string              `json:"couponCount"`
	CouponValue           string              `json:"couponValue"`
	CouponEndTime         string              `json:"couponEndTime"`
	ActivityID            string              `json:"activityId"`
	PreferentialDistinct  string              `json:"preferentialDistinct"`
	CouponPlatform        string              `json:"couponPlatform"`
	MemberAttributeList   MemberAttributeList `json:"memberAttributeList"`
	CouponStartTime       string              `json:"couponStartTime"`
	StartTime             string              `json:"startTime"`
	EndTime               string              `json:"endTime"`
	IsValidCoupon         string              `json:"isValidCoupon"`
	ChannelID             string              `json:"channelId"`
}

type MemberAttributeList struct {
	MemberLevel          string `json:"memberLevel"`
	VipMemberLimit       string `json:"vipMemberLimit"`
	MemberAttributeLimit string `json:"memberAttributeLimit"`
}
