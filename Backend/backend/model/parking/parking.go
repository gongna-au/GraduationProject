package parking

import (
	db "backend/database"
	"backend/model"
	"fmt"
	"time"
)

type ParkingLotModel struct {
	model.BaseModel
	Name      string  `json:"name" gorm:"column:name;" binding:"required"`
	Address   string  `json:"address" gorm:"column:address;" binding:"required"`
	ImageUrl  string  `json:"imageUrl" gorm:"column:imageUrl;" binding:"required"`
	Spaces    int     `json:"spaces" gorm:"column:spaces;" binding:"required"`
	Charge    float64 `json:"charge" gorm:"type:decimal(10,2);not null" binding:"required"`
	Latitude  float64 `json:"latitude" gorm:"type:decimal(10,8);not null" binding:"required"`
	Longitude float64 `json:"longitude" gorm:"type:decimal(11,8);not null" binding:"required"`
}

func (u *ParkingLotModel) TableName() string {
	return "parking_lots"
}

func (u *ParkingLotModel) Create() error {
	return db.DB.
		Table(u.TableName()).
		Create(u).Error
}

func (u *ParkingLotModel) Save() (err error) {
	return db.DB.
		Table(u.TableName()).
		Save(u).Error
}

func (u *ParkingLotModel) Update() error {
	return db.DB.
		Table(u.TableName()).
		Where("id = ?", u.ID).
		Updates(map[string]interface{}{
			"name":      u.Name,
			"address":   u.Address,
			"imageUrl":  u.ImageUrl,
			"spaces":    u.Spaces,
			"charge":    u.Charge,
			"latitude":  u.Latitude,
			"longitude": u.Longitude,
		}).Error
}

func (u *ParkingLotModel) Delete() error {
	return db.DB.
		Table(u.TableName()).
		Where("id = ?", u.ID).
		Delete(u).Error
}

// 获取所有停车场的列表。
func GetAllParkingLots() ([]ParkingLotModel, error) {
	var parkingLots []ParkingLotModel
	err := db.DB.
		Table((&ParkingLotModel{}).TableName()).
		Find(&parkingLots).Error
	return parkingLots, err
}

// 根据ID获取单个停车场的详细信息。
func GetParkingLotByID(id uint) (*ParkingLotModel, error) {
	u := &ParkingLotModel{}
	err := db.DB.
		Table(u.TableName()).
		Where("id = ?", id).First(u).Error
	return u, err
}

// 根据ID获取单个停车场的详细信息。
func GetParkingLotByName(name string) (*ParkingLotModel, error) {
	u := &ParkingLotModel{}
	err := db.DB.
		Table(u.TableName()).
		Where("name = ?", name).First(u).Error
	return u, err
}

// 定义一个结构体来接收查询结果
type ParkingLotWithDistance struct {
	ParkingLotModel
	Distance float64 `json:"distance"`
}

func FindNearbyParkingLots(userLat, userLong float64) ([]ParkingLotWithDistance, error) {
	var parkingLots []ParkingLotWithDistance
	sql := `
    SELECT *, 
           (6371 * acos(cos(radians(?)) 
                        * cos(radians(latitude)) 
                        * cos(radians(longitude) - radians(?)) 
                        + sin(radians(?)) 
                        * sin(radians(latitude)))) AS distance 
    FROM parking_lots
    ORDER BY distance ASC
    `
	// 使用GORM的Raw方法执行查询，并将结果扫描到parkingLots切片中
	err := db.DB.Raw(sql, userLat, userLong, userLat).Scan(&parkingLots).Error
	if err != nil {
		return nil, fmt.Errorf("error finding nearby parking lots: %v", err)
	}

	return parkingLots, nil
}

// 定义一个结构体来接收查询结果
func (ParkingHistoryModel) TableName() string {
	return "parking_history"
}

