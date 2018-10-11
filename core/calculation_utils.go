// ogwc (https://github.com/c-mueller/ogwc).
// Copyright (c) 2018 Christian Müller <cmueller.dev@gmail.com>.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package core

func (p *Participant) AddFleetLoss(f Fleet) {
	if p.AdditionalLosses == nil {
		p.AdditionalLosses = &f
	} else {
		r := p.AdditionalLosses.Add(f)
		p.AdditionalLosses = &r
	}
}

func (p ParticipantList) Find(name string) (int, *Participant) {
	for i, v := range p {
		if name == v.Name {
			return i, &v
		}
	}
	return -1, nil
}

func (p ParticipantList) IsPresent(name string) bool {
	for _, v := range p {
		if name == v.Name {
			return true
		}
	}
	return false
}
