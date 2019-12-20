#####  configer go 客户端使用说明

###### 配置说明
1、configer 已安装
2、替换git地址，git config --global --add url."ssh://git@git.intra.weibo.com:2222/".insteadOf "http://git.intra.weibo.com/"
3、go get -v -insecure

```go

import (
	api "git.intra.weibo.com/user_growth_common/configer_go"
)

// 初始化
path (string) configer安装路径
valLength (int) 要获取值最大长度
api.ConfigerInit(path,valLength)

```