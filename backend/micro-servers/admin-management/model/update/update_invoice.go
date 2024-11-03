package update

import (
	"fmt"
	"malai_agency/backend/services/query"

	"github.com/nleeper/goment"
)

func UpdateInvoiceQuotation(requestObj map[string]interface{}) (interface{}, error) {
	var err error
	var masterLastId interface{}
	Type := fmt.Sprint(requestObj["type"])
	operation := requestObj["operation"].(string)
	if requestObj["item_list"] == nil {
		requestObj["item_list"] = make([]interface{}, 0)
	}
	items := requestObj["item_list"].([]interface{})
	delete(requestObj, "item_list")
	delete(requestObj, "type")
	delete(requestObj, "operation")
	delete(requestObj, "date_t")
	if Type == "invoice" {
		requestObj["table_name"] = "S_Invoice"
		if operation == "create" {
			delete(requestObj, "id")
			if requestObj["ref_number"] == nil {
				requestObj["ref_number"] = getQINumber(requestObj, "invoice")
			}
			masterLastId, _ = query.InsertWithMap(requestObj)

		} else {
			masterLastId = fmt.Sprint(requestObj["id"])
			query.UpdateWithMap(requestObj)
		}
		fmt.Println(masterLastId)
		if len(items) != 0 {
			for _, i := range items {
				obj := i.(map[string]interface{})
				obj["table_name"] = "S_Invoice_Items"
				if obj["id"] != nil {
					query.UpdateWithMap(obj)
				} else {
					obj["922_id"] = masterLastId
					query.InsertWithMap(obj)
				}
			}
		}
	} else {
		requestObj["table_name"] = "S_Quotation"
		if operation == "create" {
			delete(requestObj, "id")
			if requestObj["ref_number"] == nil {
				requestObj["ref_number"] = getQINumber(requestObj, "quotation")
			}
			masterLastId, _ = query.InsertWithMap(requestObj)
		} else {
			masterLastId = fmt.Sprint(requestObj["id"])
			query.UpdateWithMap(requestObj)
		}
		if len(items) != 0 {
			for _, i := range items {
				obj := i.(map[string]interface{})
				obj["table_name"] = "S_Quotation_Items"
				if obj["id"] != nil {
					query.UpdateWithMap(obj)
				} else {
					obj["926_id"] = masterLastId
					obj["id"], _ = query.InsertWithMap(obj)
				}
			}
		}
	}

	return "", err
}

func getQINumber(requestObj map[string]interface{}, Type string) string {
	refNo := ""
	prefix := ""
	fName := ""
	if Type == "invoice" {
		prefix = "IN"
		fName = "invoice_count"
	} else {
		prefix = "QU"
		fName = "quote_number"
	}
	sql := "SELECT JSON_OBJECT('count',`" + fName + "`) FROM `S_Count_Manage` WHERE `id`=1"
	res, err := query.SqlJsonToMap(sql)
	if err != nil {

	}
	var count float64
	if res != nil && res["count"] != nil {
		count = res["count"].(float64)
	} else {
		count = 10000
	}
	count = count + 1
	date, _ := goment.New(requestObj["last_modified"])
	temp := date.Format("YY")
	refNo = prefix + "/" + temp + "/" + fmt.Sprint(count)
	sql = "UPDATE `S_Count_Manage` SET `" + fName + "`=? WHERE `id`=?"
	valueList := make([]interface{}, 0)
	valueList = append(valueList, count, 1)

	query.Update(sql, valueList)

	return refNo
}
