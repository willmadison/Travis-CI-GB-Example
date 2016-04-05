package flavor

import "testing"

func TestSourTaste(t *testing.T) {
	sours := Sour(7)

	taste := sours.Taste()

	if taste != "Sweet" {
		t.Fatal(`Expected: "Sweet", Got:`, taste)
	}
}
