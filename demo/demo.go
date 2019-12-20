package main
import (
	"fmt"
	api "git.intra.weibo.com/user_growth_common/configer_go/api"
)
func main(){
	// 初始化
	api.ConfigerInit("/data0/soft/configer",5000)

	// 获取值
	etcdValue,_ := api.EtcdGet("/tauth/234242342134")
	fmt.Printf(etcdValue)
}