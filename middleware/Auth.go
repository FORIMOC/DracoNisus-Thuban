package middleware

import (
	"net/http"

	"forimoc.DracoNisus-Thuban/common"
	"forimoc.DracoNisus-Thuban/model/ModelAdmin"
	"forimoc.DracoNisus-Thuban/util"
	"github.com/gin-gonic/gin"
)

// Auth 管理员验证中间件
// 管理员等级 =>
func Auth(adminLevel int) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// 检验票据格式
		if tokenString == "" {
			common.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			ctx.Abort()
			return
		}

		// 检验票据格式
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			common.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			ctx.Abort()
			return
		}

		// 验证通过后获取claim中的userID
		ID := claims.UserID
		DB := common.GetDB()
		var admin ModelAdmin.Admin
		// 用户不存在
		if err := DB.First(&admin, ID).Error; util.HandleNotFoundErr(err, ctx) {
			common.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			ctx.Abort()
			return
		}

		// 管理员等级过低
		if admin.Level > adminLevel {
			common.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			ctx.Abort()
			return
		}

		// 用户存在且等级足够 将admin信息写入上下文
		ctx.Set("admin", admin)

		ctx.Next()
	}
}
