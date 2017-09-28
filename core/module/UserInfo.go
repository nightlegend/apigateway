package module

type UserInfo struct {
	USERNAME string `json:"userName" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
	EMAIL    string `json:"email" binding:"required"`
}
