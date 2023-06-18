package utils

type UserInfo struct {
	UserID   string
	Username string
	Email    string
	Token    string
	Roles    []string
}

var userInfo *UserInfo

func SetTokenInfo(userInfo UserInfo) error {
	userInfo = userInfo
	return nil
}

// // Hàm để truy xuất thông tin token
func GetTokenInfo() *UserInfo {
	return userInfo
}
