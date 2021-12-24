package interface_pattern

import (
	"testing"
)

func TestParent(t *testing.T) {
	country := &Country{WithName{Name: "China"}}
	city := &City{WithName{Name: "Shanghai"}}

	country.PrintStr()
	city.PrintStr()
}
