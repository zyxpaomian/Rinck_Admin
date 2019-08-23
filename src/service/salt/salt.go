package salt

import (
	"util/log"
	"util/config"
    "bytes"
    "fmt"
    //"strings"
    "io/ioutil"
    "net/http"
    //"net/url"
    "crypto/tls"
    "encoding/json"
    //"SST/controllers"
    //"github.com/astaxie/beego"
    //"github.com/bitly/go-simplejson"
)


type AuthReutrnJson struct {
	Perms  []string `json:"perms"`
	Start  float64  `json:"start"`
	Token  string   `json:"token"`
	Expire float64  `json:"expire"`
	User   string   `json:"user"`
	Eauth  string   `json:"eauth"`
}

type AuthReutrnslice struct {
	ReturnSlice []AuthReutrnJson `json:"return"`
}

type AuthJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Eauth    string `json:"eauth"`
}

type ModelReturnSlice struct {
	Return []map[string]string
}

type ModelCExec struct {
	Fun	string `json fun` 
	Client string `json client`
	Tgt string `json string`
	Arg string `json string`
}

func GetToken() string {
	var saltauth AuthJson
	salt_url := config.GlobalConf.GetStr("salt", "salt_api_url")
	salt_api_url := fmt.Sprintf("%s%s", salt_url, "/login")

    saltauth.Username = config.GlobalConf.GetStr("salt", "salt_user")
    saltauth.Password = config.GlobalConf.GetStr("salt", "salt_passwd")
    saltauth.Eauth = config.GlobalConf.GetStr("salt", "salt_auth_mod")

    jsonstr, err := json.Marshal(saltauth)
    if err != nil {
        log.Errorf("解析salt认证json失败: ",err.Error())
    }
    req, err := http.NewRequest("POST", salt_api_url, bytes.NewBuffer(jsonstr))
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Content-Type", "application/json")

    tr := &http.Transport{
        TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
        DisableCompression: true,
    }

    client := &http.Client{Transport: tr}

    resp, err := client.Do(req)
    if err != nil {
        log.Errorf("获取Salt Token失败: ",err.Error())
    }
    defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	bodystr := string(body)
	
	var returnslice AuthReutrnslice
	json.Unmarshal([]byte(bodystr), &returnslice)
	
	token := (returnslice.ReturnSlice[0].Token)

	fmt.Println(token)
	return token
}

func ModelExec(model string,saltclient string,salttgt string,saltarg string) []map[string]string {
	var execresult []map[string]string

	token  := GetToken()
	postdata := fmt.Sprintf("{\"fun\":\"%s\",\"client\":\"%s\",\"tgt\":\"%s\",\"arg\":\"%s\",\"expr_form\":\"list\"}", model,saltclient,salttgt,saltarg)
	var jsonPostData = []byte(postdata)

    salt_api_url := config.GlobalConf.GetStr("salt", "salt_api_url")
	req, err := http.NewRequest("POST", salt_api_url, bytes.NewBuffer(jsonPostData))
	req.Header.Set("Content-type","application/json") 
    req.Header.Set("X-Auth-Token", token)
    tr := &http.Transport{
        TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
        DisableCompression: true,
    }
    client := &http.Client{Transport: tr}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	bodystr := string(body)

	var returnslice ModelReturnSlice
	json.Unmarshal([]byte(bodystr), &returnslice)

	execresult = returnslice.Return
	/*for _, t := range returnslice.Return {
		for k,v := range t {
			hostid := k
			execresult := v
		}
	}*/
	//return string(body)
	return execresult
}