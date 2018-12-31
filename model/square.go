package model

type (
	Square struct {
		Pos
		Highlighted bool
	}

)

func (s *Square) IsHighlighted() bool {
	return s.Highlighted
}