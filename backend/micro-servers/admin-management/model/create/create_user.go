package create

import (
	"malai_agency/backend/services/query"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(data map[string]interface{}) error {
	// Create hash format password string
	newHashString, err := bcrypt.GenerateFromPassword([]byte(data["password"].(string)), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	data["password"] = newHashString
	sql := "INSERT INTO `Users`(`created_on`, `last_modified`,`last_modified_by`, " +
		"`created_by`,`user_name`, `email`, `password`,`designation`) VALUES (?,?,?,?,?,?,?,?)"
	var valueList []interface{}
	valueList = append(valueList, data["current_time"], data["current_time"], data["user_id"],
		data["user_id"], data["user_name"], data["email"], data["password"])
	if data["designation"] != nil {
		valueList = append(valueList, data["designation"].(map[string]interface{})["id"])
	} else {
		valueList = append(valueList, nil)
	}
	// Insert into table
	_, err = query.Insert(sql, valueList)
	if err != nil {
		return err
	}
	return nil
}
