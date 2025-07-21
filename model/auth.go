package model

// 重新发送认证验证码 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=captchaResend
type AuthCaptchaResendReqBody struct {
	SerialNo string `json:"serialNo" dc:"认证流水号"`
}

// 重新发送认证验证码 Res
type AuthCaptchaResendResBody struct{}

// 认证验证码校验 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=captchaVerify
type AuthCaptchaVerifyReqBody struct {
	SerialNo string `json:"serialNo" dc:"认证流水号"`
	Captcha  string `json:"captcha" dc:"短信验证码"`
}

// 认证验证码校验 Res
type AuthCaptchaVerifyResBody struct {
	Result   int    `json:"result" dc:"认证结果：1.成功 2.失败"`
	SerialNo string `json:"serialNo" dc:"认证流水号"`
	Type     string `json:"type" dc:"认证类型"`
}

// 实名认证流水号查询 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=getSerialNo
type GetAuthSerialNoReqBody struct {
	UserType int    `json:"userType" dc:"认证主体类型：1.企业，2.个人"`
	UserName string `json:"userName,omitempty" dc:"姓名 或 企业名称"`
	IdNo     string `json:"idNo,omitempty" dc:"身份证 或 企业证件号"`
}

// 实名认证流水号查询 Res
type GetAuthSerialNoResBody []GetAuthSerialNoResBodyItem

type GetAuthSerialNoResBodyItem struct {
	SerialNo     string `json:"serialNo" dc:"认证流水号"`
	Result       int    `json:"result" dc:"认证结果：1.成功 2.失败"`
	UserType     int    `json:"userType" dc:"认证主体类型：1.企业，2.个人"`
	AuthTypeCode int    `json:"authTypeCode" dc:"认证类型码"`
	AuthTypeName string `json:"authTypeName" dc:"认证类型"`
}

// 实名认证信息查询 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=getAuthRecordInfo
type GetAuthRecordInfoReqBody struct {
	SerialNo string `json:"serialNo" dc:"认证流水号"`
}

// 实名认证信息查询 Res
type GetAuthRecordInfoResBody struct {
	SerialNo          string `json:"serialNo" dc:"认证流水号"`
	UserType          int    `json:"userType" dc:"认证主体类型：1.企业，2.个人"`
	Result            int    `json:"result" dc:"认证结果：1.成功 2.失败"`
	RealName          string `json:"realName,omitempty" dc:"个人姓名 / 法人姓名"`
	IdCardNo          string `json:"idCardNo,omitempty" dc:"个人身份证 / 法人身份证"`
	Mobile            string `json:"mobile,omitempty" dc:"个人手机号 / 法人手机号"`
	CompanyName       string `json:"companyName,omitempty" dc:"企业名称"`
	CreditCode        string `json:"creditCode,omitempty" dc:"社会统一信用代码"`
	AuthTypeCode      string `json:"authTypeCode,omitempty" dc:"认证类型码"`
	AuthTypeName      string `json:"authTypeName,omitempty" dc:"认证类型"`
	AgentResult       string `json:"agentResult,omitempty" dc:"办理人认证结果：0.认证中 1.认证成功 2.认证失败"`
	AgentName         string `json:"agentName,omitempty" dc:"办理人姓名"`
	AgentIdCardNo     string `json:"agentIdCardNo,omitempty" dc:"办理人身份证"`
	AgentMobile       string `json:"agentMobile,omitempty" dc:"办理人手机号"`
	AgentAuthTypeCode int    `json:"agentAuthTypeCode,omitempty" dc:"办理人认证类型码"`
	AgentAuthTypeName string `json:"agentAuthTypeName,omitempty" dc:"办理人认证类型"`
}

// 个人运营商三要素认证 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=mobile3
type AuthPersonMobile3ReqBody struct {
	RealName string `json:"realName" dc:"真实姓名"`
	IdCardNo string `json:"idCardNo" dc:"身份证号"`
	Mobile   string `json:"mobile" dc:"手机号码（限中国大陆11位手机号）"`
}

// 个人运营商三要素认证 Res
type AuthPersonMobile3ResBody struct {
	Result   int    `json:"result" dc:"认证结果：0.暂无结果/认证中 1.成功 2.失败"`
	SerialNo string `json:"serialNo" dc:"认证流水号"`
	Type     string `json:"type" dc:"认证类型"`
}

// 企业法人运营商三要素认证 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=companyMobile3
type AuthCompanyMobile3ReqBody struct {
	CompanyName string `json:"companyName" dc:"企业名称"`
	CreditCode  string `json:"creditCode" dc:"社会统一信用代码"`
	RealName    string `json:"realName" dc:"真实姓名"`
	IdCardNo    string `json:"idCardNo" dc:"身份证号"`
	Mobile      string `json:"mobile" dc:"手机号码（限中国大陆11位手机号）"`
}

// 企业法人运营商三要素认证 Res
type AuthCompanyMobile3ResBody struct {
	Result   int    `json:"result" dc:"认证结果：0.暂无结果/认证中 1.成功 2.失败"`
	SerialNo string `json:"serialNo" dc:"认证流水号"`
	Type     string `json:"type" dc:"认证类型"`
}
