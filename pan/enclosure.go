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

// Enclosure represents the definition of the resource shared.
type Enclosure struct {
	XMLName xml.Name `xml:"enclosure"`
	Length  string   `xml:"length"`
	Type    string   `xml:"type"`
	URL     string   `xml:"url"`
}

// Equal returns true if enclosure2 is equal to e, false otherwise.
func (e *Enclosure) Equal(enclosure Enclosure) bool {
	if e.Length != enclosure.Length ||
		e.Type != enclosure.Type ||
		e.URL != enclosure.URL {
		return false
	}
	return true
}

// UnmarshalYAML is the unmarshaler for Enclosure.
func (e *Enclosure) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var enclosure map[string]interface{}
	if err = unmarshal(&enclosure); err != nil {
		return
	}
	attributes := enclosure["attributes"].(map[interface{}]interface{})
	e.Length = strconv.Itoa(attributes["length"].(int))
	e.Type = attributes["type"].(string)
	e.URL = attributes["url"].(string)
	return
}
