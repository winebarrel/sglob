package sglob

type scanner struct {
	buf []rune
	pos int
}

func newScanner(s string) *scanner {
	return &scanner{
		buf: []rune(s),
		pos: 0,
	}
}

func (s *scanner) HasNext() bool {
	return s.pos < len(s.buf)
}

func (s *scanner) Next() (rune, bool) {
	if !s.HasNext() {
		return -1, false
	}

	c := s.buf[s.pos]
	s.pos++

	return c, true
}

func (s *scanner) Peek() (rune, bool) {
	if !s.HasNext() {
		return -1, false
	}

	return s.buf[s.pos], true
}

func (s *scanner) SkipWildcard() (int, bool) {
	if c, ok := s.Peek(); !ok || c != '*' {
		return -1, false
	}

	n := 0

	for i := s.pos; i < len(s.buf); i++ {
		c := s.buf[i]

		if c != '*' && c != '?' {
			break
		}

		if c == '?' {
			n++
		}

		s.pos++
	}

	return n, true
}

func (s *scanner) SkipBefore(x rune) (int, bool) {
	if !s.HasNext() {
		return -1, false
	}

	for i := len(s.buf) - 1; i >= s.pos; i-- {
		c := s.buf[i]

		if c == x {
			n := i - s.pos
			s.pos = i
			return n, true
		}
	}

	return -1, false
}

func (s *scanner) Rest() string {
	return string(s.buf[s.pos:len(s.buf)])
}

func (s *scanner) RestLen() int {
	return len(s.buf) - s.pos
}
