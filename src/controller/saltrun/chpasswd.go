package saltrun

import (
	"service/salt"
	//"dao/user"
	//"fmt"
)


func ChangePasswd(ip string,argname string)  map[string]string {
	changeResult := salt.ModelExec("hello.haha","local",ip,argname)
	return changeResult[0]
}
