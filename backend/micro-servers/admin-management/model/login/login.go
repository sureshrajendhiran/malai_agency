package login

import (
	"encoding/json"
	"malai_agency/backend/response"
	"malai_agency/backend/services/query"

	"flowpod_server/utils/auth"
	"flowpod_server/utils/custom_errors"
	"flowpod_server/utils/session"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

func Login(r *http.Request, w http.ResponseWriter) (interface{}, error) {
	var data map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&data)
	responseObj := make(map[string]interface{})
	data["current_time"] = r.Header["Current-Time"][0]
	// var email, password string
	email := data["email"].(string)
	password := data["password"].(string)

	sqlQuery := "SELECT JSON_OBJECT('id',`id`,'user_name',`user_name`,'email',`email`," +
		"'password',`password`,'status',`status`)  FROM S_Users u WHERE email='" + email + "'"
	UserData, err := query.SqlJsonToMap(sqlQuery)
	if err != nil {
		return nil, err
	}

	if UserData == nil || UserData["id"] == nil {
		assignMap(response.ErrorConst["invalid_email"], responseObj)
		responseObj["success"] = false
		return responseObj, nil
	} else {
		// Email address valid
		userInfo := session.UserInfoDetails{}
		// Scan data from DB
		userInfo.Id = int(UserData["id"].(float64))
		userInfo.UserName = UserData["user_name"].(string)
		userInfo.Email = UserData["email"].(string)
		userInfo.Password = UserData["password"].(string)
		userInfo.Status = UserData["status"].(string)

		// Password check
		if userInfo.Password == password {
			// Check user active or not

			if strings.ToLower(userInfo.Status) == "active" {
				resData := loginSuccess(r, fmt.Sprint(userInfo.Id))
				assignMap(resData, responseObj)
			} else {
				// Account blocked
				return accountBlockMsg(), nil
			}
		} else {
			// Number of attempt limit check
			return invalidPasswordMsg(int(UserData["failed_attempt"].(float64)), userInfo.Id), nil
		}

	}
	return responseObj, nil

}

func loginSuccess(r *http.Request, userId string) map[string]interface{} {
	userInfo := make(jwt.MapClaims)
	responseObj := make(map[string]interface{})
	sql := "(SELECT JSON_OBJECT('id',user.`id`,'user_name',user.`user_name`,'email',user.`email`) FROM `S_Users` user WHERE id IN(" + userId + "))"
	userData, _ := query.SqlJsonToMap(sql)
	userInfo["email"] = userData["email"]
	userInfo["id"] = userData["id"]
	userInfo["user_name"] = userData["user_name"]
	token := auth.CreateTokenEndpointNew(userInfo)
	// Login logs used to insert into Login_logs table
	userLogInput := make(map[string]interface{})
	userLogInput["user"] = userData["id"]
	userLogInput["token"] = token
	userLogInput["status"] = "success"
	userLogInput["description"] = "Successfully login"

	responseObj["token"] = token
	responseObj["user_info"] = userInfo

	return responseObj
}

// Account block message
func accountBlockMsg() interface{} {
	responseObj := make(map[string]interface{})
	assignMap(custom_errors.AccoutBlock, responseObj)
	responseObj["success"] = false
	return responseObj
}

// Incorrect Password
func invalidPasswordMsg(count int, userId int) interface{} {
	responseObj := make(map[string]interface{})
	assignMap(custom_errors.InvalidPassword, responseObj)
	// responseObj["error_message"] = responseObj["error_message"].(string) + fmt.Sprint(". ", 3-count, " more attempts remaining.")
	// Update count

	responseObj["success"] = false
	return responseObj
}

// Change accout status in-active
func accountBlock(userID int) {
	sql := "UPDATE `Users` SET `status`=? WHERE id=?"
	valueList := make([]interface{}, 0)
	valueList = append(valueList, "in-active")
	valueList = append(valueList, userID)
	query.Update(sql, valueList)
}

func assignMap(from map[string]interface{}, to map[string]interface{}) {
	for key, value := range from {
		to[key] = value
	}
	return
}