type ParkingHistoryModel struct {
	ID              int       `json:"id" gorm:"column:id;" binding:"required" `
	UserID          int       `json:"user_id" gorm:"column:user_id;" binding:"required"`
	ParkingLotID    int       `json:"parking_lot_id" gorm:"column:parking_lot_id;" binding:"required" `
	VehicleNumber   string    `json:"vehicle_number" gorm:"column:vehicle_number;" binding:"required" `
	ParkingDuration string    `json:"parking_duration"` // 这是一个计算字段，不直接存储在数据库中
	StartTime       time.Time `json:"start_time" gorm:"column:start_time;" binding:"required"`
	EndTime         time.Time `json:"end_time" gorm:"column:end_time;" binding:"required" `
	PaymentAmount   float64   `json:"payment_amount" gorm:"type:decimal(10,2);" binding:"required"`
}

type SalesByDate struct {
	TotalRevenue       float64
	Date               string
	TotalParkings      int
	MinParkingDuration time.Duration
	MaxParkingDuration time.Duration
	MinDuration        string
	MaxDuration        string
}

func (ph *ParkingHistoryModel) GetSales(endTime string) (SalesByDate, error) {
	var res []ParkingHistoryModel
	err := db.DB.Table(ph.TableName()).
		Where("DATE(end_time) = ?", endTime).
		Scan(&res).Error
	var s SalesByDate
	s.Date = endTime
	var minDuration, maxDuration time.Duration = time.Hour * 24, 0
	s.TotalParkings = len(res)
	for _, v := range res {
		s.TotalRevenue = s.TotalRevenue + v.PaymentAmount
		// 计算当前条目的停车时长
		parkingDuration := v.EndTime.Sub(v.StartTime)

		// 更新最小和最大停车时长
		if parkingDuration < minDuration {
			minDuration = parkingDuration
		}
		if parkingDuration > maxDuration {
			maxDuration = parkingDuration
		}
	}
	if s.TotalParkings > 0 {
		s.MinParkingDuration = minDuration
		s.MaxParkingDuration = maxDuration
	} else {
		// 如果没有停车记录，可能需要设置一个默认值或保持为零值
		s.MinParkingDuration = 0
		s.MaxParkingDuration = 0
	}

	s.MinDuration = formatDuration(s.MinParkingDuration)
	s.MaxDuration = formatDuration(s.MaxParkingDuration)
	return s, err
}

type ParkingHistoryWithLot struct {
	ParkingHistoryID int       `gorm:"column:id"`
	UserID           int       `gorm:"column:user_id"`
	VehicleNumber    string    `gorm:"column:vehicle_number"`
	StartTime        time.Time `gorm:"column:start_time"`
	EndTime          time.Time `gorm:"column:end_time"`
	ParkingLotName   string    `gorm:"column:name"`
	Address          string    `gorm:"column:address"`
	PaymentAmount    float64   `gorm:"column:payment_amount"` // 支付金额
	Duration         string
}

func SearchParkingHistoryByUserID(userID int) ([]ParkingHistoryWithLot, error) {
	var results []ParkingHistoryWithLot
	err := db.DB.Table("parking_history").
		Select("parking_history.id, parking_history.user_id, parking_history.vehicle_number, parking_history.start_time, parking_history.end_time, parking_history.payment_amount,parking_lots.name, parking_lots.address").
		Joins("join parking_lots on parking_lots.id = parking_history.parking_lot_id").
		Where("parking_history.user_id = ?", userID). // 示例：针对特定 user_id 进行查询
		Scan(&results).Error

	for i, _ := range results {
		duration := results[i].EndTime.Sub(results[i].StartTime)
		// 打印结果
		results[i].Duration = formatDuration(duration)
	}
	return results, err

}

func formatDuration(duration time.Duration) string {
	totalSeconds := int(duration.Seconds())
	days := totalSeconds / (24 * 3600)
	hours := (totalSeconds % (24 * 3600)) / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	return fmt.Sprintf("%d days %d hours %d minutes %d seconds", days, hours, minutes, seconds)
}

