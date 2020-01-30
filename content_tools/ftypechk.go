package content_tools

import (
	"encoding/base64"
	"errors"
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
	"github.com/kmahyyg/pb-go/config"
)

func ContentValidityCheck(data []byte, shortid string) error{
	// https://stackoverflow.com/questions/42758054/read-multipart-form-data-as-byte-in-go/42758241
	detectedType, err := filetype.Match(data)
	if detectedType != types.Unknown && err != filetype.ErrEmptyBuffer {
		return errors.New("File uploaded is illegal.")
	}
	if !config.ServConf.Content.Allow_Base64Encode {
		var abandonedDecoded []byte
		_, decodeErr1 := base64.RawStdEncoding.Decode(abandonedDecoded, data)
		_, decodeErr2 := base64.RawURLEncoding.Decode(abandonedDecoded, data)
		_, decodeErr3 := base64.StdEncoding.Decode(abandonedDecoded, data)
		_, decodeErr4 := base64.URLEncoding.Decode(abandonedDecoded, data)
		if decodeErr1 == nil || decodeErr2 == nil || decodeErr3 == nil || decodeErr4 == nil {
			return errors.New("Pure Encoded Data is not allowed to upload.")
		}
	}
	return nil
}

