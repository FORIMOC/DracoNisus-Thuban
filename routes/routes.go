package routes

import (
	"forimoc.DracoNisus-Thuban/controller/HeimboardController"
	"forimoc.DracoNisus-Thuban/controller/HeimbotController"
	"forimoc.DracoNisus-Thuban/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// 机器人监听入口
	r.POST("/", HeimbotController.MainEntry)

	// 管理员登录
	r.POST("/api/auth/login", HeimboardController.Login)
	// 管理员注册
	r.POST("/api/auth/register", HeimboardController.Register)

	/* group 数据crud */
	// 更新 单个群聊的 监听状态
	r.POST("/api/group/update_group_status", middleware.Auth(0), HeimboardController.UpdateGroupStatus)
	// 获取 所有群聊的列表
	r.POST("/api/group/get_group_list", middleware.Auth(1), HeimboardController.GetGroupList)
	// 获取 单个群聊的 基本数据
	r.POST("/api/group/get_group_info", middleware.Auth(1), HeimboardController.GetGroupInfo)
	// 获取 单个群聊的 社交网络数据
	r.POST("/api/group/get_group_graphs", middleware.Auth(1), HeimboardController.GetGroupGraph)
	// 获取 单个群聊的 用户数据
	r.POST("/api/group/get_users_list", middleware.Auth(1), HeimboardController.GetUserList)
	// 获取 单个群聊的 缓存消息数据
	r.POST("/api/group/get_group_cache_messages", middleware.Auth(1), HeimboardController.GetGroupCacheMessages)
	// 获取 单个群聊的 消息时间分布数据
	r.POST("/api/group/get_group_times", middleware.Auth(1), HeimboardController.GetGroupTimes)
	// 获取 单个群聊的 关键词频数据
	r.POST("/api/group/get_group_keywords", middleware.Auth(1), HeimboardController.GetGroupKeywords)
	// 删除 单个群聊的 所有相关数据
	r.POST("/api/group/delete_group_data", middleware.Auth(0), HeimboardController.DeleteGroupData)

	/* keyword 数据crud */
	// 增加 单个关键词
	r.POST("/api/keyword/add_keyword", middleware.Auth(0), HeimboardController.AddKeyword)
	// 删除 单个关键词
	r.POST("/api/keyword/delete_keyword", middleware.Auth(0), HeimboardController.DeleteKeyword)
	// 获取 所有关键词列表
	r.POST("/api/keyword/list_keyword", middleware.Auth(0), HeimboardController.GetKeywordList)
	// 删除 所有不在关键词列表里的关键词记录
	r.POST("/api/keyword/regulate_keyword", middleware.Auth(0), HeimboardController.RegulateKeyword)

	/* key user 数据crud */
	// 更新 单个关键用户的 监听状态
	r.POST("/api/key_user/update_key_user_status", middleware.Auth(0), HeimboardController.UpdateKeyUserStatus)
	// 增加 单个关键用户
	r.POST("/api/key_user/add_key_user", middleware.Auth(0), HeimboardController.AddKeyUser)
	// 删除 单个关键用户
	r.POST("/api/key_user/delete_key_user", middleware.Auth(0), HeimboardController.DeleteKeyUser)
	// 设置 关键用户备注
	r.POST("/api/key_user/set_key_user_info", middleware.Auth(0), HeimboardController.SetKeyUserInfo)
	// 搜索 关键用户昵称
	r.POST("/api/key_user/search_username", middleware.Auth(0), HeimboardController.SearchUsername)
	// 获取 所有关键用户列表
	r.POST("/api/key_user/get_key_user_list", middleware.Auth(0), HeimboardController.GetKeyUserList)
	// 获取 单个关键用户信息
	r.POST("/api/key_user/get_key_user_info", middleware.Auth(0), HeimboardController.GetKeyUserInfo)
	// 获取 单个关键用户发言群聊
	r.POST("/api/key_user/get_key_user_groups", middleware.Auth(0), HeimboardController.GetKeyUserGroup)
	// 获取 单个关键用户发言时间分布
	r.POST("/api/key_user/get_key_user_times", middleware.Auth(0), HeimboardController.GetKeyUserTime)
	// 获取 单个关键用户关键词数据
	r.POST("/api/key_user/get_key_user_keywords", middleware.Auth(0), HeimboardController.GetKeyUserKeywords)
	// 获取 单个关键用户缓存消息
	r.POST("/api/key_user/get_key_user_cache_messages", middleware.Auth(0), HeimboardController.GetKeyUserCacheMessages)

	/* HyperGroup 数据crud */
	// 计算 超群数据
	r.POST("/api/hyper_group/calculate_hyper_group", middleware.Auth(1), HeimboardController.CalculateHyperGroup)
	return r
}
