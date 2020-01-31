package misc

import "os"

// IsFileExist return true if filename exists
func IsFileExist(filename string) bool {
	exist := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
