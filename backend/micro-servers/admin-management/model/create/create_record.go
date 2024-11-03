package create

import (
	"fmt"
	"malai_agency/backend/services/logs"
	"malai_agency/backend/services/query"
)

func CreateNewRecord(data map[string]interface{}) (string, error) {

	sql, valueList := InsertQueryParse(data)
	fmt.Println(sql, valueList)
	lastId, err := query.Insert(sql, valueList)

	if err != nil {
		logs.Logs("(CreateNewRecord) insert value error", err)
		return "", err
	}

	return fmt.Sprint(lastId), nil
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
