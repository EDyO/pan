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

// ITunesImage represents an AtomLink tag.
type ITunesImage struct {
	XMLName xml.Name `xml:"itunes:image"`
	Href    string   `xml:"href,attr,omitempty"`
}

// ITunesImageFromMap is an ITunesImage factory from map[interface{}]interface{}.
func ITunesImageFromMap(iTunesImageMap map[interface{}]interface{}) ITunesImage {
	attributesMap := iTunesImageMap["attributes"].(map[interface{}]interface{})
	return ITunesImage{
		Href: attributesMap["href"].(string),
	}
}

// UnmarshalYAML is the YAML unmarshaler for ITunesImage.
func (i *ITunesImage) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var iTunesImageMap map[interface{}]interface{}
	if err = unmarshal(&iTunesImageMap); err != nil {
		return
	}
	iTunesImage := ITunesImageFromMap(iTunesImageMap)
	i.Href = iTunesImage.Href
	return
}
