package response

import (
	"encoding/json"
	"net/http"
)

//Response func used to return user response with status code like 200
func Response(w http.ResponseWriter, data interface{}, statusCode int) {
	res := Message(statusCode)
	if data != nil {
		res["info"] = data
	}
	jsonStr, _ := json.Marshal(res)
	w.WriteHeader(statusCode)
	w.Write(jsonStr)
	return
}

//ResponseWithCount func used to return user response with status code like 200
func ResponseWithCount(w http.ResponseWriter, data interface{}, count interface{}, statusCode int) {
	res := Message(statusCode)
	res["info"] = data
	res["count"] = count
	jsonStr, _ := json.Marshal(res)
	w.WriteHeader(statusCode)
	w.Write(jsonStr)
	return
}

//ResponseError func used to return user response with status code like 200
func ResponseError(w http.ResponseWriter, err error, errorCode string) {
	res := (ErrorConst[errorCode])
	res["error"] = err
	res["statusCode"] = ErrorConst[errorCode]["http_code"].(int)
	res["statusMessage"] = http.StatusText(ErrorConst[errorCode]["http_code"].(int))
	jsonStr, _ := json.Marshal(res)
	w.WriteHeader(ErrorConst[errorCode]["http_code"].(int))
	w.Write(jsonStr)
	return
}

// JsonError  func used to return invalid json body error
func JsonError(w http.ResponseWriter) {
	data := Message(400)
	data["error_code"] = JsonErrorCode
	data["error_message"] = "Invalid JSON Body."
	jsonStr, _ := json.Marshal(data)
	w.WriteHeader(400)
	w.Write(jsonStr)
	return
}

//Message func used to create body based on http.StatusCode like (200,201,500)
func Message(code int) map[string]interface{} {
	return map[string]interface{}{"statusCode": code, "statusMessage": http.StatusText(code)}
}
