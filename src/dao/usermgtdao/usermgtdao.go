package usermgtdao

import (
	"util/mysql"
	ce "util/error"
	"util/log"
	"structs"
	// "fmt"
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

func (authdao *AuthDao) GetRoleInfo() ([]*structs.RoleInfo ,error) {
	resultlist := []*structs.RoleInfo{}
	tx := mysql.DB.GetTx()
	if tx == nil {
		log.Errorln("MySQL 获取TX失败")
		return nil, ce.DBError()
	}
	stmt, err := tx.Prepare("select rolename from role;")
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
		result := &structs.RoleInfo{}
		err := rows.Scan(&result.Rolename)
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

func (authdao *AuthDao) RoleUpdate(username string,role string) (int,error) {
	updateid, err := mysql.DB.SimpleUpdate("update user set group_id = (select id from groups where role_id = (select id from role where rolename=?)) where username = ?",role,username)
	if err != nil || updateid == 0{
		log.Errorln("更新用户权限失败")
		return -1,ce.DBError()
	}
	return updateid, nil
}

func (authdao *AuthDao) UserDel(username string) (int,error) {
	updateid, err := mysql.DB.SimpleUpdate("delete from user where username = ?",username)
	if err != nil || updateid == 0{
		log.Errorln("删除用户失败")
		return -1,ce.DBError()
	}
	return updateid, nil
}

func (authdao *AuthDao) PasswordReset(username string,password string) (int,error) {
	updateid, err := mysql.DB.SimpleUpdate("update user set password=MD5(?) where username = ?;",password,username)
	if err != nil || updateid == 0{
		log.Errorln("重置密码失败")
		return -1,ce.DBError()
	}
	return updateid, nil
}

func (authdao *AuthDao) UserAdd(username string, password string, role string) (int,error) {
	insertid, err := mysql.DB.SimpleInsert("insert into user (`username`,`password`,create_time,group_id) values (?,md5(?),NOW(),(select id from groups where role_id = (select id from role where rolename=?)))",username,password,role)
	if err != nil || insertid == -1{
		log.Errorln("添加用户失败")
		return -1,ce.DBError()
	}
	return insertid, nil
}
