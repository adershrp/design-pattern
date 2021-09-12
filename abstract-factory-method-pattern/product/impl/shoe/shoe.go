package shoe

type Shoe struct {
	Logo string
	Size int
}

func (s *Shoe) GetLogo() string {
	return s.Logo
}

func (s *Shoe) GetSize() int {
	return s.Size
}
