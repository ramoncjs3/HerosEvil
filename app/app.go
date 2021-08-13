package app

import (
	"HerosEvil/AESUtil"
	"HerosEvil/LZStringUtil"
	"HerosEvil/RandUtil"
	"HerosEvil/ReqUtil"
	"HerosEvil/SignUtil"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"log"
	"strconv"
	"time"
)

//登陆身份认证信息
type Login struct {
	Username      string
	Password      string
	Loginid       string
	Token         string
	RoleID        float64
	LoginFlag     float64
	QuickLoginUrl string
	KillNpcUrl    string
}

var LoginStruct = &Login{
	"",
	"",
	"",
	"",
	0,
	0,
	"",
	"",
}

//待杀的NPC
var killedNpc = []string{"1162","1155","1156","1151","1150","1153","1158","1154","1161","1152","1159","1157","1160","1163","1136","1138","1139","1194","1193","1148","1147","1143","1144","1137","1142","1141","1145","1140","1146","1127","1122","1121","1130","1301","1125","1124","1123","1128","1126","1131","1129","1115","1102","1110","1111","1103","1105","1108","1230","1107","1101","1104","1114","1106","1112","1109","1113","1085","1094","1086","1087","1089","1091","1090","1095","1088","1092","1097","1093","1099","1096","1072","1081","1078","1074","1075","1079","1077","1073","1070","1069","1071","1082","1080","1053","1050","1055","1051","1059","1058","1054","1052","1060","1061","1062","1057","1056","1066","1065","1063"}

//登陆认证第一步，返回loginid与token
func ReqLogin() error {

	var inter map[string]interface{}
	reqUrl := "https://xfsdk.tfy-inc.com/app/login.php"
	m := make(map[string]string)
	m["username"] = LoginStruct.Username
	m["password"] = LoginStruct.Password
	m["tfyuuid"] = RandUtil.RandStr(16)
	m["agent"] = "app-10-02"
	m["logintime"] = fmt.Sprintf("%d", time.Now().Unix())
	m["gameid"] = "10"
	m["deviceType"] = "android"
	m["appid"] = "1001"
	m["sign"] = SignUtil.SignCheck(m)
	reqBodys := ""
	for i, j := range m {
		reqBodys = reqBodys + i
		reqBodys = reqBodys + "=" + j + "&"
	}
	resp, err := ReqUtil.ReqPostData(reqUrl, reqBodys)
	if err != nil {
		return err
	}
	err = resp.ToJSON(&inter)
	if err != nil {
		return err
	}
	LoginStruct.Loginid = inter["loginid"].(string)
	LoginStruct.Token = inter["token"].(string)
	return nil
}

//登陆认证第二步，返回loginFlag，roleID，account
func quickLogin() error {
	var tmp map[string]interface{}
	nonce := strconv.FormatInt(time.Now().Unix(), 13) + RandUtil.RandStr(8)
	originBodys := fmt.Sprintf(`{"mod":"User","do":"quicklogin","p":{"account":"%s","pwd":"123456","checkObj":{"app":"1","sdk":"meple100221","uin":"%s","sess":"%s","newPack":1},"channel":"meple100221","flag":1,"macAdress":"02:00:00:00:00:00","platform":"0","web":false,"clientVersion":{"android":"2.3.4.1628007977548"},"NeedUpdateVersion":"2.1.7","inGameTime":0,"roleID":0,"loginFlag":"1","nonce":"%s"}}`, LoginStruct.Loginid, LoginStruct.Loginid, LoginStruct.Token, nonce)
	reqBodys := AESUtil.EncryptAES(LZStringUtil.CompressToBase64(originBodys))
	resp, err := ReqUtil.ReqPostData(LoginStruct.QuickLoginUrl+"/quicklogin", reqBodys)
	if err != nil {
		return err
	}
	_ = resp.ToJSON(&tmp)
	if tmp["code"] != float64(200) {
		log.Println("deJsonRespCode!200:", tmp)
		return err
	}
	respDE, _ := LZStringUtil.DecompressFromBase64(AESUtil.DecryptAES(tmp["MsgData"].(string)))
	var respJson map[string]interface{}
	_ = json.Unmarshal([]byte(respDE), &respJson)
	if respJson["roleID"] == nil && respJson["loginFlag"] == nil {
		return errors.New("获取roleID与loginFlag失败,大概率被封号了.")
	}
	LoginStruct.RoleID = respJson["roleID"].(float64)
	LoginStruct.LoginFlag = respJson["loginFlag"].(float64)
	return nil
}

//击杀NPC
func KillNpc(str string) error {
	var tmp map[string]interface{}
	originBodys := fmt.Sprintf(`{"mod":"User","do":"killNpc","p":{"npcIDArr":[%s],"type":1,"roleID":"%s","web":false,"clientVersion":{"android":"2.3.4.1628007977548"},"NeedUpdateVersion":"2.3.3","userAccount":"%s","loginFlag":"%s","inGameTime":45344761.5610908}}`, str, strconv.FormatFloat(LoginStruct.RoleID, 'f', -1, 64), LoginStruct.Loginid, strconv.FormatFloat(LoginStruct.LoginFlag, 'f', -1, 64))
	reqBodys := AESUtil.EncryptAES(LZStringUtil.CompressToBase64(originBodys))
	resp, err := ReqUtil.ReqPostData(LoginStruct.KillNpcUrl+"/killNpc", reqBodys)
	if err != nil {
		return err
	}
	_ = resp.ToJSON(&tmp)
	if tmp["code"] != float64(200) {
		log.Println("请求失败！:", tmp)
		return err
	}
	respDE, _ := LZStringUtil.DecompressFromBase64(AESUtil.DecryptAES(tmp["MsgData"].(string)))
	log.Println(respDE)
	js, err := simplejson.NewJson([]byte(respDE))
	if err != nil {
		log.Println(err)
	}
	if data, ok := js.GetPath("User.killNpc", "note", "attrch").CheckGet("evil"); ok {
		tmp, _ := data.Int()
		log.Println(fmt.Sprintf("已击杀NpcID:%s.", str))
		log.Println(fmt.Sprintf("当前善恶值:%d.", tmp))
	} else {
		log.Println(fmt.Sprintf("这个NPC杀过了,NpcID:%s.", str))
	}
	return nil
}

func Run() {
	err := ReqLogin()
	if err != nil {
		log.Println(err)
	} else {
		err = quickLogin()
		if err != nil {
			log.Println(err)
		} else {

			for num, npcID := range killedNpc {
				log.Println(fmt.Sprintf("当前执行第%d个Npc,剩余%d个Npc任务待执行.", num+1, len(killedNpc)-num-1))
				err := KillNpc(npcID)
				if err != nil {
					log.Println(err)
				}
				t := RandUtil.RandTime()
				log.Println(fmt.Sprintf("随机等待%d秒...", t))
				time.Sleep(time.Duration(t) * time.Second)
			}
		}
	}

}
