/*
 * @Author: your name
 * @Date: 2020-02-28 11:45:14
 * @LastEditTime: 2020-03-19 11:35:38
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: /configer/driver/go/configer/demo/demo.go
 */
package main
import (
	"fmt"
	configer "configer_driver_go/driver"
)
func main(){
	configer.ConfInit("/Users/xubin6/Documents/www/go/src/configer",0,5000)
	etcdValue,_ := configer.ConfGet("/test/key21")
	fmt.Printf(etcdValue)
}