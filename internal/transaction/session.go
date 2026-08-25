package transaction

type Session struct {
	Commits   int
	Rollbacks int
}

func (s *Session) Finish(operationErr error) error {
	if operationErr != nil {
		s.Rollbacks++
		return operationErr
	}
	s.Commits++
	return nil
}
