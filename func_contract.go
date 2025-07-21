package asign

import (
	"context"
	"github.com/cyjaysong/asign-go/model"
)

// 创建待签署文件
func (the Client) ContractCreate(ctx context.Context, req *model.ContractCreateReqBody) (res *model.BaseRes[model.ContractCreateResBody], err error) {
	path := "/contract/createContract"
	res = new(model.BaseRes[model.ContractCreateResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 添加签署方
func (the Client) ContractAddSigner(ctx context.Context, req *model.ContractAddSignerReqBody) (res *model.BaseRes[model.ContractAddSignerResBody], err error) {
	path := "/contract/addSigner"
	res = new(model.BaseRes[model.ContractAddSignerResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 查询合同状态
func (the Client) GetContractStatus(ctx context.Context, req *model.GetContractStatusReqBody) (res *model.BaseRes[model.GetContractStatusResBody], err error) {
	path := "/contract/status"
	res = new(model.BaseRes[model.GetContractStatusResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 查询合同信息
func (the Client) GetContract(ctx context.Context, req *model.GetContractReqBody) (res *model.BaseRes[model.GetContractResBody], err error) {
	path := "/contract/getContract"
	res = new(model.BaseRes[model.GetContractResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 下载合同信息
func (the Client) DownloadContract(ctx context.Context, req *model.DownloadContractReqBody) (res *model.BaseRes[model.DownloadContractResBody], err error) {
	path := "/contract/downloadContract"
	res = new(model.BaseRes[model.DownloadContractResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}
