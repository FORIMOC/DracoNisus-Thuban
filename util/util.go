package util

import (
	"errors"
	"net/http"
	"os/exec"
	"strings"

	"forimoc.DracoNisus-Thuban/common"
	"forimoc.DracoNisus-Thuban/model/ModelAdmin"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
)

// HandleNotFoundErr 处理数据库未找到记录错误
// err错误, ctx上下文 => 是/否为未找到错误
func HandleNotFoundErr(err error, ctx *gin.Context) bool {
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// not found err
			return true
		} else {
			// other err
			callerInfo := GetCallerInfo(2)
			if ctx != nil {
				common.Response(ctx, http.StatusInternalServerError, 500, nil, "数据库错误")
			}
			// TODO: 发邮件
			common.LoggerTer.WithFields(logrus.Fields{"caller info": callerInfo}).Error("database error: " + err.Error())
			common.LoggerFile.WithFields(logrus.Fields{"caller info": callerInfo}).Fatal("database error: " + err.Error())
		}
	}
	// no err
	return false
}

// ReachEndpoint curl本机cqhttp的endpoint
// endpoint路由, endpoint参数 => json结果, err错误
func ReachEndpoint(uri string, query string) (string, error) {
	port := viper.GetInt("bot.port")
	domain := "http://127.0.0.1"
	u, err := CombineURL(domain, port, uri, query)
	if err != nil {
		return "", err
	}
	json, err := exec.Command("curl", u).CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(json), nil
}

// IsGroupInConfigItem 群聊是否在配置文件中的某个群聊集合中
// 目标群聊, 配置项路径 => true/false
func IsGroupInConfigItem(groupID string, configItem string) bool {
	configString := viper.GetString(configItem)
	configGroups := strings.Split(configString, ",")
	for i := 0; i < len(configGroups); i++ {
		if groupID == configGroups[i] {
			return true
		}
	}
	return false
}

// IsHighLevelAdmin 判断是否为高等级管理员(0)
// 数据库实例, 用户ID => true/false
func IsHighLevelAdmin(DB *gorm.DB, userID string) (bool, error) {
	var admin ModelAdmin.Admin
	if err := DB.Where("user_id = ?", userID).First(&admin).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return false, err
		}
	} else {
		if admin.Level < 1 {
			return true, nil
		}
	}
	return false, nil
}

// GetGroupInfo 获取群聊基本信息
// 群ID => 群名称, 群成员数, 群最大成员数
func GetGroupInfo(groupID string) (string, int, int, error) {
	json, err := ReachEndpoint("/get_group_info", "group_id="+groupID)
	if err != nil {
		return "", -1, -1, err
	}
	groupName := gjson.Get(string(json), "data.group_name").String()
	member := gjson.Get(string(json), "data.member_count").Int()
	maxMember := gjson.Get(string(json), "data.max_member_count").Int()
	return groupName, int(member), int(maxMember), nil
}

// GetGroupList 获取群聊ID列表
// => 群聊ID字符串数组
// func GetGroupList() []string {
// 	var groupList []string
// 	botPort := viper.GetString("bot.port")
// 	u := "http://127.0.0.1:" + botPort + "/get_group_list"
// 	json, _ := exec.Command("curl", u).CombinedOutput()
// 	for i := 0; i < int(gjson.Get(string(json), "data.#").Int()); i++ {
// 		groupList = append(groupList, gjson.Get(string(json), "data."+strconv.Itoa(i)+".group_id").String())
// 	}
// 	return groupList
// }
