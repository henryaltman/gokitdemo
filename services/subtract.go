package services

// Subtract implement Subtract method
func (s BasicService) Subtract(a, b int) (int, error) {
	return a - b, nil
}
