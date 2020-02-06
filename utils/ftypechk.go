package utils

import (
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
)

// ContentValidityCheck : Check if has conflicted or empty invalid config in file
func ContentValidityCheck(data []byte) bool {
	// https://stackoverflow.com/questions/42758054/read-multipart-form-data-as-byte-in-go/42758241
	detectedType, err := filetype.Match(data)
	if detectedType != types.Unknown && err != filetype.ErrEmptyBuffer {
		return false
	}
	return true
}
