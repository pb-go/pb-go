package utils

import (
	"encoding/base64"
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
	"github.com/pb-go/pb-go/config"
)

func ContentValidityCheck(data []byte) bool {
	// https://stackoverflow.com/questions/42758054/read-multipart-form-data-as-byte-in-go/42758241
	detectedType, err := filetype.Match(data)
	if detectedType != types.Unknown && err != filetype.ErrEmptyBuffer {
		return false
	}
	if !config.ServConf.Content.Allow_Base64Encode {
		var abandonedDecoded []byte
		_, decodeErr1 := base64.RawStdEncoding.Decode(abandonedDecoded, data)
		_, decodeErr2 := base64.RawURLEncoding.Decode(abandonedDecoded, data)
		_, decodeErr3 := base64.StdEncoding.Decode(abandonedDecoded, data)
		_, decodeErr4 := base64.URLEncoding.Decode(abandonedDecoded, data)
		if decodeErr1 == nil || decodeErr2 == nil || decodeErr3 == nil || decodeErr4 == nil {
			return false
		}
	}
	return true
}
