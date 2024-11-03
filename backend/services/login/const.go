package login

// LoginDetails Used to decode value
type LoginDetails struct {
	Email    string
	Password string
}

var InvalidPassword = map[string]interface{}{"error_code": "invalid_password", "error_message": "Invalid Password check again"}

var InvalidEmail = map[string]interface{}{"error_code": "invalid_email", "error_message": "Invalid Email check again"}

var UserInactiveError = map[string]interface{}{"error_code": "user_inactive", "error_message": "User account locked"}
