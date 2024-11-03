package login

import (
	"fmt"
	"malai_agency/backend/services/query"

	"golang.org/x/crypto/bcrypt"
)

// Compare hash string and password string
func PasswordCompare(hashString string, passwordString string) bool {
	errf := bcrypt.CompareHashAndPassword([]byte(hashString), []byte(passwordString))
	// && errf == bcrypt.ErrMismatchedHashAndPassword
	if errf != nil { //Password does not match!
		return false
	}
	return true
}

func passwordCallBack(id string, password string) bool {
	sql := "SELECT `password` FROM `Users` WHERE `id`='" + id + "'"
	hashPassword, err := query.QueryToId(sql)
	if err != nil {
		fmt.Println("(passwordCallBack) select error", err)
		return false
	}
	return PasswordCompare(hashPassword, password)
}
