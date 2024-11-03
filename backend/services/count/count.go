package count

import (
	"malai_agency/backend/services/query"
	"strings"
)

var DB = query.DB

func QueryToCount(sql string) (int, error) {
	// fmt.Println(sql)
	str := strings.Split(sql, " FROM ")[1]
	// fmt.Println(str)
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
