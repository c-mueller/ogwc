package repo

import (
	"bytes"
	"math/rand"
	"time"
)

const idCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const idLen = 10

var charsetRunes []rune

func init() {
	charsetRunes = bytes.Runes([]byte(idCharset))
	rand.Seed(time.Now().Unix())
}

func generateID() string {
	id := ""
	for i := 0; i < idLen; i++ {
		idx := rand.Intn(len(charsetRunes))
		id += string(charsetRunes[idx])
	}
	return id
}
