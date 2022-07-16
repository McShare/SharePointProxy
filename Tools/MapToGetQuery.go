/*
 * @Author: NyanCatda
 * @Date: 2022-07-16 15:55:46
 * @LastEditTime: 2022-07-16 15:58:38
 * @LastEditors: NyanCatda
 * @Description: Map转换为GetQuery
 * @FilePath: \SharePointProxy\Tools\MapToGetQuery.go
 */
package Tools

import "fmt"

/**
 * @description: Map转换为GetQuery
 * @param {map[string]any} Map 参数Map
 * @return {string} GetQuery
 */
func MapToGetQuery(Map map[string]any) string {
	var Query string
	for Key, Value := range Map {
		Query += Key + "=" + fmt.Sprintf("%v", Value) + "&"
	}
	return Query[:len(Query)-1]
}
