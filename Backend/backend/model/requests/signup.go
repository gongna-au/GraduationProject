package requests

type SignupPhoneExistRequest struct {
	Phone string `valid:"phone" json:"phone,omitempty" gorm:"column:phone;" binding:"required"`
} //@name SignupPhoneExistRequest

// SignupUsingPhoneRequest 通过手机注册的请求信息
type SignupUsingPhoneRequest struct {
	Phone    string `valid:"phone" json:"phone,omitempty" gorm:"column:phone;" binding:"required"`
	Password string `valid:"password" json:"password,omitempty" gorm:"column:password;" binding:"required"`
} //@name SignupUsingPhoneRequest

// AddNewCardRequest 绑定新的银行卡
type AddNewCardRequest struct {
	UserID       string `json:"user_id" binding:"required"`       // 用户ID
	CardNumber   string `json:"card_number" binding:"required"`   // 银行卡号
	CardPassword string `json:"card_password" binding:"required"` // 银行卡密码
	BankName     string `json:"bank_name" binding:"required"`     // 银行名称
	ExpiryDate   string `json:"expiry_date" binding:"required"`   // 到期日期，格式为"YYYY-MM-DD"
} //@name AddNewCardRequest

type UnbindCardRequest struct {
	UserID     string `json:"user_id" binding:"required"`     // 用户ID
	CardNumber string `json:"card_number" binding:"required"` // 银行卡号
}

type UpdatePasswordRequest struct {
	UserID      string `json:"user_id" binding:"required"`      // 用户ID
	OldPassword string `json:"old_password" binding:"required"` // 银行卡号
	NewPassword string `json:"new_password" binding:"required"` // 银行卡号
} //@name UpdatePasswordRequest

type UpdateOverageRequest struct {
	UserID         string `json:"user_id" binding:"required"`         // 用户ID
	Password       string `json:"password" binding:"required"`        // 银行密码
	CardNumber     string `json:"card_number" binding:"required"`     // 银行密码
	RechargeAmount int    `json:"recharge_amount" binding:"required"` // 银行密码
} //@name UpdateOverageRequest
