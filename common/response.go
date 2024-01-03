/* response.go 系统响应

 */

package common

import (
	"github.com/gin-gonic/gin"
)

// Response Heimboard前端通用响应包
// gin上下文, http状态码, 状态码, 响应数据, 响应消息 =>
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}
