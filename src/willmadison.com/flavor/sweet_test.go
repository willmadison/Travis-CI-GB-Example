package flavor

import "testing"

func TestSweetTest(t *testing.T) {
	sweets := Sweet(16)

	taste := sweets.Taste()

	if taste != "Sweet" {
		t.Fatal(`Expected: "Sweet", Got:`, taste)
	}
}
