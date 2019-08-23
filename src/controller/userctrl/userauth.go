package userctrl


import (
	"dao/user"
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
	result, err := user.UserDao.UserAuth(username,password)
	if err != nil {
		log.Errorln("用户认证失败，无法获取token")
		return nil, ce.AuthError()
	}
	token, _ := uuid.NewV4()
	insertid, err := user.UserDao.TokenSave(username,token.String())
	if err != nil || insertid == -1{
		log.Errorln("存入token失败")
		return nil,ce.DBError()
	}
	result.Token = token.String()
	return result, nil
}

func GetUserInfo() ([]*structs.UserInfo,error){
	// result := &structs.UserInfo{}
	resultlist, err := user.UserDao.GetUserInfo()
	if err != nil {
		log.Errorln("无法获取该用户相关信息")
		return nil, ce.DBError()
	}
	return resultlist, nil
}
