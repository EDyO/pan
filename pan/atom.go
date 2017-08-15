// Copyright © 2017 Ignasi Fosch
//
//  This file is part of pan.
//
//  pan is free software: you can redistribute it and/or modify
//  it under the terms of the GNU Lesser General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  pan is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU Lesser General Public License for more details.
//
//  You should have received a copy of the GNU Lesser General Public License
//  along with pan. If not, see <http://www.gnu.org/licenses/>.
//

package pan

import (
	"encoding/xml"
)

// AtomLink represents an AtomLink tag.
type AtomLink struct {
	XMLName xml.Name `xml:"atom:link"`
	Href    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr"`
	Type    string   `xml:"type,attr"`
}

// AtomLinkFromMap is an AtomLink factory from map[interface{}]interface{}.
func AtomLinkFromMap(atomLinkMap map[interface{}]interface{}) AtomLink {
	return AtomLink{
		Href: atomLinkMap["href"].(string),
		Rel:  atomLinkMap["rel"].(string),
		Type: atomLinkMap["type"].(string),
	}
}

// Equal returns true if atomLink is equal to a, false otherwise.
func (a *AtomLink) Equal(atomLink AtomLink) bool {
	if a.Href != atomLink.Href ||
		a.Rel != atomLink.Rel ||
		a.Type != atomLink.Type {
		return false
	}
	return true
}
