package session


type Session struct {
	ch chan []byte
	*manager.Player
}

func NewSession(ch chan []byte) *Session {
	s := &Session{}
	s.ch = ch
	return s
}
