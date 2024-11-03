package admin_management

import (
	"encoding/json"
	"fmt"
	"malai_agency/backend/micro-servers/admin-management/model/create"
	"malai_agency/backend/micro-servers/admin-management/model/delete"
	"malai_agency/backend/micro-servers/admin-management/model/get"
	"malai_agency/backend/micro-servers/admin-management/model/login"
	"malai_agency/backend/micro-servers/admin-management/model/update"
	"malai_agency/backend/response"
	"malai_agency/backend/services/logs"
	"net/http"

	"github.com/gorilla/mux"
)

func LoginCtrl(w http.ResponseWriter, r *http.Request) {
	data, err := login.Login(r, w)
	if err != nil {
		response.Response(w, data, 500)
	}

	response.Response(w, data, 200)

	return
}

func GetInvoiceCtrl(w http.ResponseWriter, r *http.Request) {
	requestObj := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&requestObj)
	if err != nil {
		response.ResponseError(w, err, "invalid_json")
		return
	}
	requestObj["type"] = mux.Vars(r)["type"]
	requestObj["current_time"] = r.Header["Current-Time"][0]
	res, count, err := get.GetInvoice(requestObj)
	if err != nil {
		return
	}
	response.ResponseWithCount(w, res, count, 200)
}

func GetMasterCtrl(w http.ResponseWriter, r *http.Request) {
	requestObj := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&requestObj)
	if err != nil {
		response.ResponseError(w, err, "invalid_json")
		return
	}
	requestObj["type"] = mux.Vars(r)["type"]
	requestObj["current_time"] = r.Header["Current-Time"][0]
	res, err := get.GetMasterData(requestObj)
	if err != nil {
		return
	}
	response.Response(w, res, 200)
}

func UpdateQICtrl(w http.ResponseWriter, r *http.Request) {
	requestObj := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&requestObj)
	requestObj["type"] = mux.Vars(r)["type"]
	requestObj["operation"] = mux.Vars(r)["operation"]
	if err != nil {
		response.ResponseError(w, err, "invalid_json")
		return
	}
	if fmt.Sprint(requestObj["operation"]) == "create" {
		requestObj["created_on"] = r.Header["Current-Time"][0]
	}
	requestObj["last_modified"] = r.Header["Current-Time"][0]
	info, err := update.UpdateInvoiceQuotation(requestObj)
	if err != nil {
		return
	}
	response.Response(w, info, 201)
}

func CreateNewRowCtrl(w http.ResponseWriter, r *http.Request) {
	requestObj := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&requestObj)
	if err != nil {
		response.ResponseError(w, err, "invalid_json")
		return
	}
	requestObj["created_on"] = r.Header["Current-Time"][0]
	requestObj["last_modified"] = r.Header["Current-Time"][0]
	res, err := create.CreateNewRecord(requestObj)
	if err != nil {
		return
	}
	response.Response(w, res, 201)
}

func UpdateCtrl(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	tableName := mux.Vars(r)["table_name"]
	// Json data decode
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		logs.Logs("(UpdateCtrl) json error ", err)
		response.ResponseError(w, err, "invalid_json")
		return
	}
	data["table_name"] = tableName
	_, err = update.UpdateWithMap(data)
	if err != nil {
		response.ResponseError(w, err, "db_error")
	}
	response.Response(w, nil, 200)
}

// DeleteCtrl func used to control the delete rowInfo
func DeleteCtrl(w http.ResponseWriter, r *http.Request) {
	tableName := mux.Vars(r)["table_name"]
	ID := mux.Vars(r)["id"]
	// Select and remove files
	// Remove row from table
	err := delete.RemoveRow(tableName, ID)
	if err != nil {
		response.ResponseError(w, err, "db_error")
		return
	}
	response.Response(w, nil, 200)
	return
}

func GetInvoiceByIdCtrl(w http.ResponseWriter, r *http.Request) {
	requestObj := make(map[string]interface{})
	requestObj["id"] = mux.Vars(r)["id"]
	requestObj["operation"] = mux.Vars(r)["operation"]
	requestObj["type"] = mux.Vars(r)["type"]
	str, err := get.GetQI(requestObj)
	if err != nil {
		response.ResponseError(w, err, "db_error")
	}
	response.Response(w, str, 200)
	return
}

func SearchOptionCtrl(w http.ResponseWriter, r *http.Request) {
	requestObj := make(map[string]interface{})
	queryPara := r.URL.Query()
	requestObj["type"] = mux.Vars(r)["type"]
	// Json data decode
	requestObj["q"] = queryPara.Get("q")
	requestObj["limit"] = queryPara.Get("limit")
	res, err := get.SearchOption(requestObj)

	if err != nil {
		response.ResponseError(w, err, "db_error")
	}
	response.Response(w, res, 200)
}

func FilterCountCtrl(w http.ResponseWriter, r *http.Request) {
	requestObj := make(map[string]interface{})
	requestObj["type"] = mux.Vars(r)["type"]
	res, err := get.GetFilterCount(requestObj)

	if err != nil {
		response.ResponseError(w, err, "db_error")
	}
	response.Response(w, res, 200)
}
