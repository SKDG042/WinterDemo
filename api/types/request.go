package types

// 首先我要构建结构体，方便后续使用BindJSON()来解析请求中的数据
type RegisterRequest struct {
	Username string `json:"username"` // json来接受json格式的数据
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdatePasswordRequest struct{
	NewPassword string `json:"newPassword"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// 根据接口文档，全部请求参数都是非必选，所以使用omitempty为空时，不返回该字段
type UpdateUserInfoRequest struct {
	Nickname 		string `json:"nickname,omitempty"`
	Avatar 			string `json:"avatar,omitempty"`
	Introduction 	string `json:"introduction,omitempty"`
	Telephone 		string `json:"telephone,omitempty"`
	QQ 				string `json:"qq,omitempty"`
	Gender 			string `json:"gender,omitempty"`
	Email 			string `json:"email,omitempty"`
	Birthday 		string `json:"birthday,omitempty"`
}

type SearchProductRequest struct {
	ProductName string `form:"product_name"` // form来接受/search?product_name=xxx这样的参数
}

type GetProductsByCategoryRequest struct {
	CategoryID int `form:"category_id"`
}