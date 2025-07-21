# 智付SDK V2

### 初始化

```
accessId := "3011"
appPrivateKeyPath := "../config/Test_appPrivateKey.cer"
asignPublicKeyPath := "../config/Test_asignPublicKey.cer"
prodEnv := false // prodEnv 为true请求爱签的生产环境，为false请求爱签的测试环境，请悉知
devMode := false // devMode 为true会打印接口请求和响应内容
client, err := globebill.NewClient(accessId, appPrivateKeyPath, asignPublicKeyPath, prodEnv, devMode)
```

### 打赏

赞赏多少都是您的心意，感谢您的支持！打赏时烦请备注一下您的github账号，以便添加感谢名单

<img src="./image/微信收款码.jpg" height="300"> <img src="./image/支付宝收款码.jpg" height="300">

### 感谢名单

| Benefactor | Channel | Amount | Time |
|------------|---------|--------|------|
