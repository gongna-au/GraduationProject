package login

import (
	"backend/model/requests"
	"backend/model/response"
	"backend/model/user"
	"backend/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary      Login By Email
// @Description  用户通过邮箱和密码登陆
// @Tags         login
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.LoginByPhoneRequest  true  "Phone--手机号 和Password--密码"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Router       /login/common  [post]
// @ApiA
// CommonUserLoginByPhone 手机号
func CommonUserLoginByPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.LoginByPhoneRequest{}
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
		return
	}
	// 2. 尝试登录
	user, err := user.GetCommonUserByPhoneAndPassword(request.Phone, request.Password)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "账号不存在或密码错误")
		return
	} else {
		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Phone)
		response.JSON(c, gin.H{
			"token":  token,
			"userId": user.GetStringID(),
		})
	}

}

// ShowAccount godoc
// @Summary      UpdateUser
// @Description  更新密码
// @Tags         login
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.UpdateUserRequest true  "Email--邮箱 和Password--密码 NewPassword-新密码"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Router       /login/user  [post]
// @ApiA
// UpdateUserRequest 邮箱登录
func UpdateUser(c *gin.Context) {

	// 1. 验证表单
	request := requests.UpdateUserRequest{}
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
		return
	}
	// 2. 尝试登录
	_, err := user.UpdateUser(request.Email, request.Password, request.NewPassword)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "更新失败")
		return
	} else {
		// 登录成功
		response.JSON(c, gin.H{
			"result": "success!",
		})
	}

}
