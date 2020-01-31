package misc

import (
	"testing"
)

func TestIsFileExist(t *testing.T) {
	filename := "./path.go"
	if !IsFileExist(filename) {
		t.Errorf("%s does not exist\n", filename)
		return
	}

	filename = "some-file-not-exist"
	if IsFileExist(filename) {
		t.Errorf("%s should not exist\n", filename)
		return
	}
}
