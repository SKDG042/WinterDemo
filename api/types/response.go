package types
// 接着为了方便返回数据，我需要按照文档构建一个结构体
type Response struct {
	Status int         `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}

// 成功的Response
func SuccessResponse(data interface{}) *Response {
	return &Response{
		Status: 10000,
		Info:   "success",
		Data:   data,
	}
}

// 失败的Response
func ErrorResponse(status int, info string) *Response {
	return &Response{
		Status: status,
		Info:   info,
		Data:   nil,
	}
}

// 登录成功返回的token
type TokenResponse struct {
	Token 			string `json:"token"`
	RefreshToken 	string `json:"refresh_token"` 
}

type UserInfoResponse struct {
	ID 				uint 	`json:"id"`
	Avatar 			string 	`json:"avatar"`
	Nickname 		string 	`json:"nickname"`
	Introduction 	string 	`json:"introduction"`
	Phone 			string 	`json:"phone"`
	QQ 				string 	`json:"qq"`
	Gender 			string 	`json:"gender"`
	Email 			string 	`json:"email"`
	Birthday 		string 	`json:"birthday"`
}

// 商品分类的响应
type CategoryResponse struct {
	ID 			uint 	`json:"category_id"`
	Name 		string 	`json:"name"`
	Description string 	`json:"description"`
}

// 查询商品响应
type ProductResponse struct {
	ProductID	uint	`json:"product_id"`
	Name		string	`json:"name"`
	Description string	`json:"description"`
	Type		string	`json:"type"`
	CommentNum	int		`json:"comment_num"`
	Price		float64	`json:"price"`
	IsAddCart	bool	`json:"is_addCart"`
	Cover		string	`json:"cover"`
	PublishTime	string	`json:"publish_time"`
	Link		string	`json:"link"`
	Categories  []CategoryResponse `json:"categories"`
}

// 商品列表响应
type ProductListResponse struct {
	Products []ProductResponse `json:"products"`
}