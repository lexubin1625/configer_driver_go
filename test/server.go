package main

import(
	"io"
	"net/http"
	configer "configer_driver_go/driver"
)

func SimpleServer(w http.ResponseWriter,request *http.Request){
	
	etcdValue,_ := configer.ConfGet("/test/key1");
	io.WriteString(w,etcdValue)
}

func main(){
	configer.ConfInit("/data0/soft/configer",0,5000);
	http.HandleFunc("/configrt/etcd/get",SimpleServer)
	if err := http.ListenAndServe(":8888",nil);err != nil{
	}
}
