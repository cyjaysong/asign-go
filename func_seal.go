package asign

import (
	"context"
	"github.com/cyjaysong/asign-go/model"
)

// 查询印章
func (the Client) GetUserSeals(ctx context.Context, req *model.GetUserSealsReqBody) (res *model.BaseRes[model.GetUserSealsResBody], err error) {
	path := "/user/getUserSeals"
	res = new(model.BaseRes[model.GetUserSealsResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}

// 印章授权（无校验）
func (the Client) NocheckAuth(ctx context.Context, req *model.NocheckAuthReqBody) (res *model.BaseRes[model.NocheckAuthResBody], err error) {
	path := "/seal/nocheckAuth"
	res = new(model.BaseRes[model.NocheckAuthResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}
