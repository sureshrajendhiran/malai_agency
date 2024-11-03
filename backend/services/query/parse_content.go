package query

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func ParseFieldBasedContent(description string) (string, string, bool) {
	originalContent := description
	noStyleString := ""
	resultContent := ""
	isQuery := false
	if description != "" {
		re := regexp.MustCompile(`{{`)
		start := re.FindAllStringSubmatchIndex(originalContent, -1)
		re = regexp.MustCompile(`}}`)
		end := re.FindAllStringSubmatchIndex(originalContent, -1)
		tableName := ""
		if len(start) == len(end) && len(start) != 0 && len(end) != 0 {
			isQuery = true
			list := make([]interface{}, 0)
			for index, _ := range start {

				startIndex := start[index]
				endIndex := end[index]

				if index == 0 {
					resultContent = "'" + originalContent[0:start[index][0]] + "' "
					noStyleString = "'" + originalContent[0:start[index][0]] + "'"
				}

				subSqlQuery, tempTableName := parseContent(originalContent, startIndex, endIndex, index)
				if tableName == "" {
					tableName = tempTableName
				}
				resultContent = resultContent + `,' <span style="font-weight: 500;">'` + ", ( " + subSqlQuery + " ) ,'</span>' "
				noStyleString = noStyleString + ",(" + subSqlQuery + ")"
				if index < len(start)-1 {
					resultContent = resultContent + ",'" + originalContent[end[index][1]:start[index+1][0]] + "'"
					noStyleString = noStyleString + ",'" + originalContent[end[index][1]:start[index+1][0]] + "'"
				}
				if index == len(start)-1 {
					resultContent = resultContent + ",'" + originalContent[end[index][1]:len(originalContent)] + "')"
					noStyleString = noStyleString + ",'" + originalContent[end[index][1]:len(originalContent)] + "')"
				}
				list = append(list, subSqlQuery)
			}
			resultContent = "SELECT CONCAT(" + resultContent + " FROM `" + tableName + "` d WHERE id IN ('*$#@!*')"
			noStyleString = "SELECT CONCAT(" + noStyleString + " FROM `" + tableName + "` d WHERE id IN ('*$#@!*')"

		} else {
			isQuery = false
			resultContent = originalContent
			noStyleString = originalContent
		}

	}

	notification := resultContent
	normalnotification := noStyleString

	// fmt.Println(notification, normalnotification, isQuery)
	return notification, normalnotification, isQuery
}

func parseContent(org string, startIndex []int, endIndex []int, occurence int) (string, string) {
	valueList := strings.Split(org[startIndex[1]:endIndex[0]], ".")
	result := ""
	tableName := fmt.Sprint(valueList[0])
	sqlQuery := "SELECT JSON_OBJECT('advanced',`advanced`,'type',`type`) FROM `meta_data` WHERE `data_table_name`='" + fmt.Sprint(valueList[0]) + "' AND " +
		" `name`='" + fmt.Sprint(valueList[1]) + "';"
	data, err := SqlJsonToMap(sqlQuery)
	if err != nil {

	}
	DataType := ""

	if data["type"] != nil {
		DataType = data["type"].(string)
	}
	if data["advanced"] != nil {
		isAdvancedObj := make(map[string]interface{})
		_ = json.Unmarshal([]byte(fmt.Sprint(data["advanced"])), &isAdvancedObj)

		if (isAdvancedObj["type"] == "depended" || isAdvancedObj["type"] == "non-depended" ||
			isAdvancedObj["type"] == "auto-fill" || isAdvancedObj["type"] == "depended and auto-fill") && DataType != "user" {
			fieldNameA := getAdvancedField(isAdvancedObj)
			if fieldNameA != "" {
				result = "SELECT GROUP_CONCAT(`" + fieldNameA + "`) FROM `" + fmt.Sprint(isAdvancedObj["table_name"]) +
					"` WHERE FIND_IN_SET (`id`,d.`" + valueList[1] + "`)"
			} else {
				result = "SELECT GROUP_CONCAT(`" + valueList[1] + "`)"
			}
		} else if DataType == "user" {
			result = "SELECT GROUP_CONCAT(`user_name`) FROM `Users` WHERE FIND_IN_SET (`id`,d.`" + valueList[1] + "`)"
		} else {
			result = "SELECT GROUP_CONCAT(`" + valueList[1] + "`)"
		}

	} else {
		result = "SELECT GROUP_CONCAT(`" + valueList[1] + "`)"
	}

	result = "IF((`" + valueList[1] + "`) IS NULL OR '' ,'',(" + result + "))"
	return result, tableName
}

func getAdvancedField(obj map[string]interface{}) string {
	fieldName := ""
	if obj["field_list"] != nil && len(obj["field_list"].([]interface{})) != 0 {
		fieldInfo := obj["field_list"].([]interface{})[0].(map[string]interface{})
		if fieldInfo["name"] != nil && fmt.Sprint(fieldInfo["name"]) != "" {
			fieldName = fmt.Sprint(fieldInfo["name"])
		}
	}
	return fieldName
}

func SqlQueryContentToString(context string, noStyleContext string, isQuery bool, rowId string) (string, string) {
	context = strings.Replace(context, "*$#@!*", rowId, -1)

	noStyleContext = strings.Replace(noStyleContext, "*$#@!*", rowId, -1)

	if isQuery {
		sqlQuery := "SELECT JSON_OBJECT('context',(" + context + "),'no_style_context',(" + noStyleContext + "))"
		data, err := SqlJsonToMap(sqlQuery)
		if err != nil {
			return "", ""
		}
		if data != nil && data["context"] != nil && data["no_style_context"] != nil {
			context = fmt.Sprint(data["context"])
			noStyleContext = fmt.Sprint(data["no_style_context"])

		}
	}

	return context, noStyleContext
}
