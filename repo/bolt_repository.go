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
	"encoding/json"
	"github.com/c-mueller/ogwc/core"
	bolt "go.etcd.io/bbolt"
)

const bucketName = "reports"

var bucketNameBytes = []byte(bucketName)

type BoltRepository struct {
	Path string
	db   *bolt.DB
}

func (b *BoltRepository) Connect() error {
	db, err := bolt.Open(b.Path, 0666, bolt.DefaultOptions)
	if err != nil {
		return err
	}

	b.db = db

	return db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketNameBytes)
		return err
	})
}

func (b *BoltRepository) Insert(calculation core.CombatReportCalculation) string {
	data, _ := json.Marshal(calculation)

	uidString := ""
	var e *core.CombatReportCalculation
	for e != nil || len(uidString) == 0 {
		uidString = generateID()
		e = b.Get(uidString)
	}

	_ = b.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketNameBytes)

		return b.Put([]byte(uidString), data)
	})

	return uidString
}

func (b *BoltRepository) Update(uid string, calculation core.CombatReportCalculation) error {
	data, _ := json.Marshal(calculation)

	return b.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketNameBytes)

		return b.Put([]byte(uid), data)
	})
}

func (b *BoltRepository) Get(uid string) *core.CombatReportCalculation {

	data := make([]byte, 0)

	err := b.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketNameBytes)

		data = b.Get([]byte(uid))

		return nil
	})

	var calc core.CombatReportCalculation

	err = json.Unmarshal(data, &calc)
	if err != nil {
		return nil
	}

	return &calc
}
