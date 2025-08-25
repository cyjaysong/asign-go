package model

// 创建待签署文件 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=createContract
type ContractCreateReqBody struct {
	ContractNo           string                    `json:"contractNo" dc:"合同唯一编码(40位之内)"`
	ContractName         string                    `json:"contractName" dc:"合同名称(120位之内),文件名称不可包含有以下特殊字符:*"`
	ValidityTime         int                       `json:"validityTime,omitempty" dc:"合同签署剩余天数(系统当前时间+该天数=在此日期之前可以签署合同日期)"`
	ValidityDate         string                    `json:"validityDate,omitempty" dc:"合同有效截止日期(在此日期之前可以签署合同)"`
	SignOrder            int                       `json:"signOrder" dc:"签约方式:1:无序签约(默认)2:顺序签约"`
	FileIds              []string                  `json:"fileIds,omitempty" dc:"上传文件的唯一标识(与合同模板、合同附件必传其一)"`
	Templates            []*ContractCreateTemplate `json:"templates,omitempty" dc:"合同模板列表(与合同附件、上传文件的唯一标识必传其一)"`
	ReadSeconds          int                       `json:"readSeconds,omitempty" dc:"强制阅读时间(秒)"`
	ReadType             int                       `json:"readType,omitempty" dc:"强制阅读设置"`
	NeedAgree            int                       `json:"needAgree,omitempty" dc:"同意协议开关,开启后表示必须同意协议才可签署合同 1:开,0:关(默认)"`
	AutoExpand           int                       `json:"autoExpand,omitempty" dc:"多文件时,是否自动展开文件列表 1:展开, 0:不展开(默认)"`
	NotifyUrl            string                    `json:"notifyUrl,omitempty" dc:"合同签署完成后回调通知地址"`
	CallbackUrl          string                    `json:"callbackUrl,omitempty" dc:"合同状态status=3:过期,4:拒签,-3:失败,回调通知地址"`
	UserNotifyUrl        string                    `json:"userNotifyUrl,omitempty" dc:"某个用户签署完成之后回调地址"`
	RedirectUrl          string                    `json:"redirectUrl,omitempty" dc:"合同签署完成后同步回调地址"`
	RefuseOn             int                       `json:"refuseOn,omitempty" dc:"合同签署页退回按钮开关:1:开启,0:关闭(默认)"`
	AutoContinue         int                       `json:"autoContinue,omitempty" dc:"当前签署人签署完成自动跳转至下一签署人签署开关,1:开启,0:关闭(默认)"`
	ViewFlg              int                       `json:"viewFlg,omitempty" dc:"合同签署完是否允许可以通过链接查看合同内容,1:不允许查看,不传值:可以查看(默认)"`
	RedirectReturnUrl    string                    `json:"redirectReturnUrl,omitempty" dc:"合同发起页面返回按钮跳转url"`
	RedirectCompletedUrl string                    `json:"redirectCompletedUrl,omitempty" dc:"合同发起页面完成后跳转url"`
	StoreCloudSign       string                    `json:"storeCloudSign,omitempty" dc:"私有云项目标识(不传默认存储公有云)"`
	EnableDownloadButton int                       `json:"enableDownloadButton,omitempty" dc:"签署页面直接下载,1:是,0:否(默认)"`
}

type ContractCreateTemplate struct {
	TemplateNo    string                       `json:"templateNo,omitempty" dc:"合同模板编号"`
	FileName      string                       `json:"fileName,omitempty" dc:"文件名称(自定义模板文件的展示名称,传值不包含文件后缀)"`
	ContractNo    string                       `json:"contractNo,omitempty" dc:"合同编号(此处可传已完成签署的合同编号,实现追加签章的场景)"`
	FillData      map[string]string            `json:"fillData,omitempty" dc:"单行文本、多行文本、日期、身份证类型参数填充"`
	ComponentData []*ContractTemplateComponent `json:"componentData,omitempty" dc:"单选、复选、勾选、下拉选择、图片类型参数填充"`
	TableDatas    []*ContractTemplateTable     `json:"tableDatas,omitempty" dc:"表格填充数据"`
}

type ContractTemplateComponent struct {
	Type         int               `json:"type" dc:"组件类型:2:单选,3:勾选,9:复选,11:图片,16:下拉选择"`
	Keyword      string            `json:"keyword" dc:"参数名称"`
	DefaultValue string            `json:"defaultValue,omitempty" dc:"当填充类型为勾选(type=3)时填写:Yes:选中 Off:不选中"`
	Options      []*ContractOption `json:"options,omitempty" dc:"选项内容"`
	ImageBytes   []byte            `json:"imageByte,omitempty" dc:"图片资源Bytes"`
	ImageBase64  string            `json:"imageBase64,omitempty" dc:"图片资源Base64"`
}

