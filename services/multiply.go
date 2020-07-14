package services

// Multiply implement Multiply method
func (s BasicService) Multiply(a, b int) (int, error) {
	return a * b, nil
}
