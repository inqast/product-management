package sender

import "fmt"

type statusMessage struct {
	UserID  int64
	OrderID int64
	Status  string
}

func (m *statusMessage) GetKey() string {
	return fmt.Sprint(m.OrderID)
}
