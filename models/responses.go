package models

// ReceiptResp is the reply after receipt submission
type ReceiptResp struct {
	ID string `json:"id"`
}

// PointsResp conveys the points earned
type PointsResp struct {
	Points int `json:"points"`
}
