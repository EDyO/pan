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
	Href    string   `xml:"href,attr,omitempty"`
	Rel     string   `xml:"rel,attr,omitempty"`
	Type    string   `xml:"type,attr,omitempty"`
}

// AtomLinkFromMap is an AtomLink factory from map[interface{}]interface{}.
func AtomLinkFromMap(atomLinkMap map[interface{}]interface{}) AtomLink {
	attributesMap := atomLinkMap["attributes"].(map[interface{}]interface{})
	return AtomLink{
		Href: attributesMap["href"].(string),
		Rel:  attributesMap["rel"].(string),
		Type: attributesMap["type"].(string),
	}
}

// UnmarshalYAML is the YAML unmarshaler for AtomLink.
func (a *AtomLink) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var atomLinkMap map[interface{}]interface{}
	if err = unmarshal(&atomLinkMap); err != nil {
		return
	}
	atomLink := AtomLinkFromMap(atomLinkMap)
	a.Href = atomLink.Href
	a.Rel = atomLink.Rel
	a.Type = atomLink.Type
	return
}