type ContractOption struct {
	Index    int  `json:"index" dc:"下标:从0开始(选项的位置)"`
	Selected bool `json:"selected" dc:"选中标记(true为选中)"`
}

type ContractTemplateTable struct {
	Keyword   string              `json:"keyword" dc:"表名称"`
	RowValues []*ContractTableRow `json:"rowValues,omitempty" dc:"表格填充数据"`
}

type ContractTableRow struct {
	InsertRow bool     `json:"insertRow,omitempty" dc:"是否新增行,默认false"`
	ColValues []string `json:"colValues,omitempty" dc:"动态表格行填充内容,第一列至最后一列依次的填充"`
}

// 创建待签署文件 Res
type ContractCreateResBody struct {
	PreviewUrl    string                    `json:"previewUrl" dc:"生成默认印章编号"`
	ContractFiles []*CreateContractResFiles `json:"contractFiles,omitempty" dc:"合同文件信息"`
}

type CreateContractResFiles struct {
	FileName string `json:"fileName" dc:"文件名称"`
	AttachNo int    `json:"attachNo" dc:"附件编号"`
	Page     int    `json:"page" dc:"	页数"`
}

// 添加签署方 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=addSigner
type ContractAddSignerReqBody []*ContractAddSigner

type ContractAddSigner struct {
	ContractNo               string                  `json:"contractNo" dc:"合同唯一编码(40位之内)"`
	Account                  string                  `json:"account" dc:"用户唯一识别码"`
	SignType                 int                     `json:"signType,omitempty" dc:"签约方式:2:无感知签约(需要开通权限)3:有感知签约"`
	SealNo                   string                  `json:"sealNo,omitempty" dc:"印章编号,若不传值,则由当前主体的默认印章进行签署"`
	AuthSignAccount          string                  `json:"authSignAccount,omitempty" dc:"指定授权签约用户"`
	NoticeMobile             string                  `json:"noticeMobile,omitempty" dc:"通知手机号"`
	SignOrder                string                  `json:"signOrder,omitempty" dc:"使用顺序签约时签约顺序编号(从1开始),无序签约都为1"`
	IsNotice                 int                     `json:"isNotice,omitempty" dc:"是否接收合同签署链接的短信通知"`
	ValidateType             int                     `json:"validateType,omitempty" dc:"签署方式指定"`
	FaceAuthMode             int                     `json:"faceAuthMode,omitempty" dc:"人脸识别方式"`
	ValidateTypeList         string                  `json:"validateTypeList,omitempty" dc:"组合签署方式指定"`
	AutoSwitch               int                     `json:"autoSwitch,omitempty" dc:"自动切换签约方式"`
	IsNoticeComplete         int                     `json:"isNoticeComplete,omitempty" dc:"合同签署完成后是否通知用户,1:是,0:否(默认)"`
	WaterMark                int                     `json:"waterMark,omitempty" dc:"是否在距底部10px中央位置添加日期水印,1:是,0:否(默认)"`
	AutoSms                  int                     `json:"autoSms,omitempty" dc:"是否自动触发验证码短信,仅短信验证码方式签署时生效"`
	CustomSignFlag           int                     `json:"customSignFlag,omitempty" dc:"签章位置策略"`
	SignStrategyList         []*SignStrategy         `json:"signStrategyList,omitempty" dc:"签章位置策略"`
	SignStrikeList           []*SignStrike           `json:"signStrikeList,omitempty" dc:"骑缝章策略"`
	ReceiverFillStrategyList []*ReceiverFillStrategy `json:"receiverFillStrategyList,omitempty" dc:"接收方模板填充策略"`
	AuthConfig               *SignerAuthConfig       `json:"authConfig,omitempty" dc:"接收方模板填充策略"`
	IsIframe                 int                     `json:"isIframe,omitempty" dc:"H5人脸需要开启无Cookie模式，或使用IFrame框架"`
	WillType                 string                  `json:"willType,omitempty" dc:"指定视频核身模式"`
	SignMark                 string                  `json:"signMark,omitempty" dc:"业务系统传递的唯一标识(自定义业务编号)"`
	ChangeSeal               int                     `json:"changeSeal,omitempty" dc:"企业更换印章,1:更换,0:不更换(默认)"`
}

