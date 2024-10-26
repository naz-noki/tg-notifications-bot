package botController

type botService interface {
	SendBooking(name string, time string, date string, phone string, quantity int, is_confirmed bool) error
}

type SendBookingRequest struct {
	Name         string `json:"name"`
	Quantity     int    `json:"quantity"`
	Is_confirmed bool   `json:"is_confirmed"`
	Time         string `json:"time"`
	Date         string `json:"date"`
	Phone        string `json:"phone"`
}
