package signup

import (
	"backend/model/requests"
	"fmt"

	"backend/model/response"

	"backend/model/user"
	"backend/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary      Sign up Using Phone
// @Description  使用邮箱和密码注册帐号
// @Tags         sign
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.SignupUsingPhoneRequest  true  "Phone--手机号||Password-- 密码|| Name--昵称"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sign/common   [post]
// CommonUserSignupUsingPhone 使用邮箱和密码进行注册
func CommonUserSignupUsingPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.SignupUsingPhoneRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "Bind Error")
		return
	}
	err := user.JudgeCommonUserPhoneExist(request.Phone)
	if err != nil {
		response.BadRequest(c, fmt.Errorf("judge Common User Phone Exist Error:%v", err))
		return
	}
	// 2. 验证成功，创建数据
	userModel := user.CommonUserModel{
		Phone:    request.Phone,
		Password: request.Password,
	}
	err = userModel.Create()
	if err != nil {
		response.Abort500(c, err.Error())
	}
	if userModel.ID > 0 {
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Phone)
		response.CreatedJSON(c, gin.H{
			"token":  token,
			"userId": userModel.GetStringID(),
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
		return
	}

}
