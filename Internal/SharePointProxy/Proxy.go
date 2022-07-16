/*
 * @Author: NyanCatda
 * @Date: 2022-07-16 16:35:56
 * @LastEditTime: 2022-07-16 17:16:01
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
 * @description: 反向代理模块(废弃方案)
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

	// 修改请求
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		ModifyRequest(req)
	}

	//出错回调
	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
		AyaLog.Error("Request", err)

		//写到body里
		rw.Write([]byte(err.Error()))
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

/**
 * @description: 修改请求
 * @param {*http.Request} req 请求
 * @return {*}
 */
func ModifyRequest(req *http.Request) {
	// 添加请求头
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("accept-encoding", "gzip, deflate, br")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("if-none-match", `"{F26DB937-4E11-40EF-BF3F-E0FEB46D3AB3},4"`)
	req.Header.Set("sec-ch-ua", `" Not;A Brand";v="99", "Microsoft Edge";v="103", "Chromium";v="103"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("service-worker-navigation-preload", "true")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62")
}
