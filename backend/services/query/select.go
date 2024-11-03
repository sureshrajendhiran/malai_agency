package query

import (
	"database/sql"
	"malai_agency/backend/services/common"
	"malai_agency/backend/services/db"
	"malai_agency/backend/services/logs"
	"strings"
)

var DB *sql.DB

func init() {
	DB = db.MysqlConObj()
}

// SqlJsonToMap sql query to map only not array
func SqlJsonToMap(sql string) (map[string]interface{}, error) {
	if DB == nil {
		DB = db.MysqlConObj()
	}
	if DB != nil {
		selectObj, err := DB.Query(sql)
		if err != nil {
			logs.Logs("(SqlJsonToMap) select error ", err, sql)
			return nil, err
		}
		defer selectObj.Close()
		data := make(map[string]interface{})
		if selectObj.Next() {
			var str string
			selectObj.Scan(&str)
			data = common.StringToMap(str)
			return data, nil
		}
	}
	return nil, nil
}

// SqlJsonToArray sql query to get Array format data
func SqlJsonToArray(sql string) ([]interface{}, error) {
	if DB == nil {
		DB = db.MysqlConObj()
	}
	if DB != nil {
		selectObj, err := DB.Query(sql)
		if err != nil {
			logs.Logs("(SqlJsonToArray) select error ", err, sql)
			return nil, err
		}
		defer selectObj.Close()
		data := make([]interface{}, 0)
		for selectObj.Next() {
			var str string
			obj := make(map[string]interface{})
			selectObj.Scan(&str)
			obj = common.StringToMap(str)
			data = append(data, obj)
		}
		return data, nil
	}
	return nil, nil
}

// SingleValueBased used to get single value select like count,id,
func SingleValueBased(sql string) (string, error) {
	if DB == nil {
		DB = db.MysqlConObj()
	}
	if DB != nil {
		selectObj, err := DB.Query(sql)
		if err != nil {
			logs.Logs("(SqlJsonToMap) select error ", err, sql)
			return "", err
		}
		defer selectObj.Close()
		if selectObj.Next() {
			var strValue string
			selectObj.Scan(&strValue)
			return strValue, nil
		}
		return "", nil
	}
	return "", nil
}

// SingleValueArrayBased used to get single array value select like count,id,
func SingleValueArrayBased(sql string) ([]interface{}, error) {
	if DB == nil {
		DB = db.MysqlConObj()
	}
	if DB != nil {
		selectObj, err := DB.Query(sql)
		if err != nil {
			logs.Logs("(SqlJsonToMap) select error ", err, sql)
			return nil, err
		}
		arrayList := make([]interface{}, 0)
		defer selectObj.Close()
		for selectObj.Next() {
			var strValue string
			selectObj.Scan(&strValue)
			arrayList = append(arrayList, strValue)
		}
		return arrayList, nil
	}
	return nil, nil
}

// SqlCount sql query to count
func SqlCount(sql string) (int, error) {
	if DB == nil {
		DB = db.MysqlConObj()
	}
	if DB != nil {
		selectObj, err := DB.Query(sql)
		if err != nil {
			logs.Logs("(SqlJsonToMap) select error ", err, sql)
			return 0, err
		}
		defer selectObj.Close()
		if selectObj.Next() {
			var count int
			selectObj.Scan(&count)
			return count, nil
		}
		return 0, nil
	}
	return 0, nil
}

func QueryToId(sql string) (string, error) {
	selectObj, err := DB.Query(sql)
	if err != nil {
		logs.Logs("(QueryToId) select error", err)
		return "", err
	}
	defer selectObj.Close()
	var idsString string
	for selectObj.Next() {
		var id string
		selectObj.Scan(&id)
		if idsString == "" {
			idsString = id
		} else {
			idsString = idsString + "," + id
		}
	}

	return idsString, nil

}

func Query(sql string) error {
	selectObj, err := DB.Query(sql)
	defer selectObj.Close()
	if err != nil {
		logs.Logs("(QueryToId) select error", err)
		return err
	}
	return nil
}

func QueryToCount(sql string) (int, error) {
	// fmt.Println(sql)
	str := strings.Split(sql, " FROM ")[1]
	sqlQuery := "select count(*) from " + strings.Split(str, " LIMIT ")[0]
	selectObj, err := DB.Query(sqlQuery)
	defer selectObj.Close()
	if err != nil {
		return 0, err
	}
	if selectObj.Next() {
		var count int
		selectObj.Scan(&count)
		return count, nil
	}
	return 0, nil
}
