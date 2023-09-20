package levelManager

import "ftl/src/enum"

type Chapter struct {
	Levels      int
	ActualLevel int
	LevelZone   enum.Zone
}

func GetLevels(c *Chapter) int {
	return c.Levels
}

func SetLevels(c *Chapter, newLevels int) {
	c.Levels = newLevels
}

func GetActualLevel(c *Chapter) int {
	return c.ActualLevel
}

func SetActualLevel(c *Chapter, newActualLevel int) {
	c.ActualLevel = newActualLevel
}

func GetLevelZone(c *Chapter) enum.Zone {
	return c.LevelZone
}

func SetLevelZone(c *Chapter, newLevelZone enum.Zone) {
	c.LevelZone = newLevelZone
}
