package get

import (
	"fmt"
	"malai_agency/backend/services/query"
	"strings"
)

func GetMasterData(requestObj map[string]interface{}) ([]interface{}, error) {
	var err error
	res := make([]interface{}, 0)
	sql := ""
	Type := strings.ToLower(fmt.Sprint(requestObj["type"]))
	if Type == "item master" {
		sql = "SELECT JSON_OBJECT('id',`id`,'hsn_code',`hsn_code`,'item_name',`item_name`,'description',`description`,'rate_per_item',`rate_per_item`) FROM `S_Item_Master` "
	} else if Type == "customer master" {
		sql = "SELECT JSON_OBJECT('id',`id`,'customer_name',`customer_name`,'address',`address`,'notes',`notes`) FROM `S_Customer_Master` "
	} else if Type == "users" {
		sql = "SELECT JSON_OBJECT('id',`id`,'user_name',`user_name`,'email',`email`,'password',`password`) FROM `S_Users` "
	}
	if sql != "" {
		sql = sql + " ORDER BY `id` DESC "
		sql = sql + Limit(requestObj)
		res, err = query.SqlJsonToArray(sql)
	}

	return res, err
}
