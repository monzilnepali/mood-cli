package utils

// Include check whether item exist in slice
func Include(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
