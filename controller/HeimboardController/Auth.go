package HeimboardController

import (
	"net/http"

	"forimoc.DracoNisus-Thuban/common"
	"forimoc.DracoNisus-Thuban/model/ModelAdmin"
	"forimoc.DracoNisus-Thuban/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login 管理员登录控制器
// user_id(用户ID), password(密码) => userID(用户ID), level:int(管理员权限), token(令牌)
func Login(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	userID := ctx.PostForm("user_id")
	password := ctx.PostForm("password")

	// 判断用户是否存在
	var admin ModelAdmin.Admin
	if err := DB.Where("user_id = ?", userID).First(&admin).Error; util.HandleNotFoundErr(err, ctx) {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		common.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	// 发放token
	token, err := common.ReleaseToken(admin)
	if err != nil {
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		common.LoggerFile.Error("release token failed: " + err.Error())
		return
	}

	// 返回结果
	common.Response(ctx, http.StatusOK, 200, gin.H{
		"userID": admin.UserID,
		"level":  admin.Level,
		"token":  token,
	}, "登录成功")
}

// Register 管理员注册控制器
// user_id(用户ID), password(密码) => userID(用户ID), level:int(管理员权限), token(令牌)
func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	userID := ctx.PostForm("user_id")
	password := ctx.PostForm("password")
	if len(userID) == 0 || len(password) == 0 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户ID和密码不能为空")
		return
	}

	// 判断用户是否存在
	var admin ModelAdmin.Admin
	if err := DB.Where("user_id = ?", userID).First(&admin).Error; util.HandleNotFoundErr(err, ctx) {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		common.LoggerTer.Error("hash password failed: " + err.Error())
		return
	}

	// 添加注册用户
	newAdmin := ModelAdmin.Admin{
		UserID:   userID,
		Password: string(hashedPassword),
		Level:    1,
	}
	if err := DB.Create(&newAdmin).Error; err != nil {
		common.LoggerFile.Error("database error: " + err.Error())
	}

	// 发放token
	token, err := common.ReleaseToken(admin)
	if err != nil {
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "token系统异常")
		common.LoggerFile.Error("release token failed: " + err.Error())
		return
	}

	// 返回结果
	common.Response(ctx, http.StatusOK, 200, gin.H{
		"userID": userID,
		"level":  2,
		"token":  token,
	}, "注册成功")
}
