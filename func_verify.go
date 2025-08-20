package asign

import (
	"context"
	"github.com/cyjaysong/asign-go/model"
)

// 个人身份证二要素比对
func (the Client) VerifyPersonIdCard2(ctx context.Context, req *model.VerifyPersonIdCard2ReqBody) (res *model.BaseRes[model.VerifyPersonResBody], err error) {
	path := "/verify/person/idcard2"
	res = new(model.BaseRes[model.VerifyPersonResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 个人运营商二要素比对
func (the Client) VerifyPersonEnt2(ctx context.Context, req *model.VerifyPersonEnt2ReqBody) (res *model.BaseRes[model.VerifyPersonResBody], err error) {
	path := "/verify/person/mobile2"
	res = new(model.BaseRes[model.VerifyPersonResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 个人运营商三要素比对
func (the Client) VerifyPersonEnt3(ctx context.Context, req *model.VerifyPersonEnt3ReqBody) (res *model.BaseRes[model.VerifyPersonResBody], err error) {
	path := "/verify/person/mobile3"
	res = new(model.BaseRes[model.VerifyPersonResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 个人银行卡三要素比对
func (the Client) VerifyPersonBank3(ctx context.Context, req *model.VerifyPersonBank3ReqBody) (res *model.BaseRes[model.VerifyPersonResBody], err error) {
	path := "/verify/person/bank3"
	res = new(model.BaseRes[model.VerifyPersonResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 个人银行卡四要素比对
func (the Client) VerifyPersonBank4(ctx context.Context, req *model.VerifyPersonBank4ReqBody) (res *model.BaseRes[model.VerifyPersonResBody], err error) {
	path := "/verify/person/bank4"
	res = new(model.BaseRes[model.VerifyPersonResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 企业三要素比对
func (the Client) VerifyCompanyEnt3(ctx context.Context, req *model.VerifyCompanyEnt3ReqBody) (res *model.BaseRes[model.VerifyCompanyResBody], err error) {
	path := "/verify/company/bizInfo"
	if req.CompanyName == "*" {
		req.CompanyName = req.CreditCode
	}
	res = new(model.BaseRes[model.VerifyCompanyResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 企业四要素比对
func (the Client) VerifyCompanyEnt4(ctx context.Context, req *model.VerifyCompanyEnt4ReqBody) (res *model.BaseRes[model.VerifyCompanyResBody], err error) {
	path := "/verify/company/ent4"
	if req.CompanyName == "*" {
		req.CompanyName = req.CreditCode
	}
	res = new(model.BaseRes[model.VerifyCompanyResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 企业工商数据查询
func (the Client) EnterpriseInfo(ctx context.Context, req *model.EnterpriseInfoReqBody) (res *model.BaseRes[model.EnterpriseInfoResBody], err error) {
	path := "/user/enterprise/info"
	res = new(model.BaseRes[model.EnterpriseInfoResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}
