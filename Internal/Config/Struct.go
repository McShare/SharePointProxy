/*
 * @Author: NyanCatda
 * @Date: 2022-07-16 15:18:13
 * @LastEditTime: 2022-07-16 16:41:04
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \SharePointProxy\Internal\Config\Struct.go
 */
package Config

type Config struct {
	Run struct {
		Debug bool   `yaml:"Debug"` //Debug模式
		Mode  string `yaml:"Mode"`  //运行模式，dev为开发者模式，输出更多信息，release为发行模式
	} `yaml:"Run"`
	SharePoint struct {
		Host string `yaml:"Host"` //SharePoint地址，例如: xxxxx-my.sharepoint.com
	} `yaml:"SharePoint"`
}
