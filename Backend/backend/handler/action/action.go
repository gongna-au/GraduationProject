package action

import (
	"backend/model/action"
	"backend/model/requests"
	"backend/model/response"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary      UpdateUserBrowse
// @Description  更新用户的浏览记录
// @Tags         recommend
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.UpdateUserBroseRequest  true  "Uid--用户ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /action/browse   [post]
// UpdateUserBrowse
func UpdateUserBrowse(c *gin.Context) {
	// 1. 验证表单
	request := requests.UpdateUserBroseRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
	}
	// 2. 验证成功，创建数据
	actionModel := action.UserActionModel{
		ResourceID: request.ResourceID,
		UserID:     request.Uid,
		ActionType: "brose",
	}
	err := actionModel.Create()
	if err != nil {
		response.Error(c, err, "请求失败")
	}
	response.Data(c, gin.H{
		"res": "创建成功",
	})
}

// ShowAccount godoc
// @Summary      GetResourceByCollect
// @Description  更新用户的点赞记录
// @Tags         recommend
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.UpdateUserCollectRequest  true  "Uid--用户ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /action/collect  [post]
// GetResourceByCollect
func UpdateUserCollect(c *gin.Context) {
	// 1. 验证表单
	request := requests.UpdateUserCollectRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
	}
	// 2. 验证成功，创建数据
	actionModel := action.UserActionModel{
		ResourceID: request.ResourceID,
		UserID:     request.Uid,
		ActionType: "collect",
	}
	err := actionModel.Create()
	if err != nil {
		response.Error(c, err, "请求失败")
	}
	response.Data(c, gin.H{
		"res": "创建成功",
	})
}

// ShowAccount godoc
// @Summary      GetResourceBySave
// @Description  更新用户分享记录
// @Tags         recommend
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.UpdateUserSaveRequest  true  "Uid--用户ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /action/save  [post]
// GetResourceBySave
func UpdateUserSave(c *gin.Context) {
	// 1. 验证表单
	request := requests.UpdateUserSaveRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
	}
	// 2. 验证成功，创建数据
	actionModel := action.UserActionModel{
		ResourceID: request.ResourceID,
		UserID:     request.Uid,
		ActionType: "save",
	}
	err := actionModel.Create()
	if err != nil {
		response.Error(c, err, "请求失败")
	}
	response.Data(c, gin.H{
		"res": "创建成功",
	})
}

// ShowAccount godoc
// @Summary      GetResourceByClick
// @Description  更新用户点击记录
// @Tags         recommend
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.UpdateUserClickRequest  true  "Uid--用户ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /action/click  [post]
// GetResourceByClick
func UpdateUserClick(c *gin.Context) {

	// 1. 验证表单
	request := requests.UpdateUserClickRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
	}
	// 2. 验证成功，创建数据
	actionModel := action.UserActionModel{
		ResourceID: request.ResourceID,
		UserID:     request.Uid,
		ActionType: "click",
	}
	err := actionModel.Create()
	if err != nil {
		response.Error(c, err, "请求失败")
	}
	response.Data(c, gin.H{
		"res": "创建成功",
	})
}

// ShowAccount godoc
// @Summary      GetResourceByLike
// @Description  更新用户点赞记录
// @Tags         recommend
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.UpdateUserLikeRequest true  "Uid--用户ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /action/like  [post]
// GetResourceByLike
func UpdateUserLike(c *gin.Context) {

	// 1. 验证表单
	request := requests.UpdateUserLikeRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
	}
	// 2. 验证成功，创建数据
	actionModel := action.UserActionModel{
		ResourceID: request.ResourceID,
		UserID:     request.Uid,
		ActionType: "like",
	}
	err := actionModel.Create()
	if err != nil {
		response.Error(c, err, "请求失败")
	}
	response.Data(c, gin.H{
		"res": "创建成功",
	})
}

// ShowAccount godoc
// @Summary      UpdateUserSearch
// @Description  更新用户点赞记录
// @Tags         recommend
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.UpdateUserLikeRequest true  "Uid--用户ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /action/search  [post]
// UpdateUserSearch
func UpdateUserSearch(c *gin.Context) {

	// 1. 验证表单
	request := requests.UpdateUserSearchRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
	}
	// 2. 验证成功，创建数据
	actionModel := action.UserActionModel{
		ResourceID: request.ResourceID,
		UserID:     request.Uid,
		ActionType: "search",
	}
	err := actionModel.Create()
	if err != nil {
		response.Error(c, err, "请求失败")
	}
	response.Data(c, gin.H{
		"res": "创建成功",
	})
}
