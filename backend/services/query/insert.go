package query

import (
	"fmt"
	"malai_agency/backend/services/logs"
)

// Insert func used to create new row
func Insert(sqlQuery string, valueList []interface{}) (int, error) {
	insert, err := DB.Prepare(sqlQuery)
	if err != nil {
		fmt.Println("sssss", err)
		logs.Logs("(InsertQuery) Insert query Prepare error", err)
		return 0, err
	}
	result, err := insert.Exec(valueList...)
	if err != nil {
		logs.Logs("(InsertQuery) Insert query Exec error", err, valueList)
		return 0, err
	}
	defer insert.Close()
	rowsAffectedCount, err := result.LastInsertId()
	return int(rowsAffectedCount), nil
}

// InsertWithMap
func InsertWithMap(reqBody map[string]interface{}) (int, error) {
	sql := "INSERT INTO `" + reqBody["table_name"].(string) + "` ("
	delete(reqBody, "table_name")
	insertField := ""
	insertFieldQ := ""
	insertList := make([]interface{}, 0)
	for key, value := range reqBody {
		if insertField != "" {
			insertField = insertField + ",`" + key + "`"
			insertFieldQ = insertFieldQ + ",?"
		} else {
			insertField = "`" + key + "`"
			insertFieldQ = insertFieldQ + "?"
		}
		insertList = append(insertList, value)
	}
	sql = sql + insertField + ") VALUES (" + insertFieldQ + ")"
	return Insert(sql, insertList)
}

// Insert Query parse based on input
func InsertQueryParse(data map[string]interface{}) (string, []interface{}) {
	sqlQuery := "INSERT INTO `" + fmt.Sprint(data["table_name"]) + "` ("
	valueList := make([]interface{}, 0)
	sqlQueryRHS := "("
	delete(data, "table_name")
	len := len(data)
	count := 0
	for key, value := range data {
		if count == len-1 {
			sqlQuery = sqlQuery + "`" + fmt.Sprint(key) + "`"
			sqlQueryRHS = sqlQueryRHS + "?"
		} else {
			sqlQuery = sqlQuery + "`" + fmt.Sprint(key) + "`,"
			sqlQueryRHS = sqlQueryRHS + "?,"
		}
		valueList = append(valueList, value)
		count = count + 1
	}
	sqlQueryRHS = sqlQueryRHS + ")"
	sqlQuery = sqlQuery + ") VALUE " + sqlQueryRHS
	// fmt.Println(sqlQuery, valueList)
	return sqlQuery, valueList
}
