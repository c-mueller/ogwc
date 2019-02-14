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

package core

func (p *Participant) AddFleetLoss(fleet Fleet) {
	if p.AdditionalLosses == nil {
		p.AdditionalLosses = &fleet
	} else {
		r := p.AdditionalLosses.Add(fleet)
		p.AdditionalLosses = &r
	}
}

func (p *Participant) SetFleetLoss(fleet Fleet) {
	p.AdditionalLosses = &fleet
}

func (p *Participant) AddResourceLoss(ressources Resources) {
	if p.AdditionalResourceLosses == nil {
		p.AdditionalResourceLosses = &ressources
	} else {
		r := p.AdditionalResourceLosses.Add(ressources)
		p.AdditionalResourceLosses = &r
	}
}

func (p *Participant) SetResourceLoss(resources Resources) {
	p.AdditionalResourceLosses = &resources
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
