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
path := "configer_path" //configer安装路径
valLength := 3*1024*1024 // 要获取值最大长度 注意单位为字节
api.ConfigerInit(path,valLength)

// 获取
	etcdValue,err := api.EtcdGet("/tauth/234242342134")
    if err != nil {
        // 异常处理
    }
	fmt.Printf(etcdValue)
```