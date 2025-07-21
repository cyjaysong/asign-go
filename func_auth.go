package asign

import (
	"context"
	"github.com/cyjaysong/asign-go/model"
)

// 重新发送认证验证码
func (the Client) AuthCaptchaResend(ctx context.Context, req *model.AuthCaptchaResendReqBody) (res *model.BaseRes[model.AuthCaptchaResendResBody], err error) {
	path := "/auth/captcha/resend"
	res = new(model.BaseRes[model.AuthCaptchaResendResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 认证验证码校验
func (the Client) AuthCaptchaVerify(ctx context.Context, req *model.AuthCaptchaVerifyReqBody) (res *model.BaseRes[model.AuthCaptchaVerifyResBody], err error) {
	path := "/auth/captcha/verify"
	res = new(model.BaseRes[model.AuthCaptchaVerifyResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 实名认证流水号查询
func (the Client) GetAuthSerialNo(ctx context.Context, req *model.GetAuthSerialNoReqBody) (res *model.BaseRes[model.GetAuthSerialNoResBody], err error) {
	path := "/auth/getSerialNo"
	res = new(model.BaseRes[model.GetAuthSerialNoResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 实名认证信息查询
func (the Client) GetAuthRecordInfo(ctx context.Context, req *model.GetAuthRecordInfoReqBody) (res *model.BaseRes[model.GetAuthRecordInfoResBody], err error) {
	path := "/auth/getAuthRecordInfo"
	res = new(model.BaseRes[model.GetAuthRecordInfoResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 个人运营商三要素认证
func (the Client) AuthPersonMobile3(ctx context.Context, req *model.AuthPersonMobile3ReqBody) (res *model.BaseRes[model.AuthPersonMobile3ResBody], err error) {
	path := "/auth/person/mobile3"
	res = new(model.BaseRes[model.AuthPersonMobile3ResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 企业法人运营商三要素认证
func (the Client) AuthCompanyMobile3(ctx context.Context, req *model.AuthCompanyMobile3ReqBody) (res *model.BaseRes[model.AuthCompanyMobile3ResBody], err error) {
	path := "/auth/company/mobile3"
	res = new(model.BaseRes[model.AuthCompanyMobile3ResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}
