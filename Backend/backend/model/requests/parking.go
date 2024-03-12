package requests

type ParkingSearchByLatitudeAndLongitude struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
} //@name ParkingSearchByLatitudeAndLongitude

type AddReserveRequest struct {
	UserID        string `json:"user_id" binding:"required"`        // 用户ID
	ParkingLotId  int    `json:"parking_lot_id" binding:"required"` // 银行密码
	StartTime     string `json:"start_time" binding:"required"`     // 银行密码
	EndTime       string `json:"end_time" binding:"required"`       // 银行密码
	VehicleNumber string `json:"vehicle_number" binding:"required"` // 银行密码
	Status        string `json:"status" `                           // 银行密码
} //@name AddReserveRequest
