package inventory

import "ftl/src/enum"

type Weapon struct {
	Type              enum.Sword
	Durability        int
	Special           enum.Special
	Use               int
	SpecialDurability int
}

func GetType(s *Weapon) enum.Sword {
	return s.Type
}

func SetType(s *Weapon, newType enum.Sword) {
	s.Type = newType
}

func GetDurability(s *Weapon) int {
	return s.Durability
}

func SetDurability(s *Weapon, newDurability int) {
	s.Durability = newDurability
}

func GetSpecial(s *Weapon) enum.Special {
	return s.Special
}

func SetSpecial(s *Weapon, newSpecial enum.Special) {
	s.Special = newSpecial
}

func GetUse(s *Weapon) int {
	return s.Use
}

func SetUse(s *Weapon, newUse int) {
	s.Use = newUse
}

func GetSpecialDurability(s *Weapon) int {
	return s.SpecialDurability
}

func SetSpecialDurability(s *Weapon, newSpecialDurability int) {
	s.SpecialDurability = newSpecialDurability
}
