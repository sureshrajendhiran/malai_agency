package query

import "malai_agency/backend/services/logs"

// Update func used to update query
func Delete(sqlQuery string, valueList []interface{}) error {
	delete, err := DB.Prepare(sqlQuery)
	if err != nil {
		logs.Logs("(Delete) delete query Prepare error", err)
		return err
	}
	_, err = delete.Exec(valueList...)
	if err != nil {
		logs.Logs("(Delete) delete query Exec error", err)
		return err
	}
	return nil
}
