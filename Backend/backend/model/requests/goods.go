package requests

type GetGoodsInStockBySupplierRequest struct {
	Supplier string `json:"supplier,omitempty" binding:"required"`
} //@name GetGoodsInStockBySupplierRequest

type GetGoodsInStockByNumRequest struct {
	Num int `json:"num,omitempty" binding:"required"`
} //@name  GetGoodsInStockByNumRequest

type GetResourceByUid struct {
	Uid int `json:"uid,omitempty" gorm:"column:uid;" binding:"required"`
} //@name GetResourceByUid

type UpdateUserBroseRequest struct {
	Uid        int `json:"uid,omitempty" gorm:"column:uid;" binding:"required"`
	ResourceID int `json:"resource_id,omitempty" gorm:"column:resource_id;" binding:"required"`
} //@name UpdateUserBroseRequest

type UpdateUserLikeRequest struct {
	Uid        int `json:"uid,omitempty" gorm:"column:uid;" binding:"required"`
	ResourceID int `json:"resource_id,omitempty" gorm:"column:resource_id;" binding:"required"`
} //@name UpdateUserLikeRequest

type UpdateUserSearchRequest struct {
	Uid        int `json:"uid,omitempty" gorm:"column:uid;" binding:"required"`
	ResourceID int `json:"resource_id,omitempty" gorm:"column:resource_id;" binding:"required"`
} //@name UpdateUserLikeRequest

type UpdateUserSaveRequest struct {
	ResourceID int `json:"resource_id,omitempty" gorm:"column:resource_id;" binding:"required"`
	Uid        int `json:"uid,omitempty" gorm:"column:uid;" binding:"required"`
} //@name UpdateUserSaveRequest

type UpdateUserClickRequest struct {
	ResourceID int `json:"resource_id,omitempty" gorm:"column:resource_id;" binding:"required"`
	Uid        int `json:"uid,omitempty" gorm:"column:uid;" binding:"required"`
} //@name UpdateUserClickRequest

type UpdateUserCollectRequest struct {
	ResourceID int `json:"resource_id,omitempty" gorm:"column:resource_id;" binding:"required"`
	Uid        int `json:"uid,omitempty" gorm:"column:uid;" binding:"required"`
} //@name UpdateUserClickRequest
