package model

// 添加个人用户V2 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=addPersonalUserNew
type AddPersonalUserV2ReqBody struct {
	Account         string `json:"account" dc:"用户唯一识别码（可用证件号、手机号等具有唯一属性的标识作为参数传递）"`
	SerialNo        string `json:"serialNo,omitempty" dc:"实名认证流水号（若实名认证服务由爱签平台提供，此值必传）"`
	Name            string `json:"name,omitempty" dc:"用户姓名"`
	IdCard          string `json:"idCard,omitempty" dc:"个人身份证、台胞证、港澳通行证等证件号；【注】此值是否需要传参，详见【特别说明】"`
	IdCardType      int    `json:"idCardType,omitempty" dc:"证件类型，不传默认为1；1：居民身份证；2：台湾居民来往大陆通行证；3：港澳居民来往内地通行证等...；更多类型见详细说明；【注】当选择非居民身份证类型时，爱签平台无法提供实名认证服务。该用户的实名认证由调用方，在同步用户之前自行完成。"`
	Mobile          string `json:"mobile,omitempty" dc:"手机号码（用于接收合同签署时的短信验证码）【注】此值若不传，签署意愿验证将无法使用短信验证码方式。"`
	SignPwd         string `json:"signPwd,omitempty" dc:"签约密码明文，如果为空我方将随机生成签约密码（当签约方式为'签约密码签约'时会使用到，可通过重置接口修改）"`
	IsSignPwdNotice int    `json:"isSignPwdNotice,omitempty" dc:"是否将签约密码以短信形式通知用户：0 - 不通知（默认），1 - 通知"`
	IsNotice        int    `json:"isNotice,omitempty" dc:"用户发起合同或需要签署时是否进行短信通知：0 - 否（默认），1 - 是"`
	SealPic         string `json:"sealPic,omitempty" dc:"base64格式的印模图片"`
}

// 添加个人用户V2 Res
type AddPersonalUserV2ResBody struct {
	SealNo string `json:"sealNo" dc:"生成默认印章编号"`
}

// 添加企业用户V2 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=addEnterpriseUserNew
type AddEnterpriseUserV2ReqBody struct {
	Account         string `json:"account" dc:"用户唯一识别码（可用证件号、手机号等具有唯一属性的标识作为参数传递）"`
	SerialNo        string `json:"serialNo,omitempty" dc:"实名认证流水号（若实名认证服务由爱签平台提供，此值必传）"`
	CompanyName     string `json:"companyName,omitempty" dc:"企业名称"`
	CreditCode      string `json:"creditCode,omitempty" dc:"企业证件号"`
	CreditType      int    `json:"creditType,omitempty" dc:"企业证件类型，不传默认为1；1：统一社会信用代码；2：表示其他证件类型"`
	Name            string `json:"name,omitempty" dc:"企业法人姓名"`
	IdCard          string `json:"idCard,omitempty" dc:"法人身份证、台胞证、港澳通行证等证件号；【注】此值是否需要传参，详见【特别说明】"`
	IdCardType      int    `json:"idCardType,omitempty" dc:"证件类型，不传默认为1；1：居民身份证；2：台湾居民来往大陆通行证；3：港澳居民来往内地通行证等...；更多类型见详细说明；【注】当选择非居民身份证类型时，爱签平台无法提供实名认证服务。该用户的实名认证由调用方，在同步用户之前自行完成。"`
	Mobile          string `json:"mobile,omitempty" dc:"签约手机，该手机号将用于企业用户合同签署时短信验证，请确保真实有效；【注】此值若不传，签署意愿验证将无法使用短信验证码方式。"`
	ContactName     string `json:"contactName,omitempty" dc:"联系人姓名（与企业的法定代表人可以是同一个人）"`
	ContactIdCard   string `json:"contactIdCard,omitempty" dc:"联系人身份证号（与企业的法定代表人可以是同一个人）"`
	SignPwd         string `json:"signPwd,omitempty" dc:"签约密码明文，如果为空我方将随机生成签约密码（当签约方式为'签约密码签约'时会使用到，可通过重置接口修改）"`
	IsSignPwdNotice int    `json:"isSignPwdNotice,omitempty" dc:"是否将签约密码以短信形式通知用户：0 - 不通知（默认），1 - 通知"`
	IsNotice        int    `json:"isNotice,omitempty" dc:"用户发起合同或需要签署时是否进行短信通知：0 - 否（默认），1 - 是"`
}

