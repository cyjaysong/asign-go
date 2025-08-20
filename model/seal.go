package model

// 查询印章 Req https://web.asign.cn/platform/openDoc/docDetail?mid=getUserSeals
type GetUserSealsReqBody struct {
	Account string `json:"account" dc:"用户唯一识别码"`
	SealNo  string `json:"sealNo" dc:"印章编号（不传则返回印章列表）"`
}

// 查询印章 Res
type GetUserSealsResBody struct {
	Name        string             `json:"name" dc:"个人用户名称/企业法人名称"`
	IdNo        string             `json:"idNo" dc:"个人用户证件号/企业法人身份证"`
	Mobile      string             `json:"mobile,omitempty" dc:"用户接收短信通知的手机号"`
	UserType    int                `json:"userType" dc:"用户类型：1：企业 2：个人"`
	CompanyName string             `json:"companyName,omitempty" dc:"企业名称（当用户类型为企业时返回）"`
	List        []GetUserSealsItem `json:"list" dc:"印章列表"`
}

// 查询印章 Item
type GetUserSealsItem struct {
	SealName  string `json:"sealName" dc:"印章名称(印章抬头文字)"`
	SealUrl   string `json:"sealUrl" dc:"印章图片访问地址(当印章还在做成中时，此值为空)"`
	IsDefault int    `json:"isDefault" dc:"是否默认章：1：是 0：否"`
	SealNo    string `json:"sealNo" dc:"印章编号"`
}

// 印章授权（无校验） Req https: //web.asign.cn/platform/openDoc/docDetail?mid=nocheckAuth
type NocheckAuthReqBody struct {
	AuthAccount string `json:"authAccount" dc:"真实姓名"`
	SealNo      string `json:"sealNo" dc:"授权印章编号（仅企业印章可授权）"`
	Expired     string `json:"expired,omitempty" dc:"到期时间，不传默认为9999年23点59分59秒，指定格式：yyyy-MM-dd"`
	IsDefault   int    `json:"isDefault" dc:"是否默认授权用户：（若为默认后续签约若使用此印章则默认当该授权用户签约）,1：是,0：否（默认）"`
}

// 印章授权（无校验） Res
type NocheckAuthResBody struct{}
