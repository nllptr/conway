package conway

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Read106 expects data to be formatted according to the
// Life 1.06 specification
func Read106(filePath string, w *World) error {
	r, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Could not open file '%s'", filePath)
	}
	scanner := bufio.NewScanner(r)
	if scanner.Scan(); scanner.Text() != "#Life 1.06" {
		return fmt.Errorf("Line 1: Wrong version string (was '%s', expected '#Life 1.06').", scanner.Text())
	}
	oX, oY := centerOffset(w)
	i := 1
	for scanner.Scan() {
		coords := strings.Fields(scanner.Text())
		if len(coords) != 2 {
			return fmt.Errorf("Invalid coordinate found in file.")
		}
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			return fmt.Errorf("Line %d: Cold not read X.", i)
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			return fmt.Errorf("Line %d: Could not read Y.", i)
		}
		if (x+oX) >= 0 && (x+oY) < len((*w)[0]) && (y+oY) >= 0 && (y+oY) < len((*w)) {
			(*w)[y+oY][x+oX] = 1
		}
		i++
	}
	return nil
}

/* this function is needed to make read coordinates (0,0)
appear in the middle of the world instead of at the upper left corner */
func centerOffset(w *World) (x, y int) {
	return (len((*w)[0]) - 1) / 2, (len((*w)) - 1) / 2
}