// 添加企业用户V2 Res
type AddEnterpriseUserV2ResBody struct {
	SealNo string `json:"sealNo" dc:"生成默认印章编号"`
}

// 查询用户信息 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=getUser
type GetUserReqBody struct {
	Account    string `json:"account,omitempty" dc:"用户唯一识别码"`
	CreditCode string `json:"creditCode,omitempty" dc:"社会统一信用代码"`
	IdCard     string `json:"idCard,omitempty" dc:"证件号码"`
}

// 查询用户信息 Res
type GetUserResBody struct {
	Account      string `json:"account" dc:"用户账号，用户唯一识别码"`
	Name         string `json:"name" dc:"个人用户姓名/企业法人姓名"`
	CompanyName  string `json:"companyName,omitempty" dc:"企业名称"`
	IdCard       string `json:"idCard" dc:"个人用户证件号/企业法人身份证号"`
	Mobile       string `json:"mobile" dc:"用户手机号（签约短信通知手机号）"`
	Email        string `json:"email,omitempty" dc:"用户邮箱号"`
	UserType     int    `json:"userType" dc:"用户类型：1：企业；2：个人"`
	CreditCode   string `json:"creditCode,omitempty" dc:"社会统一信用代码"`
	BankCard     string `json:"bankCard,omitempty" dc:"用户银行卡号"`
	PortVersion  int    `json:"portVersion,omitempty" dc:"用户添加时调用的接口版本：0：历史接口；1：V2版本接口"`
	IdentifyType int    `json:"identifyType,omitempty" dc:"认证类型"` // https://preweb.asign.cn/platform/openDoc/docDetail?mid=getUser#identifyTypeId
	AuthType     int    `json:"authType,omitempty" dc:"0:非强制认证/平台方自行认证 1:爱签平台强制认证/爱签平台认证"`
	CreateTime   string `json:"createTime,omitempty" dc:"创建时间"`
	IdentifyTime string `json:"identifyTime,omitempty" dc:"认证时间"`
	Status       int    `json:"status" dc:"用户状态：0：未激活; 1：有效; 2：锁定"`
	SerialNo     string `json:"serialNo,omitempty" dc:"实名认证流水号"`
}

// 查询企业信息 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=getCompUser
type GetCompUserReqBody struct {
	Account     string `json:"account,omitempty" dc:"用户唯一识别码"`
	CreditCode  string `json:"creditCode,omitempty" dc:"社会统一信用代码"`
	CompanyName string `json:"companyName,omitempty" dc:"企业名称"`
}

// 查询企业信息 Res
type GetCompUserResBody []GetCompUserResBodyItem

type GetCompUserResBodyItem struct {
	Account      string `json:"account" dc:"用户账号，用户唯一识别码"`
	Name         string `json:"name" dc:"个人用户姓名/企业法人姓名"`
	CompanyName  string `json:"companyName,omitempty" dc:"企业名称"`
	IdCard       string `json:"idCard" dc:"个人用户证件号/企业法人身份证号"`
	UserType     int    `json:"userType" dc:"用户类型：1：企业；2：个人"`
	Status       int    `json:"status" dc:"用户状态：0：未激活; 1：有效; 2：锁定"`
	CreditCode   string `json:"creditCode,omitempty" dc:"社会统一信用代码"`
	BankCard     string `json:"bankCard,omitempty" dc:"用户银行卡号"`
	PortVersion  int    `json:"portVersion,omitempty" dc:"用户添加时调用的接口版本：0：历史接口；1：V2版本接口"`
	IdentifyType int    `json:"identifyType,omitempty" dc:"认证类型"` // https://preweb.asign.cn/platform/openDoc/docDetail?mid=getUser#identifyTypeId
	AuthType     int    `json:"authType,omitempty" dc:"0:非强制认证/平台方自行认证 1:爱签平台强制认证/爱签平台认证"`
	CreateTime   string `json:"createTime,omitempty" dc:"创建时间"`
	IdentifyTime string `json:"identifyTime,omitempty" dc:"认证时间"`
	SerialNo     string `json:"serialNo,omitempty" dc:"实名认证流水号"`
}

