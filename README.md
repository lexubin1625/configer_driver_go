#####  configer go 客户端使用说明

###### 配置说明
 1、configer 已安装 必须安装configer c++静态库
 
 2、替换git地址，git config --global --add url."ssh://git@git.intra.weibo.com:2222/".insteadOf "http://git.intra.weibo.com/"
 
 3、go get -v -insecure

###### 代码实例 (可参考test目录下文件)
```go
  
import (
	configer "git.intra.weibo.com/user_growth_common/configer_driver_go/driver"
   // 其他库按需载入
)

    // 初始化
    path := "configer_path" //configer安装路径
    valLength := 3*1024*1024 // 要获取值最大长度 注意单位为字节
    shmKey := 0 // 共享内存首地址 为0时设置为默认地址
    configer.ConfInit(path,shmKey,valLength)

    // 获取
	etcdValue,err := configer.ConfGet("/tauth/234242342134")
    if err != nil {
        // 异常处理
    }
	fmt.Printf(etcdValue)
```