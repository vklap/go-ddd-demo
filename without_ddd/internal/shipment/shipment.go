package shipment

import "context"

type InMemoryShippingService struct {
	requests []*Request
}

func (s *InMemoryShippingService) ReceiveRequest(ctx context.Context, request *Request) error {
	s.requests = append(s.requests, request)
	return nil
}