// fmtDuration 格式化时间间隔为人类可读的形式
func fmtDuration(d time.Duration) string {
	hours := d / time.Hour
	d -= hours * time.Hour
	minutes := d / time.Minute
	return fmt.Sprintf("%dh %dm", hours, minutes)
}

type ParkingReservationModel struct {
	ID            int       `json:"id" gorm:"column:id;" binding:"required"`
	UserID        int       `json:"user_id" gorm:"column:user_id;" binding:"required"`
	ParkingLotID  int       `json:"parking_lot_id" gorm:"column:parking_lot_id;" binding:"required"`
	VehicleNumber string    `json:"vehicle_number" gorm:"column:vehicle_number;" binding:"required"`
	StartTime     time.Time `json:"start_time" gorm:"column:start_time;" binding:"required"`
	EndTime       time.Time `json:"end_time" gorm:"column:end_time;" binding:"required"`
	Status        string    `json:"status" gorm:"column:status;default:PENDING" binding:"required"`
}

func (c *ParkingReservationModel) TableName() string {
	return "parking_reservations"
}

func (pr *ParkingReservationModel) Get() (*ParkingReservationModel, error) {
	res := &ParkingReservationModel{}
	err := db.DB.Table(pr.TableName()).
		Where("id = ?", pr.ID).
		Scan(res).Error
	return res, err
}

func (pr *ParkingReservationModel) Add() error {
	// 使用GORM的Create方法新增记录
	pr.StartTime = pr.StartTime.UTC() // 调整为UTC时区
	pr.EndTime = pr.EndTime.UTC()
	result := db.DB.Table(pr.TableName()).Create(pr)
	return result.Error
}

func (pr *ParkingReservationModel) Update() error {
	return db.DB.Table(pr.TableName()).
		Model(pr).
		Where("id = ?", pr.ID).
		Updates(map[string]interface{}{
			"user_id":        pr.UserID,
			"parking_lot_id": pr.ParkingLotID,
			"vehicle_number": pr.VehicleNumber,
			"start_time":     pr.StartTime,
			"end_time":       pr.EndTime,
			"status":         pr.Status,
		}).Error
}

func (pr *ParkingReservationModel) GetByUserID() ([]ReservationResponse, error) {
	var reservations []ReservationResponse
	err := db.DB.Table(pr.TableName()).
		Select("parking_reservations.*, parking_lots.name AS parking_lot_name, parking_lots.address AS parking_lot_address, parking_lots.spaces AS parking_lot_spaces, parking_lots.charge AS parking_lot_charge").
		Joins("JOIN parking_lots ON parking_lots.id = parking_reservations.parking_lot_id").
		Where("parking_reservations.user_id = ?", pr.UserID).
		Scan(&reservations).Error

	if err != nil {
		return nil, err
	}

	return reservations, nil
}

// 假设的ParkingReservationModel，现在包括了所有需要的信息
type ReservationResponse struct {
	ParkingReservationModel
	// 来自parking_lots表
	ParkingLotName    string `json:"parking_lot_name"`
	ParkingLotAddress string `json:"parking_lot_address"`
	ParkingLotSpaces  int    `json:"parking_lot_spaces"`
	ParkingLotCharge  string `json:"parking_lot_charge"`
}

// SearchParkingReserveByUserID 根据用户ID查询停车预约记录，包括停车场信息和支付银行卡信息

func CancelReserveByReserveId(reserveId int) error {
	// 查找并删除指定ID的预定记录
	result := db.DB.Table("parking_reservations").Delete(&ParkingReservationModel{}, reserveId)
	if result.Error != nil {
		return result.Error
	}
	// 可以在这里添加逻辑来处理退款等操作，如果需要的话
	// 例如，查找对应的银行卡记录，然后执行退款逻辑

	return nil
}
