package relay

import (
	"context"
	"errors"
)

var ErrTemporary = errors.New("temporary relay failure")

type AttemptFunc func(context.Context) error

type Client struct {
	attempt AttemptFunc
}

func NewClient(attempt AttemptFunc) *Client {
	return &Client{attempt: attempt}
}

func (c *Client) Send(ctx context.Context) error {
	err := c.attempt(ctx)
	if !errors.Is(err, ErrTemporary) {
		return err
	}
	return c.attempt(context.Background())
}
