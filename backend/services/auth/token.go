package auth

import (
	"fmt"
	"malai_agency/backend/response"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

var TokenKey = "AIRLIFTUSAHASH"

func Before(w http.ResponseWriter, r *http.Request) bool {
	var authCheck bool
	token := r.Header["X-Token"]
	// Check token available or not
	if len(token) == 0 {
		response.ResponseError(w, fmt.Errorf(""), "missing_token")
		return false
	}

	authCheck, userID, userType, designationID := ProtectedEndpoint(token[0])
	if !authCheck {
		response.ResponseError(w, fmt.Errorf(""), "invalid_token")
		return false
	}
	r.Header["user_id"] = []string{userID, userType}
	r.Header["designation_id"] = []string{designationID}
	return true
}

func ProtectedEndpoint(tokenString string) (bool, string, string, string) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(TokenKey), nil
	})
	// var userInfo session.UserInfoDetails
	userInfo := make(map[string]interface{})
	designtionID := ""

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		mapstructure.Decode(claims, &userInfo)
		// Check the request user based on Internal user or customer
		userType := "user"
		if userInfo["designation"] != nil {
			designtionID = userInfo["designation"].(string)
		}
		if userInfo["refer_id"] != nil {
			userType = "customer"
		}
		return true, fmt.Sprint(userInfo["id"]), userType, designtionID
	} else {
		return false, "", "", ""
	}
	return token.Valid, "", "", ""
}
