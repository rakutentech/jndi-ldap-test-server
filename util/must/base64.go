// Package must contains function wrappers that MUST succeed without an error.
// If the wrapped function returns an error, they will panic.
package must

import (
	"encoding/base64"
	"github.com/rs/zerolog/log"
)

// DecodeBase64 ensures that base64 is decoded without an error or panics on error.
func DecodeBase64(value string) []byte {
	byteData, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		log.
			Panic().
			Err(err).
			Str("base64", value).
			Msg("Unexpected failure to decode base64 string")
	}
	return byteData
}
