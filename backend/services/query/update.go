package query

import (
	"malai_agency/backend/services/db"
	"malai_agency/backend/services/logs"
)

// Update func used to update query
func Update(sqlQuery string, valueList []interface{}) (int, error) {
	update, err := db.MysqlConObj().Prepare(sqlQuery)
	if err != nil {
		logs.Logs("(Update) update query Prepare error", err)
		return 0, err
	}
	result, err := update.Exec(valueList...)
	if err != nil {
		logs.Logs("(Update) update query Exec error", err)
		return 0, err
	}
	rowsAffectedCount, _ := result.RowsAffected()
	return int(rowsAffectedCount), nil
}

// UpdateWithMap func used to update map based value
func UpdateWithMap(reqBody map[string]interface{}) (int, error) {
	sql := "UPDATE `" + reqBody["table_name"].(string) + "` SET "
	delete(reqBody, "table_name")
	updateField := ""
	updateList := make([]interface{}, 0)
	for key, value := range reqBody {
		if updateField != "" {
			updateField = updateField + ",`" + key + "`=?"
		} else {
			updateField = "`" + key + "`=?"
		}
		updateList = append(updateList, value)
	}
	sql = sql + updateField + " WHERE `id`=?"
	updateList = append(updateList, reqBody["id"])
	return Update(sql, updateList)
}