// 用户绑定手机号(运营商三要素方案) Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=modifyMobile
type UserModifyMobileReqBody struct {
	Account string `json:"account" dc:"用户唯一识别码"`
	Name    string `json:"name" dc:"用户姓名"`
	IdCard  string `json:"idCard" dc:"个人为本人身份证号/企业为法定代表人身份证号"`
	Mobile  string `json:"mobile" dc:"用户新手机号"`
}

// 用户绑定手机号(运营商三要素方案) Res
type UserModifyMobileResBody struct{}

// 发送用户绑定手机号验证码 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=sendCode
type SendUserModifyMobileCodeReqBody struct {
	Account string `json:"account" dc:"用户唯一识别码"`
}

// 发送用户绑定手机号验证码 Res
type SendUserModifyMobileCodeResBody struct {
	CodeToken string `json:"codeToken" dc:"验证码唯一标识"`
}

// 用户绑定手机号(验证码方案) Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=modifyMobileByCode
type UserModifyMobileByCodeReqBody struct {
	Account string `json:"account" dc:"用户唯一识别码"`
	Mobile  string `json:"mobile" dc:"用户新手机号"`
	Code    string `json:"code" dc:"用户接收到的验证码"`
	CToken  string `json:"ctoken" dc:"调用发送短信接口返回的codeToken"`
}

// 用户绑定手机号(验证码方案) Res
type UserModifyMobileByCodeResBody struct{}

// 修改企业信息 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=modifyCompanyInfo
type ModifyCompanyInfoReqBody struct {
	Account     string `json:"account" dc:"用户唯一识别码"`
	CompanyName string `json:"companyName,omitempty" dc:"企业名称（修改企业名称时必传）"`
	Name        string `json:"name,omitempty" dc:"法人名称（修改法人信息时必传）"`
	IdCard      string `json:"idCard,omitempty" dc:"法人身份证（修改法人信息时必传）"`
	IdCardType  int    `json:"idCardType,omitempty" dc:"法人证件类型：（修改法人信息时必传）"`
}

// 修改企业信息 Res
type ModifyCompanyInfoResBody string // 当修改企业名称时，默认会生成一个新的印章。返回参数为新的印章编号。

// 修改个人信息 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=modifyUserName
type ModifyUserNameReqBody struct {
	Account      string `json:"account" dc:"用户唯一识别码"`
	Name         string `json:"name" dc:"姓名"`
	IdentifyType int    `json:"identifyType" dc:"认证类型：2：运营商三要素认证; 3：银行卡四要素认证"`
	Mobile       string `json:"mobile,omitempty" dc:"用户手机号（identifyType=2时必传）"`
	BankCard     string `json:"bankCard,omitempty" dc:"用户银联卡号（identifyType=3时必传）"`
}

// 修改个人信息 Res
type ModifyUserNameResBody string

// 用户重新认证 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=reAuthUser
type UserReauthReqBody struct {
	Account       string `json:"account" dc:"用户唯一识别码"`
	SerialNo      string `json:"serialNo" dc:"实名认证流水号"`
	ContactName   string `json:"contactName,omitempty" dc:"联系人姓名（与企业的法定代表人可以是同一个人）"`
	ContactIdCard string `json:"contactIdCard,omitempty" dc:"联系人身份证号（与企业的法定代表人可以是同一个人）"`
}

// 用户重新认证 Res
type UserReauthResBody string

// 用户删除 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=userRemove
type UserDeleteReqBody struct {
	Account string `json:"account" dc:"用户唯一识别码"`
}

// 用户删除 Res
type UserDeleteResBody string
