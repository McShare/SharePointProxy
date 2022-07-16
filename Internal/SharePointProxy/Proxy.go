/*
 * @Author: NyanCatda
 * @Date: 2022-07-16 16:35:56
 * @LastEditTime: 2022-07-16 16:55:37
 * @LastEditors: NyanCatda
 * @Description: 反向代理模块
 * @FilePath: \SharePointProxy\Internal\SharePointProxy\Proxy.go
 */
package SharePointProxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/McShare/SharePointProxy/Internal/Config"
	"github.com/gin-gonic/gin"
	"github.com/nyancatda/AyaLog"
)

/**
 * @description: 反向代理模块
 * @param {*gin.Context} c 请求上下文
 * @return {*}
 */
func Proxy(c *gin.Context) {
	Config := Config.Get

	u := &url.URL{}

	u.Scheme = "https"                  //转发的协议
	u.Host = Config.SharePoint.Host     //转发的主机
	u.Path = c.Request.URL.Path         //转发的路径
	u.RawQuery = c.Request.URL.RawQuery //转发的参数

	proxy := httputil.NewSingleHostReverseProxy(u)

	//重写出错回调
	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
		AyaLog.Error("Request", err)

		//写到body里
		rw.Write([]byte(err.Error()))
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
