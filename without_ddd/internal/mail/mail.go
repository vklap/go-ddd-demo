package mail

import "context"

type InMemoryMailService struct {
}

func (m *InMemoryMailService) SendOrderStatus(ctx context.Context, request *OrderNotificationRequest) error {
	return nil
}
