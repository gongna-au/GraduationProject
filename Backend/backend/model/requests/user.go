package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type PaginationRequest struct {
	Sort    string `valid:"sort" form:"sort"`
	Order   string `valid:"order" form:"order"`
	PerPage string `valid:"per_page" form:"per_page"`
} //@name PaginationRequest

//验证函数
func Pagination(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"sort":     []string{"in:id,created_at,updated_at"},
		"order":    []string{"in:asc,desc"},
		"per_page": []string{"numeric_between:2,100"},
	}
	messages := govalidator.MapData{
		"sort": []string{
			"in:排序字段仅支持 id,created_at,updated_at",
		},
		"order": []string{
			"in:排序规则仅支持 asc（正序）,desc（倒序）",
		},
		"per_page": []string{
			"numeric_between:每页条数的值介于 2~100 之间",
		},
	}
	// 配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}
	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}

type UpdatePassword struct {
	OldPassword string `valid:"old" json:"oldpassword,omitempty"`
	NewPassword string `valid:"new" json:"newpassword,omitempty"`
} //@name UpdatePassword

func UpdateUserPassword(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"old": []string{"required", "min:6"},
		"new": []string{"required", "min:6"},
	}

	messages := govalidator.MapData{
		"old": []string{
			"required:旧密码为必填项",
			"min:旧密码长度需大于 6",
		},
		"new": []string{
			"required:新密码为必填项",
			"min:新密码长度需大于 6",
		},
	}
	// 配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}
	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}

type UpdateName struct {
	Name string `valid:"name" json:"name,omitempty" gorm:"column:name;" binding:"required"`
} //@name UpdateName

func UpdateUserName(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"name": []string{"required", "alpha_num", "between:3,20"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
	}
	// 配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}
	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
