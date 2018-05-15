package module

// UserInfo is define user info struct.
type UserInfo struct {
	USERNAME string `json:"userName" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
	EMAIL    string `json:"email" binding:"required"`
}
