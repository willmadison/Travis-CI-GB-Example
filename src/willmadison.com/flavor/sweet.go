package flavor

type Flavor interface {
	Taste() string
}

type Sweet int

func (s Sweet) Taste() string {
	return "Sweet"
}
