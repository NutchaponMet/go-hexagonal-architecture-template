package schemas

type UserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserResp struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
