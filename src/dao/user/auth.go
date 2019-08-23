package user

import (
	"util/mysql"
	ce "util/error"
	"util/log"
	"structs"
)

type AuthDao struct {
}

var UserDao AuthDao

func (authdao *AuthDao) UserAuth(username string, password string) ( *structs.UserInfo ,error) {
	result := &structs.UserInfo{}
	cnt, err := mysql.DB.SimpleQuery("select a.username,a.password,c.rolename from user as a left join groups as b on a.group_id = b.id left join role as c on  b.role_id = c.id where a.username = ? and a.password=MD5(?) ;",[]interface{}{username,password},&result.Username,&result.Password,&result.Role)

	if err != nil {
		log.Errorln("用户认证失败")
		return nil,ce.DBError()
	}
	if cnt == 0 {
		return nil, ce.AuthError()
	}
	return result, nil
}

func (authdao *AuthDao) GetUserInfo() ([]*structs.UserInfo ,error) {
		resultlist := []*structs.UserInfo{}
		tx := mysql.DB.GetTx()
		if tx == nil {
			log.Errorln("MySQL 获取TX失败")
			return nil, ce.DBError()
		}
		stmt, err := tx.Prepare("select a.username,b.groupname,DATE_FORMAT(a.create_time,'%Y-%m-%d %H:%i:%S'),c.rolename from user as a left join groups as b on a.group_id = b.id left join role as c on  b.role_id = c.id ;")
		if err != nil {
			tx.Rollback()
			log.Errorln("MySQL 获取TX失败: ",err.Error())
			return nil, ce.DBError()
		}
		rows, err := stmt.Query()
		if err != nil {
			log.Errorln("MySQL 查询失败: ",err.Error())
			stmt.Close()
			return nil, ce.DBError()
		}
		for rows.Next() {
			result := &structs.UserInfo{}
			err := rows.Scan(&result.Username,&result.Group,&result.CreateTime,&result.Role)
			if err != nil {
				log.Errorln("MySQL 查询失败: ",err.Error())
				rows.Close()
				stmt.Close()
				tx.Rollback()
				return nil,ce.DBError()
			} else {
				resultlist = append(resultlist, result)
			}
		}
		rows.Close()
		stmt.Close()
		tx.Commit()
		return resultlist, nil
	}


func (authdao *AuthDao) TokenSave(username string,token string) (int,error) {
	insertid, err := mysql.DB.SimpleInsert("insert into tokens(token,username,expire_time) values (?,?,date_sub(NOW(),interval -1 day))",token,username)
	if err != nil || insertid == -1{
		log.Errorln("存入token失败")
		return -1,ce.DBError()
	}
	return insertid, nil
}