// 签章策略
type SignStrategy struct {
	AttachNo     int     `json:"attachNo,omitempty" dc:"附件编号"`
	LocationMode int     `json:"locationMode,omitempty" dc:"定位方式"`
	SealNo       int     `json:"sealNo,omitempty" dc:"印章编号"`
	CanDrag      int     `json:"canDrag,omitempty" dc:"签章位置是否可以拖动,1:可以,其他值:不可以"`
	SignKey      string  `json:"signKey,omitempty" dc:"关键字或签署区名称key"`
	SignType     int     `json:"signType,omitempty" dc:"印章类型,1:签名/签章(默认),2:时间戳"`
	SignPage     int     `json:"signPage,omitempty" dc:"签章页码(定位方式为坐标签章时必传)"`
	SignX        float64 `json:"signX,omitempty" dc:"签章位置与当前签约文件的左内边距与当前签约文件宽度的比例(精确到小数点后2位)(定位方式为坐标签章时必传)"`
	SignY        float64 `json:"signY,omitempty" dc:"签章位置与当前签约文件的上内边距与当前签约文件高度的比例(精确到小数点后2位)(定位方式为坐标签章时必传)"`
	OffsetX      float64 `json:"offsetX,omitempty" dc:"坐标偏移量(像素PX)"`
	OffsetY      float64 `json:"offsetY,omitempty" dc:"坐标偏移量(像素PX)"`
}

// 骑缝章策略
type SignStrike struct {
	AttachNo   int     `json:"attachNo" dc:"附件编号"`
	CrossStyle int     `json:"crossStyle" dc:"骑缝类型,1:上下各一半,2:左右各一半,3:上下骑缝,4:左右骑缝,5:左骑缝,6:右骑缝"`
	SignPage   string  `json:"signPage,omitempty" dc:"签约页码范围(不能超过合同总页数)"`
	OddEven    int     `json:"oddEven,omitempty" dc:"1:奇数,2:偶数(仅奇数页或偶数页盖骑缝章)"`
	OffsetX    float64 `json:"offsetX,omitempty" dc:"X坐标偏移量(像素PX，骑缝章是相对于中间位置的偏移量，非骑缝章相对于当前盖章位置的偏移量)"`
	OffsetY    float64 `json:"offsetY,omitempty" dc:"Y坐标偏移量(像素PX，骑缝章是相对于中间位置的偏移量，非骑缝章相对于当前盖章位置的偏移量)"`
	SignKey    string  `json:"signKey,omitempty" dc:"签署区名称key(定位方式为模板坐标签章时此处需传模板中设置的签署区名称)"`
	ClipNumber int     `json:"clipNumber,omitempty" dc:"图章指定切割等份数量如果页面数量大于切割数量,印章将会重复,单位份(默认以页数，为0)"`
}

// 接收方模板填充策略
type ReceiverFillStrategy struct {
	AttachNo  int                 `json:"attachNo" dc:"附件编号"`
	Key       string              `json:"key" dc:"参数名称"`
	Value     string              `json:"value,omitempty" dc:"填充值"`
	FillStage int                 `json:"fillStage,omitempty" dc:"填充场景标记"`
	Options   []*ContractOption   `json:"options,omitempty" dc:"单、复选选项内容"`
	RowValues []*ContractTableRow `json:"rowValues,omitempty" dc:"表格填充值"`
}

// 添加陌生签署人认证参数配置
type SignerAuthConfig struct {
	AuthTypeList        []int `json:"authTypeList,omitempty" dc:"选择显示给个人/办理人的页面认证方式"`
	CompanyAuthTypeList []int `json:"companyAuthTypeList,omitempty" dc:"选择显示的公司认证方式"`
	ShowAuthPage        int   `json:"showAuthPage" dc:"是否展示授权页,1:是(默认),0:否"`
	AgentSignPower      bool  `json:"agentSignPower,omitempty" dc:"是否授权经办人签署"`
	FaceAuthMode        int   `json:"faceAuthMode,omitempty" dc:"人脸识别方式"`
	NeedSignPwd         int   `json:"needSignPwd,omitempty" dc:"实名认证时是否采集签约密码参数,1:是,0:否(默认)"`
}

// 添加签署方 Res
type ContractAddSignerResBody struct {
	ContractNo   string          `json:"contractNo" dc:"合同编号"`
	ContractName string          `json:"contractName" dc:"合同名称"`
	ValidityTime string          `json:"validityTime" dc:"合同有效期"`
	PreviewUrl   string          `json:"previewUrl" dc:"合同预览链接"`
	SignUSer     []*SignUserItem `json:"signUSer,omitempty" dc:"合同用户信息"`
}

