package recommend

import (
	"backend/model/recommend"
	"backend/model/requests"
	"backend/model/response"
	"backend/pkg/helpers"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary      GetResourceByBrowse
// @Description  通过用户的浏览记录推荐
// @Tags         recommend
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.GetResourceByUid  true  "Uid--用户ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /recommend/browse   [post]
// GetResourceByBrowse
func GetResourceByBrowse(c *gin.Context) {
	// 1. 验证表单
	request := requests.GetResourceByUid{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
	}
	// 2. 验证成功，创建数据
	actionModel := recommend.UserActionModel{}
	action, err := actionModel.GetUserActionById(request.Uid, "browse")
	if err != nil {
		response.Error(c, err, "请求失败")
	}
	a := []helpers.UserActionModel{}
	for _, v := range action {
		a = append(a, helpers.UserActionModel{
			ID:         v.ID,
			UserID:     v.UserID,
			ResourceID: v.ResourceID,
			ActionType: v.ActionType,
		})
	}

	r := recommend.ResourceModel{}
	resources, err := r.GetResources()
	if err != nil {
		response.Error(c, err, "请求失败")
	}

	s := []helpers.ResourceModel{}
	for _, v := range resources {
		s = append(s, helpers.ResourceModel{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Url:         v.Url,
			Tag:         v.Tag,
			Score:       0,
		})
	}

	res := helpers.GetRecommendations(a, s, 5)
	response.Data(c, res)
}

// ShowAccount godoc
// @Summary      GetResourceByCollect
// @Description  通过用户的收藏记录推荐
// @Tags         recommend
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.GetResourceByUid  true  "Uid--用户ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /recommend/collect  [post]
// GetResourceByCollect
func GetResourceByCollect(c *gin.Context) {
	// 1. 验证表单
	request := requests.GetResourceByUid{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {

		response.Error(c, err, "c.Bind请求失败")
		return
	}
	// 2. 验证成功，创建数据
	actionModel := recommend.UserActionModel{}
	action, err := actionModel.GetUserActionById(request.Uid, "like")
	if err != nil {
		response.Error(c, err, "GetUserActionById请求失败")
		return
	}
	a := []helpers.UserActionModel{}
	for _, v := range action {
		a = append(a, helpers.UserActionModel{
			ID:         v.ID,
			UserID:     v.UserID,
			ResourceID: v.ResourceID,
			ActionType: v.ActionType,
		})
	}

	r := recommend.ResourceModel{}
	resources, err := r.GetResources()

	if err != nil {
		response.Error(c, err, "GetResources请求失败")
		return
	}

	s := []helpers.ResourceModel{}
	for _, v := range resources {
		s = append(s, helpers.ResourceModel{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Url:         v.Url,
			Tag:         v.Tag,
			Score:       0,
		})
	}
	res := helpers.GetRecommendations(a, s, 5)
	response.Data(c, res)
}

// ShowAccount godoc
// @Summary      GetResourceBySave
// @Description  通过用户的保存记录推荐
// @Tags         recommend
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.GetResourceByUid  true  "Uid--用户ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /recommend/save  [post]
// GetResourceBySave
func GetResourceBySave(c *gin.Context) {

	// 1. 验证表单
	request := requests.GetResourceByUid{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
	}
	// 2. 验证成功，创建数据
	actionModel := recommend.UserActionModel{}
	action, err := actionModel.GetUserActionById(request.Uid, "save")
	if err != nil {
		response.Error(c, err, "请求失败")
	}
	a := []helpers.UserActionModel{}
	for _, v := range action {
		a = append(a, helpers.UserActionModel{
			ID:         v.ID,
			UserID:     v.UserID,
			ResourceID: v.ResourceID,
			ActionType: v.ActionType,
		})
	}

	r := recommend.ResourceModel{}
	resources, err := r.GetResources()
	if err != nil {
		response.Error(c, err, "请求失败")
	}

	s := []helpers.ResourceModel{}
	for _, v := range resources {
		s = append(s, helpers.ResourceModel{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Url:         v.Url,
			Tag:         v.Tag,
			Score:       0,
		})
	}

	res := helpers.GetRecommendations(a, s, 5)
	response.Data(c, res)
}

// ShowAccount godoc
// @Summary      GetResourceByClick
// @Description  通过用户点击记录推荐
// @Tags         recommend
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.GetResourceByUid  true  "Uid--用户ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /recommend/click  [post]
// GetResourceByClick
func GetResourceByClick(c *gin.Context) {

	// 1. 验证表单
	request := requests.GetResourceByUid{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
	}
	// 2. 验证成功，创建数据
	actionModel := recommend.UserActionModel{}
	action, err := actionModel.GetUserActionById(request.Uid, "click")
	if err != nil {
		response.Error(c, err, "请求失败")
	}
	a := []helpers.UserActionModel{}
	for _, v := range action {
		a = append(a, helpers.UserActionModel{
			ID:         v.ID,
			UserID:     v.UserID,
			ResourceID: v.ResourceID,
			ActionType: v.ActionType,
		})
	}

	r := recommend.ResourceModel{}
	resources, err := r.GetResources()
	if err != nil {
		response.Error(c, err, "请求失败")
	}

	s := []helpers.ResourceModel{}
	for _, v := range resources {
		s = append(s, helpers.ResourceModel{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Url:         v.Url,
			Tag:         v.Tag,
			Score:       0,
		})
	}

	res := helpers.GetRecommendations(a, s, 5)
	response.Data(c, res)
}

// ShowAccount godoc
// @Summary      GetResourceByLike
// @Description  通过用户点击喜欢推荐
// @Tags         recommend
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.GetResourceByUid  true  "Uid--用户ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /recommend/like  [post]
// GetResourceByLike
func GetResourceByLike(c *gin.Context) {

	// 1. 验证表单
	request := requests.GetResourceByUid{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
		return
	}
	// 2. 验证成功，创建数据
	actionModel := recommend.UserActionModel{}
	action, err := actionModel.GetUserActionById(request.Uid, "like")
	if err != nil {
		response.Error(c, err, "请求失败")
		return
	}
	a := []helpers.UserActionModel{}
	for _, v := range action {
		a = append(a, helpers.UserActionModel{
			ID:         v.ID,
			UserID:     v.UserID,
			ResourceID: v.ResourceID,
			ActionType: v.ActionType,
		})
	}

	r := recommend.ResourceModel{}
	resources, err := r.GetResources()
	if err != nil {
		response.Error(c, err, "请求失败")
		return
	}

	s := []helpers.ResourceModel{}
	for _, v := range resources {
		s = append(s, helpers.ResourceModel{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Url:         v.Url,
			Tag:         v.Tag,
			Score:       0,
		})
	}

	res := helpers.GetRecommendations(a, s, 5)
	response.Data(c, res)
}
