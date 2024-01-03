package HeimbotController

import (
	"io"

	"forimoc.DracoNisus-Thuban/common"
	"forimoc.DracoNisus-Thuban/model/ModelGroup"
	"forimoc.DracoNisus-Thuban/model/ModelKeyUser"
	"forimoc.DracoNisus-Thuban/record"
	"forimoc.DracoNisus-Thuban/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func MainEntry(ctx *gin.Context) {
	DB := common.GetDB()

	// 获取 go-cqhttp POST 转发信息
	rawData, _ := io.ReadAll(ctx.Request.Body)
	data := string(rawData)
	postType := gjson.Get(data, "post_type").String()
	if postType != "message" {
		return
	}

	// 获取参数
	groupID := gjson.Get(data, "group_id").String()
	userID := gjson.Get(data, "user_id").String()

	// 白名单过滤
	if util.IsGroupInConfigItem(groupID, "listen.whitelistGroupID") {
		return
	}

	// 群聊存在控制
	var group ModelGroup.Group
	defaultListenMode := viper.GetString("bot.defaultListenMode")
	if err := DB.Where("group_id = ?", groupID).First(&group).Error; util.HandleNotFoundErr(err, nil) {
		groupName, member, maxMember, err := util.GetGroupInfo(groupID)
		if err != nil {
			common.LoggerFile.Error("get group info failed: " + err.Error())
		}
		if len(groupName) > 0 && len(groupID) > 0 {
			newGroup := ModelGroup.Group{
				GroupID:      groupID,
				GroupName:    groupName,
				Status:       defaultListenMode,
				SumMsgNum:    0,
				AvgMsgLen:    0,
				MemberNum:    member,
				MaxMemberNum: maxMember,
			}
			if err := DB.Create(&newGroup).Error; err != nil {
				common.LoggerFile.Error("database error: " + err.Error())
			}
		} else {
			common.LoggerFile.Error("group name and ID is null")
		}
	}

	// 群聊状态控制
	switch group.Status {
	case "3": // 全局监听
		go record.GroupUserUpdate(DB, data)
		go record.GroupUpdate(DB, data)

		go record.GroupTimeUpdate(DB, data)
		go record.GroupUserTimeUpdate(DB, data)
		go record.GroupCacheMessageUpdate(DB, data)
		go record.GroupMessageUpdate(DB, data)
		go record.GroupGraphUpdate(DB, data)
		go record.GroupKeywordUpdate(DB, data)
	case "2": // 控局监听
		go record.GroupUserUpdate(DB, data)
		go record.GroupUpdate(DB, data)

		go record.GroupTimeUpdate(DB, data)
		go record.GroupUserTimeUpdate(DB, data)
		go record.GroupCacheMessageUpdate(DB, data)
		go record.GroupGraphUpdate(DB, data)
		go record.GroupKeywordUpdate(DB, data)
	case "1": // 摘要监听
		go record.GroupUpdate(DB, data)
	// 无监听(不会发送给前端)
	case "0":
	default:
	}

	/* 关键用户 */
	// 是否属于关键用户
	var keyUser ModelKeyUser.KeyUser
	if err := DB.Table("key_users").Where("user_id = ?", userID).First(&keyUser).Error; util.HandleNotFoundErr(err, nil) {
		return
	}
	switch keyUser.Status {
	case "1": // 监听
		go record.KeyUserUpdate(DB, data)
		go record.KeyUserTimeUpdate(DB, data)
		go record.KeyUserCacheMessageUpdate(DB, data)
		go record.KeyUserGroupUpdate(DB, data)
		go record.KeyUserKeywordUpdate(DB, data)
	// 暂不监听
	case "0":
	default:
	}
}
