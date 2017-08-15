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
	"strconv"
)

// EnclosureFromMap is an Enclosure factory from map[interface{}]interface{}.
func EnclosureFromMap(enclosureMap map[interface{}]interface{}) Enclosure {
	attributes := enclosureMap["attributes"].(map[interface{}]interface{})
	if attributes["url"] == nil {
		attributes["url"] = ""
	}
	return Enclosure{
		Length: strconv.Itoa(attributes["length"].(int)),
		Type:   attributes["type"].(string),
		URL:    attributes["url"].(string),
	}
}

// Enclosure represents the definition of the resource shared.
type Enclosure struct {
	XMLName xml.Name `xml:"enclosure"`
	Length  string   `xml:"length,attr"`
	Type    string   `xml:"type,attr"`
	URL     string   `xml:"url,attr"`
}

// UnmarshalYAML is the unmarshaler for Enclosure.
func (e *Enclosure) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var enclosureMap map[interface{}]interface{}
	if err = unmarshal(&enclosureMap); err != nil {
		return
	}
	enclosure := EnclosureFromMap(enclosureMap)
	e.Length = enclosure.Length
	e.Type = enclosure.Type
	e.URL = enclosure.URL
	return
}
