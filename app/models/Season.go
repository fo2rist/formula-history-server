package models

import "fmt"

//Season is the calendar of Races for particular Year.
type Season struct {
	Year  int
	Races []Race
}

func (s *Season) String() string {
	return fmt.Sprintf("Season %v: %v", s.Year, s.Races)
}
