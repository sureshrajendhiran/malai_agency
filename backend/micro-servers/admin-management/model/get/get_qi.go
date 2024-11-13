package get

import (
	"fmt"
	"malai_agency/backend/services/query"
)

func GetQI(requestObj map[string]interface{}) (interface{}, error) {
	var err error
	data, _ := getData(requestObj)
	res := GenerateHTML(requestObj, data)
	return res, err
}

func getData(requestObj map[string]interface{}) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	var err error
	sql := ""
	if requestObj["type"].(string) == "invoice" {
		sql = "SELECT JSON_OBJECT('id',`id`,'date',`date`,'ref_number',`ref_number`,'status',`status`,'tax_type',`tax_type`," +
			"'customer_name',`customer_name`,'customer_address',`customer_address`," +
			"'total',`total`," +
			"'item_list',(SELECT JSON_ARRAYAGG(JSON_OBJECT(" +
			"'id',`id`,'item_name',`item_name`,'unit',`unit`,'rate_per_item',`rate_per_item`,'total',`total`,'hsn_code',`hsn_code`" +
			")) FROM `S_Invoice_Items` WHERE `922_id`=si.`id`)" +
			") FROM `S_Invoice` si WHERE si.`id`='" + fmt.Sprint(requestObj["id"]) + "'"
	} else {
		sql = "SELECT JSON_OBJECT('id',`id`,'date',`date`,'ref_number',`ref_number`,'status',`status`,'tax_type',`tax_type`," +
			"'customer_name',`customer_name`,'customer_address',`customer_address`," +
			"'total',`total`," +
			"'item_list',(SELECT JSON_ARRAYAGG(JSON_OBJECT(" +
			"'id',`id`,'item_name',`item_name`,'unit',`unit`,'rate_per_item',`rate_per_item`,'total',`total`,'hsn_code',`hsn_code`" +
			")) FROM `S_Quotation_Items` WHERE `926_id`=sq.`id`)" +
			") FROM `S_Quotation` sq WHERE sq.`id`='" + fmt.Sprint(requestObj["id"]) + "'"
	}

	res, err = query.SqlJsonToMap(sql)
	if res["item_list"] == nil {
		res["item_list"] = make([]interface{}, 0)
	}

	return res, err
}
