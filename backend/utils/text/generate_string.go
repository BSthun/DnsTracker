package text

import (
	"math/rand"
	"strings"
)

func GenerateString(charactors string, number int) *string {
	var generated strings.Builder
	for i := 0; i < number; i++ {
		random := rand.Intn(len(charactors))
		randomChar := charactors[random]
		generated.WriteString(string(randomChar))
	}

	var str = generated.String()
	return &str
}

var GenerateStringSet = struct {
	Num           string
	MixedAlphaNum string
	UpperAlpha    string
	UpperAlphaNum string
}{
	"0123456789",
	"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
}
