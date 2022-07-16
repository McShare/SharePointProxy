/*
 * @Author: NyanCatda
 * @Date: 2022-07-16 15:25:57
 * @LastEditTime: 2022-07-16 15:27:31
 * @LastEditors: NyanCatda
 * @Description: 路由注册
 * @FilePath: \SharePointProxy\Web\Routers\Routers.go
 */
package Routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nyancatda/AyaLog"
)

/**
 * @description: 路由注册
 * @param {*gin.Engine} r 路由引擎
 * @return {*}
 */
func SetupRouter(r *gin.Engine) *gin.Engine {
	// 注册500错误处理
	r.Use(ServerError)

	return r
}

/**
 * @description: 500错误处理
 * @param {*gin.Context} c
 * @return {*}
 */
func ServerError(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// 打印错误堆栈信息
			AyaLog.Error("Request", r.(error))

			// 500返回
			c.String(http.StatusInternalServerError, "Server Error")
			c.Abort()
			return
		}
	}()
	c.Next()
}
