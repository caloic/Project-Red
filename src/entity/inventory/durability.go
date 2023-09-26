package inventory

import "ftl/src/enum"

type Sword struct {
	Type       enum.Sword
	Durability int
}

func GetType(s *Sword) enum.Sword {
	return s.Type
}

func SetType(s *Sword, newType enum.Sword) {
	s.Type = newType
}

func GetDurability(s *Sword) int {
	return s.Durability
}

func SetDurability(s *Sword, newDurability int) {
	s.Durability = newDurability
}
