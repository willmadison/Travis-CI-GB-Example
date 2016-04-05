package flavor

import "testing"

func TestSourTaste(t *testing.T) {
	sours := Sour(7)

	taste := sours.Taste()

	if taste != "Sour" {
		t.Fatal(`Expected: "Sour", Got:`, taste)
	}
}
