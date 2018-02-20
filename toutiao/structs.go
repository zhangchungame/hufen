package toutiao

type ToutiaoLoginReturn struct {
	UserId int `json:"user_id"`
	ErrorCode int `json:"error_code"`
	RedirectUrl string `json:"redirect_url"`
	Description string `json:"description"`
}
