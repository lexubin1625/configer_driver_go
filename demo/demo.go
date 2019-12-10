package main
import (
	"fmt"
	api "configer/api"
)
func main(){
	api.ConfigerInit("/data0/soft/configer",5000);
	etcdValue,_ := api.EtcdGet("/test/key1");
	fmt.Printf(etcdValue);
}