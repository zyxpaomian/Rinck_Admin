package structs

type UserInfo struct {
	Username string `json:username`
	Password string `json:password`
	Role string `json:role`
	Token string `json:token`
	Group string `json:group`
	CreateTime string `json:createtime`
}

type RoleInfo struct {
	Rolename string `json:rolename`
	Roleid int64 `json:Roleid`
}


