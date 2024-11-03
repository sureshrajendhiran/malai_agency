package get

import (
	"database/sql"
	"fmt"
	"malai_agency/backend/services/db"
	"malai_agency/backend/services/query"
	"strconv"
	"strings"
)

var DB *sql.DB

func init() {
	DB = db.MysqlConObj()
}

func GetInvoice(requestObj map[string]interface{}) ([]interface{}, int, error) {
	res := make([]interface{}, 0)
	var err error
	sql := ""
	if requestObj["type"].(string) == "invoice" {
		sql = "SELECT JSON_OBJECT('id',`id`,'date',`date`,'ref_number',`ref_number`,'status',`status`," +
			"'customer_name',`customer_name`,'customer_address',`customer_address`," +
			"'total',`total`," +
			"'item_list',(SELECT JSON_ARRAYAGG(JSON_OBJECT(" +
			"'id',`id`,'item_name',`item_name`,'unit',`unit`,'rate_per_item',`rate_per_item`,'total',`total`,'hsn_code',`hsn_code`" +
			")) from `S_Invoice_Items` where `922_id`=si.`id`)" +
			") FROM `S_Invoice` si "
	} else {
		sql = "SELECT JSON_OBJECT('id',`id`,'date',`date`,'ref_number',`ref_number`,'status',`status`," +
			"'customer_name',`customer_name`,'customer_address',`customer_address`," +
			"'total',`total`," +
			"'item_list',(SELECT JSON_ARRAYAGG(JSON_OBJECT(" +
			"'id',`id`,'item_name',`item_name`,'unit',`unit`,'rate_per_item',`rate_per_item`,'total',`total`,'hsn_code',`hsn_code`" +
			")) from `S_Quotation_Items` where `926_id`=sq.`id`)" +
			") FROM `S_Quotation` sq "
	}
	sql = sql + filterParse(requestObj)
	res, err = query.SqlJsonToArray(sql)
	for _, i := range res {
		obj := i.(map[string]interface{})
		if obj["item_list"] == nil {
			obj["item_list"] = make([]interface{}, 0)
		}
	}
	if err != nil {

	}
	count := QueryToCount(sql)

	return res, count, err
}

func filterParse(requestObj map[string]interface{}) string {
	str := ""
	if requestObj["filter"] != nil {
		filter := requestObj["filter"].(map[string]interface{})
		name := strings.ToLower(fmt.Sprint(filter["name"]))
		str = " WHERE "
		if name == "cancelled" {
			str = str + " `status` IN ('cancelled') "
		} else if name == "all" {
			str = str + " `id` IS NOT NULL  "
		} else if name == "paid" {
			str = str + " `status` IN ('paid') "
		} else if name == "sent" {
			str = str + " `status` IN ('sent') "
		} else if name == "pending" {
			str = str + " `status` IN ('pending') "
		} else if name == "pending" {
			str = str + " `status` IN ('pending') "
		} else if name == "sent and pending" {
			str = str + " `status` IN ('sent and pending') "
		} else if name == "approved" {
			str = str + " `status` IN ('approved') "
		}
	}
	str = str + sort(requestObj)
	str = str + Limit(requestObj)
	return str
}
func Limit(requestObj map[string]interface{}) string {
	str := ""
	if requestObj["limit"] != nil {
		str = str + " LIMIT " + fmt.Sprint(requestObj["limit"]) + " "
		if requestObj["page"] != nil {
			str = str + " OFFSET " + Page(fmt.Sprint(requestObj["limit"]), fmt.Sprint(requestObj["page"]))

		}
	}
	return str
}

func sort(requestObj map[string]interface{}) string {
	str := ""
	if requestObj["sort_type"] != nil {
		s := fmt.Sprint(requestObj["sort_type"])
		if s == "" {
			str = " ORDER BY `" + fmt.Sprint(requestObj["sort_field"]) +
				"` " + fmt.Sprint(requestObj["sort_type"])
		}

	} else {
		str = " ORDER BY `id` DESC "
	}
	return str
}

func Page(limit string, page string) string {
	p, _ := strconv.Atoi(page)
	l, _ := strconv.Atoi(limit)
	return fmt.Sprint(p * l)
}

func QueryToCount(sql string) int {
	str := strings.Split(sql, " FROM ")[1]
	sqlQuery := "select JSON_OBJECT('count',count(*)) from " + strings.Split(str, " LIMIT ")[0]
	data, _ := query.SqlJsonToMap(sqlQuery)
	if data != nil {
		return int(data["count"].(float64))
	}
	return 0
}

func GetInvoiceById(requestObj map[string]interface{}) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	var err error
	sql := "SELECT JSON_OBJECT('id',`id`,'date',`date`,'invoice_no',`invoice_no`,'status',`status`," +
		"'customer_name',`customer_name`,'customer_address',`customer_address`," +
		"'total',`total`,'items'," +
		"(SELECT JSON_ARRAYAGG(JSON_OBJECT('id',`id`,'rate_description',`rate_description`,'description',`description`,'hsn_code',`hsn_code`," +
		"'quantity',`quantity`,'rate_per_item',`rate_per_item`,'total',`total`" +
		")) FROM `S_Invoice_Items` WHERE `914_id`='" + fmt.Sprint(requestObj["id"]) + "')" +
		") FROM `S_Invoice` WHERE `id`='" + fmt.Sprint(requestObj["id"]) + "'"

	res, err = query.SqlJsonToMap(sql)

	if err != nil {

	}
	if res["items"] == nil {
		res["items"] = make([]interface{}, 0)
	}
	return res, err
}
