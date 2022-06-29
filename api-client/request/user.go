package request

type UserLoginRequest struct {
	Username string `json:"username" binding:"required" form:"username"`
	Password string `json:"password" binding:"required" form:"password"`
}

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required" form:"username"`
	Password string `json:"password" binding:"required" form:"password"`
}

type UserGetRequest struct {
	UserId int64 `json:"user_id" binding:"required" form:"user_id"`
}
