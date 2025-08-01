package model

// 个人身份证二要素比对 Req https://web.asign.cn/platform/openDoc/docDetail?mid=idcard2
type VerifyPersonIdCard2ReqBody struct {
	RealName string `json:"realName" dc:"真实姓名"`
	IdCardNo string `json:"idCardNo" dc:"身份证号"`
}

// 个人运营商二要素比对 Req https://web.asign.cn/platform/openDoc/docDetail?mid=mobile2
type VerifyPersonEnt2ReqBody struct {
	RealName string `json:"realName" dc:"真实姓名"`
	Mobile   string `json:"mobile" dc:"手机号"`
}

// 个人运营商三要素比对 Req https://web.asign.cn/platform/openDoc/docDetail?mid=mobileAuth3
type VerifyPersonEnt3ReqBody struct {
	RealName string `json:"realName" dc:"真实姓名"`
	IdCardNo string `json:"idCardNo" dc:"身份证号"`
	Mobile   string `json:"mobile" dc:"手机号"`
}

// 个人银行卡三要素比对 Req https://web.asign.cn/platform/openDoc/docDetail?mid=bank3
type VerifyPersonBank3ReqBody struct {
	RealName string `json:"realName" dc:"真实姓名"`
	IdCardNo string `json:"idCardNo" dc:"身份证号"`
	BankCard string `json:"bankCard" dc:"银行卡号"`
}

// 个人银行卡四要素比对 Req https://web.asign.cn/platform/openDoc/docDetail?mid=bank4
type VerifyPersonBank4ReqBody struct {
	RealName string `json:"realName" dc:"真实姓名"`
	IdCardNo string `json:"idCardNo" dc:"身份证号"`
	Mobile   string `json:"mobile" dc:"手机号"`
	BankCard string `json:"bankCard" dc:"银行卡号"`
}

// 个人要素比对 Res
type VerifyPersonResBody struct {
	Result   int    `json:"result" dc:"认证结果：1.认证成功 2.认证失败 3.未认证"`
	SerialNo string `json:"serialNo" dc:"比对流水号，可结合（添加个人用户接口（V2））使用"`
	Type     string `json:"type" dc:"认证类型"`
}

/// 企业相关

// 企业三要素比对 Req https://web.asign.cn/platform/openDoc/docDetail?mid=bizInfo
type VerifyCompanyEnt3ReqBody struct {
	CompanyName string `json:"companyName" dc:"企业名称"`
	CreditCode  string `json:"creditCode" dc:"社会统一信用代码"`
	RealName    string `json:"realName" dc:"法人姓名"`
}

// 企业四要素比对 Req https://web.asign.cn/platform/openDoc/docDetail?mid=ent4
type VerifyCompanyEnt4ReqBody struct {
	CompanyName string `json:"companyName" dc:"企业名称"`
	CreditCode  string `json:"creditCode" dc:"社会统一信用代码"`
	IdCardNo    string `json:"idCardNo" dc:"法人身份证号"`
	RealName    string `json:"realName" dc:"法人姓名"`
}

// 企业要素比对 Res
type VerifyCompanyResBody struct {
	Result   int    `json:"result" dc:"认证结果：1.成功 2.失败"`
	SerialNo string `json:"serialNo" dc:"比对流水号，可结合（添加企业用户接口（V2））使用"`
	Type     string `json:"type" dc:"认证类型"`
}
