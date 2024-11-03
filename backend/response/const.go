package response

var JsonErrorCode = "invalid_body"
var ErrorConst = map[string]map[string]interface{}{
	"db_error":               map[string]interface{}{"error_code": "db_error", "error_message": "Database query execution Error", "http_code": 500},
	"missing_token":          map[string]interface{}{"error_code": "missing_token", "error_message": "Missing token", "http_code": 403},
	"invalid_token":          map[string]interface{}{"error_code": "invalid_token", "error_message": "Invalid token", "http_code": 403},
	"invalid_json":           map[string]interface{}{"error_code": "invalid_json", "error_message": "Invalid JSON Body ", "http_code": 400},
	"missing_table_name":     map[string]interface{}{"error_code": "missing_table_name", "error_message": "Missing update table name ", "http_code": 400},
	"missing_row_id":         map[string]interface{}{"error_code": "missing_row_id", "error_message": "Missing Respected row id ", "http_code": 400},
	"invalid_password":       map[string]interface{}{"error_code": "invalid_password", "error_message": "Invalid Password check again", "http_code": 200},
	"invalid_email":          map[string]interface{}{"error_code": "invalid_email", "error_message": "Invalid Email check again", "http_code": 200},
	"user_inactive":          map[string]interface{}{"error_code": "user_inactive", "error_message": "User account locked", "http_code": 200},
	"unknown_error":          map[string]interface{}{"error_code": "unknown_error", "error_message": "Unknown error", "http_code": 200},
	"invalid_token_link":     map[string]interface{}{"error_code": "invalid_token_link", "error_message": "Password Reset Link Expired", "http_code": 200},
	"add_comment_not_config": map[string]interface{}{"error_code": "add_comment_not_config", "error_message": "Add Comment not config. Pls check  ", "http_code": 200},
	"field_exist":            map[string]interface{}{"error_code": "field_exist", "error_message": "Field name already exists table.Pls try another name.", "http_code": 200},
	"workflow_error":         map[string]interface{}{"error_code": "workflow_error", "error_message": "Unable to Process your request", "http_code": 501},
	"already_exist":          map[string]interface{}{"error_code": "already_exist", "error_message": "Already exist in data sets", "http_code": 200},
}
