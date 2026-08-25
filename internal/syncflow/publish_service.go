package syncflow

import (
	"errors"

	"clipboard/internal/remote"
	"clipboard/internal/transaction"
)

type PublishService struct {
	session *transaction.Session
	send    func() error
}

func NewPublishService(session *transaction.Session, send func() error) *PublishService {
	return &PublishService{session: session, send: send}
}

func (s *PublishService) Publish() error {
	err := remote.Normalize(s.send())
	var failure *remote.Failure
	if errors.As(err, &failure) && failure.Temporary {
		err = remote.Normalize(s.send())
	}
	return s.session.Finish(err)
}
