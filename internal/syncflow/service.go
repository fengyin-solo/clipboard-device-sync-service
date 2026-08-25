package syncflow

import (
	"context"

	"clipboard/internal/relay"
)

type RelayService struct {
	client *relay.Client
}

func NewRelayService(client *relay.Client) *RelayService {
	return &RelayService{client: client}
}

func (s *RelayService) Sync(ctx context.Context) error {
	return s.client.Send(ctx)
}
