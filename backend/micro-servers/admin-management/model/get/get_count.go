package get

import "malai_agency/backend/services/query"

func GetFilterCount(requestObj map[string]interface{}) ([]interface{}, error) {
	res := make([]interface{}, 0)
	var err error
	sql := ""
	if requestObj["type"].(string) == "invoice" {
		sql = "SELECT JSON_OBJECT(" +
			"'All',COUNT(*)," +
			"'Pending',SUM(IF(`status` In ('pending'),1,0))," +
			"'Sent',SUM(IF(`status` In ('sent'),1,0))," +
			"'Sent and Pending',SUM(IF(`status` In ('Sent and Pending'),1,0))," +
			"'Paid',SUM(IF(`status` In ('paid'),1,0))," +
			"'Cancelled',SUM(IF(`status` In ('Cancelled'),1,0))" +
			") FROM `S_Invoice` "
	} else if requestObj["type"].(string) == "quotation" {
		sql = "SELECT JSON_OBJECT(" +
			"'All',COUNT(*)," +
			"'Pending',SUM(IF(`status` In ('pending'),1,0))," +
			"'Sent',SUM(IF(`status` In ('sent'),1,0))," +
			"'Approved',SUM(IF(`status` In ('Approved'),1,0))," +
			"'Cancelled',SUM(IF(`status` In ('Cancelled'),1,0))" +
			") FROM `S_Quotation` "
	} else {
		sql = "SELECT JSON_OBJECT(" +
			"'users',(SELECT COUNT(*) FROM `S_Users`)," +
			"'item master',(SELECT COUNT(*) FROM `S_Item_Master`)," +
			"'customer master',(SELECT COUNT(*) FROM `S_Customer_Master`)" +
			")"
	}
	if sql != "" {
		tempRes, err := query.SqlJsonToMap(sql)
		if err != nil {
		}
		for k, v := range tempRes {
			temp := make(map[string]interface{})
			temp["name"] = k
			temp["count"] = v
			res = append(res, temp)
		}
	}

	return res, err
}
