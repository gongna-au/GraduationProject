package user

import (
	"backend/model"
	"backend/model/requests"
	"backend/model/response"
	"fmt"

	"strconv"

	"backend/model/user"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary      Add New Bank Card
// @Description  添加新的银行卡
// @Tags         sign
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.SignupUsingPhoneRequest  true  "Phone--手机号||Password-- 密码|| Name--昵称"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /user/card  [post]
//
//	AddNewCard  添加新的银行卡
func AddNewCard(c *gin.Context) {
	// 1. 验证表单
	request := requests.AddNewCardRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Abort500(c, "绑定参数失败")
		return
	}
	userID, _ := strconv.Atoi(request.UserID) // 从查询参数中获取用户ID

	// 2. 验证，创建数据
	cardModel := user.BankCardsBound{
		UserId:       userID,
		CardNumber:   request.CardNumber,
		CardPassword: request.CardPassword,
		BankName:     request.BankName,
		ExpiryDate:   request.ExpiryDate,
	}
	if !cardModel.Valid() {
		response.BadRequest(c, fmt.Errorf("请求的参数不合法"))
		return
	}
	isBind, err := cardModel.IsBind()
	if isBind == true || err != nil {
		response.Abort500(c, "已经被绑定")
		return
	}
	err = cardModel.Create()
	if err != nil {
		response.Abort500(c, err.Error())
		return
	}
	if cardModel.ID > 0 {
		response.JSON(c, cardModel.CardNumber)
	} else {
		response.Abort500(c, "绑定银行卡失败，请稍后尝试~")
		return
	}

}

func UnbindCard(c *gin.Context) {
	// 1. 验证表单
	request := requests.UnbindCardRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Abort500(c, "绑定参数失败")
		return
	}
	userID, _ := strconv.Atoi(request.UserID) // 从查询参数中获取用户ID

	// 2. 验证，创建数据
	cardModel := user.BankCardsBound{
		UserId:     userID,
		CardNumber: request.CardNumber,
	}
	err := cardModel.Delete()
	if err != nil {
		response.Abort500(c, "解绑银行卡失败，请稍后尝试~")
		return
	}
	response.JSON(c, cardModel.CardNumber) // 使用你的响应工具返回数据
}

func GetBoundedCard(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("userId")) // 从查询参数中获取用户ID
	cardModel := user.BankCardsBound{
		UserId: userID,
	}
	boundedCards, err := cardModel.Get()
	if err != nil {
		response.Error(c, err, "获取绑定的银行卡失败")
		return
	}
	response.JSON(c, boundedCards) // 使用你的响应工具返回数据
}

func GetUserName(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("userId")) // 从查询参数中获取用户ID
	userModel := user.CommonUserModel{
		BaseModel: model.BaseModel{ID: userID},
	}
	user, err := userModel.GetName()
	if err != nil {
		response.Error(c, err, "用户名错误")
		return
	}

	response.JSON(c, user) // 使用你的响应工具返回数据
}

func UpdatePassword(c *gin.Context) {
	request := &requests.UpdatePasswordRequest{}
	// 2. 绑定
	if err := c.Bind(request); err != nil {
		response.Abort500(c, "绑定参数失败")
		return
	}
	userID, _ := strconv.Atoi(request.UserID) // 从查询参数中获取用户ID
	userModel := user.CommonUserModel{
		BaseModel: model.BaseModel{ID: userID},
	}
	user, err := userModel.Get()
	if err != nil {
		response.Error(c, err, "获取用户失败")
		return
	}
	if user != nil && user.Password == request.OldPassword {
		user.Password = request.NewPassword
		err := user.Update()
		if err != nil {
			response.Abort500(c, "更新失败")
		} else {
			response.JSON(c, "更新成功")
		}
		return
	}
	response.Abort500(c, "更新失败")
}

func UpdateOverage(c *gin.Context) {
	request := &requests.UpdateOverageRequest{}
	// 2. 绑定
	if err := c.Bind(request); err != nil {
		response.Abort500(c, "绑定参数失败")
		return
	}
	userID, _ := strconv.Atoi(request.UserID) // 从查询参数中获取用户ID
	userModel := user.CommonUserModel{
		BaseModel: model.BaseModel{ID: userID},
	}
	u, err := userModel.Get()
	if err != nil || u == nil {
		response.Abort500(c, "获取用户信息失败")
		return
	}
	cardModel := user.BankCardsBound{
		UserId:     userID,
		CardNumber: request.CardNumber,
	}
	card, err := cardModel.GetCardByUserIdAndNumber()
	if err != nil || card == nil {
		response.Abort500(c, "获取绑定的卡失败")
		return
	}
	if card.CardPassword != request.Password {
		response.Abort500(c, "输入的银行卡密码错误")
		return
	}
	u.Overage = u.Overage + request.RechargeAmount
	err = u.Update()
	if err != nil {
		response.Abort500(c, "充值失败")
		return
	}
	response.JSON(c, u.Overage) // 使用你的响应工具返回数据
}

type UserResponse struct {
	Overage  int `json:"overage"` // 假设Overage是float64类型
	Expenses int `json:"expenses"`
}

func GetOverageAndExpenses(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("userId")) // 从查询参数中获取用户ID
	userModel := user.CommonUserModel{
		BaseModel: model.BaseModel{ID: userID},
	}
	user, err := userModel.Get()
	if err != nil {
		response.Error(c, err, "用户名错误")
		return
	}

	response.JSON(c, UserResponse{
		Overage:  user.Overage,
		Expenses: user.Expenses,
	}) // 使用你的响应工具返回数据
}

func GetUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("userId")) // 从查询参数中获取用户ID
	userModel := user.CommonUserModel{
		BaseModel: model.BaseModel{ID: userID},
	}
	name, err := userModel.Get()
	if err != nil {
		response.Error(c, err, "用户名错误")
		return
	}

	response.JSON(c, name) // 使用你的响应工具返回数据
}
