package delete

import (
	"malai_agency/backend/services/logs"
	"malai_agency/backend/services/query"
)

// RemoveRow func used to delete row from table
func RemoveRow(tableName string, ID string) error {
	sql := "DELETE FROM `" + tableName + "` WHERE `id`=?"
	err := query.Delete(sql, []interface{}{ID})
	if err != nil {
		logs.Logs("(RemoveRow) delete row data error ", err)
	}
	return err
}
