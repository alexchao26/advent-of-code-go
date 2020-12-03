package util

import "testing"

// not a real test... used o run copy to clipboard without a main package
func TestCopyToClipboard(t *testing.T) {
	err := CopyToClipboard("asdfqwert")
	if err != nil {
		t.Errorf("Unexpected error while running CopyToClipboard: %w", err)
	}
}
