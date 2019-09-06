package userctrl


import (
	"dao/usermgtdao"
	"util/log"
	ce "util/error"
	"structs"
	"github.com/satori/go.uuid"
)

/*type AuthMod struct {
	dao user.AuthDao
}*/
//var Auth *AuthMod


func CreateToken(username string,password string) (*structs.UserInfo,error){
//func CreateToken(username string,password string) (string,error){
	result := &structs.UserInfo{}
	result, err := usermgtdao.UserDao.UserAuth(username,password)
	if err != nil {
		log.Errorln("用户认证失败，无法获取token")
		return nil, ce.AuthError()
	}
	token, _ := uuid.NewV4()
	insertid, err := usermgtdao.UserDao.TokenSave(username,token.String())
	if err != nil || insertid == -1{
		log.Errorln("存入token失败")
		return nil,ce.DBError()
	}
	result.Token = token.String()
	return result, nil
}

func UpdateUserRole(username string,role string) (int,error){
	updateid, err := usermgtdao.UserDao.RoleUpdate(username,role)
	if err != nil || updateid == -1{
		log.Errorln("更新用户权限失败")
		return -1,ce.DBError()
	}
	return 1,nil
}

func DelUser(username string) (int,error){
	updateid, err := usermgtdao.UserDao.UserDel(username)
	if err != nil || updateid == -1{
		log.Errorln("删除用户失败")
		return -1,ce.DBError()
	}
	return 1,nil
}

func AddUser(username string, password string, role string) (int,error){
	insertid, err := usermgtdao.UserDao.UserAdd(username,password,role)
	if err != nil || insertid == -1{
		log.Errorln("添加用户失败")
		return -1,ce.DBError()
	}
	return 1,nil
}

func ResetPassword(username string,password string) (int,error) {
	updateid, err := usermgtdao.UserDao.PasswordReset(username,password)
	if err != nil || updateid == -1{
		log.Errorln("重置密码失败")
		return -1,ce.DBError()
	}
	return 1,nil
}


func GetUserInfo() ([]*structs.UserInfo,error){
	// result := &structs.UserInfo{}
	resultlist, err := usermgtdao.UserDao.GetUserInfo()
	if err != nil {
		log.Errorln("无法获取该用户相关信息")
		return nil, ce.DBError()
	}
	return resultlist, nil
}

func GetRoleList() ([]*structs.RoleInfo,error){
	// result := &structs.UserInfo{}
	resultlist, err := usermgtdao.UserDao.GetRoleInfo()
	if err != nil {
		log.Errorln("无法获取RoleList")
		return nil, ce.DBError()
	}
	return resultlist, nil
}