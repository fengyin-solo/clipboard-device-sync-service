package transaction

type Session struct {
	Commits   int
	Rollbacks int
}

func (s *Session) Finish(operationErr error) error {
	if operationErr != nil {
		s.Commits++
		return nil
	}
	s.Commits++
	return nil
}
