package login

import (
	"encoding/json"
	"malai_agency/backend/response"
	"malai_agency/backend/services/auth"
	"malai_agency/backend/services/query"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// NewPasswordChangeCtrl function will give you the list of users
func NewPasswordChangeCtrl(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	//decode the request body into struct and failed if any error occur
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.ResponseError(w, err, "invalid_json")
		return
	}
	// New password model
	callbackCode, err := NewPasswordChange(data)
	if err != nil {
		response.ResponseError(w, err, callbackCode)
		return
	}

	if callbackCode != "" {
		response.ResponseError(w, nil, callbackCode)
		return
	}

	response.Response(w, nil, 200)
	return
}

func NewPasswordChange(data map[string]interface{}) (string, error) {
	// Token validate
	callback, userID, _, _ := auth.ProtectedEndpoint(data["token"].(string))
	if !callback {
		return "invalid_token_link", nil
	}

	s, err := query.SingleValueBased("SELECT `reset_token` FROM `Portal_Users`  WHERE `id`=" + userID + " AND `reset_token`='" + data["token"].(string) + "'")
	if err != nil {
		return "db_error", err
	}

	if s == "" {
		return "invalid_token_link", nil
	}

	// Generated new Hash string
	newHashString, err := bcrypt.GenerateFromPassword([]byte(data["password"].(string)), bcrypt.DefaultCost)
	if err != nil {
		return "db_error", err
	}
	// New password updated to
	sql := "UPDATE `Portal_Users` SET `password`=?,`reset_token`=?,`open_invite`=? WHERE `id` =?"
	_, err = query.Update(sql, []interface{}{newHashString, nil, 0, userID})
	if err != nil {
		return "db_error", err
	}

	return "", nil
}
