/*
 * @Author: NyanCatda
 * @Date: 2022-07-16 15:38:35
 * @LastEditTime: 2022-07-16 15:59:40
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \SharePointProxy\Web\Routers\Proxy.go
 */
package Routers

import (
	"errors"
	"net/http"

	"github.com/McShare/SharePointProxy/Internal/SharePointProxy"
	"github.com/McShare/SharePointProxy/Tools"
	"github.com/gin-gonic/gin"
	"github.com/nyancatda/AyaLog"
)

/**
 * @description: 反向代理路由注册
 * @param {*gin.Engine} r 路由引擎
 * @return {*}
 */
func ProxyRouter(r *gin.Engine) {
	// 获取文件
	r.GET("/*URLPath", func(c *gin.Context) {
		// 获取所有请求参数
		Query := c.Request.URL.Query()
		var QueryMap = make(map[string]any, len(Query))
		for k := range Query {
			QueryMap[k] = c.Query(k)
		}
		// 转换为查询字符串
		QueryString := Tools.MapToGetQuery(QueryMap)

		FileBuffer, HttpResponse, err := SharePointProxy.GetFile(c.Param("URLPath") + "?" + QueryString)
		if err != nil {
			AyaLog.Error("Request", err)
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		if HttpResponse.StatusCode != http.StatusOK {
			AyaLog.Error("Request", errors.New("Error in upstream service response code "+HttpResponse.Status))
			c.String(http.StatusInternalServerError, "Error in upstream service response code "+HttpResponse.Status)
			return
		}

		// 获取文件类型
		ContentType := http.DetectContentType(FileBuffer.Bytes())

		// 按照源响应头返回响应头
		for Header, Value := range HttpResponse.Header {
			var HeaderValue string
			if len(Value) > 1 {
				for _, v := range Value {
					HeaderValue += v + ","
				}
			} else {
				HeaderValue = Value[0]
			}
			c.Writer.Header().Set(Header, HeaderValue)
		}

		// 返回文件
		c.Data(http.StatusOK, ContentType, FileBuffer.Bytes())
	})
}
