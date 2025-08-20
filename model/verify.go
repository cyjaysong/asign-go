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

// 企业工商数据查询 Req https://web.asign.cn/platform/openDoc/docDetail?mid=enterpriseInfo
type EnterpriseInfoReqBody struct {
	CompanyName string `json:"companyName" dc:"企业名称"`
	CreditCode  string `json:"creditCode" dc:"社会统一信用代码"`
}

// 企业工商数据查询 Res
type EnterpriseInfoResBody struct {
	AbuItem             string `json:"abuItem" dc:"许可经营范围"`
	Address             string `json:"address" dc:"注册地址"`
	ApprDate            string `json:"apprDate" dc:"核准日期"`
	AreaCode            string `json:"areaCode" dc:"住所地所在行政区代码"`
	CancelDate          string `json:"cancelDate" dc:"注销日期"`
	CbuItem             string `json:"cbuItem" dc:"一般经营范围"`
	Province            string `json:"province" dc:"省"`
	City                string `json:"city" dc:"市"`
	County              string `json:"county" dc:"区县"`
	CreditCode          string `json:"creditCode" dc:"统一社会信用代码"`
	EnterpriseName      string `json:"enterpriseName" dc:"企业名称"`
	EnterpriseStatus    string `json:"enterpriseStatus" dc:"经营状态"`
	EnterpriseType      string `json:"enterpriseType" dc:"企业类型"`
	EsDate              string `json:"esDate" dc:"开业日期"`
	FrName              string `json:"frName" dc:"企业法人姓名"`
	IndustryCode        string `json:"industryCode" dc:"行业门类代码"`
	IndustryName        string `json:"industryName" dc:"国民经济行业名称"`
	IndustryPhyCode     string `json:"industryPhyCode" dc:"国民经济行业代码"`
	IndustryPhyName     string `json:"industryPhyName" dc:"行业门类名称"`
	OpenFrom            string `json:"openFrom" dc:"经营期限自"`
	OpenTo              string `json:"openTo" dc:"经营期限限至"`
	OperateScopeAndForm string `json:"operateScopeAndForm" dc:"经营范围及方式"`
	OrgCode             string `json:"orgCode" dc:"组织机构代码"`
	RegCap              string `json:"regCap" dc:"注册资金"`
	RegCapCur           string `json:"regCapCur" dc:"注册币种"`
	RegNo               string `json:"regNo" dc:"工商注册号"`
	Regorg              string `json:"regorg" dc:"登记机关"`
	RevokeDate          string `json:"revokeDate" dc:"吊销日期"`
	Orderid             string `json:"orderid" dc:"查询订单号"`
}
