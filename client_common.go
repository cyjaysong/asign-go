package asign

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/deatil/go-cryptobin/cryptobin/rsa"
	"strconv"
	"time"
)

// 公共相关接口请求
func (the Client) post(ctx context.Context, path string, reqBody, resBody any) (err error) {
	bizDataBytes, _ := sonic.Marshal(reqBody)
	node, _ := sonic.Get(bizDataBytes)
	_ = node.SortKeys(true)
	bizDataBytes, _ = node.MarshalJSON()
	bizDataStr := string(bizDataBytes)
	sign, timestamp := the.Sha1WithRsaSign(bizDataStr, bizDataBytes)
	formData := map[string]string{"appId": the.AppId, "timestamp": timestamp, "bizData": bizDataStr}
	req := the.reqClient.R().SetContext(ctx).SetHeader("sign", sign).SetHeader("timestamp", timestamp)
	response, err := req.EnableForceMultipart().SetFormData(formData).Post(path)
	if err != nil {
		return err
	}
	return response.Unmarshal(resBody)
}

// Sha1WithRsaSign 签名
func (the Client) Sha1WithRsaSign(bizDataStr string, bizDataBytes []byte) (sign, timestamp string) {
	dataMd5 := fmt.Sprintf("%x", md5.Sum(bizDataBytes))
	timestamp = strconv.FormatInt(time.Now().Add(time.Minute*10).UnixMilli(), 10)
	sign = bizDataStr + dataMd5 + the.AppId + timestamp
	obj := rsa.New().FromPrivateKey(the.appPrivateKey).SetSignHash("SHA1")
	return obj.FromString(sign).Sign().ToBase64String(), timestamp
}

// Sha1WithRsaVerify 验签
func (the Client) Sha1WithRsaVerify(bodyBytes []byte) (pass bool) {
	bodyNode, err := sonic.Get(bodyBytes)
	if err != nil || !bodyNode.Valid() {
		return false
	}
	signNode := bodyNode.Get("sign")
	if !signNode.Valid() {
		return false
	}
	sign, err := signNode.String()
	if err != nil || sign == "" {
		return false
	}
	_, _ = bodyNode.Unset("sign")
	_ = bodyNode.SortKeys(true)
	bodyBytes, _ = bodyNode.MarshalJSON()
	obj := rsa.New().FromPublicKey(the.asignPublicKey).SetSignHash("SHA1")
	return obj.FromBase64String(sign).Verify(bodyBytes).ToVerify()
}
