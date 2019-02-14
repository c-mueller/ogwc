// ogwc (https://github.com/c-mueller/ogwc).
// Copyright (C) 2018-2019 Christian Müller <dev@c-mueller.xyz>.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
