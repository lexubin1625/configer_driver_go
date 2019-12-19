package main

import(
	"io"
	"net/http"
	api "git.intra.weibo.com/user_growth_common/configer_go/api"
)

func SimpleServer(w http.ResponseWriter,request *http.Request){
	
	etcdValue,_ := api.EtcdGet("/test/key1");
	io.WriteString(w,etcdValue)
}

func main(){
	api.ConfigerInit("/data0/soft/configer",5000)
	http.HandleFunc("/configrt/etcd/get",SimpleServer)
	if err := http.ListenAndServe(":8888",nil);err != nil{
	}
}
