package short

type Short struct {
	Logo string
	Size int
}

func (s *Short) GetLogo() string {
	return s.Logo
}

func (s *Short) GetSize() int {
	return s.Size
}
