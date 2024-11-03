package login

import (
	"back_end_servers/services/users"
	"fmt"
	"net/http"
	"strings"

	"github.com/mssola/user_agent"
	"golang.org/x/crypto/bcrypt"
)

func MapToUserDetails(value map[string]interface{}) users.CustomerUserDetails {
	var data users.CustomerUserDetails
	data.Id = int(value["id"].(float64))
	data.IdStr = fmt.Sprint(value["id"])
	data.Email = fmt.Sprint(value["email"])
	data.CustomerType = fmt.Sprint(value["customer_type"])
	data.PhoneNumber = fmt.Sprint(value["phone_number"])
	data.Role = fmt.Sprint(value["role"])
	data.ReferId = fmt.Sprint(value["refer_id"])
	data.Status = strings.ToLower(fmt.Sprint(value["status"]))
	data.Password = fmt.Sprint(value["password"])
	data.UserName = fmt.Sprint(value["user_name"])
	if value["photo"] != nil {
		data.Photo = fmt.Sprint(value["photo"])
	}
	return data
}

// Compare hash string and password string
func PasswordCompare(hashString string, passwordString string) bool {
	errf := bcrypt.CompareHashAndPassword([]byte(hashString), []byte(passwordString))
	// && errf == bcrypt.ErrMismatchedHashAndPassword
	if errf != nil { //Password does not match!
		return false
	}
	return true
}

func UserAgentParse(r *http.Request, data map[string]interface{}) {
	ua := user_agent.New(r.Header.Get("User-Agent"))
	if ua.Mobile() {
		data["device_type"] = "Mobile"
	} else if ua.Bot() {
		data["device_type"] = "Other"
	} else {
		data["device_type"] = "Computer"
	}
	data["os_type"] = ua.OS()
	name, _ := ua.Browser()
	data["browser"] = name
	return
}
