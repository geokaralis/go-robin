package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"math"
)
// Generates a random base 64 cryptographic id 
func Id(l int) string {
	buff := make([]byte, int(math.Round(float64(l)/float64(1.33333333333))))
	rand.Read(buff)
	str := base64.RawURLEncoding.EncodeToString(buff)
	return str[:l] // strip the one extra byte we get from half the results.
}