package asign

import (
	"context"
	"github.com/cyjaysong/asign-go/model"
)

// 添加个人用户（V2）
func (the Client) AddPersonalUserV2(ctx context.Context, req *model.AddPersonalUserV2ReqBody) (res *model.BaseRes[model.AddPersonalUserV2ResBody], err error) {
	path := "/v2/user/addPersonalUser"
	res = new(model.BaseRes[model.AddPersonalUserV2ResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 添加企业用户（V2）
func (the Client) AddEnterpriseUserV2(ctx context.Context, req *model.AddEnterpriseUserV2ReqBody) (res *model.BaseRes[model.AddEnterpriseUserV2ResBody], err error) {
	path := "/v2/user/addEnterpriseUser"
	res = new(model.BaseRes[model.AddEnterpriseUserV2ResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 查询用户信息
func (the Client) GetUser(ctx context.Context, req *model.GetUserReqBody) (res *model.BaseRes[model.GetUserResBody], err error) {
	path := "/user/getUser"
	res = new(model.BaseRes[model.GetUserResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 查询企业信息
func (the Client) GetCompUser(ctx context.Context, req *model.GetCompUserReqBody) (res *model.BaseRes[model.GetCompUserResBody], err error) {
	path := "/user/getCompUser"
	res = new(model.BaseRes[model.GetCompUserResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 用户绑定手机号(运营商三要素方案)
func (the Client) UserModifyMobile(ctx context.Context, req *model.UserModifyMobileReqBody) (res *model.BaseRes[model.UserModifyMobileResBody], err error) {
	path := "/user/modifyMobile"
	res = new(model.BaseRes[model.UserModifyMobileResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 发送用户绑定手机号验证码
func (the Client) SendUserModifyMobileCode(ctx context.Context, req *model.SendUserModifyMobileCodeReqBody) (res *model.BaseRes[model.SendUserModifyMobileCodeResBody], err error) {
	path := "/user/sendCode"
	res = new(model.BaseRes[model.SendUserModifyMobileCodeResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 用户绑定手机号(验证码方案)
func (the Client) UserModifyMobileByCode(ctx context.Context, req *model.UserModifyMobileByCodeReqBody) (res *model.BaseRes[model.UserModifyMobileByCodeResBody], err error) {
	path := "/user/modifyMobileByCode"
	res = new(model.BaseRes[model.UserModifyMobileByCodeResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 修改企业信息
func (the Client) ModifyCompanyInfo(ctx context.Context, req *model.ModifyCompanyInfoReqBody) (res *model.BaseRes[model.ModifyCompanyInfoResBody], err error) {
	path := "/user/modifyCompanyInfo"
	res = new(model.BaseRes[model.ModifyCompanyInfoResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 修改企业信息
func (the Client) ModifyUserName(ctx context.Context, req *model.ModifyUserNameReqBody) (res *model.BaseRes[model.ModifyUserNameResBody], err error) {
	path := "/v2/user/modifyUserName"
	res = new(model.BaseRes[model.ModifyUserNameResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 用户重新认证
func (the Client) UserReauth(ctx context.Context, req *model.UserReauthReqBody) (res *model.BaseRes[model.UserReauthResBody], err error) {
	path := "/v2/user/reAuthUser"
	res = new(model.BaseRes[model.UserReauthResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 用户删除
func (the Client) UserDelete(ctx context.Context, req *model.UserDeleteReqBody) (res *model.BaseRes[model.UserDeleteResBody], err error) {
	path := "/v2/user/remove"
	res = new(model.BaseRes[model.UserDeleteResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}
