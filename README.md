<h1 align="center">Genshin Sign</h1>

原神米游社自动签到

**不提供二进制发布**

<!--

## 信息准备

### 获取 token 和 uid

程序提供两种登录方式，stoken 和 cookie token

#### cookie token 获取 (推荐)

在 [米游社网页端](https://bbs.mihoyo.com/ys/) 登录, 并打开 cookie, 记录两项内容 `account_id` 和 `cookie_token`. 
两项 cookie 均在 `mihoyo.com` 域名下

> account_id 为米游社的 uid, 后续参数配置使用这个 uid 而非游戏内 uid

#### stoken 获取

> 使用 app 抓包工具，米游社退出重新登录后，在 getMultiToken 接口获取

## docker-compose

```yaml
version: "3"

services:
  genshin:
    image: royii/genshin-sign:latest
    environment:
      uid: <uid>
      ctoken: <cookie token>
      stoken: <stoken>
```

## 使用 GitHub Action

1. fork
2. 添加 secret
  + uid
  + ctoken
  + stoken

> action 通过判断是否配置了 secret.uid 确定是否启用 job

-->