type SignUserItem struct {
	Account          string                 `json:"account" dc:"用户唯一识别码"`
	SignUrl          string                 `json:"signUrl" dc:"合同签署链接"`
	PwdSignUrl       string                 `json:"pwdSignUrl" dc:"密码签署链接"`
	SignOrder        int                    `json:"signOrder,omitempty" dc:"顺序签约的序号"`
	Name             string                 `json:"name" dc:"用户姓名"`
	IdCard           string                 `json:"idCard" dc:"用户身份证"`
	SignMark         string                 `json:"signMark,omitempty" dc:"业务系统传递的唯一标识(自定义业务编号)"`
	Mobile           string                 `json:"mobile,omitempty" dc:"手机号(签约短信通知手机号)"`
	CompanyName      string                 `json:"companyName,omitempty" dc:"企业名称"`
	UserType         int                    `json:"userType,omitempty" dc:"用户类型,1:企业,2:个人"`
	SealNo           string                 `json:"sealNo,omitempty" dc:"印章编号"`
	SignType         int                    `json:"signType,omitempty" dc:"签约类型,1:代签,2:无感知签,3:有感知签"`
	ValidateType     int                    `json:"validateType,omitempty" dc:"签约校验方式,1:验证码,2:签约密码,3:人脸识别,4:手写"`
	SignStatus       int                    `json:"signStatus,omitempty" dc:"签约状态,1:签约中,2:已签约,3:待签约,4:拒签"`
	SignFinishedTime string                 `json:"signFinishedTime,omitempty" dc:"用户签约完成时间"`
	ContractViewHis  []*ContractViewHisItem `json:"contractViewHis,omitempty" dc:"访问日志"`
}

type ContractViewHisItem struct {
	AccessTime string `json:"accessTime" dc:"访问时间"`
	Detail     string `json:"detail" dc:"访问详情"`
	Opened     int    `json:"opened" dc:"是否签署,1:已签署,0:未签署"`
}

// 合同撤销 Req https://web.asign.cn/platform/openDoc/docDetail?mid=withdraw
type WithdrawContractReqBody struct {
	ContractNo       string `json:"contractNo" dc:"合同唯一编码"`
	WithdrawReason   string `json:"withdrawReason" dc:"撤销原因，最多50字"`
	IsNoticeSignUser bool   `json:"isNoticeSignUser" dc:"是否短信通知签署用户，true 通知false不通知，默认不通知"`
}
type WithdrawContractResBody struct{}

// 查询合同状态 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=status
type GetContractStatusReqBody struct {
	ContractNo string `json:"contractNo" dc:"合同唯一编码"`
}
type GetContractStatusResBody struct {
	ContractNo   string `json:"contractNo" dc:"合同唯一编码"`
	ContractName string `json:"contractName" dc:"合同名称"`
	Status       int    `json:"status" dc:"合同状态,0:等待签约,1:签约中,2:已签约,3:过期,4:拒签,6:作废,7:撤销"`
}

// 查询合同信息 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=getContract
type GetContractReqBody struct {
	ContractNo string `json:"contractNo" dc:"合同唯一编码"`
}
type GetContractResBody struct {
	ContractNo       string          `json:"contractNo" dc:"合同唯一编码"`
	ContractName     string          `json:"contractName" dc:"合同名称"`
	Status           int             `json:"status" dc:"合同状态,0:等待签约,1:签约中,2:已签约,3:过期,4:拒签,6:作废,7:撤销"`
	ValidityTime     string          `json:"validityTime" dc:"合同有效期"`
	SignFinishedTime string          `json:"signFinishedTime,omitempty" dc:"合同签约完成时间"`
	PreviewUrl       string          `json:"previewUrl" dc:"合同预览链接"`
	EmbeddedUrl      string          `json:"embeddedUrl" dc:"合同嵌入式拖章链接"`
	Remark           string          `json:"remark" dc:"退回/拒签原因"`
	SignUsers        []*SignUserItem `json:"signUser" dc:"合同用户信息"`
}

// 下载合同 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=downloadContract
type DownloadContractReqBody struct {
	ContractNo       string `json:"contractNo" dc:"合同唯一编码"`
	Outfile          string `json:"outfile,omitempty" dc:"文件本地路径(下载的文件，写到本地，当文件数为1时，下载的文件类型是pdf，否则为zip)"`
	Force            int    `json:"force,omitempty" dc:"强制下载标识,0(默认):未签署完的无法下载,1:无论什么状态都强制下载"`
	DownloadFileType int    `json:"downloadFileType,omitempty" dc:"下载文件类型,多附件下载"`
}
type DownloadContractResBody struct {
	FileName string `json:"fileName" dc:"文件名"`
	MD5      string `json:"md5" dc:"文件md5值"`
	FileType int    `json:"fileType" dc:"文件类型,0:PDF,1:ZIP"`
	Size     int64  `json:"size" dc:"文件大小"`
	Data     string `json:"data" dc:"Base64字符串"`
}
