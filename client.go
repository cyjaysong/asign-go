package asign

import (
	reqclient "github.com/imroc/req/v3"
	"os"
	"time"
)

const testBaseUrl = "https://prev.asign.cn"
const prodBaseUrl = "https://oapi.asign.cn"

type Client struct {
	reqClient      *reqclient.Client
	AppId          string // APPID
	appPrivateKey  []byte // 应用私钥
	asignPublicKey []byte // 爱签公钥
	prodEnv        bool   // 是否生产环境
	devMode        bool   // 是否调试模式
}

func NewClient(appId, appPrivateKeyPath, asignPublicKeyPath string, prodEnv, devMode bool) (client *Client, err error) {
	privateBytes, err := os.ReadFile(appPrivateKeyPath)
	if err != nil {
		return nil, err
	}
	publicFileBytes, err := os.ReadFile(asignPublicKeyPath)
	if err != nil {
		return nil, err
	}

	client = &Client{AppId: appId, appPrivateKey: privateBytes, asignPublicKey: publicFileBytes, prodEnv: prodEnv, devMode: devMode}
	client.reqClient = reqclient.C().SetCommonRetryCount(1)
	client.reqClient.SetIdleConnTimeout(time.Second * 3)
	client.reqClient.SetTimeout(time.Second * 10)
	client.reqClient.SetUserAgent("")
	client.reqClient.SetBaseURL(testBaseUrl)
	if prodEnv {
		client.reqClient.SetBaseURL(prodBaseUrl)
	}
	if devMode {
		client.reqClient.DevMode()
	}
	return
}
