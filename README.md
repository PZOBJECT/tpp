# TPP (Third-Party Platform)

> Go SDK 统一接入：钉钉、飞书、微信公众号、微信小程序、企业微信、微信开放平台、支付宝、腾讯微卡

Forked from [leapig/tpp](https://github.com/leapig/tpp)，搬运到 PZOBJECT/Hermes-test。

## 快速开始

```shell
go get github.com/PZOBJECT/Hermes-test
```

```go
// 企微【企业内部开发】/【服务商代开发】应用
app := tpp.NewTpp().WW(ww.Config{
    CorpId:     "企业ID",
    CorpSecret: "应用的凭证密钥",
})

// 微信小程序
app := tpp.NewTpp().MP(mp.Config{...})

// 微信公众号
app := tpp.NewTpp().OA(oa.Config{...})

// 腾讯微卡
app := tpp.NewTpp().WK(wk.Config{...})

// 微信开放平台
app := tpp.NewTpp().WO(wo.Config{...})

// 支付宝
app := tpp.NewTpp().AP(ap.Config{...})

// 钉钉
app := tpp.NewTpp().DT(dt.Config{...})

// 飞书
app := tpp.NewTpp().FS(fs.Config{...})

// 调用平台接口
app.DoAnything()
```

## 支持平台

| 包名 | 平台 | 说明 |
|------|------|------|
| `ww` | 企业微信 | 企业内部开发 / 服务商代开发 |
| `mp` | 微信小程序 | 小程序登录、消息、支付 |
| `oa` | 微信公众号 | 公众号消息、菜单、网页授权 |
| `wk` | 腾讯微卡 | 微卡集成 |
| `wo` | 微信开放平台 | 第三方平台开发 |
| `ap` | 支付宝 | 支付、登录 |
| `dt` | 钉钉 | 通讯录、消息、微应用 |
| `fs` | 飞书 | 通讯录、消息、应用管理 |

## License

Apache 2.0
