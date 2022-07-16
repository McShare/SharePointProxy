/*
 * @Author: NyanCatda
 * @Date: 2022-07-16 15:32:16
 * @LastEditTime: 2022-07-16 16:54:30
 * @LastEditors: NyanCatda
 * @Description: SharePoint代理模块
 * @FilePath: \SharePointProxy\Internal\SharePointProxy\GetFile.go
 */
package SharePointProxy

import (
	"bytes"
	"net/http"

	"github.com/McShare/SharePointProxy/Internal/Config"
	"github.com/nyancatda/AyaLog"
	"github.com/nyancatda/HttpRequest"
)

/**
 * @description: 获取文件
 * @param {string} Path 路径 /xxxx/xxxx/xxxx?xxxxx=xxxxx
 * @return {*bytes.Buffer} 文件内容
 * @return {*http.Response} 响应体
 * @return {error} 错误
 */
func GetFile(Path string) (*bytes.Buffer, *http.Response, error) {
	Config := Config.Get

	// 组成URL
	URL := "https://" + Config.SharePoint.Host + Path
	AyaLog.DeBug("GetFile", URL)

	// 发起请求
	Body, HttpResponse, err := HttpRequest.GetRequest(URL, []string{})
	if err != nil {
		return nil, nil, err
	}

	// 返回信息流储存至Buffer
	PhotoBuffer := bytes.NewBuffer(Body)

	return PhotoBuffer, HttpResponse, nil
}
