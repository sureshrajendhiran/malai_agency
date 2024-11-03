package get

import (
	"fmt"
	"malai_agency/backend/services/query"
)

func SearchOption(requestObj map[string]interface{}) ([]interface{}, error) {
	res := make([]interface{}, 0)
	var err error
	sql := ""
	if requestObj["type"].(string) == "rate_items" {
		sql = "SELECT JSON_OBJECT('id',`id`,'hsn_code',`hsn_code`,'item_name',`item_name`,'description',`description`,'rate_per_item',`rate_per_item`) FROM `S_Item_Master` " +
			" WHERE `item_name` LIKE '%%" + fmt.Sprint(requestObj["q"]) + "%%'"
	} else if requestObj["type"].(string) == "customers" {
		sql = "SELECT JSON_OBJECT('id',`id`,'customer_name',`customer_name`,'address',`address`,'notes',`notes`) FROM `S_Customer_Master` " +
			" WHERE `customer_name` LIKE '%%" + fmt.Sprint(requestObj["q"]) + "%%' OR `address` LIKE '%%" + fmt.Sprint(requestObj["q"]) + "%%'"
	}
	if sql != "" {
		res, err = query.SqlJsonToArray(sql)
		if err != nil {
			return res, err
		}
	}

	return res, err
}
